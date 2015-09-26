/**
 * Created by Michael on 2015/7/30.
 */
package register
import (
	"net/http"
	log"github.com/golang/glog"
	"web/webvo"
	"web/email"
	"db"
	"web/proxy"
	"crypto/md5"
	"encoding/hex"
	"utils"
	"strconv"
	"time"
	"math/rand"
	"strings"
	"vo"
	"csv"
	"redis"
	"crypto/tls"
	"io/ioutil"
	"encoding/json"
	"web/login"
)
func init() {
	proxy.PrivilegeRegist(10100, registerAccount)// 用户注册
	proxy.PrivilegeRegist(10028, qqLogin)// QQ注册登陆
}


func qqLogin(body *[]byte, r *http.Request,w http.ResponseWriter) (bit *[]byte){
	webstoc:= webvo.WebStoC10028Data1{T:10028}
	defer func() {
		if err := recover(); err != nil {
			webstoc.E = 13108
		}
		bit, _ = webstoc.Encode()
	}()
	webctos := webvo.WebCtoS10028Data1{}
	webctos.Decode(body)

	ip:= "https://119.147.19.43/v3/user/is_login?"
	appid:= webctos.Appid
	pf:= webctos.Pf
	opendid:= webctos.Openid
	openkey:= webctos.Openkey
	url:= ip+"openid="+opendid+"&openkey=" + openkey + "&appid="+appid+"&pf=" + pf
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, _:= client.Get(url)

	if resp != nil{

		defer resp.Body.Close()
		content, _ := ioutil.ReadAll(resp.Body)
		type Ret struct  {
			Ret int
			Msg string
		}
		m:= Ret{}
		json.Unmarshal(content,&m)
		log.Infoln(m.Ret,m.Msg)
		// QQ 第三方登陆成功
		if m.Ret == 0{

			u := &db.Tb_user_passwd{Qq_uid:webctos.Uid}
			err := u.GetByQQUid()
			log.Infoln(err)
			// 没有注册
			if err != nil{
				userData := &vo.UserData{}

				userData.TerminalType = 2
				userData.Nickname =webctos.Nickname
				userData.WorldID = 1
				userData.Sex = webctos.Sex
				userData.Qq_uid = webctos.Uid

				userid,err := db.SlaveDB.GetUnregID();
				if(err != nil || userid == 0){
					log.Errorln(err)
					webstoc.E = 13108
					return bit
				}
				userData.Userid = userid
				_,err = db.MasterDB.RegisterUser(userData)
				if err != nil{
					log.Errorln(err)
					webstoc.E = 13108
					return bit
				}
				webstoc.Data = userData.Userid
			}else{
				webstoc.Data = u.Account_num
			}
			log.Infoln(webstoc.Data)
			cookie := redis.SetCookie(w, "login", u.Account_num)
			loginSession:=&redis.LoginSessionData{}
			loginSession.Userid = int64(u.Account_num)
			loginSession.SetSession(cookie.Value)

			currentIP := strings.Split(r.RemoteAddr, ":")[0]
			go login.UpdateUserLogin(currentIP,  u.Account_num)


		}else{
			webstoc.E = 13108
		}
	}else{
		webstoc.E = 13108
	}

	return bit
}

func registerAccount(body *[]byte, r *http.Request,w http.ResponseWriter) (bit *[]byte){
	webstoc:= webvo.WebStoC10100Data1{T:10100}
	defer func() {
		if err := recover(); err != nil {
			webstoc.E = 13108
			bit, _ = webstoc.Encode()
		}
	}()


	webctos := webvo.WebCtoS10100Data1{}
	webctos.Decode(body)
	var userid int
	userData := &vo.UserData{}
	if len(webctos.Pwd) == 32 {
		if checkAutoCode(webctos.Authcode) {
			if webctos.Account != "" {
				if utils.EmailRegexp(webctos.Account) {
					userData.Email = webctos.Account
					userid=db.SlaveDB.GetUserIDByEmail(webctos.Account)
				}else if (utils.PhoneRegexp(webctos.Account)) {
					userData.Phone = webctos.Account
					userid=db.SlaveDB.GetUserIDByPhone(webctos.Account)
				}else {
					webstoc.E = 13105
				}
			}else {
				webstoc.E =13104
			}
			if (userid > 0) {
				webstoc.E =13106
			}else {
				userData.TerminalType = webctos.Terminaltype
				userData.Auth = email.GenerateIdentifyCode()
				h := md5.New()
				h.Write([]byte(webctos.Pwd + userData.Auth))
				userData.Pwd = hex.EncodeToString(h.Sum(nil))
				userData.CreaterIP = utils.InetToaton(strings.Split(r.RemoteAddr,":")[0])
				worldData := csv.World.Hash
				for _, v := range worldData {
					if v.Default == 1 {
						userData.WorldID = v.Id
						break
					}
				}
				if userData.WorldID == 0{
					webstoc.E = 13108
				}else{
					userid, createTime, err := createUser(userData)
					if err != nil{
						webstoc.E = 13108
					}else{
						cookie := redis.SetCookie(w,"login",userid)
						loginSession:=&redis.LoginSessionData{}
						loginSession.Userid = int64(userid)
						loginSession.SetSession(cookie.Value)

						webstoc.Data = &webvo.WebStoC10100Data2{}
						webstoc.Data.Servertime = int(createTime)
						webstoc.Data.Userid = userid
					}
				}
			}
		}else {
			webstoc.E = 13006
		}
	}else{
		webstoc.E = 13107
	}

	b, _ := webstoc.Encode()
	return b
}

func createUser(userData *vo.UserData) (id int, createTime int64, e interface{}) {
	var userid int
	timestr := strconv.Itoa(int(time.Now().Unix()))
	nickname := "用户" + utils.SubStr(timestr,2,len(timestr)) + strconv.Itoa(rand.Intn(100))
	num,err := db.SlaveDB.CheckNickname(nickname)
	if (err != nil || num != 0) {
		log.Errorln(err)
		return userid, 0, err
	}

	userid,err = db.SlaveDB.GetUnregID();
	if(err != nil || userid == 0){
		log.Errorln(err)
		return userid, 0, err
	}
	userData.Userid = userid
	userData.Nickname = nickname
	cTime,err := db.MasterDB.RegisterUser(userData)
	if err != nil{
		log.Errorln(err)
		return userid,cTime,err
	}
	return userid, cTime, nil
}

func checkAutoCode(code string) bool {
	return true
}


