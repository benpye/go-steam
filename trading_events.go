package steam

import (
	"github.com/benpye/go-steam/protocol"
	"github.com/benpye/go-steam/protocol/steamlang"
	"github.com/benpye/go-steam/steamid"
)

type TradeProposedEvent struct {
	RequestID TradeRequestID
	Other     steamid.SteamID `json:",string"`
}

func (e *TradeProposedEvent) GetJobID() protocol.JobID {
	return 0
}

type TradeResultEvent struct {
	RequestID TradeRequestID
	Response  steamlang.EEconTradeResponse
	Other     steamid.SteamID `json:",string"`
	// Number of days Steam Guard is required to have been active
	NumDaysSteamGuardRequired uint32
	// Number of days a new device cannot trade for.
	NumDaysNewDeviceCooldown uint32
	// Default number of days one cannot trade after a password reset.
	DefaultNumDaysPasswordResetProbation uint32
	// See above.
	NumDaysPasswordResetProbation uint32
}

func (e *TradeResultEvent) GetJobID() protocol.JobID {
	return 0
}

type TradeSessionStartEvent struct {
	Other steamid.SteamID `json:",string"`
}

func (e *TradeSessionStartEvent) GetJobID() protocol.JobID {
	return 0
}
