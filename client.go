package steam

import (
	"bytes"
	"compress/gzip"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/benpye/go-steam/cryptoutil"
	"github.com/benpye/go-steam/netutil"
	"github.com/benpye/go-steam/protocol"
	"github.com/benpye/go-steam/protocol/protobuf/steam"
	"github.com/benpye/go-steam/protocol/steamlang"
	"github.com/benpye/go-steam/steamid"
)

var serverList ServerList

// Client represents a client to the Steam network.
// Always poll events from the channel returned by Events() or receiving messages will stop.
// All access, unless otherwise noted, should be threadsafe.
//
// When a FatalErrorEvent is emitted, the connection is automatically closed. The same client can be used to reconnect.
// Other errors don't have any effect.
type Client struct {
	// these need to be 64 bit aligned for sync/atomic on 32bit
	sessionID int32
	_         uint32
	steamID   uint64

	Auth          *Auth
	Social        *Social
	Notifications *Notifications
	Trading       *Trading
	GC            *GameCoordinator
	JobManager    *JobManager

	events        chan interface{}
	handlers      []PacketHandler
	handlersMutex sync.RWMutex

	tempSessionKey []byte

	ConnectionTimeout time.Duration

	mutex              sync.RWMutex // guarding conn and writeChan
	conn               connection
	currentServer      *netutil.PortAddr
	writeChan          chan protocol.IMsg
	writeBuf           *bytes.Buffer
	heartbeat          *time.Ticker
	disconnectExpected bool
}

// CallbackEvent is the interface all client callback events other than errors
// must implement.
type CallbackEvent interface {
	GetJobID() protocol.JobID
}

// PacketHandler is the packet handler interface that a type must match to
// be registered as a client packet handler.
type PacketHandler interface {
	HandlePacket(*protocol.Packet)
}

// NewClient creates a new instance of the Steam client type.
func NewClient() *Client {
	client := &Client{
		events:   make(chan interface{}, 3),
		writeBuf: new(bytes.Buffer),
	}
	client.JobManager = newJobManager(client)
	client.RegisterPacketHandler(client.JobManager)
	client.Auth = &Auth{client: client}
	client.RegisterPacketHandler(client.Auth)
	client.Social = newSocial(client)
	client.RegisterPacketHandler(client.Social)
	client.Notifications = newNotifications(client)
	client.RegisterPacketHandler(client.Notifications)
	client.Trading = &Trading{client: client}
	client.RegisterPacketHandler(client.Trading)
	client.GC = newGC(client)
	client.RegisterPacketHandler(client.GC)
	return client
}

// Events gets the event channel. By convention all events are pointers. It is never closed.
func (c *Client) Events() <-chan interface{} {
	return c.events
}

// Emit emits an event to the client events channel.
func (c *Client) Emit(event interface{}) {
	c.events <- event
	if callback, ok := event.(CallbackEvent); ok {
		c.JobManager.CompleteJob(callback)
	}
}

// Fatalf emits a FatalErrorEvent formatted with fmt.Errorf and disconnects.
func (c *Client) Fatalf(format string, a ...interface{}) {
	c.Emit(FatalErrorEvent(fmt.Errorf(format, a...)))
	c.disconnect(false)
}

// Errorf emits an ErrorEvent formatted with fmt.Errorf.
func (c *Client) Errorf(format string, a ...interface{}) {
	c.Emit(fmt.Errorf(format, a...))
}

// RegisterPacketHandler registers a PacketHandler that receives all incoming packets.
func (c *Client) RegisterPacketHandler(handler PacketHandler) {
	c.handlersMutex.Lock()
	defer c.handlersMutex.Unlock()
	c.handlers = append(c.handlers, handler)
}

// SteamID returns the SteamID for the currently logged on user.
func (c *Client) SteamID() steamid.SteamID {
	return steamid.SteamID(atomic.LoadUint64(&c.steamID))
}

