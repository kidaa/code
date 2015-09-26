/**
 * Created by Michael on 2015/7/30.
 */
package proxy
import (
	"net/http"
	"redis"
)
var hash map[int]func(uid int,body * []byte)* []byte
var hashPrivilege map[int]func( body * []byte, r *http.Request,w http.ResponseWriter)* []byte


func init(){
	hash = make(map[int]func(uid int,body * []byte)* []byte)
	hashPrivilege = make(map[int]func( body * []byte, r *http.Request,w http.ResponseWriter)* []byte)
}

// 普通协议请注册本函数
func Regist(t int,fun func(uid int,body * []byte)* []byte){
	hash[t] = fun
}

// 提供给特权协议
func PrivilegeRegist(t int,fun func(body * []byte, r *http.Request,w http.ResponseWriter)* []byte){
	hashPrivilege[t] = fun
}

// t:协议号
func Run(t int,body * []byte, r *http.Request,w http.ResponseWriter){
	fun,ok:=hash[t]

	if ok{

		cookie, err := r.Cookie("login")
		if err != nil{
			w.Write([]byte("session time out"))
			return
		}

		var userid int
		if cookie != nil && cookie.Value != ""{
			userid,err=redis.GetUseridByCookie(cookie.Value)
			if err != nil{
				w.Write([]byte("session time out"))
				return
			}
			go redis.AddUserSessionExpire(cookie.Value)
		}

		result:= fun(userid,body)
		if result != nil && len(*result) > 0{

			w.Write(* result)
		}else{
			w.Write([]byte("request error"))
		}
	}else{

		f,ok:=hashPrivilege[t]
		if ok{
			result:= f(body,r,w)
			if result != nil && len(*result) > 0{
				w.Write(* result)
			}else{
				w.Write([]byte("request error"))
			}
		}else{
			w.Write([]byte("proto code error"))
		}
	}
}