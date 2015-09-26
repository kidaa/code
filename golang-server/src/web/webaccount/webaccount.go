/**
 * Created by Michael on 2015/8/11.
 *	操作用户账户相关的信息
 *
 *
 */
package webaccount

import (
	"db"
	log "github.com/golang/glog"
	"regexp"
	"utils"
	"web/email"
	"web/proxy"
	"web/sms"
	"web/webvo"
)

func init() {
	proxy.Regist(10011, saveUserNickname)
	proxy.Regist(10012, saveUserSex)
	proxy.Regist(10013, saveUserAddress)
	proxy.Regist(10014, saveUserSign)
	proxy.Regist(10015, saveUserPwd)
	proxy.Regist(10018, getUserDialogCount)
	proxy.Regist(10020, getUsersData)
	proxy.Regist(10010, getUserData) //  获取用户的详细资料
	proxy.Regist(10021, getUsersActiveTime)
	proxy.Regist(10022, sendCode)  //给将要绑定的手机或邮箱发送验证码
	proxy.Regist(10023, isBand)    //验证是否绑定了指定的手机或邮箱
	proxy.Regist(10024, banding)   //绑定指定的手机或邮箱,code: 发送到手机或邮箱的验证码，new：新手机或邮箱号
	proxy.Regist(10026, deBanding) //解除绑定手机或邮箱,old:之前绑定的手机或邮箱号
}

func sendCode(uid int, body *[]byte) *[]byte {
	ctos := webvo.WebCtoS10022Data1{}
	ctos.Decode(body)
	log.Infoln(uid)
	stoc := webvo.WebStoC10022Data1{T: 10022}

	if utils.EmailRegexp(ctos.New) {
		identifyCode := email.GenerateIdentifyCode()
		email.Send(ctos.New, identifyCode, uid)

	} else if utils.PhoneRegexp(ctos.New) {
		identifyCode := email.GenerateIdentifyCode()
		sms.SendMSM(ctos.New, identifyCode)
	} else {
		stoc.E = 20001 // 数据格式错误
	}

	b, _ := stoc.Encode()
	return b
}

func isBand(uid int, body *[]byte) *[]byte {
	log.Infoln(uid)
	ctos := webvo.WebCtoS10023Data1{}
	ctos.Decode(body)

	stoc := webvo.WebStoC10023Data1{T: 10023}

	if utils.EmailRegexp(ctos.Id) {
		u := &db.Tb_user_passwd{Account_num: uid}
		err := u.Read()
		if err == nil {
			if ctos.Id == u.Account_email {

			} else {
				stoc.E = 13003
			}
		} else {
			log.Errorln(err)
			stoc.E = 14001 // 用户不存在
		}
	} else if utils.PhoneRegexp(ctos.Id) {
		u := &db.Tb_user_passwd{Account_num: uid}
		err := u.Read()
		if err == nil {
			if ctos.Id == u.Phone_num {

			} else {
				stoc.E = 13112
			}
		} else {
			log.Errorln(err)
			stoc.E = 14001 // 用户不存在
		}
	} else {
		stoc.E = 20001 // 数据格式错误
	}

	b, _ := stoc.Encode()
	return b
}

func banding(uid int, body *[]byte) *[]byte {
	ctos := webvo.WebCtoS10024Data1{}
	ctos.Decode(body)
	log.Infoln(uid, ctos.New)
	stoc := webvo.WebStoC10024Data1{T: 10024}

	if utils.EmailRegexp(ctos.New) {
		u := &db.Tb_user_passwd{Account_num: uid}
		err := u.Read()
		if err == nil {
			u.Account_email = ctos.New
			u.UpateEmail()

			m := &db.Tb_user_member{Account_num: uid}
			err := m.Read()
			if err == nil {
				m.Account_email = ctos.New
				m.UpateEmail()
			}
		} else {
			log.Errorln(err)
			stoc.E = 14001 // 用户不存在
		}
	} else if utils.PhoneRegexp(ctos.New) {
		u := &db.Tb_user_passwd{Account_num: uid}
		err := u.Read()
		if err == nil {
			u.Phone_num = ctos.New
			u.UpatePhone()

			m := &db.Tb_user_member{Account_num: uid}
			err := m.Read()
			if err == nil {
				m.Phone_num = ctos.New
				m.UpatePhone()
			}

		} else {
			log.Errorln(err)
			stoc.E = 14001 // 用户不存在
		}

	} else {
		stoc.E = 20001 // 数据格式错误
	}

	b, _ := stoc.Encode()
	return b
}

