// implement message type from OpenBot v11
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/event/message.md

package entity

// Message
// General message type
type Message interface {
	GetMsgType() string
	GetSender() Sender
	GetContent() string
	GetSelfId() int64
	GetRowMessage() string
}

// Sender
// General sender type
type Sender interface {
	GetId() int64
}

type AnonymousUser struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Flag string `json:"flag"`
}

func (a AnonymousUser) GetId() int64 {
	return a.Id
}

type metaSender struct {
	UserId   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int32  `json:"age"`
}

type PrivateMsgSender struct {
	metaSender
}

func (p PrivateMsgSender) GetId() int64 {
	return p.UserId
}

type GroupSender struct {
	metaSender
	Card  string `json:"card"`
	Area  string `json:"area"`
	Level string `json:"level"`
	Role  string `json:"role"`
	Title string `json:"title"`
}

func (g GroupSender) GetId() int64 {
	return g.UserId
}

type metaMessage struct {
	Time        int64  `json:"time"`
	SelfId      int64  `json:"self_id"`
	PostType    string `json:"post_type"`
	MessageType string `json:"message_type"`
	SubType     string `json:"sub_type"`
	MessageId   int32  `json:"message_id"`
	UserId      int64  `json:"user_id"`
	Message     string `json:"message"`
	RawMessage  string `json:"raw_message"`
	Font        int32  `json:"font"`
}

func (m metaMessage) GetContent() string {
	return m.Message
}

func (m metaMessage) GetMsgType() string {
	return m.MessageType
}

func (m metaMessage) GetSelfId() int64 {
	return m.SelfId
}

func (m metaMessage) GetRowMessage() string {
	return m.RawMessage
}

// PrivateMessage
// the implements of private message from OpenBot v11
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/event/message.md#%E7%A7%81%E8%81%8A%E6%B6%88%E6%81%AF
type PrivateMessage struct {
	metaMessage                  // inherit from MetaMessage
	Sender      PrivateMsgSender `json:"sender"`
	Anonymous   AnonymousUser    `json:"anonymous"`
}

func (p PrivateMessage) GetSender() Sender {
	return p.Sender
}

// GroupMessage
// the implements of group message type from OpenBot v11
// Reference: https://github.com/botuniverse/onebot/blob/master/v11/specs/event/message.md#%E7%BE%A4%E6%B6%88%E6%81%AF
type GroupMessage struct {
	metaMessage             // inherit from MetaMessage
	Sender      GroupSender `json:"sender"`
	GroupId     int64       `json:"group_id"`
}

// GetSender
// Get message sender
func (g GroupMessage) GetSender() Sender {
	return g.Sender
}

type MessageLink struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

type MessageInfo struct {
	IsAtMessage bool
	At          int64
	Images      string
	Video       string
	Voice       string
	Faces       []string
	Shares      []MessageLink
	Musics      []MessageLink
}
