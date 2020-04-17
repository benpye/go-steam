package steam

import (
	"bytes"
	"encoding/binary"
	"io"
	"sync"
	"time"

	"github.com/benpye/go-steam/protocol"
	"github.com/benpye/go-steam/protocol/protobuf/steam"
	"github.com/benpye/go-steam/protocol/steamlang"
	"github.com/benpye/go-steam/rwu"
	"github.com/benpye/go-steam/socialcache"
	"github.com/benpye/go-steam/steamid"
	"google.golang.org/protobuf/proto"
)

// Provides access to social aspects of Steam.
type Social struct {
	mutex sync.RWMutex

	name         string
	avatar       []byte
	personaState steamlang.EPersonaState

	Friends *socialcache.FriendsList
	Groups  *socialcache.GroupsList
	Chats   *socialcache.ChatsList

	client *Client
}

func newSocial(client *Client) *Social {
	return &Social{
		Friends: socialcache.NewFriendsList(),
		Groups:  socialcache.NewGroupsList(),
		Chats:   socialcache.NewChatsList(),
		client:  client,
	}
}

// Gets the local user's avatar
func (s *Social) GetAvatar() []byte {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.avatar
}

// Gets the local user's persona name
func (s *Social) GetPersonaName() string {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.name
}

// Sets the local user's persona name and broadcasts it over the network
func (s *Social) SetPersonaName(name string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.name = name
	s.client.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientChangeStatus, &steam.CMsgClientChangeStatus{
		PersonaState: proto.Uint32(uint32(s.personaState)),
		PlayerName:   proto.String(name),
	}))
}

// Gets the local user's persona state
func (s *Social) GetPersonaState() steamlang.EPersonaState {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.personaState
}

// Sets the local user's persona state and broadcasts it over the network
func (s *Social) SetPersonaState(state steamlang.EPersonaState) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.personaState = state
	s.client.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientChangeStatus, &steam.CMsgClientChangeStatus{
		PersonaState: proto.Uint32(uint32(state)),
	}))
}

// Sends a chat message to ether a room or friend
func (s *Social) SendMessage(to steamid.SteamId, entryType steamlang.EChatEntryType, message string) {
	//Friend
	if to.GetAccountType() == int32(steamlang.EAccountType_Individual) || to.GetAccountType() == int32(steamlang.EAccountType_ConsoleUser) {
		s.client.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientFriendMsg, &steam.CMsgClientFriendMsg{
			Steamid:       proto.Uint64(to.ToUint64()),
			ChatEntryType: proto.Int32(int32(entryType)),
			Message:       []byte(message),
		}))
		//Chat room
	} else if to.GetAccountType() == int32(steamlang.EAccountType_Clan) || to.GetAccountType() == int32(steamlang.EAccountType_Chat) {
		chatID := to.ClanToChat()
		s.client.Write(protocol.NewClientMsg(&steamlang.MsgClientChatMsg{
			ChatMsgType:     entryType,
			SteamIdChatRoom: chatID,
			SteamIdChatter:  s.client.SteamID(),
		}, []byte(message)))
	}
}

// Adds a friend to your friends list or accepts a friend. You'll receive a FriendStateEvent
// for every new/changed friend
func (s *Social) AddFriend(id steamid.SteamId) {
	s.client.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientAddFriend, &steam.CMsgClientAddFriend{
		SteamidToAdd: proto.Uint64(id.ToUint64()),
	}))
}

// Removes a friend from your friends list
func (s *Social) RemoveFriend(id steamid.SteamId) {
	s.client.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientRemoveFriend, &steam.CMsgClientRemoveFriend{
		Friendid: proto.Uint64(id.ToUint64()),
	}))
}

// Ignores or unignores a friend on Steam
func (s *Social) IgnoreFriend(id steamid.SteamId, setIgnore bool) {
	ignore := uint8(1) //True
	if !setIgnore {
		ignore = uint8(0) //False
	}
	s.client.Write(protocol.NewClientMsg(&steamlang.MsgClientSetIgnoreFriend{
		MySteamId:     s.client.SteamID(),
		SteamIdFriend: id,
		Ignore:        ignore,
	}, make([]byte, 0)))
}

// Requests persona state for a list of specified SteamIds
func (s *Social) RequestFriendListInfo(ids []steamid.SteamId, requestedInfo steamlang.EClientPersonaStateFlag) {
	var friends []uint64
	for _, id := range ids {
		friends = append(friends, id.ToUint64())
	}
	s.client.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientRequestFriendData, &steam.CMsgClientRequestFriendData{
		PersonaStateRequested: proto.Uint32(uint32(requestedInfo)),
		Friends:               friends,
	}))
}

