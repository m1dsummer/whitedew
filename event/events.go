package event

type Event interface {
	GetNoticeType() string
	GetTime() int64
}

type MetaEvent struct {
	Time       int64  `json:"time"`
	SelfId     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	SubType    string `json:"sub_type"`
}

func (m MetaEvent) GetNoticeType() string {
	return m.NoticeType
}

func (m MetaEvent) GetTime() int64 {
	return m.Time
}

func (m MetaEvent) GetSubType() string {
	return m.SubType
}

type Handler func(e Event)

func (h Handler) Handle(e Event) {
	h(e)
}

type GroupEvent struct {
	MetaEvent
	GroupId int64 `json:"group_id"`
	UserId  int64 `json:"user_id"`
}

type _File struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	BusId int64  `json:"busid"`
}

// GroupUploadEvent
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/event/notice.md#%E7%BE%A4%E6%96%87%E4%BB%B6%E4%B8%8A%E4%BC%A0
type GroupUploadEvent struct {
	GroupEvent
	File _File `json:"file"`
}

// GroupAdminEvent
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/event/notice.md#%E4%BA%8B%E4%BB%B6%E6%95%B0%E6%8D%AE-1
type GroupAdminEvent struct {
	GroupEvent
}

// GroupDecreaseEvent
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/event/notice.md#%E7%BE%A4%E6%88%90%E5%91%98%E5%87%8F%E5%B0%91
type GroupDecreaseEvent struct {
	GroupEvent
	OperatorId int64 `json:"operator_id"`
}

// GroupIncreaseEvent
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/event/notice.md#%E7%BE%A4%E6%88%90%E5%91%98%E5%A2%9E%E5%8A%A0
type GroupIncreaseEvent struct {
	GroupEvent
	OperatorId int64 `json:"operator_id"`
}

// GroupBanEvent
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/event/notice.md#%E7%BE%A4%E7%A6%81%E8%A8%80
type GroupBanEvent struct {
	GroupEvent
	OperatorId int64 `json:"operator_id"`
	Duration   int64 `json:"duration"`
}

// FriendAddEvent
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/event/notice.md#%E5%A5%BD%E5%8F%8B%E6%B7%BB%E5%8A%A0
type FriendAddEvent struct {
	MetaEvent
	UserId int64 `json:"user_id"`
}

// GroupRecallEvent
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/event/notice.md#%E7%BE%A4%E6%B6%88%E6%81%AF%E6%92%A4%E5%9B%9E
type GroupRecallEvent struct {
	GroupEvent
	OperatorId int64 `json:"operator_id"`
	MessageId  int64 `json:"message_id"`
}

// FriendRecallEvent
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/event/notice.md#%E5%A5%BD%E5%8F%8B%E6%B6%88%E6%81%AF%E6%92%A4%E5%9B%9E
type FriendRecallEvent struct {
	MetaEvent
	MessageId int64 `json:"message_id"`
}

// PokeEvent
//
// https://github.com/botuniverse/onebot/blob/master/v11/specs/event/notice.md#%E7%BE%A4%E5%86%85%E6%88%B3%E4%B8%80%E6%88%B3
type PokeEvent struct {
	GroupEvent
	TargetId int64 `json:"target_id"`
}

// LuckyKingEvent
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/event/notice.md#%E7%BE%A4%E7%BA%A2%E5%8C%85%E8%BF%90%E6%B0%94%E7%8E%8B
type LuckyKingEvent struct {
	GroupEvent
	TargetId int64 `json:"target_id"`
}

// HonorEvent
//
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/event/notice.md#%E7%BE%A4%E6%88%90%E5%91%98%E8%8D%A3%E8%AA%89%E5%8F%98%E6%9B%B4
type HonorEvent struct {
	GroupEvent
	HonorType string `json:"honor_type"`
}

