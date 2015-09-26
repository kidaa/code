package chat

import (
	. "db"
	log "github.com/golang/glog"
	"proxy"
	"socket"
	"strconv"
	"time"
	"vo"
	"db"
	"redis"
)

func init() {
	proxy.Regist(1012, NoReadMsgHdr)
	proxy.Regist(1013, NoDealRequestHdr)
	proxy.Regist(1009, FriendlistHdr)
	proxy.Regist(1001, MessagelistHdr)			// 离线消息列表
	proxy.Regist(1002, RequestlistHdr)
	proxy.Regist(1003, SearchaccountHdr)
	proxy.Regist(1004, AddfriendHdr)
	proxy.Regist(1006, AllowfriendHdr)
	proxy.Regist(1007, SendmsgHdr)
	proxy.Regist(1014, FriendsCountHdr)

	proxy.Regist(1016, modifyRemarksHdr)
	proxy.Regist(1018, delFriendHdr)
	proxy.Regist(1020, userMsgCountList)
}

const (
	MAX_PER_PAGE = 30
)

func userMsgCountList(msg *[]byte, c *socket.Connection) {

	stoc := vo.StoC1020Data1{T: 1020}
	//list, err := SlaveDB.MsgCountList(c.UserData.Userid)
	//	SlaveDB.GetLastMsgByOfUser(,c.UserData.Userid)
	list:= make([]*vo.MsgData,0,10)
	err:=redis.ChatMsg.GetAllOfflineChatMsg(c.UserData.Userid,&list)
	if err == nil {
		hash := make(map[int]*vo.StoC1020Data2)

		for i := 0; i < len(list); i++ {
			data, ok := hash[list[i].Userid]
			if ok {
				data.Msgnum++
			} else {
				data := &vo.StoC1020Data2{}
				hash[list[i].Userid] = data
				data.Msgnum = 1
				data.Userid = list[i].Userid
				//sendTime, msg, e := SlaveDB.GetLastMsgByOfUser(data.Userid, c.UserData.Userid)
//				d:= &vo.MsgData{}
//				e:=redis.ChatMsg.GetLastMsgByFriendid( c.UserData.Userid,data.Userid,d)
//				if e != nil {
//					log.Errorln(e)
//				}
				data.Sendtime = int(list[i].Sendtime)
				data.Msg = list[i].Msg
			}
		}

		arr := make([]*vo.StoC1020Data2, 10)
		arr = arr[:0]
		for _, v := range hash {
			arr = append(arr, v)
		}
		stoc.Data = arr

	}

	bit, err := stoc.Encode()
	if err != nil {
		log.Errorln(err)
	}
	c.Send <- *bit
}

// 删除好友
func delFriendHdr(msg *[]byte, c *socket.Connection) {
	ctos := vo.CtoS1018Data1{}
	ctos.Decode(msg)
	userid := c.UserData.Userid

	b := vo.StoC1018Data1{T: 1018}

	err := MasterDB.DeleteFriend(userid, ctos.Friendid)
	if err!= nil{
		b.E = 10001
	}
	b.Data = ctos.Friendid
	bit, _ := b.Encode()
	c.Send <- *bit
}

// 修改好友的备注
func modifyRemarksHdr(msg *[]byte, c *socket.Connection) {
	ctos := vo.CtoS1016Data1{}
	err := ctos.Decode(msg)

	b := vo.StoC1016Data1{T: ctos.T}

	if err != nil {
		b.E = 11009
	} else {
		userid := c.UserData.Userid
		if ctos.Friendid <= 0 {
			b.E = 11009
		} else {
			if ctos.Remarks == "" {
				b.E = 10002
			} else {
				affected := MasterDB.ModifyRemarks(userid, ctos.Friendid, ctos.Remarks)
				if affected == 0 {
					b.E = 10001
				}
			}
		}
	}

	bit, _ := b.Encode()
	c.Send <- *bit
}

