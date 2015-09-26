package widget
import (
	"socket"
	"proxy"
	"vo"
	log"github.com/golang/glog"
	"constant"
	"db"
	"csv"
	"errors"
	"math/rand"
)
func init() {
	proxy.Regist(6010, getUserBuff)
	proxy.Regist(6011, gainCoin)
	proxy.Regist(6012, useDefense)
	proxy.Regist(6016, attackNum)
	proxy.Regist(6017, attackList)
}

func getUserBuff(msg *[]byte, c *socket.Connection) {
	ctos := vo.CtoS6010Data1{}
	ctos.Decode(msg)
	stoc := vo.StoC6010Data1{T:ctos.T}
	db.MasterDB.DelTimeoutBuff(c.UserData.Userid)
	data, err :=db.SlaveDB.GetUserBuffList(c.UserData.Userid)
	if err == nil {
		stoc.Data = data
	}else{
		stoc.E = 15009
	}
	log.Infoln(ctos.T)

	b, err := stoc.Encode()
	if err != nil {
		log.Error(err)
	}
	c.Send <- *b

}

// 道具获取积分
func gainCoin(msg *[]byte, c *socket.Connection) {
	ctos := vo.CtoS6011Data1{}
	ctos.Decode(msg)
	stoc := vo.StoC6011Data1{T:ctos.T,Data:&vo.StoC6011Data2{}}
	stoc.Data.Userid = ctos.Userid
	stoc.Data.Widgetid = ctos.Widgetid


	if ctos.Userid != c.UserData.Userid {
		switch ctos.Widgetid {
		case constant.GainCoin1:
			fallthrough
		case constant.GainCoin2:
			fallthrough
		case constant.GainCoin3:
			fallthrough
		case constant.GainCoin4:
			fallthrough
		case constant.GainCoin5:
			score := &db.Tb_score_stat{Account_num: ctos.Userid}
			err:=score.GetCoin()
			if err == nil {
				if score.Score_total > 200 {
					defstuckID, _ := CheckDefense(ctos.Userid)
					_, err := UseWidget(c.UserData.Userid, ctos.Widgetid, 1)
					if err == nil {
						ok := getDefenseStatus(ctos.Widgetid, defstuckID)
						if ok {
							changeCoin, err := getGainCoinNum(ctos.Widgetid)
							if err == nil {
								lessenCoin, err := AddOrLessenCoin(ctos.Userid, changeCoin, 50, 2)
								if err == nil {
									send6014Msg(&c.UserData, ctos.Userid, ctos.Widgetid,defstuckID, lessenCoin)
									addCoin, err := AddOrLessenCoin(c.UserData.Userid, lessenCoin, 31, 1)
									if err == nil {
										stoc.Data.Coin = int(addCoin)
									}else {
										stoc.E = 15004
									}
								}else {
									stoc.E = 15004
								}
							}else {
								stoc.E = 15002
							}
						}else {
							send6013Msg(&c.UserData, ctos.Userid, ctos.Widgetid, defstuckID)
							stoc.Data.Success = 1
							stoc.Data.Defensedid = defstuckID
						}
					}else {
						stoc.E = 15002
					}
				}else {
					stoc.Data.Success = 2
				}
			}else {
				stoc.E = 15008
			}
		default:
			stoc.E = 15003
		}
	}else{
		stoc.E = 15012
	}
	log.Infoln(ctos.Userid, ctos.T, ctos.Widgetid)

	b, err := stoc.Encode()
	if err != nil {
		log.Error(err)
	}
	c.Send <- *b
}


