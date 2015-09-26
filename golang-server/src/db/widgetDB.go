package db

import (
	_ "github.com/go-sql-driver/mysql"
	log "github.com/golang/glog"
	"time"
	"errors"
	."vo"
	"redis"
	"strconv"
)


// 获取玩家的所有道具
func (this *Slave)  GetUserWidget(userid int) []*WidgetData {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
		}
	}()
	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
	}

	rows, err := this.db.Query("SELECT articles_id,articles_type,sum,bag_num,create_time FROM tb_user_bag WHERE account_num =?",
		userid)


	defer rows.Close()

	result := make([] *WidgetData, 10)
	result = result[:0]
	//
	for rows.Next() {
		widgetData := &WidgetData{}
		err := rows.Scan(&widgetData.Id, &widgetData.Type, &widgetData.Count, &widgetData.Position, &widgetData.CreateTime)
		if err != nil {
			log.Errorln(err)
		}
		result = append(result, widgetData)
	}
	//
	err = rows.Err()
	if err != nil {
		log.Errorln(err)
	}

	return result
}


// 判断玩家的指定道具数量,返回该道具剩余数量
func (this *Slave)  GetCountByWidgetID(userid int, widgetID int) int{
	var wcount int
	err := this.db.QueryRow("SELECT sum FROM tb_user_bag WHERE  account_num = ? AND articles_id = ?",
		userid, widgetID).Scan(&wcount)
	if err != nil {
		log.Errorln(err)
		return wcount     // 没有改道具
	}
	return wcount
}


// 删除玩家的指定道具,返回该道具剩余数量
func (this *Master)  DelUserWidget(userid int, widgetID int, count int) (int, error) {
	var wcount int = SlaveDB.GetCountByWidgetID(userid,widgetID)

	if wcount == 0 {
		return 0, errors.New("没有该道具")        // 没有该道具
	}

	var leftCount int
	if count >= wcount {
		stmt, err := this.db.Prepare("DELETE FROM tb_user_bag WHERE account_num = ? AND articles_id = ?")
		if err != nil {
			log.Errorln(err)
			return wcount,err
		}
		if stmt != nil{
			defer stmt.Close()
		}
		_, err = stmt.Exec(userid, widgetID)
		if err != nil {
			log.Errorln(err)
			return wcount,err
		}

		leftCount = 0
		MasterDB.WidgetRecord(userid, widgetID, wcount, 1)
	}else {
		stmt, err := this.db.Prepare("UPDATE tb_user_bag SET sum = ? WHERE account_num = ? AND articles_id = ?")
		if err != nil {
			log.Errorln(err)
			return wcount,err
		}
		if stmt != nil{
			defer stmt.Close()
		}
		_, err = stmt.Exec(wcount-count, userid, widgetID)
		if err != nil {
			log.Errorln(err)
			return wcount,err
		}
		leftCount = wcount - count
		MasterDB.WidgetRecord(userid, widgetID, count, 1)
	}

	return leftCount, nil
}

// 添加玩家的指定道具
func (this *Master)  AddUserWidget(userid int, widgetID int, count int) (c int, e interface{}) {
	var wcount int = -1
	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
		return wcount, e
	}

	if count <=0 {
		log.Errorln("购买道具数量必须大于0")
		return 0, errors.New("购买道具数量必须大于0")
	}

	err = SlaveDB.db.QueryRow("SELECT sum FROM tb_user_bag WHERE  account_num = ? AND articles_id = ?",
		userid, widgetID).Scan(&wcount)

	if err != nil {
		//		log.Errorln(err)
		//		return wcount,err
	}
	log.Infoln(wcount)
	// 执行alter table tb_user_bag add UNIQUE (account_num,articles_id) 建立UNIQUE类型的索引，避免插入重复数据，
	// 也可在Navicat视图下，右键设计表->切换到索引栏->右键添加索引，建立成功记得保持

	if wcount == -1 {
		stmt, err := this.db.Prepare("INSERT  INTO tb_user_bag(account_num, articles_id,articles_type,sum,bag_num,create_time) VALUES(?,?,?,?,?,?)")
		if err != nil {
			log.Errorln(err)
			return wcount, err
		}
		defer stmt.Close()
		_, err = stmt.Exec(userid, widgetID, 0, count, 0, time.Now().Unix())
		if err != nil {
			log.Errorln(err)
			return wcount, err
		}
		wcount = count
		log.Infoln(wcount)
	}else {
		stmt, err := this.db.Prepare("UPDATE tb_user_bag SET sum = ? WHERE account_num = ? AND articles_id = ?")
		if err != nil {
			log.Errorln(err)
			return wcount, err
		}
		defer stmt.Close()
		_, err = stmt.Exec(wcount+count, userid, widgetID)
		if err != nil {
			log.Errorln(err)
			return wcount, err
		}
		wcount =  count +wcount
		log.Infoln(wcount)
	}

	MasterDB.WidgetRecord(userid, widgetID, count, 0)

