package db

import (
	"strconv"
	"time"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/golang/glog"
	"vo"

	"redis"
)


// 添加一个加好友请求
func (this *Master)InsertFriendRequst(req_account int, incept_account int, msg string) (e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()

	err := this.db.Ping()
	if err != nil {
		return err
	}

	// 执行alter table tb_friend_request add UNIQUE (incept_account,req_account) 建立UNIQUE类型的索引，避免插入重复数据，
	// 也可在Navicat视图下，右键设计表->切换到索引栏->右键添加索引，建立成功记得保持
	// 有则更新新无则插入
	stmt, err := this.db.Prepare("INSERT INTO tb_friend_request(req_account,incept_account,req_content,req_time) VALUES(?,?,?,?)" +
	"ON DUPLICATE KEY UPDATE req_status=?,req_content=?,req_time=?")
	if stmt != nil {
		defer stmt.Close()
		systime:= time.Now().Unix()
		_, err = stmt.Exec(req_account, incept_account, msg, systime,0,msg,systime)
	}

	if err != nil{
		return err
	}

	return e
}

// 添加一个好友
func (this *Master)InsertAddFriend(account_num int, friend_id int) (e interface{}){
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()

	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
		return  err
	}

	// 执行alter table tb_friend add UNIQUE (account_num,friend_id) 建立UNIQUE类型的索引，避免插入重复数据，
	// 也可在Navicat视图下，右键设计表->切换到索引栏->右键添加索引，建立成功记得保持
	stmt, err := this.db.Prepare("INSERT ignore INTO tb_friend(account_num, friend_id,create_time) VALUES(?, ?,?)")
	if stmt != nil{
		defer stmt.Close()

		systime:= int(time.Now().Unix())
		_, err = stmt.Exec(account_num, friend_id, systime)
		if err != nil {
			log.Errorln(err)
		}else{
			// redis同步
			syncFiendToRedis(account_num,friend_id)
		}

		_, err = stmt.Exec(friend_id, account_num, systime)
		if err != nil {
			log.Errorln(err)
		}else{
			// redis同步
			syncFiendToRedis(friend_id,account_num)

		}
	}


	if err != nil {
		log.Errorln(err)
		return  err
	}

	return  nil
}

func syncFiendToRedis(account_num int, friend_id int) {

	if redis.UserFriendCache.Exists(strconv.Itoa(friend_id)){
		systime:= int(time.Now().Unix())
		data := &vo.Relationship{}
		data.FriendID = account_num
		data.CreateTime = systime
		go  redis.UserFriendCache.AddFriend(friend_id,data)
	}

}


// 获取好友的最后一条离线消息
/*func (this *Slave) GetLastMsgByOfUser(send_id int,receive_id int) (sendTime int,msg string,e interface{}){
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()

	err := this.db.QueryRow("SELECT content,create_time FROM tb_msg WHERE receive_status = 0 AND send_id =? AND receive_id = ? ORDER BY create_time DESC LIMIT 0,1",send_id,receive_id).Scan(&msg,&sendTime)
	if err != nil {
		log.Errorln(err)

		return sendTime,msg,err
	}
	return sendTime,msg,nil
}*/


// 查询未读，聊天消息数量
/*
func (this *Slave) NoReadMsgCount(receive_id int)int{
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
		}
	}()
	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
		return  0
	}

	var count int
	err = this.db.QueryRow("SELECT COUNT(*) from tb_msg where receive_status = 0 AND receive_id =?",
		receive_id).Scan(&count)
	return count
}
*/


// 插入离线消息
/*func (this *Master)AddMsgToTable(send_id int, receive_id int, msg string) (e interface{}){
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()

	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
		return  err
	}

	stmt, err := this.db.Prepare("INSERT INTO tb_msg(send_id,receive_id,content,create_time) VALUES(?,?,?,?)")
	if stmt != nil{
		defer stmt.Close()
		_, err = stmt.Exec(send_id, receive_id, msg, time.Now().Unix())

		if err != nil {
			log.Errorln(err)
			return  err
		}
	}

	if err != nil {
		log.Errorln(err)
		return  err
	}

	return  nil
}*/
/*

// 当前用户对应某好友的聊天消息列表
func (this *Slave)MsgByUserList(receive_id int, send_id int) (d []*vo.MsgData,e interface{}) {
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
	var result = make([]*vo.MsgData,0, 10)
	rows, err := this.db.Query("SELECT send_id,content ,receive_status,create_time FROM tb_msg WHERE receive_status = 0 AND send_id = ? AND receive_id=?",
		send_id,receive_id)

	if rows != nil{
		defer rows.Close()


		for rows.Next() {
			var data *vo.MsgData = new(vo.MsgData)
			err := rows.Scan(&data.Userid, &data.Msg, &data.Status,&data.Sendtime)
			if err != nil {
				log.Errorln(err)
			}
			result = append(result, data)
		}

		//读取完消息，删除
		stmt, err := MasterDB.db.Prepare("DELETE FROM tb_msg WHERE send_id=? AND receive_id = ?")
		if err != nil{
			log.Errorln(err)
		}

		if stmt != nil{
			defer stmt.Close()
			_, err = stmt.Exec(send_id,receive_id)
		}
	}

	if err != nil {
		log.Errorln(err)
		return nil,err
	}
	return result,nil
}
*/