// Requests persona state for a specified SteamId
func (s *Social) RequestFriendInfo(id steamid.SteamId, requestedInfo steamlang.EClientPersonaStateFlag) {
	s.RequestFriendListInfo([]steamid.SteamId{id}, requestedInfo)
}

// Requests profile information for a specified SteamId
func (s *Social) RequestProfileInfo(id steamid.SteamId) {
	s.client.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientFriendProfileInfo, &steam.CMsgClientFriendProfileInfo{
		SteamidFriend: proto.Uint64(id.ToUint64()),
	}))
}

// Requests all offline messages and marks them as read
func (s *Social) RequestOfflineMessages() {
	s.client.Write(protocol.NewClientMsgProtobuf(steamlang.EMsg_ClientChatGetFriendMessageHistoryForOfflineMessages, &steam.CMsgClientChatGetFriendMessageHistoryForOfflineMessages{}))
}

// Attempts to join a chat room
func (s *Social) JoinChat(id steamid.SteamId) {
	chatID := id.ClanToChat()
	s.client.Write(protocol.NewClientMsg(&steamlang.MsgClientJoinChat{
		SteamIdChat: chatID,
	}, make([]byte, 0)))
}

// Attempts to leave a chat room
func (s *Social) LeaveChat(id steamid.SteamId) {
	chatID := id.ClanToChat()
	payload := new(bytes.Buffer)
	binary.Write(payload, binary.LittleEndian, s.client.SteamID().ToUint64())                 // ChatterActedOn
	binary.Write(payload, binary.LittleEndian, uint32(steamlang.EChatMemberStateChange_Left)) // StateChange
	binary.Write(payload, binary.LittleEndian, s.client.SteamID().ToUint64())                 // ChatterActedBy
	s.client.Write(protocol.NewClientMsg(&steamlang.MsgClientChatMemberInfo{
		SteamIdChat: chatID,
		Type:        steamlang.EChatInfoType_StateChange,
	}, payload.Bytes()))
}

// Kicks the specified chat member from the given chat room
func (s *Social) KickChatMember(room steamid.SteamId, user steamid.SteamId) {
	chatID := room.ClanToChat()
	s.client.Write(protocol.NewClientMsg(&steamlang.MsgClientChatAction{
		SteamIdChat:        chatID,
		SteamIdUserToActOn: user,
		ChatAction:         steamlang.EChatAction_Kick,
	}, make([]byte, 0)))
}

// Bans the specified chat member from the given chat room
func (s *Social) BanChatMember(room steamid.SteamId, user steamid.SteamId) {
	chatID := room.ClanToChat()
	s.client.Write(protocol.NewClientMsg(&steamlang.MsgClientChatAction{
		SteamIdChat:        chatID,
		SteamIdUserToActOn: user,
		ChatAction:         steamlang.EChatAction_Ban,
	}, make([]byte, 0)))
}

// Unbans the specified chat member from the given chat room
func (s *Social) UnbanChatMember(room steamid.SteamId, user steamid.SteamId) {
	chatID := room.ClanToChat()
	s.client.Write(protocol.NewClientMsg(&steamlang.MsgClientChatAction{
		SteamIdChat:        chatID,
		SteamIdUserToActOn: user,
		ChatAction:         steamlang.EChatAction_UnBan,
	}, make([]byte, 0)))
}

func (s *Social) HandlePacket(packet *protocol.Packet) {
	switch packet.EMsg {
	case steamlang.EMsg_ClientPersonaState:
		s.handlePersonaState(packet)
	case steamlang.EMsg_ClientClanState:
		s.handleClanState(packet)
	case steamlang.EMsg_ClientFriendsList:
		s.handleFriendsList(packet)
	case steamlang.EMsg_ClientFriendMsgIncoming:
		s.handleFriendMsg(packet)
	case steamlang.EMsg_ClientAccountInfo:
		s.handleAccountInfo(packet)
	case steamlang.EMsg_ClientAddFriendResponse:
		s.handleFriendResponse(packet)
	case steamlang.EMsg_ClientChatEnter:
		s.handleChatEnter(packet)
	case steamlang.EMsg_ClientChatMsg:
		s.handleChatMsg(packet)
	case steamlang.EMsg_ClientChatMemberInfo:
		s.handleChatMemberInfo(packet)
	case steamlang.EMsg_ClientChatActionResult:
		s.handleChatActionResult(packet)
	case steamlang.EMsg_ClientChatInvite:
		s.handleChatInvite(packet)
	case steamlang.EMsg_ClientSetIgnoreFriendResponse:
		s.handleIgnoreFriendResponse(packet)
	case steamlang.EMsg_ClientFriendProfileInfoResponse:
		s.handleProfileInfoResponse(packet)
	case steamlang.EMsg_ClientFSGetFriendMessageHistoryResponse:
		s.handleFriendMessageHistoryResponse(packet)
	}
}

