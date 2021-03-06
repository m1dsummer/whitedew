package main

import (
	"github.com/m1dsummer/whitedew/entity"
	"github.com/m1dsummer/whitedew/server"
	"github.com/m1dsummer/whitedew/utils/chain"
	"strings"

	"github.com/m1dsummer/whitedew"
)

type Plugin struct{}

func (p Plugin) Init(w *whitedew.WhiteDew) {
	w.SetRowMsgHandler(Handler)
}

func Handler(session *server.Session) {
	ban := "涩图"
	content := session.Message.GetContent()
	if strings.Contains(content, ban) {
		if session.Message.GetMsgType() == "group" {
			msg := session.Message.(entity.GroupMessage)
			chain := chain.MessageChain{}
			str := chain.Prepare().At(msg.Sender.GetId()).Plain("敏感词警告!").String()
			session.PostGroupMessage(msg.GroupId, str)
		} else {
			session.PostPrivateMessage(session.Message.GetSender().GetId(), "敏感词警告!")
		}
	}
}

func main() {
	w := whitedew.New()
	w.SetCQServer("http://localhost:60001", "access-key")
	w.AddPlugin(Plugin{})
	w.Run("/event", 60000)
}
