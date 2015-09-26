/**
 * Created by Michael on 2015/7/29.
 */
package backpack
import (
	"socket"
	"proxy"
	"db"
	"vo"
	"widget"
)
func init(){
	proxy.Regist(2000,widgetList)
	proxy.Regist(2001,delWidget)
	proxy.Regist(2002,addWidget)
	proxy.Regist(2003,useWidget)
	proxy.Regist(2004,getCellCount)
}
func getCellCount(msg *[]byte,c *socket.Connection){

}

// 获取用户背包里的所以道具
func widgetList(msg *[]byte,c *socket.Connection){
	stoc:= vo.StoC2000Data1{T:2000}
	defer func() {
		if err := recover(); err != nil {

		}
		bit, _ := stoc.Encode()
		c.Send <- *bit
	}()

	userid:= c.UserData.Userid
	list:=db.SlaveDB.GetUserWidget(userid)

	for i:=0;i<len(list);i++{
		soC2000Data2:= vo.StoC2000Data2{}
		soC2000Data2.Count = list[i].Count
		soC2000Data2.Id = list[i].Id
		stoc.Data = append(stoc.Data,&soC2000Data2)
	}
}

func delWidget(msg *[]byte,c *socket.Connection){

}

func addWidget(msg *[]byte,c *socket.Connection){

}

func useWidget(msg *[]byte,c *socket.Connection){
	stoc := vo.StoC2003Data1{T:2003}
	defer func() {
		if err := recover(); err != nil {
			stoc.E = 15012
		}
		bit, _ := stoc.Encode()
		c.Send <- *bit
	}()

	ctos:= vo.CtoS2003Data1{}
	ctos.Decode(msg)
	userid:= c.UserData.Userid
	widget.UseWidget(userid,ctos.Id,ctos.Count)

}