func deBanding(uid int, body *[]byte) *[]byte {
	ctos := webvo.WebCtoS10026Data1{}
	ctos.Decode(body)
	log.Infoln(uid)
	stoc := webvo.WebStoC10026Data1{T: 10026}

	if utils.EmailRegexp(ctos.Old) {
		u := &db.Tb_user_passwd{Account_num: uid}
		err := u.Read()
		if err == nil {
			if ctos.Old == u.Account_email {
				u.Account_email = ""
				u.UpateEmail()

				m := &db.Tb_user_member{Account_num: uid}
				err := m.Read()
				if err == nil {
					m.Account_email = ""
					m.UpateEmail()
				}
			} else {
				stoc.E = 13003
			}
		} else {
			log.Errorln(err)
			stoc.E = 14001 // 用户不存在
		}
	} else if utils.PhoneRegexp(ctos.Old) {
		u := &db.Tb_user_passwd{Account_num: uid}
		err := u.Read()
		if err == nil {
			if ctos.Old == u.Phone_num {
				u.Phone_num = ""
				u.UpatePhone()

				m := &db.Tb_user_member{Account_num: uid}
				err := m.Read()
				if err == nil {
					m.Phone_num = ""
					m.UpatePhone()
				}

			} else {
				stoc.E = 13112
			}
		} else {
			log.Errorln(err)
			stoc.E = 14001 // 用户不存在
		}

	} else {
		stoc.E = 20001 // 数据格式错误
	}

	b, _ := stoc.Encode()
	return b
}

// 获取用户在线时长和活跃天数
func getUsersActiveTime(uid int, body *[]byte) *[]byte {
	ctos := webvo.WebCtoS10021Data1{}
	ctos.Decode(body)
	result := webvo.WebStoC10021Data1{T: 10021}
	result.Data = &webvo.WebStoC10021Data2{}
	tb := &db.Tb_user_active{Account_num: uid}

	if tb.Request() != nil {
		result.E = 14005
	} else {
		result.Data.Day = tb.Day_num
		result.Data.Active = int(tb.Active_num * 100)
		result.Data.Time = tb.Online_time
	}

	b, _ := result.Encode()
	return b
}

