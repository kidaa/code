/**
 * Created by Michael on 2015/8/7.
 *针对缘分纸条数据库存取接口
 *
 *
 */
package db

import (
	_ "github.com/go-sql-driver/mysql"
	log "github.com/golang/glog"
	"time"
	. "vo"
)

// 添加一条缘分纸条
func (this *Master) AddKarma(send_id int, widget_id int, msg string, sex int, birth int64, provinceid int, cityid int) (e interface{}) {
	var err error
	defer func() {
		if err := recover(); err != nil {
			e = err
		}
	}()
	stmt, err := this.db.Prepare("INSERT INTO tb_karma(send_id,widget_id,content,sex,birth,province_id,city_id,create_time) VALUES(?,?,?,?,?,?,?,?)")
	defer stmt.Close()

	if err != nil {
		log.Errorln(err)
	}
	_, err = stmt.Exec(send_id, widget_id, msg, sex, birth, provinceid, cityid, time.Now().Unix())
	if err != nil {
		log.Errorln(err)
	}
	return e
}

//  根据ID查询指定纸条的数据
func (this *Slave) QuestKarmaByid(id int) (karmaData *KarmaData, e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()
	data := &KarmaData{}
	err := this.db.QueryRow("SELECT k.send_id,k.status,k.widget_id,k.read_times,k.content,k.sex,k.birth,k.province_id,k.city_id,k.create_time FROM tb_karma AS k WHERE k.id = ?",
		id).Scan(&data.SendID, &data.Status, &data.WidgetID, &data.ReadTimes, &data.Content, &data.Sex, &data.Birthday, &data.Provinceid, &data.Cityid, &data.CreateTime)

	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	data.ID = id
	return data, nil

}

func (this *Master) DelKarma(id int) (e interface{}) {
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

	stmt, err := MasterDB.db.Prepare(`DELETE FROM tb_karma WHERE id=?`)
	if stmt != nil {
		defer stmt.Close()
		_, err = stmt.Exec(id)
		if err != nil {
			log.Errorln(err)
			return err
		}
	}

	return nil
}

// 获取一条缘分纸条, 排除已经建立缘分的纸条再次获取，自己不能获取自己的缘分纸条。权重排序：获取异性>读取次数最少>发布时间最老
func (this *Slave) GetKarma(userid int, sex int, birth int64, provinceid int, cityid int) (karmaData *KarmaData, e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()
	data := &KarmaData{}
	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
		return data, err
	}

	err = this.db.QueryRow("SELECT k.id,k.status,k.send_id,k.widget_id,k.read_times,k.content,k.sex,k.birth,k.province_id,k.city_id,k.create_time FROM tb_karma AS k "+
		" WHERE k.status =0 AND  k.send_id != ? AND k.send_id NOT IN(SELECT r.friend_id FROM tb_karma_relation AS r WHERE r.account_num = ?) ORDER BY if(k.sex=?,1,0), k.read_times,k.create_time DESC limit 0, 1 ",
		userid, userid, sex).Scan(&data.ID, &data.Status, &data.SendID, &data.WidgetID, &data.ReadTimes, &data.Content, &data.Sex, &data.Birthday, &data.Provinceid, &data.Cityid, &data.CreateTime)
	if err != nil {
		log.Errorln(err)
		return data, err
	}

	// 删除24 小时的纸条
	if (time.Now().Unix() - data.CreateTime) > (24 * 60 * 60) {
//		if data.Status == 0{
			stmt, err := MasterDB.db.Prepare("UPDATE tb_karma SET status = 1 WHERE id =?")
			if stmt != nil {
				defer stmt.Close()
				_, err = stmt.Exec(data.ID)
				if err != nil {
					log.Errorln(err)
				}
			}
//		}
	} else {
		// 更新读取次数
		stmt, err := MasterDB.db.Prepare("UPDATE tb_karma SET read_times = read_times+1 WHERE id =?")
		if stmt != nil {
			defer stmt.Close()
			_, err = stmt.Exec(data.ID)
			if err != nil {
				log.Errorln(err)
			}
		}
	}

	return data, nil
}

// 插入缘分关系
func (this *Master) AddKarmaRelation(account_num int, friend_id int) (e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()

	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
	}

	stmt, err := this.db.Prepare("INSERT INTO tb_karma_relation(account_num,friend_id,create_time) VALUES(?,?,?)")
	if stmt != nil {
		defer stmt.Close()

		_, err = stmt.Exec(account_num, friend_id, time.Now().Unix())

		_, err = stmt.Exec(friend_id, account_num, time.Now().Unix())
	}

	return e
}

