package db

import (
	_ "github.com/go-sql-driver/mysql"
	log "github.com/golang/glog"
	"time"
	"vo"
)

/**
 * 创建积分变化记录
 * @param int userid 用户ID
 * @param int coinNum 积分数
 * @param int typeid 积分数
 * @param int recordtype 积分数
 */
func (this *Master) CoinRecord(userid int, coinNum int64, typeid int, recordtype int) (e interface{}) {
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

	log.Infoln(coinNum)

	stmt, err := this.db.Prepare("INSERT  INTO tb_score_record (account_num, score_num, type_id, record_type, create_time) VALUES(?,?,?,?,?)")
	if err != nil {
		log.Errorln(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userid, coinNum, typeid, recordtype, time.Now().Unix())
	if err != nil {
		log.Errorln(err)
		return err
	}
	return nil
}

/**
 * 获取用户积分防御Buff
 * @param int account_num 用户ID
 * @return int prop_num 道具ID
 * @return int end_time 结束时间
 */
func (this *Slave) GetCoinBuff(account_num int) (prop_num int, e interface{})  {
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
	var propNum int
	err = this.db.QueryRow("SELECT prop_number FROM tb_user_buff WHERE account_num = ? AND prop_type = 2",
		account_num).Scan(&propNum)

	if err != nil {
		log.Errorln(err)
		return 0,nil
	}
	return propNum,nil
}

/**
 * 更新用户积分防御Buff
 * @param int account_num 用户ID
 * @param int widgetid 道具ID
 * @param int duration 持续时间
 * @return int endtime 结束时间
 */
func (this *Master) UpdateCoinBuff(account_num int, widgetid int, duration int) (endtime int ,e interface{})  {
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

	startTime := time.Now().Unix()
	endTime := startTime + int64(duration)
	stmt, err := this.db.Prepare("UPDATE tb_user_buff SET prop_number = ?, start_time = ?, end_time = ? WHERE account_num = ? AND prop_type = 2")
	if err != nil {
		log.Errorln(err)
		return 0,err
	}
	defer stmt.Close()
	_, err = stmt.Exec(widgetid, startTime, endTime, account_num)
	if err != nil {
		log.Errorln(err)
		return 0,err
	}
	return int(endTime),nil
}

/**
 * 插入用户积分防御Buff
 * @param int account_num 用户ID
 * @param int widgetid 道具ID
 * @param int duration 持续时间
 * @return int endtime 结束时间
 */
func (this *Master) InsertCoinBuff(account_num int, widgetid int, duration int) (endtime int ,e interface{})  {
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

	startTime := time.Now().Unix()
	endTime := startTime + int64(duration)
	stmt, err := this.db.Prepare("INSERT INTO tb_user_buff(account_num, prop_number,prop_type, start_time, end_time) VALUES(?,?,?,?,?)")
	if err != nil {
		log.Errorln(err)
		return 0,err
	}
	defer stmt.Close()
	_, err = stmt.Exec(account_num, widgetid, 2, startTime, endTime)
	if err != nil {
		log.Errorln(err)
		return 0,err
	}
	return int(endTime),nil
}

/**
 * 删除用户积分防御Buff
 * @param int account_num 用户ID
 * @return int count 积分总数
 */
func (this *Master) DelCoinBuff(account_num int) (e interface{})  {
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

	stmt, err := this.db.Prepare("DELETE FROM tb_user_buff WHERE account_num = ? AND prop_type = 2 AND  end_time <= ? ")
	if err != nil {
		log.Errorln(err)
		return err
	}
	defer stmt.Close()
	_,err = stmt.Exec(account_num, time.Now().Unix())
	if err != nil {
		log.Errorln(err)
		return err
	}
	return nil
}

/**
 * 获取用户Buff
 * @param int account_num 用户ID
 * @return int prop_num 道具ID
 * @return int end_time 结束时间
 */
func (this *Slave) GetUserBuffList(account_num int) (list []*vo.StoC6010Data2, e interface{})  {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
		}
	}()

	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
		return nil,err
	}

	rows, err := this.db.Query("SELECT prop_number, end_time FROM tb_user_buff WHERE account_num = ?",
		account_num)
	defer rows.Close()

	var result = make([]*vo.StoC6010Data2, 0,5)
	for rows.Next() {
		data:= &vo.StoC6010Data2{}
		err := rows.Scan(&data.Widgetid, &data.Endtime)
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
 * 删除过期Buff
 * @param int account_num 用户ID
 */
func (this *Master) DelTimeoutBuff(account_num int) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
		}
	}()

	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
	}

	stmt, err := this.db.Prepare("DELETE FROM tb_user_buff WHERE account_num = ? AND end_time <= ? ")
	if err != nil {
		log.Errorln(err)
	}

	defer stmt.Close()
	_,err = stmt.Exec(account_num, time.Now().Unix())
	if err != nil {
		log.Errorln(err)
	}
}

