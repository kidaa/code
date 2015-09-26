/**
 * Created by Michael on 2015/7/30.
 */
package login

import (
	"db"
	"net/http"
	"redis"
	"strconv"
	"strings"
	"time"
	"utils"
	"web/email"
	"web/proxy"
	"web/sms"
	"web/weather"
	"web/webvo"
	log"github.com/golang/glog"
)

func init() {
	proxy.PrivilegeRegist(10001, receivePNGCode)         // 向账户绑定的邮箱发送验证码
	proxy.PrivilegeRegist(10004, receiveIdentifyingCode) // 用户获取到6位手机验证码，请求修改用户账户的密码,前端MD5加密
	proxy.PrivilegeRegist(10009, receiveChangePWD)       // 向账户绑定的手机发送验证码
	proxy.PrivilegeRegist(10000, userlogin)              // 用户登陆
}

// 用户登陆，登陆成功后返回用户的详细账户资料，并返回Cookie,用户根据这个cookie,登陆socket服务器
func userlogin(body *[]byte, r *http.Request, w http.ResponseWriter) (bit *[]byte) {

	stoc := webvo.WebStoC10000Data1{T: 10000}
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			stoc.E = 13001
			bit, _ = stoc.Encode()
		}
	}()

	webctos := webvo.WebCtoS10000Data1{}
	webctos.Decode(body)

	stoc.E = 13001
	var userid int
	var err error

	if webctos.Account != "" {
		if utils.EmailRegexp(webctos.Account) {
			userid = db.SlaveDB.GetUserIDByEmail(webctos.Account)
		} else if utils.PhoneRegexp(webctos.Account) {
			userid = db.SlaveDB.GetUserIDByPhone(webctos.Account)
		} else {
			userid, err = strconv.Atoi(webctos.Account)
			if err != nil {
				stoc.E = 13001
			}
		}
	} else {
		stoc.E = 13104
	}

	if userid > 0 {
		user := db.Tb_user_passwd{Account_num: userid}
		//  密码正确
		if len(webctos.Pwd) == 32 && user.PWDIsOK(webctos.Pwd) {
			// 获取用户资料
			cookie := redis.SetCookie(w, "login", userid)
			loginSession:=&redis.LoginSessionData{}
			loginSession.Userid = int64(userid)
			loginSession.SetSession(cookie.Value)
			currentIP := strings.Split(r.RemoteAddr, ":")[0]
			go UpdateUserLogin(currentIP, userid)
			stoc.E = 0
			stoc.Data = &webvo.WebStoC10000Data2{}
			stoc.Data.Userid = userid
			stoc.Data.Servertime = int(time.Now().Unix())
			stoc.Data.Ens = cookie.Value
		}else{
			stoc.E = 13001
		}
	}

	b, err := stoc.Encode()
	return b
}

// 用户提交绑定的手机或者邮箱号，并正确输入图片验证码，服务器判断图标验证码正确，即向绑定的邮箱或在手机发送带有更改账户密码的验证码
func receivePNGCode(body *[]byte, r *http.Request, w http.ResponseWriter) *[]byte {
	webctos := webvo.WebCtoS10001Data1{}
	webctos.Decode(body)

	webstoc := webvo.WebStoC10001Data1{T: 10001}
	cookie, _ := r.Cookie("pngverify")
	if cookie == nil || cookie.Value == "" {
		webstoc.E = 13009
		b, _ := webstoc.Encode()

		return b
	}

	sessonData := &redis.SessionData{}
	err := sessonData.Decode(cookie.Value)

	if err == nil {
		serverCode := sessonData.PngCode
		if serverCode == webctos.Pngcode {
			identifyCode := email.GenerateIdentifyCode()
			var userid int

			if utils.EmailRegexp(webctos.PhoneORmail) == false {
				userid = db.SlaveDB.GetUserIDByPhone(webctos.PhoneORmail)
				sessonData.Phone = webctos.PhoneORmail
				sms.SendMSM(webctos.PhoneORmail, identifyCode)
			} else {
				userid = db.SlaveDB.GetUserIDByEmail(webctos.PhoneORmail)
				sessonData.Email = webctos.PhoneORmail
				email.Send(webctos.PhoneORmail, identifyCode, userid)
			}

			sessonData.IdentifyCode = identifyCode
			sessonData.Userid = userid

			go sessonData.Encode(cookie.Value)

		} else {
			webstoc.E = 13008
		}
	} else {
		webstoc.E = 13009
	}

	b, _ := webstoc.Encode()
	return b
}

// 用户从手机或者邮箱里获取验证码，提交到服务器验证，如果验证通过，即进入等待输入密码状态
func receiveIdentifyingCode(body *[]byte, r *http.Request, w http.ResponseWriter) *[]byte {
	webctos := webvo.WebCtoS10004Data1{}
	webctos.Decode(body)
	webstoc := webvo.WebStoC10004Data1{T: 10004}

	cookie, _ := r.Cookie("pngverify")

	if cookie == nil || cookie.Value == "" {
		webstoc.E = 13009
		b, _ := webstoc.Encode()
		return b
	}

	sessonData := &redis.SessionData{}
	err := sessonData.Decode(cookie.Value)
	if err == nil {
		if sessonData.IdentifyCode == webctos.Code {
			sessonData.IdentifyCode = ""
			sessonData.PasswordChangeStatus = true
			go sessonData.Encode(cookie.Value)

		} else {
			webstoc.E = 13006
		}
	} else {
		webstoc.E = 13009
	}
	b, _ := webstoc.Encode()
	return b
}

// 前面修改密码验证码通过验证，这里用户提交密码后，用户的账户密码即被更改
func receiveChangePWD(body *[]byte, r *http.Request, w http.ResponseWriter) *[]byte {
	webctos := webvo.WebCtoS10009Data1{}
	webctos.Decode(body)
	webstoc := webvo.WebStoC10009Data1{T: 10009}

	cookie, _ := r.Cookie("pngverify")
	if cookie == nil || cookie.Value == "" {
		webstoc.E = 13009
		b, _ := webstoc.Encode()
		return b
	}

	sessonData := &redis.SessionData{}
	err := sessonData.Decode(cookie.Value)
	if err == nil {
		if sessonData.PasswordChangeStatus {
			pwd := db.Tb_user_passwd{Account_num: sessonData.Userid}
			pwd.Passwd = webctos.Newpwd
			pwd.UpdatePWD()

			go sessonData.Clear(cookie.Value)
			go redis.ClearCookie(w, r, "pngverify")
		}
	} else {
		webstoc.E = 13009
	}

	b, _ := webstoc.Encode()
	return b
}

func UpdateUserLogin(currentIP string, userid int) {

	ip := utils.InetToaton(currentIP)

	tb := db.Tb_user_active{Account_num: userid}
	tb.GetIP()
	// 上次登陆地址跟现在的地址不一致，则从新获取用户的物理位置
	if tb.Login_ip != ip {
		cityVO := weather.IP2City(currentIP)
		if cityVO != nil {
			//			userData.City = cityVO.City
			//			userData.Province= cityVO.Region
		}
		tb.Account_num = userid
		tb.Login_ip = ip
		tb.UpdateIP()
	}
}