//	if err != nil {
//		log.Errorln(err)
//		return wcount, err
//	}
	return wcount, nil
}

/**
*道具获取使用记录
* @param int account_num 用户ID
* @return int widgetID 道具ID
* @return int count 道具数量
* @return int recordType 记录类型 0:购买 1:消耗
*/
func (this *Master) WidgetRecord(account_num int, widgetID int, count int, recordType int) (e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e =err
		}
	}()
	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
	}

	stmt, err := this.db.Prepare("INSERT  INTO tb_prop_records (account_num, prop_number, count_num, record_type, create_time) VALUES(?,?,?,?,?)")
	if err != nil {
		log.Errorln(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(account_num, widgetID, count, recordType, time.Now().Unix())
	if err != nil {
		log.Errorln(err)
		return err
	}
	return nil
}

/**
*道具交战记录
* @param int attack_user 进攻者ID
* @param int defensed_user 防御者ID
* @param int attack_prop 进攻道具ID
* @param int defensed_prop 防御道具ID
* @param int change_num 改变数值
* @param int change_type 改变数值类型
* @param int status 交战结果  0:成功 1:被防御 2:反弹
* @param int msg_status 消息状态 0：未读 1：已读
*/
func (this *Master) AttackRecord(attack_user int, defensed_user int, attack_prop int, defensed_prop int, change_num int64, change_type int, status int, msg_status int) (e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e =err
		}
	}()
	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
	}

	stmt, err := this.db.Prepare("INSERT  INTO tb_attack_records (attack_user, defensed_user, attack_prop, defensed_prop, change_number,change_type,status,msg_status,create_time) VALUES(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Errorln(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(attack_user, defensed_user, attack_prop, defensed_prop, change_num, change_type, status, msg_status, time.Now().Unix())
	if err != nil {
		log.Errorln(err)
		return err
	}
	return nil
}

/**
* 离线被攻击次数
* @param int account_num 用户ID
*/
func (this *Slave) OffLineAttackNum(account_num int) (count int,e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()

	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
		return 0,err
	}
	var total int
	err = this.db.QueryRow("SELECT COUNT(id) FROM tb_attack_records WHERE defensed_user = ? AND msg_status = 0",
		account_num).Scan(&total)

	if err != nil {
		log.Errorln(err)
		return 0,nil
	}
	return total,nil
}

/**
* 离线被攻击列表
* @param int account_num 用户ID
*/
func (this *Slave) OffLineAttackList(account_num int) (list []*StoC6017Data2,e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()

	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
		return nil,err
	}
	rows, err := this.db.Query("SELECT attack_user, attack_prop, defensed_prop, change_number,change_type,status,create_time FROM tb_attack_records WHERE defensed_user = ? AND msg_status = 0",
		account_num)
	defer rows.Close()

	var result = make([]*StoC6017Data2, 0,10)
	for rows.Next() {
		data:= &StoC6017Data2{}
		err := rows.Scan(&data.Attackuser, &data.Attackprop, &data.Defensedprop, &data.Changenum, &data.Changetype, &data.Status, &data.Createtime )
		if err != nil {
			log.Errorln(err)
		}
		result = append(result, data)
	}
	err = rows.Err()
	if err != nil {
		log.Errorln(err)
	}

	return result,nil
}

/**
* 更新被攻击列表为已读
* @param int account_num 用户ID
*/
func (this *Master) UpdateOffLineAttack(account_num int) (e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()

	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
		return err
	}
	stmt, err := this.db.Prepare("UPDATE tb_attack_records SET msg_status = 1 WHERE defensed_user = ? AND msg_status = 0")
	if err != nil {
		log.Errorln(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(account_num)
	if err != nil {
		log.Errorln(err)
		return err
	}
	return nil
}

func checkErr(err error) {
	if err != nil {
		log.Errorln(err)
	}
}

/**
* 获取用户烟花状态
* @param int account_num 用户ID
* @return int status 0：可以播放，1:自己不能播放，2:对方不能播放
*/
func (this *Slave) GetBothFireworksStatus(account_num int,other_id int) (status int) {
	err := redis.GetTmpExpires("Firework"+strconv.Itoa(account_num))
	if err == nil || account_num == 0 {
		return 1
	}
	err = redis.GetTmpExpires("Firework"+strconv.Itoa(other_id))
	if err == nil || other_id == 0 {
		return 2
	}
	return 0
}

