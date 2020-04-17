package netutil

import (
	"net"
	"strconv"
	"strings"

	"github.com/benpye/go-steam/protocol/protobuf/steam"
)

// An addr that is neither restricted to TCP nor UDP, but has an IP and a port.
type PortAddr struct {
	IP   net.IP
	Port uint16
}

// Parses an IP address with a port, for example "209.197.29.196:27017".
// If the given string is not valid, this function returns nil.
func ParsePortAddr(addr string) *PortAddr {
	parts := strings.Split(addr, ":")
	if len(parts) != 2 {
		return nil
	}
	ip := net.ParseIP(parts[0])
	if ip == nil {
		return nil
	}
	port, err := strconv.ParseUint(parts[1], 10, 16)
	if err != nil {
		return nil
	}
	return &PortAddr{ip, uint16(port)}
}

func (p *PortAddr) ToTCPAddr() *net.TCPAddr {
	return &net.TCPAddr{IP: p.IP, Port: int(p.Port)}
}

func (p *PortAddr) ToUDPAddr() *net.UDPAddr {
	return &net.UDPAddr{IP: p.IP, Port: int(p.Port)}
}

func (p *PortAddr) String() string {
	return p.IP.String() + ":" + strconv.FormatUint(uint64(p.Port), 10)
}

func ReadIPv4(ip uint32) net.IP {
	r := make(net.IP, 4)
	r[3] = byte(ip)
	r[2] = byte(ip >> 8)
	r[1] = byte(ip >> 16)
	r[0] = byte(ip >> 24)
	return r
}

func ParseIPAddress(addr *steam.CMsgIPAddress) net.IP {
	if addr.GetV6() != nil {
		return addr.GetV6()
	}

	return ReadIPv4(addr.GetV4())
}