// 获取缘分关系
func (this *Slave) GetKarmaRelation(account_num int) (relaList []*RelationKarma, e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()

	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
	}

	var result = make([]*RelationKarma, 0, 10)
	rows, err := this.db.Query("SELECT account_num,friend_id,status,create_time FROM tb_karma_relation WHERE account_num=?",
		account_num)
	if rows != nil {
		defer rows.Close()

		for rows.Next() {
			var data *RelationKarma = new(RelationKarma)
			err := rows.Scan(&data.Userid, &data.Friendid, &data.Status, &data.CreateTime)
			if err != nil {
				log.Errorln(err)
			}
			result = append(result, data)
		}

		err = rows.Err()
		if err != nil {
			log.Errorln(err)
		}
	}

	return result, e
}

// 删除一条缘分关系
func (this *Master) DeleteKarmaRelation(userid int, friendid int) int64 {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
		}
	}()
	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
	}
	var num int64
	stmt, err := this.db.Prepare(`DELETE FROM tb_karma_relation WHERE account_num=? AND friend_id= ?`)
	if stmt != nil {
		defer stmt.Close()

		res, err := stmt.Exec(userid, friendid)
		if err != nil {
			log.Errorln(err)
		}
		num, err = res.RowsAffected()
		if err != nil {
			log.Errorln(err)
		}
	}

	return num
}

// 检索缘分纸条是否存在
func (this *Slave) ExistKarma(id int) (e interface{}) {
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

	var count int
	err = this.db.QueryRow("SELECT id FROM tb_karma WHERE id =?",
		id).Scan(&count)
	if err != nil {
		return err
	}

	return nil
}

// 插入缘分离线消息
func (this *Master) AddKarmaMsg(send_id int, receive_id int, msg string) (e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()

	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
	}

	stmt, err := this.db.Prepare("INSERT INTO tb_karma_msg(send_id,receive_id,content,create_time) VALUES(?,?,?,?)")

	if stmt != nil {
		defer stmt.Close()
		_, err = stmt.Exec(send_id, receive_id, msg, time.Now().Unix())
	}

	return err
}

// 查询未读，缘分聊天消息数量
func (this *Slave) NoReadKarmaMsgCount(receive_id int) int {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
		}
	}()
	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
	}

	var count int
	err = this.db.QueryRow("SELECT COUNT(*) from tb_karma_msg where status = 0 AND receive_id =?",
		receive_id).Scan(&count)
	return count
}

// 当前用户对应某缘分好友的聊天消息列表
func (this *Slave) KarmaMsgByUserList(receive_id int, send_id int) (d []*MsgData, e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()

	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	var result = make([]*MsgData, 0, 10)
	rows, err := this.db.Query("SELECT content ,create_time FROM tb_karma_msg WHERE status = 0 AND send_id = ? AND receive_id=?",
		send_id, receive_id)
	if rows != nil {
		defer rows.Close()

		for rows.Next() {
			var data *MsgData = new(MsgData)
			err := rows.Scan(&data.Msg, &data.Sendtime)
			if err != nil {
				log.Errorln(err)
			}
			result = append(result, data)
		}

		err = rows.Err()
		if err != nil {
			log.Errorln(err)
			return nil, err
		}
	}

	//读取完消息，已读状态置1
	stmt, err := MasterDB.db.Prepare("UPDATE tb_karma_msg SET status = ? WHERE send_id=? AND receive_id = ? AND status= 0")
	if stmt != nil {
		defer stmt.Close()
		_, err = stmt.Exec(1, send_id, receive_id)

	}

	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	return result, nil
}

// 获取缘分好友的最后一般离线消息
func (this *Slave) GetLastKarmaMsgByOfUser(send_id int, receive_id int) (sendTime int, msg string, e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()

	err := this.db.QueryRow("SELECT content,create_time FROM tb_karma_msg WHERE status = 0 AND send_id =? AND receive_id = ? ORDER BY create_time DESC LIMIT 0,1", send_id, receive_id).Scan(&msg, &sendTime)
	if err != nil {
		log.Errorln(err)

		return sendTime, msg, err
	}
	return sendTime, msg, nil
}

// 所有缘分好友对应的聊天消息数量，返回ID和对应的数量
func (this *Slave) UserKarmaMsgCountList(userid int) (list []int, e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()
	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	result := make([]int, 0, 10)
	rows, err := this.db.Query("SELECT send_id from tb_karma_msg WHERE receive_id =? AND status = 0", userid)
	if rows != nil {
		defer rows.Close()

		for rows.Next() {
			var id int
			err = rows.Scan(&id)
			if err != nil {
				log.Errorln(err)
			}
			result = append(result, id)
		}

		if err != nil {
			log.Errorln(err)
			return nil, err
		}
	}

	if err != nil {
		log.Errorln(err)
		return nil, err
	}

	return result, nil
}



// 验证是否是缘分好友
func (this *Slave) IsKarmaRelation(userid int,otherid int) (boo bool, e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()
	result := false
	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
		return result, err
	}

	var id int
	err = this.db.QueryRow("SELECT id FROM tb_karma_relation WHERE account_num =? AND friend_id = ?",userid,otherid).Scan(&id)
	if err!= nil{
		return result, err
	}

	if id >0 {
		result = true
	}

	return result, nil
}