// 所有好友对应的聊天消息数量，返回ID和对应的数量
func (this *Slave) MsgCountList(userid int) (list []int ,e interface{}){
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

	result:= make([]int,10)
	result=result[:0]
	rows, err := this.db.Query("SELECT send_id from tb_msg WHERE receive_id =? AND receive_status = 0",userid)

	if rows != nil{
		defer rows.Close()
		for rows.Next() {
			var id int
			err = rows.Scan(&id)
			if err != nil {
				log.Errorln(err)
			}
			result = append(result, id)
		}
	}

	if err != nil {
		log.Errorln(err)
		return nil,err
	}

	return result,nil
}


// 请求列表
func (this *Slave)RequestList(account_num int, limit0 int, limit1 int) []*vo.StoC1002Data2 {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
		}
	}()
	var result = make([]*vo.StoC1002Data2,0, 10)
	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
		return  result
	}

	rows, err := this.db.Query("SELECT tb_user_member.account_num,tb_user_member.nickname,tb_user_member.sex,tb_user_member.create_time ,tb_friend_request.req_content,tb_friend_request.req_status FROM tb_user_member INNER JOIN tb_friend_request ON tb_friend_request.req_account=tb_user_member.account_num AND tb_friend_request.incept_account =? And tb_friend_request.req_status < 2 LIMIT ?, ?",
		account_num, limit0, limit1)

	if rows != nil{
		defer rows.Close()
		for rows.Next() {
			var data *vo.StoC1002Data2 = new(vo.StoC1002Data2)
			err := rows.Scan(&data.Userid, &data.Nickname, &data.Sex, &data.Headpic, &data.Msg, &data.Status)
			if err != nil {
				log.Errorln(err)
			}
			result = append(result, data)


		}
	}


	//读取完消息，已读状态置1
	stmt, err :=  MasterDB.db.Prepare("UPDATE tb_friend_request SET req_status = ? WHERE req_status = 0 AND incept_account = ?")
	if stmt != nil{
		defer stmt.Close()
		_, err = stmt.Exec(1, account_num)
		if err != nil {
			log.Errorln(err)
		}
	}

	return result

}

func (this *Slave) IsFriend(userid int,otherid int) *vo.Relationship {

	// 好友关系判断，是否为好友
	if otherid != userid {
		m, _ := this.GetFriendListByUserid(userid)
		for _, v := range m {
			if v.FriendID ==otherid  {
				return v
			}
		}
	}
	return  nil
}



// 获取指定用户的好友关系列表
func (this *Slave) GetFriendListByUserid(account_num int) (userList []* vo.Relationship, e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()

	relationship:= make([]* vo.Relationship,0,10)
	err:=redis.UserFriendCache.GetFriendByID(account_num,&relationship)
	if err != nil{
		err := this.db.Ping()
		if err != nil {
			log.Errorln(err)
			return relationship,err
		}

		rows, err := this.db.Query("SELECT friend_id,remark,group_id,friend_status,create_time FROM tb_friend WHERE account_num =?",account_num)
		if rows != nil{
			defer rows.Close()

			for rows.Next() {
				data:= &vo.Relationship{}
				err := rows.Scan(&data.FriendID, &data.Remarks,&data.GroupID, &data.Status, &data.CreateTime)
				if err != nil {
					log.Errorln(err)
				}
				relationship = append(relationship,data)
			}
		}

		if err != nil {
			log.Errorln(err)
			return relationship,err
		}

		go redis.UserFriendCache.SetFriendByID(account_num,&relationship)
	}
	return relationship,nil
}

