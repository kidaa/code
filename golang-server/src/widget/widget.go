/**
 * Created by Michael on 2015/7/29.
 * 一些比较简单的道具使用
 *
 *
 *
 */
package widget

import (
	"constant"
	"db"
	log "github.com/golang/glog"
	"proxy"
	"socket"
	"vo"
	"utils"
	"csv"
	"errors"
	"redis"
	"strconv"
)

func init() {
	proxy.Regist(6001, useBalloon)
	proxy.Regist(6023, useFireworks)
}

//  使用烟花道具
func useFireworks(msg *[]byte, c *socket.Connection) {
	ctos := vo.CtoS6023Data1{}
	ctos.Decode(msg)
	stoc := vo.StoC6023Data1{T: ctos.T}

	if ctos.Userid > 0 {
		if ctos.Widgetid == constant.Fireworks {
			// 判断对方是否在线
//			otherUser, ok := socket.Hub.Connections[ctos.Userid]
			if socket.Hub.Exists(ctos.Userid) {
				msg, strLen:= utils.GetTrimStrLen(ctos.Message)
				if strLen >0 && strLen <= 30{
					status := db.SlaveDB.GetBothFireworksStatus(c.UserData.Userid, ctos.Userid);
					if status == 0{
						_, err := UseWidget(c.UserData.Userid, ctos.Widgetid, 1);
						if err == nil {
							UpdateBothFireworksTime(c.UserData.Userid, ctos.Userid,ctos.Widgetid)
							send := vo.StoC6024Data1{T: 6024}
							send.Data = &vo.StoC6024Data2{Userid:c.UserData.Userid,Nickname:c.UserData.Nickname,Widgetid:ctos.Widgetid,Message:msg,Sex:c.UserData.Sex,Headpic:c.UserData.Headpic}
							stoc.Data = &vo.StoC6023Data2{Userid:c.UserData.Userid,Nickname:c.UserData.Nickname,Widgetid:ctos.Widgetid,Message:msg,Sex:c.UserData.Sex,Headpic:c.UserData.Headpic}
							b, err := send.Encode()
							if err != nil {
								log.Errorln(err)
							}
//							otherUser.Send <- *b
							m := vo.Broadcast{}
							m.Channel = ctos.Userid
							m.Msg = *b
							socket.Hub.Broadcast <-m

						} else {
							stoc.E = 15002 // 没有道具使用失败
						}
					}else{
						if status == 1{
							stoc.E = 15012//你正在欣赏烟花
						}else{
							stoc.E = 15014//对方正在和别人欣赏烟花
						}
					}
				}else{
					stoc.E = 15017//烟花消息长度错误
				}
			} else {
				stoc.E = 11005 // 对方不在线
			}
		} else {
			stoc.E = 15003 //  道具类型错误
		}
	} else {
		stoc.E = 11008 // 提交的用户ID为0
	}

	b, err := stoc.Encode()
	if err != nil {
		log.Errorln(err)
	}
	c.Send <- *b
}
// 广播气球消息
func useBalloon(msg *[]byte,c *socket.Connection){
	ctos:= vo.CtoS6001Data1{}
	ctos.Decode(msg)
	stoc:= vo.StoC6001Data1{T:6001}

	world:=c.UserData.WorldID

	var count = 0;
	ctos.Msg,count = utils.GetTrimStrLen(ctos.Msg)
	if count > 0 && count <= 20{
		if constant.Balloon1 == ctos.Id || constant.Balloon2 == ctos.Id {
			_,err:= UseWidget( c.UserData.Userid,ctos.Id,1)
			log.Infoln(err)
			if err == nil {
				var m = vo.Broadcast{}
				m.Channel = world
				stocbroadcard:= vo.StoC6002Data1{T:6002,Data:&vo.StoC6002Data2{}}
				stocbroadcard.Data.Msg = ctos.Msg
				stocbroadcard.Data.Id = ctos.Id
				stocbroadcard.Data.Nickname = c.UserData.Nickname
				b,_:= stocbroadcard.Encode()
				m.Msg = *b
				socket.Hub.Broadcast <- m
			}else{
				stoc.E = 15002
			}

		}else{
			stoc.E = 15001
		}
	}else{
		stoc.E = 15011
	}
	b,err:=stoc.Encode()
	if err != nil {
		log.Errorln(err)
	}
	c.Send<- *b
}

//使用道具,socket推送前端，更新道具数量
func UseWidget(userID int,widgetID int,count int) (num int, e interface{}){
	leftCount,err:=db.MasterDB.DelUserWidget(userID, widgetID, count)
	if err == nil {
		UserWidgetChange(userID,widgetID,leftCount);
	}
	return leftCount, err
}

//更新双方烟花结束时间
func UpdateBothFireworksTime(userID int,otherID int, widgetID int) (e interface{}){
	data, ok := csv.Prop.Hash[widgetID]
	if ok {
//		db.MasterDB.UpdateBothFireworksTime(userID, otherID,data.Tag);
		redis.SetTmpExpires("Firework"+strconv.Itoa(userID),data.Tag)
		redis.SetTmpExpires("Firework"+strconv.Itoa(otherID),data.Tag)
		return nil
	}else {
		return errors.New("无该道具持续时间")
	}
}

//服务器改变用户道具数量推送
func UserWidgetChange(userID int,widgetID int,count int) (e interface{}){
	del := vo.StoC2005Data1{T: 2005,Data:&vo.StoC2005Data2{}}
	del.Data.Id = widgetID
	del.Data.Count =count
	b,err:=del.Encode()
	if err != nil {
		log.Error(err)
		return err
	}
//	targetConn, ok := socket.Hub.Connections[userID]
	if socket.Hub.Exists(userID) {

		m := vo.Broadcast{}
		m.Channel = userID
		m.Msg = *b
		socket.Hub.Broadcast <-m
//		targetConn.Send <- *b
	}
	return nil
}

