package user

import (
	. "vo"

	"db"
	log "github.com/golang/glog"
	"proxy"
	"socket"
)

func init() {
	proxy.Regist(3001, getUserDataHdr)
	proxy.Regist(3002, checkUserOnline)
}

// 获取用户的详细资料
func getUserDataHdr(msg *[]byte, c *socket.Connection) {
	result := StoC3001Data1{T: 3001}
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			result.E = 14001
		}
		bit, _ := result.Encode()
		c.Send <- *bit
	}()

	webctos := CtoS3001Data1{}
	webctos.Decode(msg)

	if webctos.Userid <= 0 {
		result.E = 14001
	} else {
		member := &db.Tb_user_member{Account_num: webctos.Userid}
		err := member.Read()

		if err == nil {

			result.Data = &StoC3001Data2{}

			world := &db.Tb_world_currently{Account_num: webctos.Userid}
			err = world.Read()

			score := &db.Tb_score_stat{Account_num: webctos.Userid}
			err = score.Read()

			rel := db.SlaveDB.IsFriend(c.UserData.Userid, webctos.Userid)
			// 好友关系判断，是否为好友
			if rel != nil {
				result.Data.Isfriend = 1
				result.Data.Remark = rel.Remarks
			}

			result.Data.World_id = world.World_id

			result.Data.Currency_total = int(score.Currency_total)
			result.Data.Score_total = int(score.Score_total)

			result.Data.Account_email = member.Account_email
			result.Data.Account_num = member.Account_num
			result.Data.Birth = int(member.Birth)
			result.Data.City_id = member.City_id
			result.Data.Creater_ip = int(member.Creater_ip)
			result.Data.Grade = member.Grade
			result.Data.Grade_exp = int(member.Grade_exp * 100)
			result.Data.Headpic = int(member.Create_time)
			result.Data.Terminal_type = member.Terminal_type
			result.Data.Province_id = member.Province_id
			result.Data.Nickname = member.Nickname
			result.Data.Phone_num = member.Phone_num
			result.Data.Sign = member.Sign
			result.Data.Validate_status = member.Status
			result.Data.Sex = member.Sex
			result.Data.User_type = member.User_type

		} else {
			result.E = 14001
		}
	}

	log.Infoln(webctos.Userid)
}

func checkUserOnline(msg *[]byte, c *socket.Connection) {
	ctos := CtoS3002Data1{}
	ctos.Decode(msg)
	stoc := StoC3002Data1{T: ctos.T}
	ok := socket.Hub.Exists(ctos.Userid)
	if ok {
		stoc.Data = 1
	}
	b, err := stoc.Encode()
	if err != nil {
		log.Errorln(err)
	}
	c.Send <- *b
}
