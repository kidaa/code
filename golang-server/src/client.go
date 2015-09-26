package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"flag"
	log "fmt"
	"github.com/Unknwon/goconfig"
	"github.com/gorilla/websocket"
	"io"
	"os"
	"regexp"
	"socket"
	"strconv"
	"strings"
	"time"
	"vo"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var cstDialer = websocket.Dialer{
	Subprotocols:    []string{"p1", "p2"},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var c Connection
var config *goconfig.ConfigFile

func main() {
	flag.Parse()

	var err error
	config, err = goconfig.LoadConfigFile("./client.ini")
	if err != nil {
		log.Println(err)
		<-time.After(time.Second * 2)
		return
	}
	go connectAndLogin()

	<-time.After(time.Hour)
}

func connectAndLogin() {

	ipaddr, err := config.GetValue("user", "ip")
	ws, _, err := cstDialer.Dial(ipaddr, nil)

	if err != nil {
		log.Println(err)
		<-time.After(time.Second * 2)
		return
	}

	username, err := config.GetValue("user", "username")
	userData := vo.UserData{}
	id, err := strconv.Atoi(username)
	userData.Userid = id
	c = Connection{Send: make(chan []byte, 256), Ws: ws, UserData: userData, Connected: true}

	defer func() {
		if err := recover(); err != nil {

		}
	}()
	defer func() { c.Unregister() }()

	ctos := vo.CtoS1000Data1{T: 1000}
	ctos.Userid = userData.Userid
	ctos.Worldid = 1
	password, err := config.GetValue("user", "password")
	h := md5.New()
	h.Write([]byte(password))                                                     // 需要加密的字符串为 123456
	log.Println("user: ", userData.Userid, "pwd: ", password, "remote: ", ipaddr) // 输出加密结果

	ctos.Pwd = hex.EncodeToString(h.Sum(nil))
	b, _ := ctos.Encode()

	go stream_copy(os.Stdin)
	go c.writePump()
	c.Send <- *b

	c.readPump()
}

func stream_copy(src io.Reader) {
	buf := make([]byte, 1024)
	go func() {
		defer func() {
		}()
		for {
			var nBytes int
			var err error
			nBytes, err = src.Read(buf)

			if err != nil {
				if err != io.EOF {
					log.Printf("Read error: %s\n", err)
				}
				break
			}

			rel := strings.Replace(string(buf[0:nBytes]), "\n", "", -1)

			reg := regexp.MustCompile(`login`)
			s := reg.Find([]byte(rel))
			reg = regexp.MustCompile(`close`)
			clo := reg.Find([]byte(rel))

			reg = regexp.MustCompile(`"t":1000`)
			l := reg.Find([]byte(rel))

			if c.Connected {
				if len(clo) > 0 {
					c.Unregister()
					continue
				}
			}
			if len(s) > 0 {
				connectAndLogin()

			} else if len(l) > 0 {
				if c.Connected {

					m := make(map[string]interface{})
					err := json.Unmarshal([]byte(rel), &m)

					log.Println(" len(l)", len(l))

					if err != nil {
						log.Println(":input json error")
					} else if m["t"] == nil {
						log.Println("t is nil")
					} else {

						ctos := vo.CtoS1000Data1{T: 1000}

						h := md5.New()
						h.Write([]byte(m["pwd"].(string))) // 需要加密的字符串为 123456

						ctos.Pwd = hex.EncodeToString(h.Sum(nil))
						ctos.Userid = m["userid"].(int)
						b, _ := ctos.Encode()

						log.Println(string(*b))
						c.Send <- *b
					}
				} else {
					log.Println("socket closed input <login> to relogin")
				}

			} else {
				if c.Connected {

					m := make(map[string]interface{})
					err := json.Unmarshal([]byte(rel), &m)

					if err != nil {
						log.Println(":input json error")
					} else if m["t"] == nil {
						log.Println("t is nil")
					} else {

//						for i := 0; i < 10000; i++ {
							c.Send <- []byte(rel)
//						}
					}

				} else {
					log.Println("socket closed input <login> to relogin")
				}
			}
		}
	}()
}

type Connection struct {
	Ws        *websocket.Conn
	Send      chan []byte
	Connected bool
	UserData  vo.UserData
}

func (c *Connection) Unregister() {
	defer func() {

		if e := recover(); e != nil {
			log.Println("Unregister", e)
		}
	}()

	if c.Connected {
		close(c.Send)
		c.Ws.Close()
	}

	c.Connected = false
}

func (c *Connection) write(mt int, payload []byte) error {
	c.Ws.SetWriteDeadline(time.Now().Add(writeWait))

	payload = socket.Packet(payload)

	return c.Ws.WriteMessage(mt, payload)
}

func (c *Connection) reader(readerChannel chan []byte) {
	for {
		select {
		// 如果管道关闭则退出for循环，因为管道关闭不会阻塞导致for进入死循环
		case message, ok := <-readerChannel:
			if !ok {
				break
			}

			// 获取协议号
			reg := regexp.MustCompile(`"t":[\d]+`)
			s := reg.Find(message)
			reg = regexp.MustCompile(`[\d]+`)
			i, err := strconv.Atoi(string(reg.Find(s)))

			if err == nil {
				go ProxyHandle(i, &message, c)
			} else {
				log.Println("解包出错")
				break
			}

		}
	}
}

func (c *Connection) readPump() {
	defer func() {
		c.Unregister()
		if e := recover(); e != nil {
			log.Println("socket closed on read input <login> to relogin")
		}
	}()
	c.Ws.SetReadLimit(maxMessageSize)
	c.Ws.SetReadDeadline(time.Now().Add(pongWait))
	c.Ws.SetPongHandler(func(string) error { c.Ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	//声明一个临时缓冲区，用来存储被截断的数据
	tmpBuffer := make([]byte, 0, socket.HeaderLen+1)

	//声明一个管道用于接收解包的数据
	readerChannel := make(chan []byte, socket.HeaderLen+1)
	go c.reader(readerChannel)

	for {
		_, message, err := c.Ws.ReadMessage()
		if err != nil {
			break
		}

		var e interface{}
		tmpBuffer ,e= socket.Unpack(append(tmpBuffer, message...), readerChannel)
		if e != nil{
			return
		}
		// 获取协议号
		/*	reg := regexp.MustCompile(`"t":[\d]+`)
			s := reg.Find(message)
			reg = regexp.MustCompile(`[\d]+`)
			i,err:= strconv.Atoi(string(reg.Find(s)))

			if err == nil {
				go ProxyHandle(i,&message,c)
			}else{
				log.Println("解包出错")
				break
			}*/
	}
}

// 消息分发
func ProxyHandle(t int, message *[]byte, c *Connection) {
	defer func() {

		if err := recover(); err != nil {
			c.Unregister()
			log.Println("proxy", err)
		}
	}()
	log.Println("Recive Data-> ", string(*message))
}

func (c *Connection) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		c.Unregister()
		ticker.Stop()
		log.Println("socket closed on write input <login> to relogin")
	}()

	for {
		select {
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