func (s *Social) handleAccountInfo(packet *protocol.Packet) {
	//Just fire the personainfo, Auth handles the callback
	flags := steamlang.EClientPersonaStateFlag_PlayerName | steamlang.EClientPersonaStateFlag_Presence | steamlang.EClientPersonaStateFlag_SourceID
	s.RequestFriendInfo(s.client.SteamID(), steamlang.EClientPersonaStateFlag(flags))
}

func (s *Social) handleFriendsList(packet *protocol.Packet) {
	list := new(steam.CMsgClientFriendsList)
	packet.ReadProtoMsg(list)
	var friends []steamid.SteamId
	for _, friend := range list.GetFriends() {
		steamID := steamid.SteamId(friend.GetUlfriendid())
		isClan := steamID.GetAccountType() == int32(steamlang.EAccountType_Clan)

		if isClan {
			rel := steamlang.EClanRelationship(friend.GetEfriendrelationship())
			if rel == steamlang.EClanRelationship_None {
				s.Groups.Remove(steamID)
			} else {
				s.Groups.Add(socialcache.Group{
					SteamId:      steamID,
					Relationship: rel,
				})

			}
			if list.GetBincremental() {
				s.client.Emit(&GroupStateEvent{steamID, rel})
			}
		} else {
			rel := steamlang.EFriendRelationship(friend.GetEfriendrelationship())
			if rel == steamlang.EFriendRelationship_None {
				s.Friends.Remove(steamID)
			} else {
				s.Friends.Add(socialcache.Friend{
					SteamId:      steamID,
					Relationship: rel,
				})

			}
			if list.GetBincremental() {
				s.client.Emit(&FriendStateEvent{steamID, rel})
			}
		}
		if !list.GetBincremental() {
			friends = append(friends, steamID)
		}
	}
	if !list.GetBincremental() {
		s.RequestFriendListInfo(friends, protocol.EClientPersonaStateFlag_DefaultInfoRequest)
		s.client.Emit(&FriendsListEvent{})
	}
}

