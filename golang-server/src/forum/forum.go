/**
 * Created by Michael on 2015/8/5.
 *	论坛的接口
 *
 */
package forum

import (
	"db"
	log "github.com/golang/glog"
	"proxy"
	"socket"
	"strconv"
	"strings"
	"time"
	"utils"
	"vo"
	"redis"
	"csv"
)

func init() {
	proxy.Regist(8004, publicDialog)     // 发布对白
	proxy.Regist(8018, topstick)       // 顶贴
	proxy.Regist(8019, downstick)      // 踩贴
	proxy.Regist(8020, topstickFirst)   // 用户首次顶贴
	proxy.Regist(8021, downstickFirst) // 用户首次踩贴
	proxy.Regist(8030, getDialogs)     // 获取指定位置的前30条对白
	proxy.Regist(8031, getTopDialog)   //	获取指定位置的置顶对白
	proxy.Regist(8032, getUserDialogs) //	获取指定用户的前30条对白
	proxy.Regist(8034, delUserDialog)  //	删除指定用户的指定位置一条对白

	proxy.Regist(8036, publicComment)  //发布评论
	proxy.Regist(8038, getComment)  	//获取指定对白的前30条评论
}

func getComment(msg *[]byte, c *socket.Connection) {
	webctos := vo.CtoS8038Data1{}
	webctos.Decode(msg)
	stoc := vo.StoC8038Data1{T: webctos.T}

	if webctos.Dialogid < 0 || webctos.Page < 0{
		stoc.E = 13115
	}else{
		subdialogs := []*db.Tb_subdialog{}
		qs := db.OrmerSlave.QueryTable("Tb_subdialog")
		qs = qs.Filter("Parent_id",webctos.Dialogid)
		qs = qs.OrderBy("-Create_time")
		qs = qs.Limit(db.PAGE_NUM,webctos.Page*db.PAGE_NUM)	// 第一个参数表示获取几条数据，第二个参数表示起始位置
		_, err := qs.All(&subdialogs)
		if err != nil {
			stoc.E = 13115
		}else {
			log.Errorln(len(subdialogs))
			for i:=0;i<len(subdialogs);i++{
				data:=&vo.StoC8038Data2{}
				data.Anno = subdialogs[i].Is_anno
				data.Content = subdialogs[i].Dialog_title
				data.Userid = subdialogs[i].Account_num
				data.Time = subdialogs[i].Create_time
				data.Otherid =  subdialogs[i].Other_account
				user := &db.Tb_user_member{Account_num:subdialogs[i].Account_num}
				user.ReadNameAndHeadPic()
				data.Headpic =int( user.Create_time)
				data.Nickname = user.Nickname

				if data.Otherid> 0{
					user := &db.Tb_user_member{Account_num:data.Otherid}
					user.ReadNameAndHeadPic()
					data.Otherpic =int( user.Create_time)
					data.Othername = user.Nickname
				}
				stoc.Data = append(stoc.Data,data)
			}
		}
	}

	b, _ := stoc.Encode()
	c.Send <- *b
}

func publicComment(msg *[]byte, c *socket.Connection) {
	webctos := vo.CtoS8036Data1{}
	webctos.Decode(msg)
	stoc := vo.StoC8036Data1{T: webctos.T}

	if len(webctos.Content) > 0 ||  webctos.Dialogid >0{
		unixTime := time.Now().Unix()
		dialog := &db.Tb_subdialog{}
		dialog.Create_time = int(unixTime)
		dialog.Parent_id = webctos.Dialogid
		dialog.Account_num = c.UserData.Userid
		dialog.Is_anno = webctos.Anno
		dialog.Other_account = webctos.Otherid
		dialog.Dialog_title = webctos.Content
		_, err := dialog.Insert()
		if err != nil {
			stoc.E = 13114
		}
	}else{
		stoc.E = 13114
	}

	b, _ := stoc.Encode()
	c.Send <- *b
}

