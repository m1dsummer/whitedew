# WhiteDew
自用 QQ 机器人框架，99% 兼容 [OneBot v11](https://github.com/botuniverse/onebot).

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
	"log"
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

在这个例子中，WhiteDew 监听地址 `http://localhost:60000`，cq-http 监听在本地 60001 端口并将消息和事件上报给 WhiteDew， 由 PluginPing 插件对包含 `/ping` 字段的消息进行处理。