// 好友数量
func FriendsCountHdr(msg *[]byte, c *socket.Connection) {
	b := vo.StoC1014Data1{T: 1014}
	b.Data = SlaveDB.FriendsCount(c.UserData.Userid)
	bit, _ := b.Encode()
	c.Send <- *bit
}

// 未处理的加好友请求数量
func NoDealRequestHdr(msg *[]byte, c *socket.Connection) {
	b := vo.StoC1013Data1{T: 1013}
	b.Data = SlaveDB.NoDealRequestCount(c.UserData.Userid)
	bit, _ := b.Encode()
	c.Send <- *bit
}

// 未处理的聊天消息数量
func NoReadMsgHdr(msg *[]byte, c *socket.Connection) {
	b := vo.StoC1012Data1{T: 1012}
	b.Data,_= redis.ChatMsg.GetOfflineChatMsgCount(c.UserData.Userid)
	bit, _ := b.Encode()
	c.Send <- *bit
}

// 好友列表
func FriendlistHdr(msg *[]byte, c *socket.Connection) {
	ctoS1009Data1 := vo.CtoS1009Data1{}
	err := ctoS1009Data1.Decode(msg)
	if err != nil {
		log.Errorln(err)
	}
	ctoC1009 := vo.StoC1009Data1{T: ctoS1009Data1.T}
	userList,e:=SlaveDB.GetFriendListByUserid(c.UserData.Userid)
	if e == nil{

		list := make([]*vo.StoC1009Data2,0, 10)
		for _,datav:= range userList {
			data := &vo.StoC1009Data2{}

			member := &db.Tb_user_member{Account_num:datav.FriendID}
			member.Read()
			data.Userid = member.Account_num
			data.Nickname = member.Nickname
			data.Sex = member.Sex
			data.Headpic = int(member.Create_time)
			data.Remark = datav.Remarks
			list = append(list, data)
		}
		ctoC1009.Data = list
	}else{
		ctoC1009.E = 11010
	}

	b, _ := ctoC1009.Encode()
	c.Send <- *b
}

//消息列表(完成)
func MessagelistHdr(msg *[]byte, c *socket.Connection) {
	ctoS1001Data1 := vo.CtoS1001Data1{}
	err := ctoS1001Data1.Decode(msg)
	if err != nil {
		log.Errorln(err)
	}
	ctoC := vo.StoC1001Data1{T: ctoS1001Data1.T}

	messagelist := make([]*vo.StoC1001Data2, 10)
	messagelist = messagelist[:0]

	msgList := make([]*vo.MsgData,0,10)
//	msgList, e := SlaveDB.MsgByUserList(c.UserData.Userid, ctoS1001Data1.Userid)
	e:= redis.ChatMsg.GetOfflineChatMsgByFriendid(c.UserData.Userid,ctoS1001Data1.Userid,0,&msgList)
	if e == nil {
		for i := 0; i < len(msgList); i++ {
			m := msgList[i]
			msgData := &vo.StoC1001Data2{}
			msgData.Userid = m.Userid
			msgData.Msg = m.Msg
			msgData.Sendtime = int(m.Sendtime)
			messagelist = append(messagelist, msgData)
		}
	} else {
		ctoC.E = 11006
	}

	ctoC.Data = messagelist
	b, _ := ctoC.Encode()
	c.Send <- *b
}

// 请求列表
func RequestlistHdr(msg *[]byte, c *socket.Connection) {
	ctoS1002Data1 := vo.CtoS1002Data1{}
	err := ctoS1002Data1.Decode(msg)
	if err != nil {
		log.Errorln(err)
	}
	userData := c.UserData
	num := userData.Userid
	toc := vo.StoC1002Data1{T: ctoS1002Data1.T}
	requestList := SlaveDB.RequestList(num, MAX_PER_PAGE*ctoS1002Data1.Page, MAX_PER_PAGE)

	toc.Data = requestList
	b, _ := toc.Encode()
	c.Send <- *b

}

