package steam

import (
	"github.com/benpye/go-steam/protocol/steamlang"
	"github.com/benpye/go-steam/steamid"
)

type LoggedOnEvent struct {
	Result                    steamlang.EResult
	ExtendedResult            steamlang.EResult
	OutOfGameSecsPerHeartbeat int32
	InGameSecsPerHeartbeat    int32
	PublicIP                  uint32
	ServerTime                uint32
	AccountFlags              steamlang.EAccountFlags
	ClientSteamID             steamid.SteamID `json:",string"`
	EmailDomain               string
	CellID                    uint32
	CellIDPingThreshold       uint32
	Steam2Ticket              []byte
	UsePics                   bool
	WebAPIUserNonce           string
	IPCountryCode             string
	VanityURL                 string
	NumLoginFailuresToMigrate int32
	NumDisconnectsToMigrate   int32
}

type LogOnFailedEvent struct {
	Result steamlang.EResult
}

type LoginKeyEvent struct {
	UniqueID uint32
	LoginKey string
}

type LoggedOffEvent struct {
	Result steamlang.EResult
}

type MachineAuthUpdateEvent struct {
	Hash []byte
}

type AccountInfoEvent struct {
	PersonaName          string
	Country              string
	CountAuthedComputers int32
	AccountFlags         steamlang.EAccountFlags
	FacebookID           uint64 `json:",string"`
	FacebookName         string
}
