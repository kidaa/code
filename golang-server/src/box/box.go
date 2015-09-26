package box
import (
	"socket"
	"proxy"
	"vo"
	"db"
	"csv"
	"strings"
	"math/rand"
	"strconv"
	"widget"
)
func init() {
	proxy.Regist(4000, getWorldBoxNum)
	proxy.Regist(4001, openBox)

}

// 获取世界宝箱数量
func getWorldBoxNum(msg *[]byte, c *socket.Connection) {
	ctos := vo.CtoS4000Data1{}
	ctos.Decode(msg)

	stoc := vo.StoC4000Data1{T:ctos.T}
	total, err := db.SlaveDB.GetWorldBoxNum(c.UserData.Userid, c.UserData.WorldID, ctos.Boxid)
	if err != nil {
		stoc.E = 18002
	}else {
		stoc.Data = total
	}
	b, _ := stoc.Encode()
	c.Send <- *b
}

// 找宝获得物品
func openBox(msg *[]byte, c *socket.Connection) {
	ctos := vo.CtoS4001Data1{}
	ctos.Decode(msg)

	stoc := vo.StoC4001Data1{T:ctos.T}
	if ctos.Boxid <= 0 {
		stoc.E = 18003
	}else {
		count, err := db.SlaveDB.GetWorldBoxNum(c.UserData.Userid, ctos.Boxid, c.UserData.WorldID)
		if err != nil || count <= 0 {
			stoc.E = 18004
		}else {
			oddsVO, ok := csv.BoxOdds.Hash[ctos.Boxid]
			if ok {
				var oddsArr, thingArr []string
				var thingID, thingType int
				oddsNum := rand.Intn(100) // 随机确认获得的是道具还是实物
				if oddsNum < oddsVO.WidgetOdds {//获得道具
					thingArr = strings.Split(oddsVO.Widget, ",")
					oddsArr = strings.Split(oddsVO.WidgetEach, ",")
					thingType = 1
				}else {// 获得实物
					thingArr = strings.Split(oddsVO.Goods, ",")
					oddsArr = strings.Split(oddsVO.GoodsEach, ",")
					thingType = 2
				}
				eachOdds := 0
				oddsNum = rand.Intn(100) // 随机获得thingArr中的物品
				for i := 0; i<len(oddsArr); i++ {
					eachNum, _ := strconv.Atoi(oddsArr[i])
					eachOdds += eachNum
					if eachOdds > oddsNum {
						eachNum, _ := strconv.Atoi(thingArr[i])
						thingID = eachNum
						break
					}
				}
				if thingID <=0 {
					stoc.E = 18003
				}else {
					if thingType == 1 {
						total, err := db.MasterDB.UpdateWorldBoxNum(c.UserData.Userid, c.UserData.WorldID, ctos.Boxid, count - 1)
						if err != nil {
							stoc.E = 18003
						}else {
							count, err := db.MasterDB.AddUserWidget(c.UserData.Userid, thingID, 1)
							if err != nil {
								stoc.E = 18003
							}else {
								stoc.Data = &vo.StoC4001Data2{Id:thingID, Num:1, Genre:thingType, Count:1, Boxtotal:total}
								widget.UserWidgetChange(c.UserData.Userid,thingID,count)
							}
						}
					}else {
						//todo:实物奖品
					}
				}
			}else {
				stoc.E = 18003
			}
		}
	}
	b, _ := stoc.Encode()
	c.Send <- *b
}