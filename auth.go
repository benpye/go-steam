package steam

import (
	"context"
	"crypto/sha1"
	"sync/atomic"
	"time"

	"github.com/benpye/go-steam/netutil"
	"github.com/benpye/go-steam/protocol"
	"github.com/benpye/go-steam/protocol/protobuf/steam"
	"github.com/benpye/go-steam/protocol/steamlang"
	"github.com/benpye/go-steam/steamid"
	"google.golang.org/protobuf/proto"
)

type Auth struct {
	client  *Client
	details *LogOnDetails
}

type SentryHash []byte

type LogOnDetails struct {
	Username string

	// If logging into an account without a login key, the account's password.
	Password string

	// If you have a Steam Guard email code, you can provide it here.
	AuthCode string

	// If you have a Steam Guard mobile two-factor authentication code, you can provide it here.
	TwoFactorCode  string
	SentryFileHash SentryHash
	LoginKey       string

	// true if you want to get a login key which can be used in lieu of
	// a password for subsequent logins. false or omitted otherwise.
	ShouldRememberPassword bool
}

// LogOn logs on to Steam with the given details. You must always specify username and
// password OR username and loginkey. For the first login, don't set an authcode or a hash and you'll
// receive an error (EResult_AccountLogonDenied)
// and Steam will send you an authcode. Then you have to login again, this time with the authcode.
// Shortly after logging in, you'll receive a MachineAuthUpdateEvent with a hash which allows
// you to login without using an authcode in the future.
//
// If you don't use Steam Guard, username and password are enough.
//
// After the event EMsg_ClientNewLoginKey is received you can use the LoginKey
// to login instead of using the password.
func (a *Auth) LogOn(details *LogOnDetails) {
	if details.Username == "" {
		panic("Username must be set!")
	}
	if details.Password == "" && details.LoginKey == "" {
		panic("Password or LoginKey must be set!")
	}

	logOn := new(steam.CMsgClientLogon)
	logOn.AccountName = &details.Username
	logOn.Password = &details.Password
	if details.AuthCode != "" {
		logOn.AuthCode = proto.String(details.AuthCode)
	}
	if details.TwoFactorCode != "" {
		logOn.TwoFactorCode = proto.String(details.TwoFactorCode)
	}
	logOn.ClientLanguage = proto.String("english")
	logOn.ProtocolVersion = proto.Uint32(steamlang.MsgClientLogon_CurrentProtocol)
	logOn.ShaSentryfile = details.SentryFileHash
	if details.LoginKey != "" {
		logOn.LoginKey = proto.String(details.LoginKey)
	}
	if details.ShouldRememberPassword {
		logOn.ShouldRememberPassword = proto.Bool(details.ShouldRememberPassword)
	}

	atomic.StoreUint64(&a.client.steamID, uint64(steamid.NewIDAdv(0, 1, int32(steamlang.EUniverse_Public), int32(steamlang.EAccountType_Individual))))

	a.client.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientLogon, logOn))
}

// LogOff logs off of Steam cleanly and disconnects.
func (a *Auth) LogOff() {
	logOff := new(steam.CMsgClientLogOff)
	a.client.disconnectExpected = true
	a.client.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientLogOff, logOff))
}

// AcceptNewLoginKey accepts a new login key from a LoginKeyEvent.
func (a *Auth) AcceptNewLoginKey(event *LoginKeyEvent) {
	a.client.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientNewLoginKeyAccepted, &steam.CMsgClientNewLoginKeyAccepted{
		UniqueId: proto.Uint32(event.UniqueID),
	}))
}

// RequestWebAPIUserNonce requests a nonce to authenticate the user against the web API.
func (a *Auth) RequestWebAPIUserNonce(ctx context.Context) (*WebAPIUserNonceEvent, error) {
	body := new(steam.CMsgClientRequestWebAPIAuthenticateUserNonce)
	msg := protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientRequestWebAPIAuthenticateUserNonce, body)

	job := a.client.JobManager.NewJob()
	msg.SetSourceJobID(job.JobID)
	a.client.Write(msg)

	ev, err := job.Wait(ctx)
	return ev.(*WebAPIUserNonceEvent), err
}

// HandlePacket recieves all packets from the Steam3 network.
func (a *Auth) HandlePacket(packet *protocol.Packet) {
	switch packet.EMsg {
	case steamlang.EMsg_ClientLogOnResponse:
		a.handleLogOnResponse(packet)
	case steamlang.EMsg_ClientNewLoginKey:
		a.handleNewLoginKey(packet)
	case steamlang.EMsg_ClientSessionToken:
	case steamlang.EMsg_ClientLoggedOff:
		a.handleLoggedOff(packet)
	case steamlang.EMsg_ClientUpdateMachineAuth:
		a.handleUpdateMachineAuth(packet)
	case steamlang.EMsg_ClientAccountInfo:
		a.handleAccountInfo(packet)
	case steamlang.EMsg_ClientRequestWebAPIAuthenticateUserNonceResponse:
		a.handleAuthNonceResponse(packet)
	}
}