func (s *Social) handlePersonaState(packet *protocol.Packet) {
	list := new(steam.CMsgClientPersonaState)
	packet.ReadProtoMsg(list)
	flags := steamlang.EClientPersonaStateFlag(list.GetStatusFlags())
	for _, friend := range list.GetFriends() {
		id := steamid.SteamId(friend.GetFriendid())
		if id == s.client.SteamID() { //this is our client id
			s.mutex.Lock()
			if friend.GetPlayerName() != "" {
				s.name = friend.GetPlayerName()
			}
			avatar := friend.GetAvatarHash()
			if protocol.ValidAvatar(avatar) {
				s.avatar = avatar
			}
			s.mutex.Unlock()
		} else if id.GetAccountType() == int32(steamlang.EAccountType_Individual) {
			if (flags & steamlang.EClientPersonaStateFlag_PlayerName) == steamlang.EClientPersonaStateFlag_PlayerName {
				if friend.GetPlayerName() != "" {
					s.Friends.SetName(id, friend.GetPlayerName())
				}
			}
			if (flags & steamlang.EClientPersonaStateFlag_Presence) == steamlang.EClientPersonaStateFlag_Presence {
				avatar := friend.GetAvatarHash()
				if protocol.ValidAvatar(avatar) {
					s.Friends.SetAvatar(id, avatar)
				}
				s.Friends.SetPersonaState(id, steamlang.EPersonaState(friend.GetPersonaState()))
				s.Friends.SetPersonaStateFlags(id, steamlang.EPersonaStateFlag(friend.GetPersonaStateFlags()))
			}
			if (flags & steamlang.EClientPersonaStateFlag_GameDataBlob) == steamlang.EClientPersonaStateFlag_GameDataBlob {
				s.Friends.SetGameAppId(id, friend.GetGamePlayedAppId())
				s.Friends.SetGameId(id, friend.GetGameid())
				s.Friends.SetGameName(id, friend.GetGameName())
			}
		} else if id.GetAccountType() == int32(steamlang.EAccountType_Clan) {
			if (flags & steamlang.EClientPersonaStateFlag_PlayerName) == steamlang.EClientPersonaStateFlag_PlayerName {
				if friend.GetPlayerName() != "" {
					s.Groups.SetName(id, friend.GetPlayerName())
				}
			}
			if (flags & steamlang.EClientPersonaStateFlag_Presence) == steamlang.EClientPersonaStateFlag_Presence {
				avatar := friend.GetAvatarHash()
				if protocol.ValidAvatar(avatar) {
					s.Groups.SetAvatar(id, avatar)
				}
			}
		}
		s.client.Emit(&PersonaStateEvent{
			StatusFlags:            flags,
			FriendID:               id,
			State:                  steamlang.EPersonaState(friend.GetPersonaState()),
			StateFlags:             steamlang.EPersonaStateFlag(friend.GetPersonaStateFlags()),
			GameAppID:              friend.GetGamePlayedAppId(),
			GameID:                 friend.GetGameid(),
			GameName:               friend.GetGameName(),
			GameServerIP:           friend.GetGameServerIp(),
			GameServerPort:         friend.GetGameServerPort(),
			QueryPort:              friend.GetQueryPort(),
			SourceSteamID:          steamid.SteamId(friend.GetSteamidSource()),
			GameDataBlob:           friend.GetGameDataBlob(),
			Name:                   friend.GetPlayerName(),
			Avatar:                 friend.GetAvatarHash(),
			LastLogOff:             friend.GetLastLogoff(),
			LastLogOn:              friend.GetLastLogon(),
			ClanRank:               friend.GetClanRank(),
			ClanTag:                friend.GetClanTag(),
			OnlineSessionInstances: friend.GetOnlineSessionInstances(),
			PersonaSetByUser:       friend.GetPersonaSetByUser(),
		})
	}
}

func (s *Social) handleClanState(packet *protocol.Packet) {
	body := new(steam.CMsgClientClanState)
	packet.ReadProtoMsg(body)
	var name string
	var avatar []byte
	if body.GetNameInfo() != nil {
		name = body.GetNameInfo().GetClanName()
		avatar = body.GetNameInfo().GetShaAvatar()
	}
	var totalCount, onlineCount, chattingCount, ingameCount uint32
	if body.GetUserCounts() != nil {
		usercounts := body.GetUserCounts()
		totalCount = usercounts.GetMembers()
		onlineCount = usercounts.GetOnline()
		chattingCount = usercounts.GetChatting()
		ingameCount = usercounts.GetInGame()
	}
	var events, announcements []ClanEventDetails
	for _, event := range body.GetEvents() {
		events = append(events, ClanEventDetails{
			ID:         event.GetGid(),
			EventTime:  event.GetEventTime(),
			Headline:   event.GetHeadline(),
			GameID:     event.GetGameId(),
			JustPosted: event.GetJustPosted(),
		})
	}
	for _, announce := range body.GetAnnouncements() {
		announcements = append(announcements, ClanEventDetails{
			ID:         announce.GetGid(),
			EventTime:  announce.GetEventTime(),
			Headline:   announce.GetHeadline(),
			GameID:     announce.GetGameId(),
			JustPosted: announce.GetJustPosted(),
		})
	}

	//Add stuff to group
	clanid := steamid.SteamId(body.GetSteamidClan())
	if body.NameInfo != nil {
		info := body.NameInfo
		s.Groups.SetName(clanid, info.GetClanName())
		s.Groups.SetAvatar(clanid, info.GetShaAvatar())
	}
	if body.GetUserCounts() != nil {
		s.Groups.SetMemberTotalCount(clanid, totalCount)
		s.Groups.SetMemberOnlineCount(clanid, onlineCount)
		s.Groups.SetMemberChattingCount(clanid, chattingCount)
		s.Groups.SetMemberInGameCount(clanid, ingameCount)
	}
	s.client.Emit(&ClanStateEvent{
		ClanID:              clanid,
		AccountFlags:        steamlang.EAccountFlags(body.GetClanAccountFlags()),
		ClanName:            name,
		Avatar:              avatar,
		MemberTotalCount:    totalCount,
		MemberOnlineCount:   onlineCount,
		MemberChattingCount: chattingCount,
		MemberInGameCount:   ingameCount,
		Events:              events,
		Announcements:       announcements,
	})
}

