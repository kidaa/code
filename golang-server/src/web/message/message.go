package message
import (
	"web/proxy"
	"web/webvo"
	"db"
	"redis"
)
func init() {
	proxy.Regist(13000, getMessageTotal)
}

// 获取未读系统消息总数
func getMessageTotal(uid int, body *[]byte) *[]byte {
	ctos := webvo.WebCtoS13000Data1{}
	ctos.Decode(body)
	stoc := webvo.WebStoC13000Data1{T:ctos.T}
	if uid<= 0{
		stoc.E = 18001
	}else {
		stoc.Data = &webvo.WebStoC13000Data2{}

		shop := &webvo.WebStoC13000Data3{}//todo:等待商城完善
		stoc.Data.Shop = shop
		stoc.Data.Total += shop.New + shop.Sales

		prop := &webvo.WebStoC13000Data4{}//todo:等待道具完善
		stoc.Data.Prop = prop
		stoc.Data.Total += prop.Obtain

		friend := &webvo.WebStoC13000Data5{}
		friend.Chat,_= redis.ChatMsg.GetOfflineChatMsgCount(uid)
		friend.Request = db.SlaveDB.NoDealRequestCount(uid)
		stoc.Data.Friend = friend
		stoc.Data.Total += friend.Chat + friend.Request

		sys := &webvo.WebStoC13000Data6{}
		sys.Attack, _ = db.SlaveDB.OffLineAttackNum(uid)
//		sys.Dialog, _ = db.SlaveDB.OffLineAttackNum(uid)//todo:等待对白完善
		stoc.Data.Sys = sys
		stoc.Data.Total += sys.Attack + sys.Dialog
	}
	b, _ := stoc.Encode()
	return b
}