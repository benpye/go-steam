package steam

import (
	"time"

	"github.com/benpye/go-steam/protocol/steamlang"
	"github.com/benpye/go-steam/steamid"
)

type FriendsListEvent struct{}

type FriendStateEvent struct {
	SteamID      steamid.SteamId `json:",string"`
	Relationship steamlang.EFriendRelationship
}

func (f *FriendStateEvent) IsFriend() bool {
	return f.Relationship == steamlang.EFriendRelationship_Friend
}

type GroupStateEvent struct {
	SteamID      steamid.SteamId `json:",string"`
	Relationship steamlang.EClanRelationship
}

func (g *GroupStateEvent) IsMember() bool {
	return g.Relationship == steamlang.EClanRelationship_Member
}

// Fired when someone changing their friend details
type PersonaStateEvent struct {
	StatusFlags            steamlang.EClientPersonaStateFlag
	FriendID               steamid.SteamId `json:",string"`
	State                  steamlang.EPersonaState
	StateFlags             steamlang.EPersonaStateFlag
	GameAppID              uint32
	GameID                 uint64 `json:",string"`
	GameName               string
	GameServerIP           uint32
	GameServerPort         uint32
	QueryPort              uint32
	SourceSteamID          steamid.SteamId `json:",string"`
	GameDataBlob           []byte
	Name                   string
	Avatar                 []byte
	LastLogOff             uint32
	LastLogOn              uint32
	ClanRank               uint32
	ClanTag                string
	OnlineSessionInstances uint32
	PersonaSetByUser       bool
}

// Fired when a clan's state has been changed
type ClanStateEvent struct {
	ClanID              steamid.SteamId `json:",string"`
	AccountFlags        steamlang.EAccountFlags
	ClanName            string
	Avatar              []byte
	MemberTotalCount    uint32
	MemberOnlineCount   uint32
	MemberChattingCount uint32
	MemberInGameCount   uint32
	Events              []ClanEventDetails
	Announcements       []ClanEventDetails
}

type ClanEventDetails struct {
	ID         uint64 `json:",string"`
	EventTime  uint32
	Headline   string
	GameID     uint64 `json:",string"`
	JustPosted bool
}

// Fired in response to adding a friend to your friends list
type FriendAddedEvent struct {
	Result      steamlang.EResult
	SteamID     steamid.SteamId `json:",string"`
	PersonaName string
}

// Fired when the client receives a message from either a friend or a chat room
type ChatMsgEvent struct {
	ChatRoomID steamid.SteamId `json:",string"` // not set for friend messages
	ChatterID  steamid.SteamId `json:",string"`
	Message    string
	EntryType  steamlang.EChatEntryType
	Timestamp  time.Time
	Offline    bool
}

// Whether the type is ChatMsg
func (c *ChatMsgEvent) IsMessage() bool {
	return c.EntryType == steamlang.EChatEntryType_ChatMsg
}

// Fired in response to joining a chat
type ChatEnterEvent struct {
	ChatRoomID    steamid.SteamId `json:",string"`
	FriendID      steamid.SteamId `json:",string"`
	ChatRoomType  steamlang.EChatRoomType
	OwnerID       steamid.SteamId `json:",string"`
	ClanID        steamid.SteamId `json:",string"`
	ChatFlags     byte
	EnterResponse steamlang.EChatRoomEnterResponse
	Name          string
}

// Fired in response to a chat member's info being received
type ChatMemberInfoEvent struct {
	ChatRoomID      steamid.SteamId `json:",string"`
	Type            steamlang.EChatInfoType
	StateChangeInfo StateChangeDetails
}

type StateChangeDetails struct {
	ChatterActedOn steamid.SteamId `json:",string"`
	StateChange    steamlang.EChatMemberStateChange
	ChatterActedBy steamid.SteamId `json:",string"`
}

// Fired when a chat action has completed
type ChatActionResultEvent struct {
	ChatRoomID steamid.SteamId `json:",string"`
	ChatterID  steamid.SteamId `json:",string"`
	Action     steamlang.EChatAction
	Result     steamlang.EChatActionResult
}

// Fired when a chat invite is received
type ChatInviteEvent struct {
	InvitedID    steamid.SteamId `json:",string"`
	ChatRoomID   steamid.SteamId `json:",string"`
	PatronID     steamid.SteamId `json:",string"`
	ChatRoomType steamlang.EChatRoomType
	FriendChatID steamid.SteamId `json:",string"`
	ChatRoomName string
	GameID       uint64 `json:",string"`
}

// Fired in response to ignoring a friend
type IgnoreFriendEvent struct {
	Result steamlang.EResult
}

// Fired in response to requesting profile info for a user
type ProfileInfoEvent struct {
	Result      steamlang.EResult
	SteamID     steamid.SteamId `json:",string"`
	TimeCreated uint32
	RealName    string
	CityName    string
	StateName   string
	CountryName string
	Headline    string
	Summary     string
}
