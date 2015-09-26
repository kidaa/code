/*
*
*socket 协议处理
*
 */

package proxy

import (
	log "github.com/golang/glog"
	"socket"
)

var hash map[int]func(msg *[]byte, c *socket.Connection)


func init() {
	hash = make(map[int]func(msg *[]byte, c *socket.Connection))
	socket.Hub.Handle = ProxyHandle
}

func Regist(key int, fun func(msg *[]byte, c *socket.Connection)) {
	hash[key] = fun
}

// 消息分发
func ProxyHandle(t int, message *[]byte, c *socket.Connection) {
	defer func() {
//		close(c.Send)
		if err := recover(); err != nil {
//			close(c.Send)
			c.Close()
			log.Errorln("proxy", err)
		}
	}()

	fun, ok := hash[t]
	if ok {
		// 没有登陆，直接断开
		if t == 1000{
			fun(message, c)
		}else{
			if c.IsLogined == false{
//				close(c.Send)
				c.Close()
			}else{
				fun(message, c)
			}
		}
	} else {
//		close(c.Send)
		c.Close()
	}
}
