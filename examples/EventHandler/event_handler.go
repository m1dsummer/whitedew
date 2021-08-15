package main

import (
	"github.com/m1dsummer/whitedew"
	"github.com/m1dsummer/whitedew/api"
	"github.com/m1dsummer/whitedew/event"
	"github.com/m1dsummer/whitedew/utils/chain"
)

type Plugin struct{}

func (p Plugin) Init(w *whitedew.WhiteDew) {
	w.SetEventHandler("poke", Handler)
}

func Handler(evt event.Event) {
	chain := chain.MessageChain{}
	pokeEvent := evt.(*event.PokeEvent)
	str := chain.Prepare().At(pokeEvent.UserId).Plain("不要戳来戳去的").String()
	api.PostGroupMessage(pokeEvent.GroupId, str)
}

func main() {
	w := whitedew.New()
	w.SetCQServer("http://localhost:60001", "access-key")
	w.AddPlugin(Plugin{})
	w.Run("/event", 60000)
}
