package whitedew

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"log"
)

type Agent struct {
	URL string
}

func NewAgent(URL string) *Agent {
	return &Agent{URL: URL}
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
	_, body, errs := gorequest.New().Post(uri).Set("Content-Type", "application/json").Send(string(data)).EndBytes()
	if errs != nil {
		log.Fatalln(errs)
	}
	return body
}

func (a *Agent) PostPrivateMessage(receiver int64, msg string) {
	param := map[string]interface{}{
		"user_id": receiver,
		"message": msg,
	}
	a.PostMessage("send_private_msg", param)
}

func (a *Agent) PostGroupMessage(receiver int64, msg string) {
	param := map[string]interface{}{
		"group_id": receiver,
		"message":  msg,
	}
	a.PostMessage("send_group_msg", param)
}

func (a *Agent) DeleteMsg(mid int64) {
	param := map[string]interface{}{
		"message_id": mid,
	}
	a.PostMessage("delete_msg", param)
}

func (a *Agent) GetMessage(mid int64) {
	param := map[string]interface{}{
		"message_id": mid,
	}
	a.PostMessage("get_msg", param)
}

func (a *Agent) GetForwardMsg(mid int64) {
	param := map[string]interface{}{
		"message_id": mid,
	}
	a.PostMessage("get_forward_msg", param)

}

func (a *Agent) SendLike(uid int64, times int64) {
	param := map[string]interface{}{
		"user_id": uid,
		"times":   times,
	}
	a.PostMessage("get_forward_msg", param)
}

func (a *Agent) SetGroupKick(gid, uid int64, reject bool) {
	param := map[string]interface{}{
		"group_id":           gid,
		"user_id":            uid,
		"reject_add_request": reject,
	}
	a.PostMessage("set_group_kick", param)
}

func (a *Agent) SetGroupBan(gid, uid, time int64) {
	param := map[string]interface{}{
		"group_id": gid,
		"user_id":  uid,
		"duration": time,
	}
	a.PostMessage("set_group_ban", param)
}

func (a *Agent) SetGroupAnonymousBan(gid, duration int64, anonymous AnonymousUser, flag string) {
	param := map[string]interface{}{
		"group_id":       gid,
		"anonymous":      anonymous,
		"duration":       duration,
		"anonymous_flag": flag,
	}
	a.PostMessage("set_group_anonymous_ban", param)
}

func (a *Agent) SetGroupWholeBan(gid int64, enable bool) {
	param := map[string]interface{}{
		"group_id": gid,
		"enable":   enable,
	}
	a.PostMessage("set_group_whole_ban", param)
}

func (a *Agent) SetGroupAdmin(gid, uid int64, enable bool) {
	param := map[string]interface{}{
		"group_id": gid,
		"user_id":  uid,
		"enable":   enable,
	}
	a.PostMessage("set_group_admin", param)
}

func (a *Agent) SetGroupAnonymous(gid int64, enable bool) {
	param := map[string]interface{}{
		"group_id": gid,
		"enable":   enable,
	}
	a.PostMessage("set_group_anonymous", param)
}

func (a *Agent) SetGroupCard(gid, uid int64, card string) {
	param := map[string]interface{}{
		"group_id": gid,
		"user_id":  uid,
		"card":     card,
	}
	a.PostMessage("set_group_card", param)
}

func (a *Agent) SetGroupName(gid int64, name string) {
	param := map[string]interface{}{
		"group_id":   gid,
		"group_name": name,
	}
	a.PostMessage("set_group_name", param)
}

func (a *Agent) SetGroupLeave(gid int64, dismiss bool) {
	param := map[string]interface{}{
		"group_id":   gid,
		"is_dismiss": dismiss,
	}
	a.PostMessage("set_group_leave", param)
}

func (a *Agent) SetGroupSpecialTitle(gid, uid int64, title string, duration int64) {
	param := map[string]interface{}{
		"group_id":      gid,
		"user_id":       uid,
		"special_title": title,
		"duration":      duration,
	}
	a.PostMessage("set_group_special_title", param)
}

func (a *Agent) SetFriendAddRequest(flag string, approve bool, remark string) {
	param := map[string]interface{}{
		"flag":    flag,
		"approve": approve,
		"remark":  remark,
	}
	a.PostMessage("set_friend_add_request", param)
}

func (a *Agent) SetGroupAddRequest(flag string, _type string, approve bool, reason string) {
	param := map[string]interface{}{
		"flag":    flag,
		"type":    _type,
		"approve": approve,
		"reason":  reason,
	}
	a.PostMessage("set_group_add_request", param)
}

func (a *Agent) GetLoginInfo() {
	a.PostMessage("get_login_info", nil)
}

func (a *Agent) GetStrangerInfo(uid int64, noCache bool) {
	param := map[string]interface{}{
		"user_id":  uid,
		"no_cache": noCache,
	}
	a.PostMessage("get_stranger_info", param)
}

func (a *Agent) GetFriendList() {
	a.PostMessage("get_friend_list", nil)
}

func (a *Agent) GetGroupInfo(gid int64, noCache bool) {
	param := map[string]interface{}{
		"group_id": gid,
		"no_cache": noCache,
	}
	a.PostMessage("get_group_info", param)
}

func (a *Agent) GetGroupList() {
	a.PostMessage("get_group_list", nil)
}

func (a *Agent) GetGroupMemberInfo(gid int64, uid int64, noCache bool) {
	param := map[string]interface{}{
		"group_id": gid,
		"user_id":  uid,
		"no_cache": noCache,
	}
	a.PostMessage("get_group_member_info", param)
}

func (a *Agent) GetGroupMemberList(gid int64) {
	param := map[string]interface{}{
		"group_id": gid,
	}
	a.PostMessage("get_group_member_list", param)
}

func (a *Agent) GetGroupHonorInfo(gid int64, _type string) {
	param := map[string]interface{}{
		"group_id": gid,
		"type":     _type,
	}
	a.PostMessage("get_group_honor_info", param)
}

func (a *Agent) GetCookies(domain string) {
	param := map[string]interface{}{
		"domain": domain,
	}
	a.PostMessage("get_cookies", param)
}

func (a *Agent) GetCsrfToken() {
	a.PostMessage("get_csrf_token", nil)
}

func (a *Agent) GetCredentials(domain string) {
	param := map[string]interface{}{
		"domain": domain,
	}
	a.PostMessage("get_credentials", param)
}

func (a *Agent) GetRecord(file string, outFormat string) {
	param := map[string]interface{}{
		"file":       file,
		"out_format": outFormat,
	}
	a.PostMessage("get_record", param)
}

func (a *Agent) GetImage(file string) {
	param := map[string]interface{}{
		"file": file,
	}
	a.PostMessage("get_image", param)
}

func (a *Agent) CanSendImage() {
	a.PostMessage("can_send_image", nil)
}

func (a *Agent) CanSendRecord() {
	a.PostMessage("can_send_record", nil)
}

func (a *Agent) GetStatus() {
	a.PostMessage("get_status", nil)
}

func (a *Agent) GetVersionInfo() {
	a.PostMessage("get_version_info", nil)
}

func (a *Agent) SetRestart() {
	param := map[string]interface{}{
		"delay": 0,
	}
	a.PostMessage("set_restart", param)
}

func (a *Agent) CleanCache() {
	a.PostMessage("clean_cache", nil)
}