// 使用防御道具
func useDefense(msg *[]byte, c *socket.Connection) {
	ctos := vo.CtoS6012Data1{}
	ctos.Decode(msg)
	stoc := vo.StoC6012Data1{T:ctos.T}
	switch ctos.Widgetid {
	case constant.Defense1:
		fallthrough
	case constant.Defense2:
		fallthrough
	case constant.Defense3:
		fallthrough
	case constant.Defense4:
		fallthrough
	case constant.Defense5:
		_, err := UseWidget(c.UserData.Userid, ctos.Widgetid, 1)
		if err == nil {
			duration, err := getDefenseTime(ctos.Widgetid)
			if err == nil {
				propNum, err := CheckDefense(c.UserData.Userid)
				if err == nil {
					if propNum == 0 {
						endtime, err := db.MasterDB.InsertCoinBuff(c.UserData.Userid, ctos.Widgetid, duration);
						if err == nil {
							stoc.Data = &vo.StoC6012Data2{}
							stoc.Data.Widgetid = ctos.Widgetid
							stoc.Data.Endtime = endtime
						}
					}else {
						endtime, err := db.MasterDB.UpdateCoinBuff(c.UserData.Userid, ctos.Widgetid, duration);
						if err == nil {
							stoc.Data = &vo.StoC6012Data2{}
							stoc.Data.Widgetid = ctos.Widgetid
							stoc.Data.Endtime = endtime
						}
					}
				}
			}else {
				stoc.E = 15002
			}
		}else {
			stoc.E = 15002
		}
	default:
		stoc.E = 15003
	}
	log.Infoln(ctos.T, ctos.Widgetid)
	b, err := stoc.Encode()
	if err != nil {
		log.Error(err)
	}
	c.Send <- *b
}

//获取离线被攻击次数
func attackNum(msg *[]byte, c *socket.Connection) {
	ctos := vo.CtoS6016Data1{}
	ctos.Decode(msg)
	stoc := vo.StoC6016Data1{T:ctos.T}
	data, err :=db.SlaveDB.OffLineAttackNum(c.UserData.Userid)
	if err == nil {
		stoc.Data = data
	}else{
		stoc.E = 15010
	}
	log.Infoln(ctos.T)

	b, err := stoc.Encode()
	if err != nil {
		log.Error(err)
	}
	c.Send <- *b
}

//获取离线被攻击详细列表
func attackList(msg *[]byte, c *socket.Connection) {
	ctos := vo.CtoS6017Data1{}
	ctos.Decode(msg)
	stoc := vo.StoC6017Data1{T:ctos.T}
	data, err :=db.SlaveDB.OffLineAttackList(c.UserData.Userid)
	if err == nil {
		stoc.Data = data
		db.MasterDB.UpdateOffLineAttack(c.UserData.Userid)
	}else{
		stoc.E = 15010
	}
	log.Infoln(ctos.T)

	b, err := stoc.Encode()
	if err != nil {
		log.Error(err)
	}
	c.Send <- *b
}

/**
 * 增加和消耗积分
 * @param int userid 用户账号
 * @param int coinnum 积分数
 * @param int typeid 积分类型 ；详见积分类型表
 * @param int recordtype 积分记录类型：1：增加积分；2：消耗积分
 * @return int coinnum 改变积分数
 */
func AddOrLessenCoin(userid int, coinnum int64, typeid int, recordtype int) (change int64, e interface{}) {
	score := &db.Tb_score_stat{Account_num:userid}
	err := score.Read()

	if err != nil {
		log.Errorln(err)
		return 0, err
	}
	count := score.Score_total
	if recordtype == 2 {
		if score.Score_total == 0 {
			return 0, nil;
		}
		if score.Score_total < coinnum {
			coinnum = score.Score_total
		}
		count = score.Score_total - coinnum
	}else {
		count = score.Score_total + coinnum
	}

	err1 := db.MasterDB.CoinRecord(userid, coinnum, typeid, recordtype)
	if err1 != nil {
		log.Errorln(err1)
		return 0, err1
	}

	score.Score_total = count
	err2:=score.UpdateCoin()
	if err2 != nil {
		log.Errorln(err2)
		return 0, err2
	}

	content := vo.StoC6015Data1{T: 6015, Data:int(count)}
//	targetConn, ok := socket.Hub.Connections[userid]

	if socket.Hub.Exists(userid) {
		b, _ := content.Encode()
//		targetConn.Send <- *b

		m := vo.Broadcast{}
		m.Channel = userid
		m.Msg = *b
		socket.Hub.Broadcast <-m
	}

	return coinnum, nil;
}

