package whitedew

import (
	"github.com/m1dsummer/whitedew/api"
	"github.com/m1dsummer/whitedew/event"
	"github.com/m1dsummer/whitedew/server"
	"github.com/m1dsummer/whitedew/utils/chain"
	"testing"
)

type PluginPing struct{}

func (p PluginPing) Init(w *WhiteDew) {
	w.SetActionHandler("/ping", Callback)
	w.SetEventHandler("poke", PokeHandler)
}

func Callback(session *server.Session) {
	api.PostPrivateMessage(session.Sender.GetId(), "pong!")
}

func PokeHandler(evt event.Event) {
	pokeEvent := evt.(*event.PokeEvent)
	msgChain := chain.MessageChain{}
	api.PostPrivateMessage(pokeEvent.UserId, msgChain.Prepare().At(pokeEvent.UserId).Plain("戳一戳").String())
}

func TestRun(t *testing.T) {
	w := New()
	w.SetCQServer("http://127.0.0.1:60001", "tQwUiHbnJEEZ7aHb9F8B2BvujWMciyyu")
	w.AddPlugin(PluginPing{})
	w.Run("/event", 60000)
}
