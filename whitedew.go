package whitedew

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type Plugin interface {
	Init(w *WhiteDew)
}

type HandlerFunc func(s *Session)

func (h HandlerFunc) Handle(s *Session) {
	h(s)
}

type Server struct {
	engine *gin.Engine
}

func NewServer() *Server {
	s := Server{}
	s.engine = gin.New()
	s.engine.Use(gin.Recovery())
	return &s
}

type Config struct {
	CQHost   string
	Secret   string
	CacheDir string
}

type WhiteDew struct {
	Config         Config
	server         *Server
	sessionManager SessionManager
	pluginManager  []Plugin
	actionPool     map[string][]HandlerFunc
	eventPool      map[string][]Handler
	rowMsgHandlers []HandlerFunc
}

func New() *WhiteDew {
	w := WhiteDew{}
	w.Config = Config{}
	w.server = NewServer()
	w.eventPool = make(map[string][]Handler)
	w.actionPool = make(map[string][]HandlerFunc)
	return &w
}

func (w *WhiteDew) SetCQServer(url string) {
	w.Config.CQHost = url
}

func (w *WhiteDew) SetAuth(secret string) {
	w.Config.Secret = secret
}

func (w *WhiteDew) SetCacheDir(dir string) {
	w.Config.CacheDir = dir
}

func (w *WhiteDew) AddPlugin(plugin Plugin) {
	w.pluginManager = append(w.pluginManager, plugin)
}

func (w *WhiteDew) messageEventHandler(msgStr []byte, session *Session) {
	action := session.Action
	var handlers []HandlerFunc
	if action != "" {
		handlers = w.actionPool[action]
	} else {
		handlers = w.rowMsgHandlers
	}
	for _, handler := range handlers {
		handler.Handle(session)
	}
}

func parseNotice(msgStr []byte) MetaEvent {
	var tmp MetaEvent
	_ = json.Unmarshal(msgStr, &tmp)
	return tmp
}

func (w *WhiteDew) defaultEventHandler(msgStr []byte, _ *Session) {
	log.Println("meta_message" + string(msgStr))
}

func (w *WhiteDew) metaEventHandler(msgStr []byte, session *Session) {

}

func (w *WhiteDew) universalEventHandler(msgStr []byte) {
	evt := parseNotice(msgStr)
	var noticeType string
	if evt.GetNoticeType() == "notify" {
		noticeType = evt.GetSubType()
	}
	eventType := EventMap[noticeType]
	if eventType == nil {
		return
	}
	rowEvent := reflect.New(eventType).Interface()
	_ = json.Unmarshal(msgStr, &rowEvent)
	handlers := w.eventPool[noticeType]
	if handlers != nil {
		for _, handler := range handlers {
			agent := NewAgent(w.Config.CQHost)
			handler.Handle(agent, rowEvent.(Event))
		}
	}
}

func (w *WhiteDew) SetRowMsgHandler(handlerFunc HandlerFunc) {
	w.rowMsgHandlers = append(w.rowMsgHandlers, handlerFunc)
}

func (w *WhiteDew) SetActionHandler(action string, handler HandlerFunc) {
	w.actionPool[action] = append(w.actionPool[action], handler)
}

func (w *WhiteDew) SetEventHandler(event string, handler Handler) {
	w.eventPool[event] = append(w.eventPool[event], handler)
}

func (w *WhiteDew) dispatchEvent(msgStr []byte) {
	postType := GetEventPostType(msgStr)
	msg := ParseMsg(msgStr)
	var session *Session
	switch postType {
	case "message":
		session = w.sessionManager.NewSession(w.Config.CQHost, msg)
		w.messageEventHandler(msgStr, session)
	case "notice", "request":
		w.universalEventHandler(msgStr)
	case "meta_event":
		w.metaEventHandler(msgStr, session)
	default:
		w.defaultEventHandler(msgStr, session)
	}
}

func (w *WhiteDew) eventHandler(c *gin.Context) {

	jsonData, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "bad requests"})
		return
	}

	// 验证 secret
	if w.Config.Secret != "" {

		signature := c.Request.Header["X-Signature"]
		if signature == nil {
			c.JSON(http.StatusBadRequest, gin.H{"reason": "Signature is required"})
			return
		}

		mac := hmac.New(sha1.New, []byte(w.Config.Secret))
		_, err = mac.Write(jsonData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		sig := hex.EncodeToString(mac.Sum(nil))
		if sig != signature[0][5:] {
			c.JSON(http.StatusForbidden, gin.H{"reason": "authenticate failed"})
			return
		}
	}

	log.Println(jsonData)
	go w.dispatchEvent(jsonData)
}

func (w *WhiteDew) Run(path string, port int) {
	for _, plugin := range w.pluginManager {
		plugin.Init(w)
	}

	w.server.engine.POST(path, w.eventHandler)

	log.Fatalln(w.server.engine.Run(fmt.Sprintf(":%d", port)))
}