/**
 * 检测用户积分防御
 * @param int userID 用户账号
 * @return int num 防御道具ID
 */
func CheckDefense(userID int) (num int, e interface{}) {
	db.MasterDB.DelCoinBuff(userID)
	propNum, err := db.SlaveDB.GetCoinBuff(userID)
	if err != nil {
		log.Infoln(err)
		return propNum, nil
	}
	return propNum, nil
}

/**
 * 获取道具吸分数值
 * @param int widgetID 吸分道具ID
 * @return int coinNum 吸取分数
 */
func getGainCoinNum(widgetID int) (coinNum int64, e interface{}) {
	data, ok := csv.Suck.Hash[widgetID]
	if ok {
		num := data.Lowerlimit + rand.Intn(data.Toplimit - data.Lowerlimit + 1)
		return int64(num), nil
	}else {
		return 0, errors.New("无该道具吸分数值")
	}
}

/**
 * 获取防御道具时间
 * @param int widgetID 吸分道具ID
 * @return int duration 持续时间
 */
func getDefenseTime(widgetID int) (duration int, e interface{}) {
	data, ok := csv.Defsuck.Hash[widgetID]
	if ok {
		return data.Duration * 60, nil
	}else {
		return 0, errors.New("无该道具防御时间")
	}
}

/**
 * 获取吸分行为结果
 * @param int attackID 吸分道具ID
 * @return int defensedID 防御道具ID
 */
func getDefenseStatus(attackID int, defensedID int) bool {
	if defensedID == 0{
		return true
	}
	suck, ok := csv.Suck.Hash[attackID]
	defsuck, ok1 := csv.Defsuck.Hash[defensedID]
	if ok && ok1 {
		base := defsuck.Def + (defsuck.Grade - suck.Grade) * 10
		if rand.Intn(101) > base {
			return true
		}else {
			return false
		}
	}else {
		return false
	}
}

/**
 * 发送防御生效消息
 * @param int attackUser 进攻者ID
 * @param int defensedUser 防御者ID
 * @param int widgetID 吸分道具ID
 */
func send6013Msg(attackUser *vo.UserData, defensedUser int, widgetID int,defstuckID int) {
//	attack, _ := socket.Hub.Get(attackUser)
//	defensed, ok := socket.Hub.Connections[defensedUser]
	status := 0
	if socket.Hub.Exists(defensedUser) {
		content := vo.StoC6013Data1{T: 6013,Data:&vo.StoC6013Data2{}}
		content.Data.Userid =attackUser.Userid
		content.Data.Widgetid =widgetID
		content.Data.Nickname =attackUser.Nickname

		b, _ := content.Encode()
//		defensed.Send <- *b
		m := vo.Broadcast{}
		m.Channel = defensedUser
		m.Msg = *b
		socket.Hub.Broadcast <-m

		status = 1
	}
	db.MasterDB.AttackRecord(attackUser.Userid,defensedUser,widgetID,defstuckID,0,0,1,status)
}

/**
 * 发送被吸分消息
 * @param int attackUser 进攻者ID
 * @param int defensedUser 防御者ID
 * @param int widgetID 吸分道具ID
 * @return int coin 被吸分数
 */
func send6014Msg(attackUser *vo.UserData, defensedUser int, widgetID int,defstuckID int, coin int64) {
//	attack,_:=socket.Hub.Get(attackUser)
//	defensed, ok := socket.Hub.Connections[defensedUser]/
	status := 0
	if socket.Hub.Exists(defensedUser) {
		content := vo.StoC6014Data1{T: 6014,Data:&vo.StoC6014Data2{}}
		content.Data.Userid =attackUser.Userid
		content.Data.Widgetid =widgetID
		content.Data.Nickname =attackUser.Nickname
		content.Data.Coin =int(coin)
		b, _ := content.Encode()
//		defensed.Send <- *b
		m := vo.Broadcast{}
		m.Channel = defensedUser
		m.Msg = *b
		socket.Hub.Broadcast <-m
		status = 1
	}
	db.MasterDB.AttackRecord(attackUser.Userid,defensedUser,widgetID,defstuckID,coin,0,0,status)
}