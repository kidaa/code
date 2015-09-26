
/**
 * Created by Michael on 2015/8/5.
 */
package rpc
/*

import (
	log"github.com/golang/glog"
	"web/webvo"
	"strings"
	"socket"
	"strconv"
	"vo"
)


type BroadcastPublic int
// #desc广播当前世界所有人新帖消息（冒星星）id:大楼中位置ID
func (this * BroadcastPublic)Broadcast(args *webvo.WebCtoS20004Data1, reply *([]string)) error {
	defer func() {
		if err:=recover();err!=nil{
			log.Errorln(err)
		}
	}()

	*reply = append(*reply, "test")
	log.Infoln(*args)

	arr:=strings.Split(args.Id,"_")
	worldid:= arr[0]

	wid,err:= strconv.Atoi(worldid)
	if err != nil{
		log.Errorln(err)
	}

	scto:= vo.StoC5002Data1{T:5002,Data:&vo.StoC5002Data2{}}
	scto.Data.Id = args.Id
	b,err:= scto.Encode()

	var m = vo.Broadcast{}
	m.Channel = wid
	m.Msg = *b

	socket.Hub.Broadcast <- m

	return nil
}
*/



