/*
*
*Adobe Flash套接字安全策略端口监听
*
*/


package policy

import(
    "bufio"  
    "net"  
    "time"  
	log"github.com/golang/glog"
)

var (  
    ClientMap map[int]net.Conn = make(map[int]net.Conn)  
)  
  
func Create(port string) {   
	defer func(){
		if e:=recover();e != nil{
			log.Errorln(e)
		}	
	}()
	 listener, err := net.Listen("tcp", ":" +port)  
    checkError(err)  
    clientIndex := 0  
  
    for {  
        clientIndex++  
        conn, err := listener.Accept()  
        if err != nil {  
            continue  
        }  
        go handleClient(conn, clientIndex)  
    }  	
}  
  
func handleClient(conn net.Conn, index int) {  
	defer func(){
		if e:=recover();e != nil{
			log.Errorln(e)
		}	
	}()
    ClientMap[index] = conn  
  
   Log("come from: ", conn.RemoteAddr(), "index: ", index)  
    fc := func() {  
        time.Sleep(time.Second) //给客户端1秒的响应的时间，否则客户端有可能读不到数据就提前Close了  
        conn.Close()  
        delete(ClientMap, index)  
       
    }  
    defer fc()  
    sendFirstMsg(conn)  
}  
func sendFirstMsg(conn net.Conn) {   
		 str :=	`<cross-domain-policy>
				<allow-access-from domain="*"to-ports="*"/>
				</cross-domain-policy>`
    writer := bufio.NewWriter(conn)  
    writer.WriteString(str)  
    writer.Flush()  
    
}  


func checkError(err error) {  
    if err != nil {  
        Log(err.Error())   
    }  
}  

func Log(v ... interface{}){
	log.Errorln(v)
}
