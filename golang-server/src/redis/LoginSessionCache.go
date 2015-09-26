/**
 * Created by Michael on 2015/8/17.
 * 用户登陆session
 */
package redis

import (
	"github.com/garyburd/redigo/redis"
	"bytes"
	"encoding/gob"
	"errors"
	"crypto/md5"
	"encoding/hex"
	log "github.com/golang/glog"
	"net/http"
	"strconv"
	"time"
)

const ForgotPWDSession  = "ForgotPWDSession:"
// 玩家密码更改session
func (this *SessionData)Encode(s string) error {

	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err:= enc.Encode(this)
	b:=buffer.Bytes()

	err = Send("SET", ForgotPWDSession+s, string(b))
	if err != nil {
		return err
	}

	err = Send("EXPIRE", ForgotPWDSession+s, Expires)
	if err != nil {
		return err
	}
	return nil
}

func (this *SessionData)Clear(s string) error {
	return Send("DEL", ForgotPWDSession+s)
}

// 获取玩家密码更改session
func (this *SessionData)Decode(s string) error {
	b,err:=redis.String(Do("GET", ForgotPWDSession+s))

	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	buffer.Write([]byte(b))
	dec := gob.NewDecoder(&buffer)
	return dec.Decode(this)
}

// 用户web请求session数据

type SessionData struct {
	PngCode              string // 邮箱或手机验证修改用户密码图片验证码
	Email                string // 邮箱账户
	Phone                string // 用户绑定的手机号
	Userid               int    // 用户唯一id
	IdentifyCode         string // 6位修改密码验证码
	PasswordChangeStatus bool   // 验证码验证通过，等待用户更改密码
}

const (
	Expires = 60 * 30
)

type LoginSessionData struct {
	Userid int64
	//	UnixTime int64
}

// 玩家登陆redis里记录session
func (this *LoginSessionData) GetSession(s string) error {
	if s != "" {
		b, err := Do("GET", "UserSession:"+s)
		if err != nil {
			return err
		}
		bit,err:=redis.Bytes(b,err)
		if err != nil {
			return err
		}

		err = this.Decode(bit)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("session nill")
}

// 玩家登陆redis里记录session
func (this *LoginSessionData) SetSession(s string) error {
	if this.Userid > 0 {

		b, err := this.Encode()
		if err != nil {
			return err
		}
		err = Send("SET", "UserSession:"+s, b)
		if err != nil {
			return err
		}

		err = Send("EXPIRE", "UserSession:"+s, Expires)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("session nill")
}

// 加密session
func (this *LoginSessionData) Encode() ([]byte, error) {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(this)
	if err != nil {
		return []byte{}, err
	}
	return buffer.Bytes(), nil
}

// 解密session
func (this *LoginSessionData) Decode(b []byte) error {
	var buffer bytes.Buffer
	buffer.Write(b)
	dec := gob.NewDecoder(&buffer)
	return dec.Decode(this)
}

// set cookie一定要在内容返回之前设置，才有效
func SetCookie(w http.ResponseWriter, name string, userid int) *http.Cookie {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
		}
	}()

	h := md5.New()
	h.Write([]byte(strconv.Itoa(userid) + strconv.Itoa(int(time.Now().Unix()))))
	c := &http.Cookie{
		Name:  name,
		Value: hex.EncodeToString(h.Sum(nil)),
		Path:  "/",
	}

	http.SetCookie(w, c)
	return c
}

// 解压base64Cookie值成userid
func GetUseridByCookie(cookieValue string) (int, error) {
	if cookieValue == "" {
		return 0, errors.New("string is nill")
	}

	loginSession := &LoginSessionData{}
	err := loginSession.GetSession(cookieValue)

	if err != nil {
		return 0, err
	}

	return int(loginSession.Userid), nil
}

// 清除Cookie
func ClearCookie(w http.ResponseWriter, r *http.Request, name string) *http.Cookie {
	c := &http.Cookie{
		Name:  name,
		Value: "",
		Path:  "/",
	}
	http.SetCookie(w, c)
	return c
}

// redis里是否记录玩家登陆
func IsExitsUserSession(s string) (bool) {
	if s != ""{
		v, err := redis.Bool(Do("EXISTS", "UserSession:"+s))
		if err != nil {
			return false
		}
		return v
	}

	return false
}



// 延迟玩家登陆session过期时间
func AddUserSessionExpire(s string) error {
	if s != ""{
		return Send("EXPIRE", "UserSession:"+s, Expires)
	}
	return errors.New("session nill")
}







