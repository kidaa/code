/**
* Created by Michael on 2015/8/7.
*	缘分纸条系统
*
*
*
 */
package karma

import (
	"constant"
	"db"
	log "github.com/golang/glog"
	"proxy"
	"socket"
	"time"
	"utils"
	"vo"
	"widget"
)

func init() {

	proxy.Regist(6105, userKarmaWidgetHdr)
	proxy.Regist(6106, getKarmaNotes)
	proxy.Regist(6107, replyKarmaNotes)
	proxy.Regist(6109, sendkarmaMSG)
	proxy.Regist(6110, deleRelation)
	proxy.Regist(6113, getkarmaRelationList)

	proxy.Regist(6118, noReadMsgCountHdr) // 所有未读的聊天消息数量
	proxy.Regist(6119, usersMsgCountList) // 所有缘分好友，未读的消息数量，好友ID跟未读数量对应
	proxy.Regist(6120, userMsgList)       // 指定缘分好友，未读的消息内容
}

func userMsgList(msg *[]byte, c *socket.Connection) {
	ctos := vo.CtoS6120Data1{}
	err := ctos.Decode(msg)
	if err != nil {
		log.Errorln(err)
	}
	stoc := vo.StoC6120Data1{T: ctos.T}
	messagelist := make([]*vo.StoC6120Data2, 10)
	messagelist = messagelist[:0]

	msgList, e := db.SlaveDB.KarmaMsgByUserList(c.UserData.Userid, ctos.Userid)
	if e == nil {
		for i := 0; i < len(msgList); i++ {
			m := msgList[i]
			msgData := &vo.StoC6120Data2{}
			msgData.Userid = m.Userid
			msgData.Msg = m.Msg
			msgData.Sendtime = int(m.Sendtime)
			messagelist = append(messagelist, msgData)
		}
	} else {
		stoc.E = 11004
	}

	stoc.Data = messagelist
	b, _ := stoc.Encode()
	c.Send <- *b
}

func usersMsgCountList(msg *[]byte, c *socket.Connection) {
	stoc := vo.StoC6119Data1{T: 6119}
	list, err := db.SlaveDB.UserKarmaMsgCountList(c.UserData.Userid)
	if err == nil {
		hash := make(map[int]*vo.StoC6119Data2)

		for i := 0; i < len(list); i++ {
			data, ok := hash[list[i]]
			if ok {
				data.Msgnum++
			} else {
				data := &vo.StoC6119Data2{}
				hash[list[i]] = data
				data.Msgnum = 1
				data.Userid = list[i]
				sendTime, msg, e := db.SlaveDB.GetLastKarmaMsgByOfUser(data.Userid, c.UserData.Userid)
				if e != nil {
					log.Infoln(e)
				}
				data.Sendtime = sendTime
				data.Msg = msg
			}
		}

		arr := make([]*vo.StoC6119Data2, 10)
		arr = arr[:0]
		for _, v := range hash {
			arr = append(arr, v)
		}
		stoc.Data = arr
	}

	bit, _ := stoc.Encode()

	c.Send <- *bit
}

func noReadMsgCountHdr(msg *[]byte, c *socket.Connection) {
	b := vo.StoC6118Data1{T: 6118}
	b.Data = db.SlaveDB.NoReadKarmaMsgCount(c.UserData.Userid)
	bit, _ := b.Encode()
	c.Send <- *bit
}

// 使用缘分纸条道具
func userKarmaWidgetHdr(msg *[]byte, c *socket.Connection) {
	ctos := vo.CtoS6105Data1{}
	ctos.Decode(msg)
	stoc := vo.StoC6105Data1{T: ctos.T}
	if constant.Karma == ctos.Id {
		var strLen = 0
		ctos.Msg, strLen = utils.GetTrimStrLen(ctos.Msg)
		if strLen <= 0 {
			stoc.E = 11011
		} else if strLen > 100 {
			stoc.E = 11012
		} else {
			// 从用户背包里删除一个纸条道具
			count, err := db.MasterDB.DelUserWidget(c.UserData.Userid, ctos.Id, 1)
			if err == nil {
				widget.UserWidgetChange(c.UserData.Userid, ctos.Id, count)
			} else {
				stoc.E = 15002 // 用户背包里删没有纸条道具
			}
			userData := c.UserData
			db.MasterDB.AddKarma(userData.Userid, ctos.Id, ctos.Msg, userData.Sex, userData.Birthday, userData.Provinceid, userData.Cityid)
		}
	} else {
		stoc.E = 15003 // 用户提交的id不是纸条道具
	}

	b, err := stoc.Encode()
	if err != nil {
		log.Errorln(err)
	}
	c.Send <- *b
}