func (s *Social) handleFriendResponse(packet *protocol.Packet) {
	body := new(steam.CMsgClientAddFriendResponse)
	packet.ReadProtoMsg(body)
	s.client.Emit(&FriendAddedEvent{
		Result:      steamlang.EResult(body.GetEresult()),
		SteamID:     steamid.SteamId(body.GetSteamIdAdded()),
		PersonaName: body.GetPersonaNameAdded(),
	})
}

func (s *Social) handleFriendMsg(packet *protocol.Packet) {
	body := new(steam.CMsgClientFriendMsgIncoming)
	packet.ReadProtoMsg(body)
	message := string(bytes.Split(body.GetMessage(), []byte{0x0})[0])
	s.client.Emit(&ChatMsgEvent{
		ChatterID: steamid.SteamId(body.GetSteamidFrom()),
		Message:   message,
		EntryType: steamlang.EChatEntryType(body.GetChatEntryType()),
		Timestamp: time.Unix(int64(body.GetRtime32ServerTimestamp()), 0),
	})
}

func (s *Social) handleChatMsg(packet *protocol.Packet) {
	body := new(steamlang.MsgClientChatMsg)
	payload := packet.ReadClientMsg(body).Payload
	message := string(bytes.Split(payload, []byte{0x0})[0])
	s.client.Emit(&ChatMsgEvent{
		ChatRoomID: steamid.SteamId(body.SteamIdChatRoom),
		ChatterID:  steamid.SteamId(body.SteamIdChatter),
		Message:    message,
		EntryType:  steamlang.EChatEntryType(body.ChatMsgType),
	})
}

func (s *Social) handleChatEnter(packet *protocol.Packet) {
	body := new(steamlang.MsgClientChatEnter)
	payload := packet.ReadClientMsg(body).Payload
	reader := bytes.NewBuffer(payload)
	name, _ := rwu.ReadString(reader)
	rwu.ReadByte(reader) //0
	count := body.NumMembers
	chatID := steamid.SteamId(body.SteamIdChat)
	clanID := steamid.SteamId(body.SteamIdClan)
	s.Chats.Add(socialcache.Chat{SteamId: chatID, GroupId: clanID})
	for i := 0; i < int(count); i++ {
		id, chatPerm, clanPerm := readChatMember(reader)
		rwu.ReadBytes(reader, 6) //No idea what this is
		s.Chats.AddChatMember(chatID, socialcache.ChatMember{
			SteamId:         steamid.SteamId(id),
			ChatPermissions: chatPerm,
			ClanPermissions: clanPerm,
		})
	}
	s.client.Emit(&ChatEnterEvent{
		ChatRoomID:    steamid.SteamId(body.SteamIdChat),
		FriendID:      steamid.SteamId(body.SteamIdFriend),
		ChatRoomType:  steamlang.EChatRoomType(body.ChatRoomType),
		OwnerID:       steamid.SteamId(body.SteamIdOwner),
		ClanID:        steamid.SteamId(body.SteamIdClan),
		ChatFlags:     byte(body.ChatFlags),
		EnterResponse: steamlang.EChatRoomEnterResponse(body.EnterResponse),
		Name:          name,
	})
}

func (s *Social) handleChatMemberInfo(packet *protocol.Packet) {
	body := new(steamlang.MsgClientChatMemberInfo)
	payload := packet.ReadClientMsg(body).Payload
	reader := bytes.NewBuffer(payload)
	chatID := steamid.SteamId(body.SteamIdChat)
	if body.Type == steamlang.EChatInfoType_StateChange {
		actedOn, _ := rwu.ReadUint64(reader)
		state, _ := rwu.ReadInt32(reader)
		actedBy, _ := rwu.ReadUint64(reader)
		rwu.ReadByte(reader) //0
		stateChange := steamlang.EChatMemberStateChange(state)
		if stateChange == steamlang.EChatMemberStateChange_Entered {
			_, chatPerm, clanPerm := readChatMember(reader)
			s.Chats.AddChatMember(chatID, socialcache.ChatMember{
				SteamId:         steamid.SteamId(actedOn),
				ChatPermissions: chatPerm,
				ClanPermissions: clanPerm,
			})
		} else if stateChange == steamlang.EChatMemberStateChange_Banned || stateChange == steamlang.EChatMemberStateChange_Kicked ||
			stateChange == steamlang.EChatMemberStateChange_Disconnected || stateChange == steamlang.EChatMemberStateChange_Left {
			s.Chats.RemoveChatMember(chatID, steamid.SteamId(actedOn))
		}
		stateInfo := StateChangeDetails{
			ChatterActedOn: steamid.SteamId(actedOn),
			StateChange:    steamlang.EChatMemberStateChange(stateChange),
			ChatterActedBy: steamid.SteamId(actedBy),
		}
		s.client.Emit(&ChatMemberInfoEvent{
			ChatRoomID:      steamid.SteamId(body.SteamIdChat),
			Type:            steamlang.EChatInfoType(body.Type),
			StateChangeInfo: stateInfo,
		})
	}
}