func publicDialog(msg *[]byte, c *socket.Connection) {
	webctos := vo.CtoS8004Data1{}
	webctos.Decode(msg)
	stoc := vo.StoC8004Data1{T: webctos.T}

	if webctos.Content == "" || webctos.Abspath == "" {
		stoc.E = 12009
	} else {
		var count = 0
		webctos.Content, count = utils.GetTrimStrLen(webctos.Content)
		if count > 23  {
			stoc.E = 12010
		} else {
			arr := strings.Split(webctos.Abspath, "_")
			worldid := arr[0]
			buildingid := arr[1]
			unixTime := time.Now().Unix()

			wid, _ := strconv.Atoi(worldid)
			bid, _ := strconv.Atoi(buildingid)

			dialog := &db.Tb_dialog{}
			dialog.Create_time = int(unixTime)
			dialog.World_id = wid
			dialog.Build_id = bid
			dialog.Abs_path = webctos.Abspath
			dialog.Account_num = c.UserData.Userid
			dialog.Dialog_type = 1
			dialog.Terminal_type = 1
			dialog.Dialog_title = webctos.Content
			lastid, err := dialog.Insert()
			if err != nil {
				log.Errorln(err)
				stoc.E = 12010
			}else{
				stat := &db.Tb_dialog_stat{Dialog_id: int(lastid)}
				stat.Create_time = int(unixTime)
				stat.Abs_path = webctos.Abspath
				stat.Dialog_type = dialog.Dialog_type
				stat.Insert()

				// 广播到该世界
				scto := vo.StoC5002Data1{T: 5002, Data: &vo.StoC5002Data2{}}
				scto.Data.Id = webctos.Abspath
				b, _ := scto.Encode()
				var m = vo.Broadcast{}
				m.Channel = wid
				m.Msg = *b
				socket.Hub.Broadcast <- m

				// 添加离线对白
				go redis.OfflineDlg.AddOfflineDialg(int(lastid),webctos.Abspath)
			}
		}
	}

	b, _ := stoc.Encode()
	c.Send <- *b
}

func topstick(msg *[]byte, c *socket.Connection) {
	webctos := vo.CtoS8018Data1{}
	webctos.Decode(msg)
	stoc := vo.StoC8018Data1{T: webctos.T}

	if (webctos.Dialogtype != 1 && webctos.Dialogtype != 2) || webctos.Abspath == ""{
		stoc.E = 15005
	}else{

		widgetData, ok := csv.Prop.Hash[webctos.Widgetid]
		if ok && widgetData.PropType == 10 {
			widgetCount, err := db.MasterDB.DelUserWidget(c.UserData.Userid, webctos.Widgetid, 1)
			if err != nil {
				stoc.E = 15002
			} else {
				err:=db.MasterDB.UseToptick(c.UserData.Userid, webctos.Dialogid, webctos.Dialogtype, webctos.Abspath, widgetData.Tag, 0)
				if err == nil{
					stoc.Data = & vo.StoC8018Data2{}
					stoc.Data.Widgetid = webctos.Widgetid
					stoc.Data.Count = widgetCount
					stoc.Data.Abspath = webctos.Abspath
					stoc.Data.Dialogid = webctos.Dialogid
					stoc.Data.Dialogtype = webctos.Dialogtype

					var topValue int
					var err1 interface{}
					if webctos.Dialogtype == 1{
						topValue, err1 = db.MasterDB.UpdateDialogTopstick(webctos.Dialogid, webctos.Abspath, widgetData.Tag)
					}else if webctos.Dialogtype == 2{
						topValue, err1 = db.MasterDB.UpdateSysDialogTopstick(webctos.Abspath, widgetData.Tag)
					}
					if err1 == nil {
						stoc.Data.Upworth = topValue
					}
				}else{
					stoc.E = 15002
				}
			}
		} else {
			stoc.E = 15003
		}
	}

	b, _ := stoc.Encode()
	c.Send <- *b
}

func downstick(msg *[]byte, c *socket.Connection) {
	webctos := vo.CtoS8019Data1{}
	webctos.Decode(msg)
	stoc := vo.StoC8019Data1{T: webctos.T}
	if (webctos.Dialogtype != 1 && webctos.Dialogtype != 2) || webctos.Abspath == ""{
		stoc.E = 15006
	}else{
		widgetData, ok := csv.Prop.Hash[webctos.Widgetid]
		if ok && widgetData.PropType == 11 {
			widgetCount, err := db.MasterDB.DelUserWidget(c.UserData.Userid, webctos.Widgetid, 1)
			if err != nil  {
				stoc.E = 15002
			} else {
				err:=db.MasterDB.UseToptick(c.UserData.Userid, webctos.Dialogid, webctos.Dialogtype, webctos.Abspath, 0, widgetData.Tag)
				if err == nil{

					stoc.Data = &vo.StoC8019Data2{}
					stoc.Data.Widgetid = webctos.Widgetid
					stoc.Data.Count = widgetCount

					stoc.Data.Abspath = webctos.Abspath
					stoc.Data.Dialogid = webctos.Dialogid
					stoc.Data.Dialogtype = webctos.Dialogtype
					var topValue int
					var err1 interface{}
					if webctos.Dialogtype == 1{
						topValue, err1 = db.MasterDB.UpdateDialogTopstick(webctos.Dialogid, webctos.Abspath, widgetData.Tag)
					}else if webctos.Dialogtype == 2{
						topValue, err1 = db.MasterDB.UpdateSysDialogTopstick(webctos.Abspath, widgetData.Tag)
					}
					if err1 == nil {
						stoc.Data.Upworth = topValue
					}
				}else {
					stoc.E = 15005
				}
			}
		} else {
			stoc.E = 15003
		}
	}

	b, _ := stoc.Encode()
	c.Send<- *b
}

