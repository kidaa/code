/*
*
*负责用户登入登出和用户间广播数据
*
 */

package socket

import (
	log "github.com/golang/glog"
	"vo"
)

// 判断玩家是否在线
func (self *hub) Exists(key int) bool {
	_, ok := self.connections[key]
	return ok
}

// 判断玩家是否重复登陆
func (self *hub) IsRelogin(key int, session string) bool {
	_, ok := self.connections[key]
	if ok {
		if self.connections[key].UserData.Session == session {
			return true
		}
	}
	return false
}

type hub struct {
	connections map[int]*Connection
	Broadcast   chan vo.Broadcast
	Register    chan *Connection
	Handle      func(i int, msg *[]byte, c *Connection)
	Unregister  chan *Connection
}

var Hub = hub{
	Broadcast:   make(chan vo.Broadcast),
	connections: make(map[int]*Connection),
	Register:    make(chan *Connection),
	Unregister:  make(chan *Connection),
}

func (h *hub) Run() {
	defer func() {
		if e := recover(); e != nil {
			log.Errorln(e)
		}
	}()

	for {
		select {
		case c := <-h.Register:
			h.connections[c.UserData.Userid] = c
			c.noLoginTimeout <- []byte{}
		case c := <-h.Unregister:
			delete(h.connections, c.UserData.Userid)
		case m := <-h.Broadcast:
			// 全服广播
			if m.Channel == -1 {
				for _, c := range h.connections {
					select {
					case c.Send <- []byte(m.Msg):
					default:
					}
				}
				// 指定世界广播
			} else if m.Channel > 0 && m.Channel < 10000 {
				for _, c := range h.connections {
					if c.UserData.WorldID == m.Channel {
						select {
						case c.Send <- m.Msg:
						default:
						}
					}
				}
			} else {
				// 给其他用户发消息
				if m.Kick {
					if con, ok := h.connections[m.Channel]; ok {
						con.Close()
					}
				} else {
					if con, ok := h.connections[m.Channel]; ok {
						select {
						case con.Send <- m.Msg:
						default:
						}
					}
				}
			}
		}
	}
}
