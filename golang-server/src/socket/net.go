/*
*
*socket 通道读写数据，断线处理
*
 */

package socket

import (
	"regexp"

	"db"
	log "github.com/golang/glog"
	"github.com/gorilla/websocket"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"time"
	"utils"
	"vo"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
	//连接建立后五秒内没有收到登陆请求，断开socket
	waitForLogin = time.Second * 5
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Connection struct {
	Send           chan []byte
	UserData       vo.UserData
	IsLogined      bool
	ws             *websocket.Conn
	readerChannel  chan []byte
	noLoginTimeout chan []byte
}

func (c *Connection)Close() {
	c.ws.Close()
}

func (c *Connection) reader(readerChannel chan []byte) {
	defer func() {
		c.ws.Close()
		log.Errorln("reader")
		if err := recover(); err != nil {
			log.Errorln(err)
		}
	}()
	for {
		select {
		// 如果管道关闭则退出for循环，因为管道关闭不会阻塞导致for进入死循环
		case message, ok := <-readerChannel:
			if !ok {
				return
			}
			// 获取协议号
			reg := regexp.MustCompile(`"t":[\d]+`)
			s := reg.Find(message)
			reg = regexp.MustCompile(`[\d]+`)
			i, err := strconv.Atoi(string(reg.Find(s)))

			if err == nil {
				go Hub.Handle(i, &message, c)
			} else {
//				close(c.Send)
				return
			}
		}
	}
}

func (c *Connection) readPump() {
	defer func() {
		c.ws.Close()
		close(c.Send)
		close(c.readerChannel)
		log.Errorln("readPump")
		if err := recover(); err != nil {
			log.Errorln(err)
		}
	}()
	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error { c.ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	//声明一个临时缓冲区，用来存储被截断的数据
	tmpBuffer := make([]byte, 0, HeaderLen+1)
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			return
		}
		var erro interface{}
		tmpBuffer, erro = Unpack(append(tmpBuffer, message...), c.readerChannel)
		if erro != nil {
			log.Errorln(erro)
			return
		}
	}
}

func (c *Connection) write(mt int, payload []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	payload = Packet(payload)
	return c.ws.WriteMessage(mt, payload)
}

func (c *Connection) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
		log.Errorln("writePump")
		if err := recover(); err != nil {
			log.Errorln(err)
		}
	}()

	for {
		select {
		// 如果管道关闭则退出for循环，因为管道关闭不会阻塞导致for进入死循环
		case message, ok := <-c.Send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (c *Connection) loginTimeout() {
	defer func() { close(c.noLoginTimeout) }()
	for {
		select {
		case <-c.noLoginTimeout:
			return
		case <-time.After(waitForLogin):
			return
		}
	}
}

// 保存退出时间，用于在线时长计算
func saveLogoutUinxTime(userid int) {
	if userid > 0 {
		tb := db.Tb_user_active{Account_num: userid}
		tb.Is_online = 0
		tb.Logout()
	}
}

// websocket 连接请求
func WSHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
		}
	}()
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	socket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorln(err)
		return
	}
	ipaddress := utils.InetToaton(strings.Split(r.RemoteAddr, ":")[0])
	log.Infoln("Remote IP: ", strings.Split(r.RemoteAddr, ":")[0], ipaddress)
	c := &Connection{Send: make(chan []byte, 256), ws: socket, noLoginTimeout: make(chan []byte), readerChannel: make(chan []byte, HeaderLen+1)}
	defer saveLogoutUinxTime(c.UserData.Userid)
	defer func(userid int) { log.Errorln("用户退出", runtime.NumGoroutine(), userid) }(len(Hub.connections))
	go c.reader(c.readerChannel)
	go c.loginTimeout()
	go c.writePump()
	defer func() { Hub.Unregister <- c }()
	c.readPump()
}
