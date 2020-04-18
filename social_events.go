package steam

import (
	"time"

	"github.com/benpye/go-steam/protocol"
	"github.com/benpye/go-steam/protocol/steamlang"
	"github.com/benpye/go-steam/steamid"
)

type FriendsListEvent struct{}

func (e *FriendsListEvent) GetJobID() protocol.JobID {
	return 0
}

type FriendStateEvent struct {
	SteamID      steamid.SteamID `json:",string"`
	Relationship steamlang.EFriendRelationship
}

func (f *FriendStateEvent) GetJobID() protocol.JobID {
	return 0
}

func (f *FriendStateEvent) IsFriend() bool {
	return f.Relationship == steamlang.EFriendRelationship_Friend
}

type GroupStateEvent struct {
	SteamID      steamid.SteamID `json:",string"`
	Relationship steamlang.EClanRelationship
}

func (g *GroupStateEvent) GetJobID() protocol.JobID {
	return 0
}

func (g *GroupStateEvent) IsMember() bool {
	return g.Relationship == steamlang.EClanRelationship_Member
}

// Fired when someone changing their friend details
type PersonaStateEvent struct {
	StatusFlags            steamlang.EClientPersonaStateFlag
	FriendID               steamid.SteamID `json:",string"`
	State                  steamlang.EPersonaState
	StateFlags             steamlang.EPersonaStateFlag
	GameAppID              uint32
	GameID                 uint64 `json:",string"`
	GameName               string
	GameServerIP           uint32
	GameServerPort         uint32
	QueryPort              uint32
	SourceSteamID          steamid.SteamID `json:",string"`
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

func (e *PersonaStateEvent) GetJobID() protocol.JobID {
	return 0
}

// Fired when a clan's state has been changed
type ClanStateEvent struct {
	ClanID              steamid.SteamID `json:",string"`
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

func (e *ClanStateEvent) GetJobID() protocol.JobID {
	return 0
}

type ClanEventDetails struct {
	ID         uint64 `json:",string"`
	EventTime  uint32
	Headline   string
	GameID     uint64 `json:",string"`
	JustPosted bool
}

func (e *ClanEventDetails) GetJobID() protocol.JobID {
	return 0
}

// Fired in response to adding a friend to your friends list
type FriendAddedEvent struct {
	Result      steamlang.EResult
	SteamID     steamid.SteamID `json:",string"`
	PersonaName string
}

func (e *FriendAddedEvent) GetJobID() protocol.JobID {
	return 0
}

// Fired when the client receives a message from either a friend or a chat room
type ChatMsgEvent struct {
	ChatRoomID steamid.SteamID `json:",string"` // not set for friend messages
	ChatterID  steamid.SteamID `json:",string"`
	Message    string
	EntryType  steamlang.EChatEntryType
	Timestamp  time.Time
	Offline    bool
}

func (e *ChatMsgEvent) GetJobID() protocol.JobID {
	return 0
}

// Whether the type is ChatMsg
func (c *ChatMsgEvent) IsMessage() bool {
	return c.EntryType == steamlang.EChatEntryType_ChatMsg
}

// Fired in response to joining a chat
type ChatEnterEvent struct {
	ChatRoomID    steamid.SteamID `json:",string"`
	FriendID      steamid.SteamID `json:",string"`
	ChatRoomType  steamlang.EChatRoomType
	OwnerID       steamid.SteamID `json:",string"`
	ClanID        steamid.SteamID `json:",string"`
	ChatFlags     byte
	EnterResponse steamlang.EChatRoomEnterResponse
	Name          string
}

func (e *ChatEnterEvent) GetJobID() protocol.JobID {
	return 0
}

// Fired in response to a chat member's info being received
type ChatMemberInfoEvent struct {
	ChatRoomID      steamid.SteamID `json:",string"`
	Type            steamlang.EChatInfoType
	StateChangeInfo StateChangeDetails
}

func (e *ChatMemberInfoEvent) GetJobID() protocol.JobID {
	return 0
}

type StateChangeDetails struct {
	ChatterActedOn steamid.SteamID `json:",string"`
	StateChange    steamlang.EChatMemberStateChange
	ChatterActedBy steamid.SteamID `json:",string"`
}

func (e *StateChangeDetails) GetJobID() protocol.JobID {
	return 0
}

// Fired when a chat action has completed
type ChatActionResultEvent struct {
	ChatRoomID steamid.SteamID `json:",string"`
	ChatterID  steamid.SteamID `json:",string"`
	Action     steamlang.EChatAction
	Result     steamlang.EChatActionResult
}

func (e *ChatActionResultEvent) GetJobID() protocol.JobID {
	return 0
}

// Fired when a chat invite is received
type ChatInviteEvent struct {
	InvitedID    steamid.SteamID `json:",string"`
	ChatRoomID   steamid.SteamID `json:",string"`
	PatronID     steamid.SteamID `json:",string"`
	ChatRoomType steamlang.EChatRoomType
	FriendChatID steamid.SteamID `json:",string"`
	ChatRoomName string
	GameID       uint64 `json:",string"`
}

func (e *ChatInviteEvent) GetJobID() protocol.JobID {
	return 0
}

// Fired in response to ignoring a friend
type IgnoreFriendEvent struct {
	Result steamlang.EResult
}

func (e *IgnoreFriendEvent) GetJobID() protocol.JobID {
	return 0
}

// Fired in response to requesting profile info for a user
type ProfileInfoEvent struct {
	Result      steamlang.EResult
	SteamID     steamid.SteamID `json:",string"`
	TimeCreated uint32
	RealName    string
	CityName    string
	StateName   string
	CountryName string
	Headline    string
	Summary     string
}

func (e *ProfileInfoEvent) GetJobID() protocol.JobID {
	return 0
}
