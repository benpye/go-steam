package steam

import (
	"github.com/benpye/go-steam/protocol/steamlang"
	"github.com/benpye/go-steam/steamid"
)

type TradeProposedEvent struct {
	RequestID TradeRequestID
	Other     steamid.SteamId `json:",string"`
}

type TradeResultEvent struct {
	RequestID TradeRequestID
	Response  steamlang.EEconTradeResponse
	Other     steamid.SteamId `json:",string"`
	// Number of days Steam Guard is required to have been active
	NumDaysSteamGuardRequired uint32
	// Number of days a new device cannot trade for.
	NumDaysNewDeviceCooldown uint32
	// Default number of days one cannot trade after a password reset.
	DefaultNumDaysPasswordResetProbation uint32
	// See above.
	NumDaysPasswordResetProbation uint32
}

type TradeSessionStartEvent struct {
	Other steamid.SteamId `json:",string"`
}