// 获取用户的详细资料
func getUserData(uid int, body *[]byte) *[]byte {
	webctos := webvo.WebCtoS10010Data1{}
	webctos.Decode(body)
	result := webvo.WebStoC10010Data1{T: webctos.T}
	if webctos.Otherid <= 0 {
		result.E = 14001
	} else {
		member := &db.Tb_user_member{Account_num: webctos.Otherid}
		err := member.Read()

		if err == nil {
			result.Data = &webvo.WebStoC10010Data2{}

			world := &db.Tb_world_currently{Account_num: webctos.Otherid}
			err = world.Read()

			score := &db.Tb_score_stat{Account_num: webctos.Otherid}
			err = score.Read()

			rel := db.SlaveDB.IsFriend(webctos.Userid, webctos.Otherid)
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

	b, _ := result.Encode()
	return b
}

// 获取一组用户基础数据
func getUsersData(uid int, body *[]byte) *[]byte {
	ctos := webvo.WebCtoS10020Data1{}
	ctos.Decode(body)
	stoc := webvo.WebStoC10020Data1{T: ctos.T}

	if len(ctos.Userids) > 0 {
		//		list, e := db.SlaveDB.GetUsersData(ctos.Userids)
		//		if e != nil {
		//			stoc.E = 14004
		//		} else {
		for i := 0; i < len(ctos.Userids); i++ {
			data := &webvo.WebStoC10020Data2{}
			m := db.Tb_user_member{Account_num: ctos.Userids[i]}
			err := m.Read()
			if err == nil {
				data.Userid = m.Account_num
				data.Cityid = m.City_id
				data.Provinceid = m.Province_id
				data.Headpic = int(m.Create_time)
				data.Nickname = m.Nickname
				data.Sex = m.Sex
				stoc.Data = append(stoc.Data, data)
			}
		}
		//		}
	} else {
		stoc.E = 14002
	}

	b, _ := stoc.Encode()
	return b
}

// 更改用户昵称
func saveUserNickname(uid int, body *[]byte) *[]byte {
	ctos := webvo.WebCtoS10011Data1{}
	ctos.Decode(body)
	stoc := webvo.WebStoC10011Data1{T: ctos.T}
	//todo:在cookie中检测自身nickname
	nickname, nicksum := utils.GetTrimStrLen(ctos.Nickname)
	if nicksum < 2 || nicksum > 10 {
		log.Errorln(nicksum)
		stoc.E = 13211
	} else {
		has := checkNickname(nickname)
		if has {
			log.Errorln("nickname已存在")
			stoc.E = 13210
		} else {
			user := db.Tb_user_member{Account_num: ctos.Userid}
			user.Nickname = nickname
			e := user.UpateNickname()
			if e != nil {
				stoc.E = 13209
				log.Errorln(e)
			}
		}
	}

	b, _ := stoc.Encode()
	return b
}

// 更改用户性别
func saveUserSex(uid int, body *[]byte) *[]byte {
	ctos := webvo.WebCtoS10012Data1{}
	ctos.Decode(body)
	stoc := webvo.WebStoC10012Data1{T: ctos.T}

	if ctos.Sex < 1 || ctos.Sex > 3 {
		stoc.E = 13212
		log.Errorln("性别类型错误")
	} else {
		user := db.Tb_user_member{Account_num: ctos.Userid}
		user.Sex = ctos.Sex
		e := user.UpateSex()
		if e != nil {
			stoc.E = 13207
			log.Errorln(e)
		}
	}
	b, _ := stoc.Encode()
	return b
}

// 更改用户地址
func saveUserAddress(uid int, body *[]byte) *[]byte {
	ctos := webvo.WebCtoS10013Data1{}
	ctos.Decode(body)
	stoc := webvo.WebStoC10013Data1{T: ctos.T}
	if ctos.Provinceid < 1 || ctos.Cityid < 1 {
		stoc.E = 13214
	} else {
		user := db.Tb_user_member{Account_num: ctos.Userid}
		user.Province_id = ctos.Provinceid
		user.City_id = ctos.Cityid
		e := user.UpateAddress()
		if e != nil {
			stoc.E = 13208
			log.Errorln(e)
		}
	}
	b, _ := stoc.Encode()
	return b
}

// 更改用户签名
func saveUserSign(uid int, body *[]byte) *[]byte {
	ctos := webvo.WebCtoS10014Data1{}
	ctos.Decode(body)
	stoc := webvo.WebStoC10014Data1{T: ctos.T}

	sign, signsum := utils.GetTrimStrLen(ctos.Sign)
	if signsum > 20 {
		log.Errorln(signsum)
		stoc.E = 13213
	} else {
		user := db.Tb_user_member{Account_num: ctos.Userid}
		user.Sign = sign
		e := user.UpateSign()
		if e != nil {
			stoc.E = 13205
			log.Errorln(e)
		}
	}
	b, _ := stoc.Encode()
	return b
}

// 更改用户密码
func saveUserPwd(uid int, body *[]byte) *[]byte {
	ctos := webvo.WebCtoS10015Data1{}
	ctos.Decode(body)
	stoc := webvo.WebStoC10015Data1{T: ctos.T}
	if ctos.Pwd == ctos.Newpwd {
		log.Errorln(ctos.Pwd, ctos.Newpwd)
		stoc.E = 13110
	} else {
		if len(ctos.Pwd) == 32 {
			if len(ctos.Newpwd) == 32 {
				if ctos.Userid > 0 {
					pwd := db.Tb_user_passwd{Account_num: ctos.Userid}
					//  密码正确
					if pwd.PWDIsOK(ctos.Pwd) {
						// 获取用户资料
						pwd.Passwd = ctos.Newpwd
						pwd.UpdatePWD()
						stoc.E = 0
					} else {
						stoc.E = 13111
					}
				}
			} else {
				stoc.E = 13007
			}
		} else {
			stoc.E = 13109
		}
	}
	b, _ := stoc.Encode()
	return b
}

// 检测用户昵称是否存在
func checkNickname(nickname string) bool {
	reg := regexp.MustCompile(`^用户\d{10}$`)
	if reg.FindString(nickname) != "" {
		return true
	}
	num, err := db.SlaveDB.CheckNickname(nickname)
	if err != nil || num != 0 {
		return true
	}
	return false
}

// 用户的对白数量
func getUserDialogCount(uid int, body *[]byte) *[]byte {
	ctos := webvo.WebCtoS10018Data1{}
	ctos.Decode(body)
	stoc := webvo.WebStoC10018Data1{T: ctos.T}

	count, e := db.SlaveDB.UserDialogCount(ctos.Userid)
	if e != nil {
		stoc.E = 13206
		log.Errorln(e)
	}
	stoc.Data = count

	b, _ := stoc.Encode()
	return b
}