func (a *Auth) handleLogOnResponse(packet *protocol.Packet) {
	if !packet.IsProto {
		a.client.Fatalf("Got non-proto logon response!")
		return
	}

	body := new(steam.CMsgClientLogonResponse)
	msg := packet.ReadProtoMsg(body)

	result := steamlang.EResult(body.GetEresult())
	if result == steamlang.EResult_OK {
		serverList.UpdateServerState(a.client.currentServer, ServerStateGood)

		atomic.StoreInt32(&a.client.sessionID, msg.Header.Proto.GetClientSessionid())
		atomic.StoreUint64(&a.client.steamID, msg.Header.Proto.GetSteamid())

		go a.client.heartbeatLoop(time.Duration(body.GetOutOfGameHeartbeatSeconds()))
	} else if result == steamlang.EResult_TryAnotherCM || result == steamlang.EResult_ServiceUnavailable {
		serverList.UpdateServerState(a.client.currentServer, ServerStateBad)
	}

	var parentalSettings *steam.ParentalSettings
	if body.GetParentalSettings() != nil {
		parentalSettings = new(steam.ParentalSettings)
		proto.Unmarshal(body.GetParentalSettings(), parentalSettings)
	}

	a.client.Emit(&LoggedOnEvent{
		Result:                    steamlang.EResult(body.GetEresult()),
		ExtendedResult:            steamlang.EResult(body.GetEresultExtended()),
		OutOfGameSecsPerHeartbeat: body.GetOutOfGameHeartbeatSeconds(),
		InGameSecsPerHeartbeat:    body.GetInGameHeartbeatSeconds(),
		PublicIP:                  netutil.ParseIPAddress(body.GetPublicIp()),
		ServerTime:                time.Unix(int64(body.GetRtime32ServerTime()), 0),
		AccountFlags:              steamlang.EAccountFlags(body.GetAccountFlags()),
		ClientSteamID:             steamid.SteamID(body.GetClientSuppliedSteamid()),
		EmailDomain:               body.GetEmailDomain(),
		CellID:                    body.GetCellId(),
		CellIDPingThreshold:       body.GetCellIdPingThreshold(),
		Steam2Ticket:              body.GetSteam2Ticket(),
		UsePICS:                   body.GetUsePics(),
		WebAPIUserNonce:           body.GetWebapiAuthenticateUserNonce(),
		IPCountryCode:             body.GetIpCountryCode(),
		VanityURL:                 body.GetVanityUrl(),
		NumLoginFailuresToMigrate: body.GetCountLoginfailuresToMigrate(),
		NumDisconnectsToMigrate:   body.GetCountDisconnectsToMigrate(),
		ParentalSettings:          parentalSettings,
	})
}

func (a *Auth) handleNewLoginKey(packet *protocol.Packet) {
	body := new(steam.CMsgClientNewLoginKey)
	packet.ReadProtoMsg(body)
	a.client.Emit(&LoginKeyEvent{
		UniqueID: body.GetUniqueId(),
		LoginKey: body.GetLoginKey(),
	})
}

func (a *Auth) handleLoggedOff(packet *protocol.Packet) {
	result := steamlang.EResult_Invalid
	if packet.IsProto {
		body := new(steam.CMsgClientLoggedOff)
		packet.ReadProtoMsg(body)
		result = steamlang.EResult(body.GetEresult())
	} else {
		body := new(steamlang.MsgClientLoggedOff)
		packet.ReadClientMsg(body)
		result = body.Result
	}
	a.client.Emit(&LoggedOffEvent{Result: result})
}

func (a *Auth) handleUpdateMachineAuth(packet *protocol.Packet) {
	body := new(steam.CMsgClientUpdateMachineAuth)
	packet.ReadProtoMsg(body)
	hash := sha1.New()
	hash.Write(packet.Data)
	sha := hash.Sum(nil)

	msg := protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientUpdateMachineAuthResponse, &steam.CMsgClientUpdateMachineAuthResponse{
		ShaFile: sha,
	})
	msg.SetTargetJobID(packet.SourceJobID)
	a.client.Write(msg)

	a.client.Emit(&MachineAuthUpdateEvent{sha})
}

func (a *Auth) handleAccountInfo(packet *protocol.Packet) {
	body := new(steam.CMsgClientAccountInfo)
	packet.ReadProtoMsg(body)
	a.client.Emit(&AccountInfoEvent{
		PersonaName:          body.GetPersonaName(),
		Country:              body.GetIpCountry(),
		CountAuthedComputers: body.GetCountAuthedComputers(),
		AccountFlags:         steamlang.EAccountFlags(body.GetAccountFlags()),
		FacebookID:           body.GetFacebookId(),
		FacebookName:         body.GetFacebookName(),
	})
}

func (a *Auth) handleAuthNonceResponse(packet *protocol.Packet) {
	// this has to be the best name for a message yet.
	msg := new(steam.CMsgClientRequestWebAPIAuthenticateUserNonceResponse)
	packet.ReadProtoMsg(msg)

	msg.GetTokenType()

	a.client.Emit(&WebAPIUserNonceEvent{
		JobID:  packet.TargetJobID,
		Result: steamlang.EResult(msg.GetEresult()),
		Nonce:  msg.GetWebapiAuthenticateUserNonce(),
	})
}
