package whitedew


import (
	"encoding/json"
	"regexp"
	"strconv"
)

// GetEventPostType
// Get the post_type field from event message
// Reference https://github.com/botuniverse/onebot/blob/master/v11/specs/event/README.md
func GetEventPostType(msg []byte) string {
	var t struct {
		PostType string `json:"post_type"`
	}
	_ = json.Unmarshal([]byte(msg), &t)
	return t.PostType
}

// ParseMsg
// Convert row message to corresponding message type (GroupMessage and PrivateMessage)
func ParseMsg(rowMsg []byte) Message {
	var m struct {
		MsgType string `json:"message_type"`
	}
	_ = json.Unmarshal(rowMsg, &m)
	msgType := m.MsgType
	if msgType == "group" {
		var t GroupMessage
		_ = json.Unmarshal(rowMsg, &t)
		return t
	} else if msgType == "private" {
		var t PrivateMessage
		_ = json.Unmarshal(rowMsg, &t)
		return t
	} else {
		return nil
	}
}

func fetchNumber(pattern string, msgStr string, size int) int64 {
	r := regexp.MustCompile(pattern)
	matches := r.FindSubmatch([]byte(msgStr))
	if matches == nil {
		return 0
	}
	id, err := strconv.ParseInt(string(matches[1]), 10, size)
	if err != nil {
		return 0
	}
	return id
}

func fetchStr(pattern string, msgStr string) []string {
	r := regexp.MustCompile(pattern)
	matches := r.FindAllSubmatch([]byte(msgStr),-1)
	if matches != nil {
		var tmp []string
		for _,m := range matches {
			tmp = append(tmp, string(m[0]))
		}
		return tmp
	}
	return []string{}
}

func AnalyzeMsg(msg Message) MessageInfo {
	var tmp MessageInfo
	msgStr := msg.GetRowMessage()

	id := fetchNumber("\\[CQ:at,qq=(\\d+)\\]", msgStr, 64)
	tmp.At = id
	tmp.IsAtMessage = id != 0

	tmp.Faces = fetchStr("\\[CQ:image,file=.+?\\]", msgStr)

	return tmp
}

