package api

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"log"
)

type Agent struct {
	URL         string
	AccessToken string
}

var (
	_agent Agent
)

func GenAgent(url string, accessToken string) {
	_agent = Agent{URL: url, AccessToken: accessToken}
}

type MessageTemplate struct {
	Action string                 `json:"action"`
	Params map[string]interface{} `json:"params"`
	Echo   string                 `json:"echo"`
}

func (a *Agent) NewPostMessage(action string, params map[string]interface{}) *MessageTemplate {
	return &MessageTemplate{Action: action, Params: params, Echo: "success"}
}

func (a *Agent) PostMessage(action string, param map[string]interface{}, autoEscape ...bool) []byte {
	length := len(autoEscape)
	switch length {
	case 1:
		param["auto_escape"] = autoEscape[0]
	case 0:
		param["auto_escape"] = false
	default:
		panic("too many arguments")
	}
	uri := fmt.Sprintf("%s/%s", a.URL, action)
	data, _ := json.Marshal(param)
	_, body, errs := gorequest.
		New().
		Post(uri).
		Set("Content-Type", "application/json").
		Set("Authorization", a.AccessToken).
		Send(string(data)).
		EndBytes()
	if errs != nil {
		log.Fatalln(errs)
	}
	return body
}