// SessionID returns the current client session ID.
func (c *Client) SessionID() int32 {
	return atomic.LoadInt32(&c.sessionID)
}

// Connected indicates if the client is connected to teh Steam network.
func (c *Client) Connected() bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.conn != nil
}

// Connect Connects to a random Steam server and returns its address.
// If this client is already connected, it is disconnected first.
// This method tries to use an address from the Steam Directory and falls
// back to the built-in server list if the Steam Directory can't be reached.
// If you want to connect to a specific server, use `ConnectTo`.
func (c *Client) Connect() (*netutil.PortAddr, error) {
	server, err := serverList.GetServerCandidate()
	if err != nil {
		return nil, err
	}

	c.ConnectTo(server)
	return server, nil
}

// ConnectTo connects to a specific server.
// If this client is already connected, it is disconnected first.
func (c *Client) ConnectTo(addr *netutil.PortAddr) {
	c.ConnectToBind(addr, nil)
}

// ConnectToBind connects to a specific server, and binds to a specified local IP
// If this client is already connected, it is disconnected first.
func (c *Client) ConnectToBind(addr *netutil.PortAddr, local *net.TCPAddr) {
	c.disconnect(false)

	c.disconnectExpected = false
	conn, err := dialTCP(local, addr.ToTCPAddr())
	if err != nil {
		c.Fatalf("Connect failed: %v", err)
		return
	}
	c.currentServer = addr
	c.conn = conn
	c.writeChan = make(chan protocol.IMsg, 5)

	go c.readLoop()
	go c.writeLoop()
}

// Disconnect disconnects the client from the Steam network - if currently connected.
func (c *Client) Disconnect() {
	c.disconnect(true)
}

func (c *Client) disconnect(userRequested bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if !userRequested && !c.disconnectExpected {
		serverList.UpdateServerState(c.currentServer, ServerStateBad)
	}

	if c.conn == nil {
		return
	}

	c.conn.Close()
	c.conn = nil
	if c.heartbeat != nil {
		c.heartbeat.Stop()
	}
	close(c.writeChan)

	c.Emit(&DisconnectedEvent{
		UserRequested: userRequested || c.disconnectExpected,
	})

}

// Write adds a message to the send queue. Modifications to the given message
// after writing are not allowed (possible race conditions).
//
// Writes to this client when not connected are ignored.
func (c *Client) Write(msg protocol.IMsg) {
	if cm, ok := msg.(protocol.IClientMsg); ok {
		cm.SetSessionID(c.SessionID())
		cm.SetSteamID(c.SteamID())
	}
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	if c.conn == nil {
		return
	}
	c.writeChan <- msg
}

func (c *Client) readLoop() {
	for {
		// This *should* be atomic on most platforms, but the Go spec doesn't guarantee it
		c.mutex.RLock()
		conn := c.conn
		c.mutex.RUnlock()
		if conn == nil {
			return
		}
		packet, err := conn.Read()

		if err != nil {
			// Failed to read - connection closed?
			break
		}
		c.handlePacket(packet)
	}

	c.disconnect(false)
}

func (c *Client) writeLoop() {
	for {
		c.mutex.RLock()
		conn := c.conn
		c.mutex.RUnlock()
		if conn == nil {
			return
		}

		msg, ok := <-c.writeChan
		if !ok {
			return
		}

		err := msg.Serialize(c.writeBuf)
		if err != nil {
			c.writeBuf.Reset()
			c.Fatalf("Error serializing message %v: %v", msg, err)
			return
		}

		err = conn.Write(c.writeBuf.Bytes())

		c.writeBuf.Reset()

		if err != nil {
			// Failed to write - connection closed?
			// We don't disconnect here - the read side will handle it
			break
		}
	}
}

func (c *Client) heartbeatLoop(seconds time.Duration) {
	if c.heartbeat != nil {
		c.heartbeat.Stop()
	}
	c.heartbeat = time.NewTicker(seconds * time.Second)
	for {
		_, ok := <-c.heartbeat.C
		if !ok {
			break
		}
		c.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientHeartBeat, new(steam.CMsgClientHeartBeat)))
	}
	c.heartbeat = nil
}

