package main

import "github.com/m1dsummer/whitedew"

type Plugin struct{}

func (p Plugin) Init(w *whitedew.WhiteDew) {
	w.SetEventHandler("poke", Handler)
}

func Handler(agent *whitedew.Agent, event whitedew.Event) {
	chain := whitedew.MessageChain{}
	pokeEvent := event.(*whitedew.PokeEvent)
	str := chain.Prepare().At(pokeEvent.UserId).Plain("不要戳来戳去的").String()
	agent.PostGroupMessage(pokeEvent.GroupId, str)
}

func main() {
	w := whitedew.New()
	w.SetCQServer("http://localhost:60001", "access-key")
	w.AddPlugin(Plugin{})
	w.Run("/event", 60000)
}