func readChatMember(r io.Reader) (steamid.SteamId, steamlang.EChatPermission, steamlang.EClanPermission) {
	rwu.ReadString(r) // MessageObject
	rwu.ReadByte(r)   // 7
	rwu.ReadString(r) //steamid
	id, _ := rwu.ReadUint64(r)
	rwu.ReadByte(r)   // 2
	rwu.ReadString(r) //Permissions
	chat, _ := rwu.ReadInt32(r)
	rwu.ReadByte(r)   // 2
	rwu.ReadString(r) //Details
	clan, _ := rwu.ReadInt32(r)
	return steamid.SteamId(id), steamlang.EChatPermission(chat), steamlang.EClanPermission(clan)
}

func (s *Social) handleChatActionResult(packet *protocol.Packet) {
	body := new(steamlang.MsgClientChatActionResult)
	packet.ReadClientMsg(body)
	s.client.Emit(&ChatActionResultEvent{
		ChatRoomID: steamid.SteamId(body.SteamIdChat),
		ChatterID:  steamid.SteamId(body.SteamIdUserActedOn),
		Action:     steamlang.EChatAction(body.ChatAction),
		Result:     steamlang.EChatActionResult(body.ActionResult),
	})
}

func (s *Social) handleChatInvite(packet *protocol.Packet) {
	body := new(steam.CMsgClientChatInvite)
	packet.ReadProtoMsg(body)
	s.client.Emit(&ChatInviteEvent{
		InvitedID:    steamid.SteamId(body.GetSteamIdInvited()),
		ChatRoomID:   steamid.SteamId(body.GetSteamIdChat()),
		PatronID:     steamid.SteamId(body.GetSteamIdPatron()),
		ChatRoomType: steamlang.EChatRoomType(body.GetChatroomType()),
		FriendChatID: steamid.SteamId(body.GetSteamIdFriendChat()),
		ChatRoomName: body.GetChatName(),
		GameID:       body.GetGameId(),
	})
}

func (s *Social) handleIgnoreFriendResponse(packet *protocol.Packet) {
	body := new(steamlang.MsgClientSetIgnoreFriendResponse)
	packet.ReadClientMsg(body)
	s.client.Emit(&IgnoreFriendEvent{
		Result: steamlang.EResult(body.Result),
	})
}

func (s *Social) handleProfileInfoResponse(packet *protocol.Packet) {
	body := new(steam.CMsgClientFriendProfileInfoResponse)
	packet.ReadProtoMsg(body)
	s.client.Emit(&ProfileInfoEvent{
		Result:      steamlang.EResult(body.GetEresult()),
		SteamID:     steamid.SteamId(body.GetSteamidFriend()),
		TimeCreated: body.GetTimeCreated(),
		RealName:    body.GetRealName(),
		CityName:    body.GetCityName(),
		StateName:   body.GetStateName(),
		CountryName: body.GetCountryName(),
		Headline:    body.GetHeadline(),
		Summary:     body.GetSummary(),
	})
}

func (s *Social) handleFriendMessageHistoryResponse(packet *protocol.Packet) {
	body := new(steam.CMsgClientChatGetFriendMessageHistoryResponse)
	packet.ReadProtoMsg(body)
	steamid := steamid.SteamId(body.GetSteamid())
	for _, message := range body.GetMessages() {
		if !message.GetUnread() {
			continue // Skip already read messages
		}
		s.client.Emit(&ChatMsgEvent{
			ChatterID: steamid,
			Message:   message.GetMessage(),
			EntryType: steamlang.EChatEntryType_ChatMsg,
			Timestamp: time.Unix(int64(message.GetTimestamp()), 0),
			Offline:   true, // GetUnread is true
		})
	}
}