// 搜索用户数据，如果搜索添加未数据且大于5位数则直接搜索用户ID，否则模糊搜索用户昵称(完成)
func SearchaccountHdr(msg *[]byte, c *socket.Connection) {
	stoc := vo.StoC1003Data1{T:1003}
	defer func() {
		if err := recover(); err != nil {
			stoc.E = 20001
			bit, _ := stoc.Encode()
			c.Send <- *bit
			log.Infoln(err)
		}
	}()

	var data vo.CtoS1003Data1
	err := data.Decode(msg)

	var arr []*vo.StoC1003Data2
	if err == nil {
		k, err := strconv.Atoi(data.Key)
		if err == nil && len(data.Key) >= 5 {
			arr = make([]*vo.StoC1003Data2, 0,1)

			userdata:= Tb_user_member{Account_num:k}
			err:=userdata.Read()
			if err == nil {
				stoC1003Data2 := &vo.StoC1003Data2{}
				stoC1003Data2.Headpic = int(userdata.Create_time)
				stoC1003Data2.Nickname = userdata.Nickname
				stoC1003Data2.Sex = userdata.Sex
				stoC1003Data2.Userid = userdata.Account_num
				stoC1003Data2.Sign = userdata.Sign
				stoC1003Data2.Provinceid = userdata.Province_id
				stoC1003Data2.Cityid = userdata.City_id
				if SlaveDB.IsFriend(c.UserData.Userid, k) != nil{
					stoC1003Data2.Isfriend = 1
				}

				arr = append(arr,stoC1003Data2)
			} else {
				stoc.E = 15008
			}
		} else {
			user := SlaveDB.BlurSearchUserNickname(c.UserData.Userid, data.Key, data.Page*MAX_PER_PAGE, MAX_PER_PAGE)
			arr = make([]*vo.StoC1003Data2,0, 10)
			for i := 0; i < len(user); i++ {
				userData := user[i]
				stoC1003Data2 := &vo.StoC1003Data2{}
				stoC1003Data2.Headpic = userData.Headpic
				stoC1003Data2.Nickname = userData.Nickname
				stoC1003Data2.Sex = userData.Sex
				stoC1003Data2.Userid = userData.Userid
				stoC1003Data2.Isfriend = userData.Isfriend
				arr = append(arr, stoC1003Data2)
			}
		}
	}else{
		stoc.E = 20001
	}

	stoc.Data = arr
	bit, _ := stoc.Encode()
	c.Send <- *bit
}

// 发起加好友请求
func AddfriendHdr(msg *[]byte, c *socket.Connection) {

	var data vo.CtoS1004Data1
	err := data.Decode(msg)
	if err != nil {
		log.Errorln(err)
	}

	toc := &vo.StoC1004Data1{T: data.T}
	incept_account := data.Userid
	if incept_account > 0 {
		req_account := c.UserData.Userid
		request := &vo.StoC1005Data1{T: 1005}
		request.Data = &vo.StoC1005Data2{}
		request.Data.Userid = c.UserData.Userid
		request.Data.Nickname = c.UserData.Nickname
		request.Data.Msg = data.Msg

		//自己不能添加未自己的好友
		if req_account != incept_account {
			// 已经是好友
			if SlaveDB.IsFriend(req_account, incept_account) != nil && SlaveDB.IsFriend(incept_account, req_account) != nil {
				toc.E = 11001
			} else {
				// 往数据库插入添加好友请求信息
				MasterDB.InsertFriendRequst(req_account, incept_account, data.Msg)
				// 已经是好友提示不能添加（为完成）
//				targetConn, ok := socket.Hub.Connections[data.Userid]



				if socket.Hub.Exists(data.Userid) {
					b, _ := request.Encode()
//					targetConn.Send <- *b
					m := vo.Broadcast{}
					m.Channel = data.Userid
					m.Msg = *b
					socket.Hub.Broadcast <-m
				}
			}
		} else {
			toc.E = 11002
		}
	} else {
		toc.E = 11008 // 好友ID为0
	}
	b, _ := toc.Encode()
	c.Send <- *b
}

