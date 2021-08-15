package server

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/m1dsummer/whitedew/event"
	"log"
	"net/http"
	"reflect"
)

type SessionHandler func(s *Session)

func (h SessionHandler) Handle(s *Session) {
	h(s)
}

type EventHandler func(evt event.Event)

func (e EventHandler) Handle(evt event.Event) {
	e(evt)
}

type Server struct {
	Engine *gin.Engine
	ActionPool     map[string][]SessionHandler
	EventPool      map[string][]EventHandler
	RowMsgPool []SessionHandler
	Secret string
	SessionManager SessionManager
}

func NewServer() *Server {
	s := Server{}
	s.Engine = gin.New()
	s.Engine.Use(gin.Recovery())
	s.ActionPool = make(map[string][]SessionHandler)
	s.EventPool  = make(map[string][]EventHandler)
	s.RowMsgPool = make([]SessionHandler, 32)
	s.SessionManager = SessionManager{}
	return &s
}

func (s *Server) messageEventHandler(msgStr []byte, session *Session) {
	action := session.Action
	var handlers []SessionHandler
	if action != "" {
		handlers = s.ActionPool[action]
	} else {
		handlers = s.RowMsgPool
	}
	for _, handler := range handlers {
		handler.Handle(session)
	}
}

func (w *Server)Run(path string, port int) {
	w.Engine.POST(path, w.HandleRequests)
	log.Fatalln(w.Engine.Run(fmt.Sprintf(":%d", port)))
}

func parseNotice(msgStr []byte) event.MetaEvent {
	var tmp event.MetaEvent
	_ = json.Unmarshal(msgStr, &tmp)
	return tmp
}

func (w *Server) defaultEventHandler(msgStr []byte, _ *Session) {
	log.Println("meta_message" + string(msgStr))
}

func (w *Server) metaEventHandler(msgStr []byte, session *Session) {

}

func (w *Server) universalEventHandler(msgStr []byte) {
	evt := parseNotice(msgStr)
	var noticeType string
	if evt.GetNoticeType() == "notify" {
		noticeType = evt.GetSubType()
	}
	eventType := event.EventMap[noticeType]
	if eventType == nil {
		return
	}
	rowEvent := reflect.New(eventType).Interface()
	_ = json.Unmarshal(msgStr, &rowEvent)
	handlers := w.EventPool[noticeType]
	if handlers != nil {
		for _, handler := range handlers {
			handler.Handle(rowEvent.(event.Event))
		}
	}
}

func (w *Server) dispatchEvent(msgStr []byte) {
	postType := GetEventPostType(msgStr)
	msg := ParseMsg(msgStr)
	var session *Session
	switch postType {
	case "message":
		session = w.SessionManager.NewSession(msg)
		w.messageEventHandler(msgStr, session)
	case "notice", "request":
		w.universalEventHandler(msgStr)
	case "meta_event":
		w.metaEventHandler(msgStr, session)
	default:
		w.defaultEventHandler(msgStr, session)
	}
}

func (w *Server) HandleRequests(c *gin.Context) {

	jsonData, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"reason": "bad requests"})
		return
	}

	// 验证 secret
	if w.Secret != "" {

		signature := c.Request.Header["X-Signature"]
		if signature == nil {
			c.JSON(http.StatusBadRequest, gin.H{"reason": "Signature is required"})
			return
		}

		mac := hmac.New(sha1.New, []byte(w.Secret))
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

	go w.dispatchEvent(jsonData)
}