func (c *Client) handlePacket(packet *protocol.Packet) {
	switch packet.EMsg {
	case steamlang.EMsg_ChannelEncryptRequest:
		c.handleChannelEncryptRequest(packet)
	case steamlang.EMsg_ChannelEncryptResult:
		c.handleChannelEncryptResult(packet)
	case steamlang.EMsg_Multi:
		c.handleMulti(packet)
	case steamlang.EMsg_ClientCMList:
		c.handleClientCMList(packet)
	}

	c.handlersMutex.RLock()
	defer c.handlersMutex.RUnlock()
	for _, handler := range c.handlers {
		handler.HandlePacket(packet)
	}
}

func (c *Client) handleChannelEncryptRequest(packet *protocol.Packet) {
	body := steamlang.NewMsgChannelEncryptRequest()
	packet.ReadMsg(body)

	if body.Universe != steamlang.EUniverse_Public {
		c.Fatalf("Invalid univserse %v!", body.Universe)
	}

	c.tempSessionKey = make([]byte, 32)
	rand.Read(c.tempSessionKey)
	encryptedKey := cryptoutil.RSAEncrypt(GetPublicKey(steamlang.EUniverse_Public), c.tempSessionKey)

	payload := new(bytes.Buffer)
	payload.Write(encryptedKey)
	binary.Write(payload, binary.LittleEndian, crc32.ChecksumIEEE(encryptedKey))
	payload.WriteByte(0)
	payload.WriteByte(0)
	payload.WriteByte(0)
	payload.WriteByte(0)

	c.Write(protocol.NewMsg(steamlang.NewMsgChannelEncryptResponse(), payload.Bytes()))
}

func (c *Client) handleChannelEncryptResult(packet *protocol.Packet) {
	body := steamlang.NewMsgChannelEncryptResult()
	packet.ReadMsg(body)

	if body.Result != steamlang.EResult_OK {
		c.Fatalf("Encryption failed: %v", body.Result)
		return
	}
	c.conn.SetEncryptionKey(c.tempSessionKey)
	c.tempSessionKey = nil

	c.Emit(&ConnectedEvent{})
}

func (c *Client) handleMulti(packet *protocol.Packet) {
	body := new(steam.CMsgMulti)
	packet.ReadProtoMsg(body)

	payload := body.GetMessageBody()

	if body.GetSizeUnzipped() > 0 {
		r, err := gzip.NewReader(bytes.NewReader(payload))
		if err != nil {
			c.Errorf("handleMulti: Error while decompressing: %v", err)
			return
		}

		payload, err = ioutil.ReadAll(r)
		if err != nil {
			c.Errorf("handleMulti: Error while decompressing: %v", err)
			return
		}
	}

	pr := bytes.NewReader(payload)
	for pr.Len() > 0 {
		var length uint32
		binary.Read(pr, binary.LittleEndian, &length)
		packetData := make([]byte, length)
		pr.Read(packetData)
		p, err := protocol.NewPacket(packetData)
		if err != nil {
			c.Errorf("Error reading packet in Multi msg %v: %v", packet, err)
			continue
		}
		c.handlePacket(p)
	}
}

func (c *Client) handleClientCMList(packet *protocol.Packet) {
	body := new(steam.CMsgClientCMList)
	packet.ReadProtoMsg(body)

	addresses := body.GetCmAddresses()
	ports := body.GetCmPorts()

	servers := make([]*netutil.PortAddr, 0, len(addresses))
	for idx, ip := range addresses {
		servers = append(servers, &netutil.PortAddr{
			IP:   netutil.ReadIPv4(ip),
			Port: uint16(ports[idx]),
		})
	}

	serverList.UpdateList(servers)

	c.Emit(&ClientCMListEvent{servers})
}
