package socketlogin

/*
*
*socket  websocket 登陆协议处理
*
 */

import (
	"db"
	"proxy"
	"redis"
	"socket"
	"time"
	"vo"
)

const (
	Login        = 1000
	ReceiveLogin = 1010
)

func init() {
	proxy.Regist(Login, login)
}

// 登陆
func login(msg *[]byte, c *socket.Connection) {
	stoc := vo.StoC1000Data1{T: Login}

	data := &vo.CtoS1000Data1{}
	err := data.Decode(msg)
	var userid int
	if err != nil {
		stoc.Data = 13001
	} else {
		userid, err = redis.GetUseridByCookie(data.Session)
		if err != nil || userid == 0 {
			stoc.Data = 13009
		}
	}

	if userid > 0 && socket.Hub.IsRelogin(userid,data.Session){
		stoc.Data = 13113
	}

	if stoc.Data > 0 {
		b, _ := stoc.Encode()
		c.Send <- *b
		<-time.After(time.Millisecond * 100)
		c.Close()
		return
	}

	c.UserData.Userid = userid
	//已经登录了
	if socket.Hub.Exists(userid) {
		cmd := vo.StoC1010Data1{T: ReceiveLogin}
		reLogin, _ := cmd.Encode()
		// 先等消息发送完下断开连接
		m := vo.Broadcast{}
		m.Channel = userid
		m.Msg = *reLogin
		socket.Hub.Broadcast <- m
		<-time.After(time.Millisecond * 100)

		m = vo.Broadcast{}
		m.Channel = userid
		m.Channel = userid
		m.Kick = true
		socket.Hub.Broadcast <- m
		<-time.After(time.Millisecond * 100)
	}
	member := &db.Tb_user_member{Account_num: c.UserData.Userid}
	member.Read()
	c.UserData.Nickname = member.Nickname
	c.UserData.Email = member.Account_email
	c.UserData.Phone = member.Phone_num
	c.UserData.Headpic = int(member.Create_time)
	world := db.Tb_world_currently{Account_num: c.UserData.Userid}
	world.GetWorldid()
	c.UserData.WorldID = world.World_id
	c.UserData.Session = data.Session
	socket.Hub.Register <- c
	b, _ := stoc.Encode()
	c.Send <- *b
	// 保存登陆时间，用于在线时长计算
	c.IsLogined = true
	go saveLoginUinxTime(c.UserData.Userid)
}

// 保存登陆时间，用于在线时长计算
func saveLoginUinxTime(userid int) {
	tb := db.Tb_user_active{Account_num: userid}
	tb.Is_online = 1
	tb.Login()
}
