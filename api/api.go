package api

import "github.com/m1dsummer/whitedew/entity"

func PostPrivateMessage(receiver int64, msg string) {
	param := map[string]interface{}{
		"user_id": receiver,
		"message": msg,
	}
	_agent.PostMessage("send_private_msg", param)
}

func PostGroupMessage(receiver int64, msg string) {
	param := map[string]interface{}{
		"group_id": receiver,
		"message":  msg,
	}
	_agent.PostMessage("send_group_msg", param)
}

func DeleteMsg(mid int64) {
	param := map[string]interface{}{
		"message_id": mid,
	}
	_agent.PostMessage("delete_msg", param)
}

func GetMessage(mid int64) {
	param := map[string]interface{}{
		"message_id": mid,
	}
	_agent.PostMessage("get_msg", param)
}

func GetForwardMsg(mid int64) {
	param := map[string]interface{}{
		"message_id": mid,
	}
	_agent.PostMessage("get_forward_msg", param)

}

func SendLike(uid int64, times int64) {
	param := map[string]interface{}{
		"user_id": uid,
		"times":   times,
	}
	_agent.PostMessage("get_forward_msg", param)
}

func SetGroupKick(gid, uid int64, reject bool) {
	param := map[string]interface{}{
		"group_id":           gid,
		"user_id":            uid,
		"reject_add_request": reject,
	}
	_agent.PostMessage("set_group_kick", param)
}

func SetGroupBan(gid, uid, time int64) {
	param := map[string]interface{}{
		"group_id": gid,
		"user_id":  uid,
		"duration": time,
	}
	_agent.PostMessage("set_group_ban", param)
}

func SetGroupAnonymousBan(gid, duration int64, anonymous entity.AnonymousUser, flag string) {
	param := map[string]interface{}{
		"group_id":       gid,
		"anonymous":      anonymous,
		"duration":       duration,
		"anonymous_flag": flag,
	}
	_agent.PostMessage("set_group_anonymous_ban", param)
}

func SetGroupWholeBan(gid int64, enable bool) {
	param := map[string]interface{}{
		"group_id": gid,
		"enable":   enable,
	}
	_agent.PostMessage("set_group_whole_ban", param)
}

func SetGroupAdmin(gid, uid int64, enable bool) {
	param := map[string]interface{}{
		"group_id": gid,
		"user_id":  uid,
		"enable":   enable,
	}
	_agent.PostMessage("set_group_admin", param)
}

func SetGroupAnonymous(gid int64, enable bool) {
	param := map[string]interface{}{
		"group_id": gid,
		"enable":   enable,
	}
	_agent.PostMessage("set_group_anonymous", param)
}

func SetGroupCard(gid, uid int64, card string) {
	param := map[string]interface{}{
		"group_id": gid,
		"user_id":  uid,
		"card":     card,
	}
	_agent.PostMessage("set_group_card", param)
}

func SetGroupName(gid int64, name string) {
	param := map[string]interface{}{
		"group_id":   gid,
		"group_name": name,
	}
	_agent.PostMessage("set_group_name", param)
}

func SetGroupLeave(gid int64, dismiss bool) {
	param := map[string]interface{}{
		"group_id":   gid,
		"is_dismiss": dismiss,
	}
	_agent.PostMessage("set_group_leave", param)
}

func SetGroupSpecialTitle(gid, uid int64, title string, duration int64) {
	param := map[string]interface{}{
		"group_id":      gid,
		"user_id":       uid,
		"special_title": title,
		"duration":      duration,
	}
	_agent.PostMessage("set_group_special_title", param)
}

func SetFriendAddRequest(flag string, approve bool, remark string) {
	param := map[string]interface{}{
		"flag":    flag,
		"approve": approve,
		"remark":  remark,
	}
	_agent.PostMessage("set_friend_add_request", param)
}

func SetGroupAddRequest(flag string, _type string, approve bool, reason string) {
	param := map[string]interface{}{
		"flag":    flag,
		"type":    _type,
		"approve": approve,
		"reason":  reason,
	}
	_agent.PostMessage("set_group_add_request", param)
}

func GetLoginInfo() {
	_agent.PostMessage("get_login_info", nil)
}

func GetStrangerInfo(uid int64, noCache bool) {
	param := map[string]interface{}{
		"user_id":  uid,
		"no_cache": noCache,
	}
	_agent.PostMessage("get_stranger_info", param)
}

func GetFriendList() {
	_agent.PostMessage("get_friend_list", nil)
}

func GetGroupInfo(gid int64, noCache bool) {
	param := map[string]interface{}{
		"group_id": gid,
		"no_cache": noCache,
	}
	_agent.PostMessage("get_group_info", param)
}

func GetGroupList() {
	_agent.PostMessage("get_group_list", nil)
}

func GetGroupMemberInfo(gid int64, uid int64, noCache bool) {
	param := map[string]interface{}{
		"group_id": gid,
		"user_id":  uid,
		"no_cache": noCache,
	}
	_agent.PostMessage("get_group_member_info", param)
}

func GetGroupMemberList(gid int64) {
	param := map[string]interface{}{
		"group_id": gid,
	}
	_agent.PostMessage("get_group_member_list", param)
}

func GetGroupHonorInfo(gid int64, _type string) {
	param := map[string]interface{}{
		"group_id": gid,
		"type":     _type,
	}
	_agent.PostMessage("get_group_honor_info", param)
}

func GetCookies(domain string) {
	param := map[string]interface{}{
		"domain": domain,
	}
	_agent.PostMessage("get_cookies", param)
}

func GetCsrfToken() {
	_agent.PostMessage("get_csrf_token", nil)
}

func GetCredentials(domain string) {
	param := map[string]interface{}{
		"domain": domain,
	}
	_agent.PostMessage("get_credentials", param)
}

func GetRecord(file string, outFormat string) {
	param := map[string]interface{}{
		"file":       file,
		"out_format": outFormat,
	}
	_agent.PostMessage("get_record", param)
}

func GetImage(file string) {
	param := map[string]interface{}{
		"file": file,
	}
	_agent.PostMessage("get_image", param)
}

func CanSendImage() {
	_agent.PostMessage("can_send_image", nil)
}

func CanSendRecord() {
	_agent.PostMessage("can_send_record", nil)
}

func GetStatus() {
	_agent.PostMessage("get_status", nil)
}

func GetVersionInfo() {
	_agent.PostMessage("get_version_info", nil)
}

func SetRestart() {
	param := map[string]interface{}{
		"delay": 0,
	}
	_agent.PostMessage("set_restart", param)
}

func CleanCache() {
	_agent.PostMessage("clean_cache", nil)
}