// 允许加好友(完成)
func AllowfriendHdr(msg *[]byte, c *socket.Connection) {

	data := vo.CtoS1006Data1{}
	err := data.Decode(msg)
	toC := vo.StoC1006Data1{T: data.T}
	if err == nil {

		if data.Userid > 0 {

			userdata:= Tb_user_member{Account_num:data.Userid}
			err:=userdata.Read()
			if err == nil {

				// 互相添加好友到数据库
				MasterDB.InsertAddFriend(data.Userid, c.UserData.Userid)

				// 成功添加好友，删除双方的加好友请求
				MasterDB.DeleteRequestAddFriend(data.Userid, c.UserData.Userid)
				MasterDB.DeleteRequestAddFriend(c.UserData.Userid, data.Userid)

//				targetConn, ok := socket.Hub.Connections[data.Userid]
				if socket.Hub.Exists(data.Userid) {
					send := vo.StoC1011Data1{T: 1011,Data:&vo.StoC1011Data2{}}
					send.Data = &vo.StoC1011Data2{}
					send.Data.Nickname = c.UserData.Nickname
					send.Data.Userid = c.UserData.Userid
					b, _ := send.Encode()
//					targetConn.Send <- *b
					m := vo.Broadcast{}
					m.Channel = data.Userid
					m.Msg = *b
					socket.Hub.Broadcast <-m
				}



				toC.Data = & vo.StoC1006Data2{}
				toC.Data.Nickname = userdata.Nickname
				toC.Data.Headpic =int(userdata.Create_time)
				toC.Data.Sex = userdata.Sex
				toC.Data.Userid = userdata.Account_num
			} else {
				toC.E = 15008 // 用户不存在
			}

		} else {
			toC.E = 11008 // 好友ID为0
		}
	} else {
		toC.E = 20001 // 提交的数据格式错误
	}

	b, _ := toC.Encode()
	c.Send <- *b
}

// 聊天消息处理 (完成)
func SendmsgHdr(msg *[]byte, c *socket.Connection) {
	data := vo.CtoS1007Data1{}
	err := data.Decode(msg)

	if err != nil {
		log.Errorln(err)
	}
	toc := vo.StoC1007Data1{T: data.T}

	if data.Userid > 0 {
		if SlaveDB.IsFriend(data.Userid, c.UserData.Userid) != nil {
//			targetConn, ok := socket.Hub.Connections[data.Userid]
			if socket.Hub.Exists(data.Userid) {
				send := vo.StoC1008Data1{T: 1008,Data:&vo.StoC1008Data2{}}

				send.Data.Msg = data.Msg
				send.Data.Userid = c.UserData.Userid
				send.Data.Sendtime = int(time.Now().Unix())
				b, _ := send.Encode()
//				targetConn.Send <- *b

				m := vo.Broadcast{}
				m.Channel = data.Userid
				m.Msg = *b
				socket.Hub.Broadcast <-m
			} else {
				//插入离线消息
				d:= &vo.MsgData{}
				d.Sendtime = time.Now().Unix()
				d.Userid = c.UserData.Userid
				d.Msg = data.Msg
				redis.ChatMsg.AddOfflineChatMsg(data.Userid,c.UserData.Userid,d)
			}
		} else {
			toc.E = 11003  //  对方把自己删了
		}
	} else {
		toc.E = 11008 // 好友ID为0
	}

	toc.Data = &vo.StoC1007Data2{}
	toc.Data.Sendtime = int(time.Now().Unix())
	toc.Data.Msg = data.Msg
	b, _ := toc.Encode()
	c.Send <- *b
}
