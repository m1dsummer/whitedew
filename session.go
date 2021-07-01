package whitedew

import (
	"regexp"
	"time"
)

type ArgParser func(s *Session)

func (a ArgParser) Parse(s *Session) {
	a(s)
}

type Session struct {
	Manager    *SessionManager
	Sender     Sender
	StartTime  time.Time
	Message    Message
	Words      []string
	Env        map[string]interface{}
	IsFirstRun bool
	Action     string
	ArgParser  ArgParser
	Agent      *Agent
}

func (s Session) Destroy() {
	s.Manager.Destroy(s.Sender.GetId())
}

func (s *Session) PostPrivateMessage(receiver int64, msg string, autoEscape ...bool) {
	s.Agent.PostPrivateMessage(s.Sender.GetId(), msg)
}

func (s *Session) PostGroupMessage(receiver int64, msg string, autoEscape ...bool) {
	s.Agent.PostGroupMessage(receiver, msg)
}

func (s Session) Get(arg string, prompt string) interface{} {
	if prompt != "" {
		s.Agent.PostGroupMessage(s.Sender.GetId(), prompt)
	}
	return s.Env[arg]
}

func (s *Session) RegisterArgParser(parser ArgParser) {
	s.ArgParser = parser
}

type SessionManager struct {
	Pool map[int64]*Session
}

func (s SessionManager) Get(uid int64) *Session {
	return s.Pool[uid]
}

func (s SessionManager) Destroy(uid int64) {
	delete(s.Pool, uid)
}

func (s SessionManager) NewSession(url string, msg Message) *Session {
	oldSession := s.Get(msg.GetSender().GetId())
	if oldSession != nil {
		oldSession.IsFirstRun = false
		return oldSession
	}

	session := Session{}
	session.StartTime = time.Now()
	session.Message = msg
	session.IsFirstRun = true
	session.Env = make(map[string]interface{})
	session.Agent = NewAgent(url)
	session.Sender = msg.GetSender()
	session.Action = ParseAction(msg.GetContent())


	return &session
}

func ParseAction(msgStr string) string {
	r := regexp.MustCompile("(/\\w+?)(\\W|$)")
	matches := r.FindSubmatch([]byte(msgStr))
	if matches == nil {
		return ""
	}
	return string(matches[1])
}
