package steam

import (
	"net"
	"time"

	"github.com/benpye/go-steam/protocol"
	"github.com/benpye/go-steam/protocol/protobuf/steam"
	"github.com/benpye/go-steam/protocol/steamlang"
	"github.com/benpye/go-steam/steamid"
)

type LoggedOnEvent struct {
	Result                    steamlang.EResult
	ExtendedResult            steamlang.EResult
	OutOfGameSecsPerHeartbeat int32
	InGameSecsPerHeartbeat    int32
	PublicIP                  net.IP
	ServerTime                time.Time
	AccountFlags              steamlang.EAccountFlags
	ClientSteamID             steamid.SteamID `json:",string"`
	EmailDomain               string
	CellID                    uint32
	CellIDPingThreshold       uint32
	Steam2Ticket              []byte
	UsePICS                   bool
	WebAPIUserNonce           string
	IPCountryCode             string
	VanityURL                 string
	NumLoginFailuresToMigrate int32
	NumDisconnectsToMigrate   int32
	ParentalSettings          *steam.ParentalSettings
}

func (e *LoggedOnEvent) GetJobID() protocol.JobID {
	return 0
}

type LoginKeyEvent struct {
	UniqueID uint32
	LoginKey string
}

func (e *LoginKeyEvent) GetJobID() protocol.JobID {
	return 0
}

type LoggedOffEvent struct {
	Result steamlang.EResult
}

func (e *LoggedOffEvent) GetJobID() protocol.JobID {
	return 0
}

type MachineAuthUpdateEvent struct {
	Hash []byte
}

func (e *MachineAuthUpdateEvent) GetJobID() protocol.JobID {
	return 0
}

type AccountInfoEvent struct {
	PersonaName          string
	Country              string
	CountAuthedComputers int32
	AccountFlags         steamlang.EAccountFlags
	FacebookID           uint64 `json:",string"`
	FacebookName         string
}

func (e *AccountInfoEvent) GetJobID() protocol.JobID {
	return 0
}

type WebAPIUserNonceEvent struct {
	JobID  protocol.JobID
	Result steamlang.EResult
	Nonce  string
}

func (e *WebAPIUserNonceEvent) GetJobID() protocol.JobID {
	return e.JobID
}
