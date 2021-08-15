package main

import (
	"github.com/m1dsummer/whitedew"
	"github.com/m1dsummer/whitedew/api"
	"github.com/m1dsummer/whitedew/server"
)

type PluginPing struct{}

func (p PluginPing) Init(w *whitedew.WhiteDew) {
	w.SetActionHandler("/ping", Callback)
}

func Callback(session *server.Session) {
	api.PostPrivateMessage(session.Sender.GetId(), "pong!")
}

func main() {
	w := whitedew.New()
	w.SetCQServer("http://localhost:60001", "access-key")
	w.AddPlugin(PluginPing{})
	w.Run("/event", 60000)
}
