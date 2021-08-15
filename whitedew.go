package whitedew

import (
	"github.com/m1dsummer/whitedew/api"
	"github.com/m1dsummer/whitedew/server"
)

type Plugin interface {
	Init(w *WhiteDew)
}

type WhiteDew struct {
	server         *server.Server
	sessionManager server.SessionManager
	pluginManager  []Plugin
}

func New() *WhiteDew {
	w := WhiteDew{}
	w.server = server.NewServer()
	return &w
}

func (w *WhiteDew) SetCQServer(url string, accessToken string) {
	api.GenAgent(url, accessToken)
}

func (w *WhiteDew) SetAuth(secret string) {
	w.server.Secret = secret
}

func (w *WhiteDew) AddPlugin(plugin Plugin) {
	w.pluginManager = append(w.pluginManager, plugin)
}

func (w *WhiteDew) SetRowMsgHandler(handlerFunc server.SessionHandler) {
	w.server.RowMsgPool = append(w.server.RowMsgPool, handlerFunc)
}

func (w *WhiteDew) SetActionHandler(action string, handler server.SessionHandler) {
	w.server.ActionPool[action] = append(w.server.ActionPool[action], handler)
}

func (w *WhiteDew) SetEventHandler(event string, handler server.EventHandler) {
	w.server.EventPool[event] = append(w.server.EventPool[event], handler)
}

func (w *WhiteDew) Run(path string, port int) {
	for _, plugin := range w.pluginManager {
		plugin.Init(w)
	}
	w.server.Run(path, port)
}
