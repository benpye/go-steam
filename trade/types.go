package trade

import (
	"github.com/benpye/go-steam/trade/tradeapi"
)

type TradeEndedEvent struct {
	Reason TradeEndReason
}

type TradeEndReason uint

const (
	TradeEndReason_Complete  TradeEndReason = 1
	TradeEndReason_Cancelled                = 2
	TradeEndReason_Timeout                  = 3
	TradeEndReason_Failed                   = 4
)

func newItem(event *tradeapi.Event) *Item {
	return &Item{
		event.AppID,
		event.ContextID,
		event.AssetID,
	}
}

type Item struct {
	AppID     uint32
	ContextID uint64
	AssetID   uint64
}

type ItemAddedEvent struct {
	Item *Item
}

type ItemRemovedEvent struct {
	Item *Item
}

type ReadyEvent struct{}
type UnreadyEvent struct{}

func newCurrency(event *tradeapi.Event) *Currency {
	return &Currency{
		event.AppID,
		event.ContextID,
		event.CurrencyID,
	}
}

type Currency struct {
	AppID      uint32
	ContextID  uint64
	CurrencyID uint64
}

type SetCurrencyEvent struct {
	Currency  *Currency
	OldAmount uint64
	NewAmount uint64
}

type ChatEvent struct {
	Message string
}
