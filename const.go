package whitedew

import (
	"reflect"
)

var EventMap map[string]reflect.Type

func init() {
	EventMap = map[string]reflect.Type{
		"group_upload":   reflect.TypeOf(GroupUploadEvent{}),
		"group_admin":    reflect.TypeOf(GroupAdminEvent{}),
		"group_decrease": reflect.TypeOf(GroupDecreaseEvent{}),
		"group_increase": reflect.TypeOf(GroupIncreaseEvent{}),
		"group_ban":      reflect.TypeOf(GroupBanEvent{}),
		"friend_add":     reflect.TypeOf(FriendAddEvent{}),
		"group_recall":   reflect.TypeOf(GroupRecallEvent{}),
		"friend_recall":  reflect.TypeOf(FriendRecallEvent{}),
		"poke":           reflect.TypeOf(PokeEvent{}),
		"lucky_king":     reflect.TypeOf(LuckyKingEvent{}),
		"honor":          reflect.TypeOf(HonorEvent{}),
	}
}
