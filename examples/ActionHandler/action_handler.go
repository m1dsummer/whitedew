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