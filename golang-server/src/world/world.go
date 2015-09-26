/**
 * Created by Michael on 2015/8/5.
 */
package world
import (
	"socket"
	"proxy"
	"vo"
	"db"
	"csv"
	"strings"
	"strconv"
	"redis"
)
func init() {
	proxy.Regist(5005, changeMapWorld)
	proxy.Regist(5006, throughWorld)
	proxy.Regist(5007, getWorldList)

	proxy.Regist(5010, getOfflineDialog)


}

// 获取离线对白列表
func getOfflineDialog(msg *[]byte, c *socket.Connection) {
	stoc := vo.StoC5010Data1{T:5010}
	user := db.Tb_user_active{Account_num:c.UserData.Userid}
	user.GetOfflineTime()

	offlineTime := user.Logout_time
	list,err:=redis.OfflineDlg.GetOfflineDialg(offlineTime)
	if err != nil{
		stoc.E = 15027
	}else{
		stoc.Data = list
	}

	b, _ := stoc.Encode()
	c.Send <- *b
}

// 地图切换世界
func changeMapWorld(msg *[]byte, c *socket.Connection) {
	ctos := vo.CtoS5005Data1{}
	ctos.Decode(msg)

	stoc := vo.StoC5005Data1{T:ctos.T}
	if ctos.Worldid <= 0 {
		stoc.E = 15013
	}else if (ctos.Worldid ==c.UserData.WorldID) {
		stoc.E = 15018
	}else{
		userWorld, err:= db.SlaveDB.CheckUserWorld(c.UserData.Userid,ctos.Worldid)
		if err != nil || userWorld == 0{
			stoc.E = 15024
		}else{
			stoc.E = ChangeWorldCurrently(c.UserData.Userid,ctos.Worldid)
			if stoc.E == 0{
				stoc.Data = ctos.Worldid
				c.UserData.WorldID = ctos.Worldid
			}
		}
	}
	b, _ := stoc.Encode()
	c.Send <- *b
}

// 穿越世界
func throughWorld(msg *[]byte, c *socket.Connection) {
	ctos := vo.CtoS5006Data1{}
	ctos.Decode(msg)

	stoc := vo.StoC5006Data1{T:ctos.T}
	if ctos.Worldid <= 0 {
		stoc.E = 15013
	}else if (ctos.Worldid ==c.UserData.WorldID) {
		stoc.E = 15018
	}else {
		_, err:= db.SlaveDB.CheckUserWorld(c.UserData.Userid,ctos.Worldid)
		if err != nil{
			worldData,ok := csv.World.Hash[ctos.Worldid]
			if ok && worldData.Status == 1{
				woldVO := &vo.UserWorld{}
				woldVO.WorldID = worldData.Id
				woldVO.AccountNum = c.UserData.Userid
				buildArr := strings.Split(worldData.Builds, ",")
				if(len(buildArr) > 0){
					buildPartCloud := make([]*vo.BuildPartCloud, 0, 10)
					for _,value := range buildArr{
						buildid,_ := strconv.Atoi(value)
						data, ok := csv.Build.Hash[buildid]
						if ok {
							buildVO := &vo.BuildPartCloud{}
							buildVO.Buildid = data.Id
							buildVO.AccountNum = c.UserData.Userid
							buildVO.Part = data.Part
							buildVO.Cloud = data.Cloud
							buildPartCloud = append(buildPartCloud, buildVO)
						}

					}
					err := db.MasterDB.OpenUserWorld(woldVO,buildPartCloud)
					if err != nil {
						stoc.E = 15013
					}else{
						stoc.E = ChangeWorldCurrently(c.UserData.Userid,ctos.Worldid)
						if stoc.E == 0{
							stoc.Data = ctos.Worldid
							c.UserData.WorldID = ctos.Worldid
						}
					}
				}else{
					stoc.E = 15024
				}
			}else{
				stoc.E = 15024
			}

		}else{
			stoc.E = 15025
		}
	}
	b, _ := stoc.Encode()
	c.Send <- *b
}

func ChangeWorldCurrently(account_num int,worldid int) (num int){
	errNum := 0
	world:=db.Tb_world_currently{Account_num:account_num}
	err:= world.GetWorldid()
	if (err == nil) {
		if (worldid == world.World_id) {
			errNum = 15018
		}else {
			world.Account_num = account_num
			world.World_id = worldid
			err:= world.UpdateWorldID()
			if err != nil {
				errNum = 15013
			}
		}
	}else {
		errNum = 15013
	}
	return errNum
}

// 获取地图列表
func getWorldList(msg *[]byte, c *socket.Connection) {
	ctos := vo.CtoS5007Data1{}
	ctos.Decode(msg)

	stoc := vo.StoC5007Data1{T:ctos.T}
	list,err := db.SlaveDB.GetWorldList(c.UserData.Userid)
	if err == nil{
		stoc.Data = make([]*vo.StoC5007Data2,0,len(list))
		for _,value:= range list{
			worldVO := &vo.StoC5007Data2{}
			worldVO.Worldid = value.WorldID
			if value.Currently > 0{
				worldVO.Currently = 1
			}else{
				worldVO.Currently = 0
			}
			stoc.Data = append(stoc.Data,worldVO)
		}
	}else{
		stoc.E = 15026
	}
	b, _ := stoc.Encode()
	c.Send <- *b
}