// 获取缘分纸条
func getKarmaNotes(msg *[]byte, c *socket.Connection) {
	ctos := vo.CtoS6106Data1{}
	ctos.Decode(msg)
	userData := c.UserData
	stoc := vo.StoC6106Data1{T: 6106}
	karmaData, err := db.SlaveDB.GetKarma(userData.Userid, userData.Sex, userData.Birthday, userData.Provinceid, userData.Cityid)
	if err == nil {
		stoc.Data = &vo.StoC6106Data2{}
		stoc.Data.Userid = karmaData.SendID
		stoc.Data.Id = karmaData.ID
		stoc.Data.Msg = karmaData.Content
		stoc.Data.Birth = int(karmaData.Birthday)
		stoc.Data.Provinceid = karmaData.Provinceid
		stoc.Data.Cityid = karmaData.Cityid
	} else {
		stoc.E = 31007 //系统没有缘分纸条
	}
	b, err := stoc.Encode()
	if err != nil {
		log.Errorln(err)
	}
	c.Send <- *b
}

func sendmsgfor6122(targetConn int, c *socket.Connection, karmaData *vo.KarmaData, currentContent string) {
	send := vo.StoC6114Data1{T: 6114, Data: &vo.StoC6114Data2{}}
	userdata := c.UserData
	if targetConn == c.UserData.Userid {
		send.Data.Userid = karmaData.SendID
		send.Data.Birth = int(karmaData.Birthday)
		send.Data.Sex = karmaData.Sex
		send.Data.Provinceid = karmaData.Provinceid
		send.Data.Cityid = karmaData.Cityid

	} else {
		send.Data.Userid = userdata.Userid
		send.Data.Birth = int(userdata.Birthday)
		send.Data.Sex = userdata.Sex
		send.Data.Provinceid = userdata.Provinceid
		send.Data.Cityid = userdata.Cityid
	}

	b, _ := send.Encode()
	//	targetConn.Send <- *b
	m := vo.Broadcast{}
	m.Channel = targetConn
	m.Msg = *b
	socket.Hub.Broadcast <- m

	stoC6122 := vo.StoC6122Data1{T: 6122}

	data := &vo.StoC6122Data2{}
	data.Msg = karmaData.Content
	data.Send = karmaData.SendID
	data.Receive = userdata.Userid
	data.Sendtime = int(karmaData.CreateTime)
	stoC6122.Data = append(stoC6122.Data, data)

	data = &vo.StoC6122Data2{}
	data.Msg = currentContent
	data.Send = userdata.Userid
	data.Receive = karmaData.SendID
	data.Sendtime = int(time.Now().Unix())
	stoC6122.Data = append(stoC6122.Data, data)

	b, _ = stoC6122.Encode()
	//	targetConn.Send <- *b

	m = vo.Broadcast{}
	m.Channel = targetConn
	m.Msg = *b
	socket.Hub.Broadcast <- m
}