// 更新请求数据
func (this *Master) UpdateRequest(req_account int, incept_account int) (e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()
	err := this.db.Ping() // 数据库连接有效性检查
	if err != nil {
		log.Errorln(err)
		return  err
	}

	stmt, err := this.db.Prepare("UPDATE tb_friend_request SET req_status = ? WHERE req_account = ? AND incept_account = ?")
	if stmt != nil{
		defer stmt.Close()
		_, err = stmt.Exec(2, req_account, incept_account)
	}

	if err != nil {
		return  err
	}
	return  nil
}

// 删除一条添加好友请求
func (this *Master)DeleteRequestAddFriend(req_account int, incept_account int) int64 {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
		}
	}()
	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
		return  0
	}
	var num int64
	stmt, err := this.db.Prepare(`DELETE FROM tb_friend_request WHERE req_account=? AND incept_account= ?`)
	if stmt !=nil{
		defer stmt.Close()

		res, err := stmt.Exec(req_account, incept_account)
		if err != nil {
			return  0
		}
		num, err = res.RowsAffected()
	}

	if err != nil {
		log.Errorln(err)
	}

	return num
}

//  模糊搜索用户昵称
func (this *Slave)BlurSearchUserNickname(selfid int,condiction string, limit0 int, limit1 int) []*vo.UserData {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
		}
	}()

	var result = make([]*vo.UserData,0, 10)
	// 搜索条件为空，则返回空数组
	if condiction == "" {
		return result
	}

	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
		return result
	}

	rows, err := this.db.Query("SELECT u.account_num,u.nickname,u.sex,u.province_id,u.city_id,u.sign,u.create_time, (SELECT COUNT(*) FROM tb_friend AS f WHERE f.account_num = " + strconv.Itoa(selfid)+" AND f.friend_id = u.account_num) FROM tb_user_member AS u where u.nickname LIKE '%" + condiction + "%' LIMIT " + strconv.Itoa(limit0) + "," + strconv.Itoa(limit1))
	if rows != nil{
		defer rows.Close()
		for rows.Next() {
			data:= &vo.UserData{}
			err := rows.Scan(&data.Userid, &data.Nickname, &data.Sex, &data.Provinceid, &data.Cityid, &data.Sign, &data.Headpic,&data.Isfriend)
			if err != nil {
				log.Errorln(err)
			}
			result = append(result, data)
		}
	}

	if err != nil {
		return result
	}

	return result
}

// 查询未读，请求叫好友数量
func (this *Slave)NoDealRequestCount(incept_account int) int{
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
		}
	}()
	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
		return  0
	}
	
	var count int
	err = this.db.QueryRow("SELECT COUNT(*) FROM tb_friend_request where req_status= 0 AND incept_account =?",
		incept_account).Scan(&count)
	return count
	
}


// 好友数量
func (this *Slave) FriendsCount(userid int)int{
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
		}
	}()
	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
		return 0
	}

	var count int
	err = this.db.QueryRow("SELECT COUNT(*) from tb_friend WHERE account_num =?",
		userid).Scan(&count)
	return count
}

// 修改好友备注
func (this *Master)  ModifyRemarks(userid int,friendid int,newRemarks string)  int64{
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
		}
	}()
	err := this.db.Ping()
	if err != nil {
		return 0
	}
	stmt, err := this.db.Prepare("UPDATE tb_friend SET remark = ? WHERE account_num = ? AND friend_id = ?")
	var affected int64
	if stmt != nil{
		defer stmt.Close()
		result, err := stmt.Exec(newRemarks, userid, friendid)
		if err != nil {
			log.Errorln(err)
			return  0
		}
		affected, err =result.RowsAffected()

		go redis.UserFriendCache.ModifyFriendRemark(userid,friendid,newRemarks)
	}
	if err != nil {
		return  0
	}

	return affected
}


// 删除好友
func (this *Master)DeleteFriend(userid int, friendid int) (e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()
	err := this.db.Ping()
	if err != nil {
		return 0
	}
	stmt, err := this.db.Prepare(`DELETE FROM tb_friend WHERE account_num=? AND friend_id= ?`)
	if stmt !=nil{
		defer stmt.Close()

		_, err := stmt.Exec(userid, friendid)
		if err != nil {
			return err
		}

		go redis.UserFriendCache.RemoveFriend(userid,friendid)
	}

	if err != nil {
		return err
	}

	return nil
}