func topstickFirst(msg *[]byte, c *socket.Connection) {
	webctos := vo.CtoS8020Data1{}
	webctos.Decode(msg)
	userid:= c.UserData.Userid
//	userid:= 60229
	stoc := vo.StoC8020Data1{T: webctos.T}
	if (webctos.Dialogtype != 1 && webctos.Dialogtype != 2) || webctos.Abspath == ""{
		stoc.E = 15005
	}else{
		//	对白位置:世界_栋_楼_房_情景
		boolean, _ := db.SlaveDB.IsUserUseToptick(userid, webctos.Dialogid, webctos.Dialogtype, webctos.Abspath)
		if boolean {
			stoc.E = 13215
		} else {
			err:=db.MasterDB.UseToptick(userid, webctos.Dialogid, webctos.Dialogtype, webctos.Abspath, 1, 0)
			if err == nil{
				stoc.Data = &vo.StoC8020Data2{}
				stoc.Data.Abspath = webctos.Abspath
				stoc.Data.Dialogid = webctos.Dialogid
				stoc.Data.Dialogtype = webctos.Dialogtype
				var topValue int
				var err1 interface{}
				if webctos.Dialogtype == 1{
					topValue, err1 = db.MasterDB.UpdateDialogTopstick(webctos.Dialogid, webctos.Abspath, 1)
				}else if webctos.Dialogtype == 2{
					topValue, err1 = db.MasterDB.UpdateSysDialogTopstick(webctos.Abspath, 1)
				}

				if err1 == nil {
					stoc.Data.Upworth = topValue
				}
			}else{
				stoc.E = 15005
			}

		}
	}


	b, _ := stoc.Encode()
//	log.Errorln(*b)
	c.Send <- *b
}
func downstickFirst(msg *[]byte, c *socket.Connection) {
	webctos := vo.CtoS8021Data1{}
	webctos.Decode(msg)
	stoc := vo.StoC8021Data1{T: webctos.T}
	if (webctos.Dialogtype != 1 && webctos.Dialogtype != 2) || webctos.Abspath == ""{
		stoc.E = 15006
	}else{
		boolean, _ := db.SlaveDB.IsUserUseDowntick(c.UserData.Userid, webctos.Dialogid, webctos.Dialogtype, webctos.Abspath)
		if boolean {
			stoc.E = 13216
		} else {
			err:=db.MasterDB.UseToptick(c.UserData.Userid, webctos.Dialogid, webctos.Dialogtype, webctos.Abspath, 0, -1)
			if err == nil{
				stoc.Data = &vo.StoC8021Data2{}
				stoc.Data.Abspath = webctos.Abspath
				stoc.Data.Dialogid = webctos.Dialogid
				stoc.Data.Dialogtype = webctos.Dialogtype
				var topValue int
				var err1 interface{}
				if webctos.Dialogtype == 1{
					topValue, err1 = db.MasterDB.UpdateDialogTopstick(webctos.Dialogid, webctos.Abspath, -1)
				}else if webctos.Dialogtype == 2{
					topValue, err1 = db.MasterDB.UpdateSysDialogTopstick(webctos.Abspath, -1)
				}


				if err1 == nil {
					stoc.Data.Upworth = topValue
				}
			}else{
				stoc.E = 15006
			}

		}
	}
	b, _ := stoc.Encode()
	c.Send <- *b
}
func getDialogs(msg *[]byte, c *socket.Connection) {
	webctos := vo.CtoS8030Data1{}
	webctos.Decode(msg)
	stoc := vo.StoC8030Data1{T: webctos.T}

	list, e := db.SlaveDB.GetDialogs(webctos.Abspath, webctos.Page)
	if e != nil || len(list) == 0 {
		stoc.E = 12001
	} else {

		for i := 0; i < len(list); i++ {
			d := &vo.StoC8030Data2{}
			d.Upworth = list[i].Upworth
			d.Abspath = list[i].Abspath
			d.Dialogid = list[i].Dialogid
			d.Dialogtype = list[i].Dialogtype
			d.Userid = list[i].Userid
			d.Content = list[i].Content
			d.Time = int(list[i].CreateTime)


			userdata:= db.Tb_user_member{Account_num:d.Userid}
			userdata.Read()

			d.Headpic = int(userdata.Create_time)
			d.Sex = userdata.Sex
			d.Provinceid = userdata.Province_id
			d.Cityid = userdata.City_id
			d.Nickname = userdata.Nickname

			stoc.Data = append(stoc.Data, d)
		}
	}
	b, _ := stoc.Encode()
	c.Send <- *b
}
func getTopDialog(msg *[]byte, c *socket.Connection) {
	webctos := vo.CtoS8031Data1{}
	webctos.Decode(msg)
	stoc := vo.StoC8031Data1{T: webctos.T}

	original, e := db.SlaveDB.GetSysDialogs(webctos.Abspath)
	if e == nil {
		data2 := &vo.StoC8031Data2{}
		data2.Abspath = original.Abspath
		data2.Dialogid = original.Dialogid
		data2.Dialogtype = original.Dialogtype
		data2.Upworth = original.Upworth
		data2.Content = original.Content
		data2.Time = int(original.CreateTime)
		stoc.Data = append(stoc.Data, data2)


		data, e := db.SlaveDB.GetTopDialog(webctos.Abspath)

		if e == nil {
			if data2.Upworth < data.Upworth {
				d := &vo.StoC8031Data2{}
				d.Abspath = data.Abspath
				d.Userid = data.Userid
				d.Dialogid = data.Dialogid
				d.Dialogtype = data.Dialogtype
				d.Upworth = data.Upworth
				d.Content = data.Content
				d.Time = int(data.CreateTime)

				userdata:= db.Tb_user_member{Account_num:d.Userid}
				userdata.Read()
				d.Headpic = int(userdata.Create_time)
				d.Sex = userdata.Sex
				d.Provinceid = userdata.Province_id
				d.Cityid = userdata.City_id
				d.Nickname = userdata.Nickname

				stoc.Data = append(stoc.Data, d)
			}
		}

	} else {
		stoc.E = 12002
	}

	b, _ := stoc.Encode()
	c.Send <- *b
}
func getUserDialogs(msg *[]byte, c *socket.Connection) {
	webctos := vo.CtoS8032Data1{}
	webctos.Decode(msg)
	stoc := vo.StoC8032Data1{T: webctos.T}

	list, e := db.SlaveDB.GetUserDialogs(c.UserData.Userid, webctos.Page)
//	list, e := db.SlaveDB.GetUserDialogs(60229, webctos.Page)
	if e != nil || len(list) == 0 {
		stoc.E = 12003
	} else {
		for i := 0; i < len(list); i++ {
			d := &vo.StoC8032Data2{}
			d.Upworth = list[i].Upworth
			d.Abspath = list[i].Abspath
			d.Content = list[i].Content
			d.Time = int(list[i].CreateTime)
			stoc.Data = append(stoc.Data, d)
			log.Errorln(*d)
		}
	}

	b, _ := stoc.Encode()
	c.Send <- *b
}

func delUserDialog(msg *[]byte, c *socket.Connection) {
	webctos := vo.CtoS8034Data1{}
	webctos.Decode(msg)
	stoc := vo.StoC8034Data1{T: webctos.T}

	if webctos.Abspath == "" || webctos.Dialogid <= 0 {
		stoc.E = 12004
	} else {
		qs := db.OrmerMaster.QueryTable("Tb_dialog")
		qs = qs.Filter("Dialog_id",webctos.Dialogid)
		qs = qs.Filter("Account_num",c.UserData.Userid)
		_,err:= qs.Delete()
		log.Errorln(err)
		if err != nil {
			stoc.E = 12004
		}else{
			qs := db.OrmerMaster.QueryTable("Tb_dialog_stat")
			qs = qs.Filter("Dialog_id",webctos.Dialogid)
			qs = qs.Filter("Abs_path",webctos.Abspath)
			qs.Delete()

			stoc.Data = &vo.StoC8034Data2{}
			stoc.Data.Abspath = webctos.Abspath
			stoc.Data.Dialogid = webctos.Dialogid
			stoc.Data.Dialogtype = webctos.Dialogtype
		}
	}

	b, _ := stoc.Encode()
	c.Send <- *b
}
