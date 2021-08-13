package whitedew

import (
	"testing"
)

type PluginPing struct{}

func (p PluginPing) Init(w *WhiteDew) {
	w.SetActionHandler("/ping", Callback)
}

func Callback(session *Session) {
	session.PostPrivateMessage(session.Sender.GetId(), "pong!")
}

func TestRun(t *testing.T) {
	w := New()
	w.SetCQServer("http://127.0.0.1:60001", "tQwUiHbnJEEZ7aHb9F8B2BvujWMciyyu")
	w.AddPlugin(PluginPing{})
	w.Run("/event", 60000)
}
