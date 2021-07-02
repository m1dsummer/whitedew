# WhiteDew
自用 QQ 机器人框架，99% 兼容 [OneBot v11](https://github.com/botuniverse/onebot).

* [快速开始](#快速开始)
* [示例](#示例)
* [核心概念](#核心概念)
	* [会话](#会话)
	* [消息](#消息)
	* [消息链](#消息链)
	* [事件](#事件)
* [事件处理](#事件处理)

## 特点
* 插件化应用管理
* 支持对话状态管理
* 简单易上手
* ...

## 快速开始

```go
package main

import (
	"github.com/m1dsummer/whitedew"
)

type PluginPing struct {}

func (p PluginPing)Init(w *whitedew.WhiteDew) {
	w.SetActionHandler("/ping", Callback)
}

func Callback(session *whitedew.Session) {
	session.PostPrivateMessage(session.Sender.GetId(), "pong!")
}

func main() {
	w := whitedew.New()
	w.SetCQServer("http://localhost:60001")
	w.AddPlugin(PluginPing{})
	w.Run("/event", 60000)
}
```

上面的代码创建了一个叫 PluginPing 的插件，当打开私聊窗口向 Bot 发送 `/ping` 时， Bot 会回复 `pong!`

框架本身并不包含 QQ 消息收发功能的实现，需要配合其他 cqhttp 框架使用，比如[go-cqhttp]()

在这个例子中，WhiteDew 监听地址 `http://localhost:60000/event`，cq-http 监听在本地 60001 端口并将消息和事件上报给 WhiteDew， 由 PluginPing 插件对包含 `/ping` 字段的消息进行处理。s

## 示例
* [使用 ActionHandler](./examples/ActionHandler)
* [使用 EventHandler](./examples/ActionHandler)
* [使用 RowMessageHandler](./examples/ActionHandler)

## 核心概念

### 会话

* `whitedew.Session`

Session 是对消息的封装，Session中的字段均与 OneBot v11 协议中保持一致。参考 OveBot v11 中[消息](https://github.com/botuniverse/onebot/blob/master/v11/specs/event/message.md) 事件。

```go
type Session struct {
	Manager    *SessionManager
	Sender     Sender
	StartTime  time.Time
	Message    Message
	Env        map[string]interface{}
	IsFirstRun bool
	Action     string
	ArgParser  ArgParser
	Agent      *Agent
}
```

方法：

* `PostPrivateMessage(receiver int64, msg string, autoEscape ...bool)`
  
  发送私聊消息

* `PostGroupMessage(receiver int64, msg string, autoEscape ...bool)`
  
  发送群聊消息
  
* `Destory()`
  销毁一个 Session。Session 提供了会话管理用于处理连续的对话，当 QQ 用户首次向 Bot 发送消息时，whitedew 会为这个用户创建一个 Session，并将 IsFirstRun 属性设置为 true，此后与这个用户的所有聊天都将使用这个 Session。
  
### 消息
* `whitedew.Message`
一个接口类型，被 PrivteMessage 和 GroupMessage 实现。
```go
type Message interface {
	GetMsgType() string
	GetSender() Sender
	GetContent() string
	GetSelfId() int64
	GetRowMessage() string
}
```
* `whitedew.metaMessage`
私聊消息与群聊消息的共同部分
```go
type metaSender struct {
	UserId   int64  `json:"user_id"`
	Nickname string `json:"nickname"`
	Sex      string `json:"sex"`
	Age      int32  `json:"age"`
}
```

* `whitedew.PrivateMessage`
私聊消息的封装
```go
type PrivateMessage struct {
	metaMessage                  // inherit from MetaMessage
	Sender      PrivateMsgSender `json:"sender"`
	Anonymous   AnonymousUser    `json:"anonymous"`
}
```
* `whitedew.GroupMessage`
群聊消息的封装
```go
type GroupMessage struct {
	metaMessage             // inherit from MetaMessage
	Sender      GroupSender `json:"sender"`
	GroupId     int64       `json:"group_id"`
}
```

### 消息链
使用 `MessageChain` 可以方便构造出复杂的消息类型。参考 OneBot v11 [消息段类型](https://github.com/botuniverse/onebot/blob/master/v11/specs/message/segment.md)。

```go
chain := whitedew.MessageChain{}
str := chain.Prepare().At(012211221).Plain("普通消息").String()
fmt.Println(str)
// [CQ:at,qq=012211221]普通消息
```

方法:
* `Prepare() *MessageChain`
* `Plain(str string) *MessageChain`
* `Face(faceId int) *MessageChain`
* `Image(url string) *MessageChain`
* `Record(url string) *MessageChain`
* `Video(url string) *MessageChain`
* `At(uid int64) *MessageChain`
* `Rps() *MessageChain`
* `Dice() *MessageChain`
* `Shake() *MessageChain`
* `Poke(_type int, id int64) *MessageChain`
* `AtAll() *MessageChain`
* `Anonymous() *MessageChain`
* `Share(title string, url string) *MessageChain`
* `Contact(_type string, uid int64) *MessageChain`
* `Location(lat float64, lon float64) *MessageChain`
* `Music(_type string, id int64) *MessageChain`
* `CustomMusic(url, audio, title string) *MessageChain`
* `Reply(mid int64) *MessageChain`
* `Forward(mid int64) *MessageChain`
* `XML(msg string) *MessageChain`
* `JSON(msg string) *MessageChain`
* `String() strin`


### 事件

参考 OneBot v11 [事件模型](https://github.com/botuniverse/onebot/blob/master/v11/specs/event/README.md)。

* `whitedew.Event`
事件接口类型。
```go
type Event interface {
	GetNoticeType() string
	GetTime() int64
}
```

* `whitedew.MetaEvent`
元事件类型，所有事件类型的共同部分。
```go
type MetaEvent struct {
	Time       int64  `json:"time"`
	SelfId     int64  `json:"self_id"`
	PostType   string `json:"post_type"`
	NoticeType string `json:"notice_type"`
	SubType    string `json:"sub_type"`
}
```

事件列表:
* `GroupUploadEvent`
* `GroupAdminEvent`
* `GroupDecreaseEvent`
* `GroupIncreaseEvent`
* `GroupBanEvent`
* `FriendAddEvent`
* `GroupRecallEvent`
* `FriendRecallEvent`
* `PokeEvent`
* `LuckyKingEvent`
* `HonorEvent`


## 事件处理

whitedew 将事件划分为消息事件与通知事件两类，通知事件可看做 OneBot v11 中[通知事件](https://github.com/botuniverse/onebot/blob/master/v11/specs/event/notice.md)与[请求事件](https://github.com/botuniverse/onebot/blob/master/v11/specs/event/request.md)的集合。

通知事件的事件处理函数定义为 `Handler`
```go
type Handler func(agent *Agent, e Event)
```

消息事件的事件处理函数被定义为 `whitedew.HandlerFunc`
```go
type HandlerFunc func(s *Session)
```
消息可分为带 action 的消息和普通消息，如果一条消息中包含 `/xxx` 格式的内容，则该消息会被分析为带 action 的消息，如 `/echo 123`， `/echo` 将会被保存在 session 变量的 Action 字段里。

使用 whitedew 对象的 SetActionHandler 方法添加 action 消息处理器，使用 SetRowMsgHandler 方法添加 普通消息处理器。

如果一条消息被 action 消息处理器处理，它将**不会**被普通消息处理器处理。