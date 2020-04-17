package steam

import (
	"github.com/benpye/go-steam/protocol"
	"github.com/benpye/go-steam/protocol/protobuf/steam"
	"github.com/benpye/go-steam/protocol/steamlang"
	"github.com/benpye/go-steam/steamid"
	"google.golang.org/protobuf/proto"
)

// Provides access to the Steam client's part of Steam Trading, that is bootstrapping
// the trade.
// The trade itself is not handled by the Steam client itself, but it's a part of
// the Steam website.
//
// You'll receive a TradeProposedEvent when a friend proposes a trade. You can accept it with
// the RespondRequest method. You can request a trade yourself with RequestTrade.
type Trading struct {
	client *Client
}

type TradeRequestID uint32

func (t *Trading) HandlePacket(packet *protocol.Packet) {
	switch packet.EMsg {
	case steamlang.EMsg_EconTrading_InitiateTradeProposed:
		msg := new(steam.CMsgTrading_InitiateTradeRequest)
		packet.ReadProtoMsg(msg)
		t.client.Emit(&TradeProposedEvent{
			RequestID: TradeRequestID(msg.GetTradeRequestId()),
			Other:     steamid.SteamId(msg.GetOtherSteamid()),
		})
	case steamlang.EMsg_EconTrading_InitiateTradeResult:
		msg := new(steam.CMsgTrading_InitiateTradeResponse)
		packet.ReadProtoMsg(msg)
		t.client.Emit(&TradeResultEvent{
			RequestID: TradeRequestID(msg.GetTradeRequestId()),
			Response:  steamlang.EEconTradeResponse(msg.GetResponse()),
			Other:     steamid.SteamId(msg.GetOtherSteamid()),

			NumDaysSteamGuardRequired:            msg.GetSteamguardRequiredDays(),
			NumDaysNewDeviceCooldown:             msg.GetNewDeviceCooldownDays(),
			DefaultNumDaysPasswordResetProbation: msg.GetDefaultPasswordResetProbationDays(),
			NumDaysPasswordResetProbation:        msg.GetPasswordResetProbationDays(),
		})
	case steamlang.EMsg_EconTrading_StartSession:
		msg := new(steam.CMsgTrading_StartSession)
		packet.ReadProtoMsg(msg)
		t.client.Emit(&TradeSessionStartEvent{
			Other: steamid.SteamId(msg.GetOtherSteamid()),
		})
	}
}

// Requests a trade. You'll receive a TradeResultEvent if the request fails or
// if the friend accepted the trade.
func (t *Trading) RequestTrade(other steamid.SteamId) {
	t.client.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_EconTrading_InitiateTradeRequest, &steam.CMsgTrading_InitiateTradeRequest{
		OtherSteamid: proto.Uint64(uint64(other)),
	}))
}

// Responds to a TradeProposedEvent.
func (t *Trading) RespondRequest(requestID TradeRequestID, accept bool) {
	var resp uint32
	if accept {
		resp = 0
	} else {
		resp = 1
	}

	t.client.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_EconTrading_InitiateTradeResponse, &steam.CMsgTrading_InitiateTradeResponse{
		TradeRequestId: proto.Uint32(uint32(requestID)),
		Response:       proto.Uint32(resp),
	}))
}

// This cancels a request made with RequestTrade.
func (t *Trading) CancelRequest(other steamid.SteamId) {
	t.client.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_EconTrading_CancelTradeRequest, &steam.CMsgTrading_CancelTradeRequest{
		OtherSteamid: proto.Uint64(uint64(other)),
	}))
}