// 回应缘分纸条
func replyKarmaNotes(msg *[]byte, c *socket.Connection) {
	stoc := vo.StoC6107Data1{T: 6107}
	defer func() {
		if err := recover(); err != nil {
			stoc.E = 11013
			bit, _ := stoc.Encode()
			c.Send <- *bit
			log.Infoln(err)
		}
	}()
	ctos := vo.CtoS6107Data1{}
	ctos.Decode(msg)

	var strLen = 0
	ctos.Msg, strLen = utils.GetTrimStrLen(ctos.Msg)
	if strLen <= 0 {
		stoc.E = 11011
	} else if strLen > 100 {
		stoc.E = 11012
	} else if db.SlaveDB.ExistKarma(ctos.Id) == nil {
		karmaData, _ := db.SlaveDB.QuestKarmaByid(ctos.Id)
		// 纸条过去删除
		if karmaData.Status == 1 {
			db.MasterDB.DelKarma(karmaData.ID)
		}
		//		targetConn, ok := socket.Hub.Connections[ctos.Userid]

		if socket.Hub.Exists(ctos.Userid) {
			sendmsgfor6122(ctos.Userid, c, karmaData, ctos.Msg)
		} else {
			db.MasterDB.AddKarmaMsg(c.UserData.Userid, karmaData.SendID, ctos.Msg)
		}

		db.MasterDB.AddKarmaRelation(c.UserData.Userid, ctos.Userid)
		sendmsgfor6122(c.UserData.Userid, c, karmaData, ctos.Msg)
	} else {
		stoc.E = 11013
	}

	b, err := stoc.Encode()
	if err != nil {
		log.Errorln(err)
	}
	c.Send <- *b
}

// 发送缘分消息
func sendkarmaMSG(msg *[]byte, c *socket.Connection) {
	stoc := vo.StoC6109Data1{T: 6109}
	defer func() {
		if e := recover(); e != nil {
			stoc.E = 20002
			log.Errorln(e)
		}
		b, _ := stoc.Encode()
		c.Send <- *b
	}()

	ctos := vo.CtoS6109Data1{}
	ctos.Decode(msg)
	boo, err := db.SlaveDB.IsKarmaRelation(ctos.Userid, c.UserData.Userid)
	if boo && err == nil {
		var strLen = 0
		ctos.Msg, strLen = utils.GetTrimStrLen(ctos.Msg)
		if strLen <= 0 {
			stoc.E = 11011
		} else if strLen > 100 {
			stoc.E = 11012
		} else {
//			targetConn, ok := socket.Hub.Connections[ctos.Userid]
			if socket.Hub.Exists(ctos.Userid) {
				send := vo.StoC6108Data1{T: 6108}
				send.Data = &vo.StoC6108Data2{}
				send.Data.Msg = ctos.Msg
				send.Data.Userid = c.UserData.Userid
				send.Data.Sendtime = int(time.Now().Unix())
				b, _ := send.Encode()
//				targetConn.Send <- *b

				m := vo.Broadcast{}
				m.Channel = ctos.Userid
				m.Msg = *b
				socket.Hub.Broadcast <-m

			} else {
				//插入离线消息
				db.MasterDB.AddKarmaMsg(c.UserData.Userid, ctos.Userid, ctos.Msg)
			}
			stoc.Data = &vo.StoC6109Data2{}
			stoc.Data.Sendtime = int(time.Now().Unix())
			stoc.Data.Msg = ctos.Msg
		}
	} else {
		stoc.E = 11003
	}

}

// 删除对方缘分
func deleRelation(msg *[]byte, c *socket.Connection) {
	ctos := vo.CtoS6110Data1{}
	ctos.Decode(msg)
	userData := c.UserData

	count := db.MasterDB.DeleteKarmaRelation(userData.Userid, ctos.Userid)
	stoc := vo.StoC6110Data1{T: 6110}
	if count == 0 {
		stoc.E = 13011
	}
	b, err := stoc.Encode()
	if err != nil {
		log.Errorln(err)
	}
	c.Send <- *b

}

// 缘分好友列表
func getkarmaRelationList(msg *[]byte, c *socket.Connection) {

	ctos := vo.CtoS6113Data1{}
	ctos.Decode(msg)
	userData := c.UserData

	list, _ := db.SlaveDB.GetKarmaRelation(userData.Userid)
	stoc := vo.StoC6113Data1{T: 6113}

	for i := 0; i < len(list); i++ {
		stoC6113Data2 := &vo.StoC6113Data2{}
		stoC6113Data2.Userid = list[i].Friendid

		stoc.Data = append(stoc.Data, stoC6113Data2)
	}
	b, err := stoc.Encode()
	if err != nil {
		log.Errorln(err)
	}
	c.Send <- *b
}
