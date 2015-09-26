/**
 * Created by Michael on 2015/8/5.
 *	论坛对白的所以接口
 *
 */
package db

import (
	"errors"
	log "github.com/golang/glog"
	"strconv"
	"strings"
	"time"
	. "vo"
)


// 用户的最新对白,按最新时间降序
func (this *Slave) GetUserDialogs(userid int, page int) (result []*DialogData, e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()
	result = make([]*DialogData, 0, 20)

	if userid <= 0 {
		return result, nil
	}
		err := this.db.Ping()
		if err != nil {
			log.Errorln(err)
			return result, err
		}
		rows, err := this.db.Query("SELECT dialog_id, abs_path,dialog_title,is_anno,status,create_time "+
		"FROM tb_dialog WHERE  account_num =? ORDER BY create_time DESC LIMIT ?,?", userid, page*PAGE_NUM, PAGE_NUM)
		if rows != nil {
			defer rows.Close()


			for rows.Next() {
				data := &DialogData{}

				data.Dialogtype = 1
				data.Userid = userid
				err := rows.Scan(&data.Dialogid, &data.Abspath,&data.Content, &data.Isanno, &data.Status, &data.CreateTime)
				if err != nil {
					log.Errorln(err)
					continue
				}else {
					err := this.GetDialogAttribute(data.Dialogid, data)
					if err != nil {
						continue
					}
				}
				arr := strings.Split(data.Abspath, "_")
				buildid, err := strconv.Atoi(arr[1])
				worldid, err := strconv.Atoi(arr[0])
				data.Worldid = worldid
				data.Buildid = buildid
				result = append(result, data)
			}
		}

		if err != nil {
			log.Errorln(err)
			return result, err
		}

	return result, nil
}

// 获取指定位置的顶贴
func (this *Slave) GetTopDialog(abs_path string) (d *DialogData, e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()

	arr := strings.Split(abs_path, "_")
	buildid, err := strconv.Atoi(arr[1])
	worldid, err := strconv.Atoi(arr[0])
	d = &DialogData{}
	d.Abspath = abs_path
	d.Worldid = worldid
	d.Buildid = buildid
	d.Dialogtype = 1

	err= this.db.QueryRow("SELECT dialog_id,up_worth,up_total_time,up_end_time,up_remain_time,up_flag,follow_num,hot_num "+
	" FROM tb_dialog_stat  WHERE abs_path = ? ORDER BY up_worth DESC LIMIT 0 ,1",
		abs_path).Scan(&d.Dialogid, &d.Upworth, &d.UpTotalTime, &d.UpEndTime, &d.UpRemainTime, &d.UpFlag, &d.FollowNum, &d.HotNum)
	if err != nil {
		log.Errorln(err)
		return d, err
	}else {
		err= this.db.QueryRow("SELECT account_num,dialog_title,terminal_type,is_anno,status,create_time"+
		" FROM tb_dialog  WHERE dialog_id = ?",
			d.Dialogid).Scan(&d.Userid, &d.Content, &d.TerminalType, &d.Isanno, &d.Status, &d.CreateTime)
		if err != nil {
			log.Errorln(err)
		}
	}

	return d, nil

}

// 系统对白:获取指定位置的系统对白
func (this *Slave) GetSysDialogs(abs_path string) (d *DialogData, e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()

	d = &DialogData{}
	if abs_path == "" {
		return d, nil
	}

	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
		return d, err
	}

	err= this.db.QueryRow("SELECT dialog_id,abs_path,up_worth,up_total_time,up_end_time,up_remain_time,up_flag,follow_num,hot_num "+
	" FROM tb_sys_dialog_stat  WHERE abs_path = ?", abs_path).Scan(&d.Dialogid, &d.Abspath, &d.Upworth, &d.UpTotalTime, &d.UpEndTime, &d.UpRemainTime, &d.UpFlag, &d.FollowNum, &d.HotNum)

	if err != nil {
		log.Errorln(err)
		return d, err
	}

	arr := strings.Split(abs_path, "_")

	buildid, err := strconv.Atoi(arr[1])
	worldid, err := strconv.Atoi(arr[0])
	d.Abspath = abs_path
	d.Worldid = worldid
	d.Buildid = buildid
	d.Dialogtype = 2

	return d, nil
}

// 更改指定系统对白的顶贴值
func (this *Master) UpdateSysDialogTopstick(abs_path string, up_worth int) (count int, e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()

	if abs_path == "" {
		return count, errors.New("paramter value error")
	}

	err := this.db.QueryRow("SELECT up_worth FROM tb_sys_dialog_stat  WHERE abs_path =? LIMIT 0,1",
		abs_path).Scan(&count)

	if err != nil {
		log.Errorln(err)
		return count, err
	}

	stmt, err := MasterDB.db.Prepare("UPDATE tb_sys_dialog_stat SET up_worth=up_worth+? WHERE abs_path =?")
	if stmt != nil {
		defer stmt.Close()
		_, err = stmt.Exec(up_worth, abs_path)
	}

	if err != nil {
		log.Errorln(err)
		return count, err
	}
	return count + up_worth, nil
}


func (this *Slave) GetDialogAttribute(dialog_id int, data *DialogData) (e interface{}) {
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

	in := this.db.QueryRow("SELECT up_worth,up_total_time,up_end_time,up_remain_time,up_flag,follow_num,hot_num "+
	"FROM tb_dialog_stat WHERE  dialog_id =? ", dialog_id).Scan(&data.Upworth, &data.UpTotalTime, &data.UpEndTime, &data.UpRemainTime, &data.UpFlag, &data.FollowNum, &data.HotNum)

	if in != nil {
		log.Errorln(in)
		return in
	}
	return nil
}

// 获取指定位置的最新对白,按最新时间降序
func (this *Slave) GetDialogs(abs_path string, page int) (result []*DialogData, e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()
	result = make([]*DialogData, 0, 20)

	if abs_path == "" {
		return result, nil
	}
		err := this.db.Ping()
		if err != nil {
			log.Errorln(err)
			return result, err
		}
		arr := strings.Split(abs_path, "_")

		buildid, err := strconv.Atoi(arr[1])
		worldid, err := strconv.Atoi(arr[0])

		rows, err := this.db.Query("SELECT dialog_id,account_num, dialog_title,is_anno,status,create_time "+
		"FROM tb_dialog WHERE  abs_path =? ORDER BY create_time DESC LIMIT ?,?", abs_path, page*PAGE_NUM, PAGE_NUM)
		if rows != nil {
			defer rows.Close()


			for rows.Next() {
				data := &DialogData{}
				data.Abspath = abs_path
				data.Worldid = worldid
				data.Buildid = buildid
				data.Dialogtype = 1
				err := rows.Scan(&data.Dialogid, &data.Userid, &data.Content, &data.Isanno, &data.Status, &data.CreateTime)


				if err != nil {
					log.Errorln(err)
					continue
				}else {
					err := this.GetDialogAttribute(data.Dialogid, data)
					if err != nil {
						continue
					}
				}
				result = append(result, data)
			}
		}

		if err != nil {
			log.Errorln(err)
			return result, err
		}

	return result, nil
}

// 搜索用户是否踩过指定的对白
func (this *Slave) IsUserUseDowntick(account_num int, dialogid int, dialogtype int, abs_path string) (b bool, e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()
	if dialogid <= 0 {
		return false, errors.New("paramter value error")
	}
	count := 0
	this.db.QueryRow("SELECT down_value FROM tb_user_topstick  WHERE account_num =? AND dialog_id = ? AND abs_path=? AND dialog_type=? ",
		account_num, dialogid, abs_path, dialogtype).Scan(&count)

	return count < 0, nil
}

// 搜索用户是否顶过指定的对白
func (this *Slave) IsUserUseToptick(account_num int, dialogid int, dialogtype int, abs_path string) (b bool, e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()
	if dialogid <= 0 {
		return false, errors.New("paramter value error")
	}
	count := 0
	this.db.QueryRow("SELECT top_value FROM tb_user_topstick  WHERE account_num =? AND dialog_id = ? AND abs_path=? AND dialog_type=?",
		account_num, dialogid, abs_path, dialogtype).Scan(&count)
	return count > 0, nil
}

// 记录用户对指定对白顶贴踩贴（包括使用道具）,如果是系统对白，dialogid等于0
func (this *Master) UseToptick(account_num int, dialogid int, dialogtype int, abs_path string, topValue int, downValue int) (e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()
	if dialogid <= 0 {
		return errors.New("paramter value error")
	}

	// 有则更新新无则插入(必须先在数据库建立一个字段或者多个字段的唯一索引)
	stmt, err := this.db.Prepare("INSERT INTO tb_user_topstick(account_num,dialog_id,dialog_type,abs_path,top_value,down_value,create_time) VALUES(?,?,?,?,?,?,?)" +
	"ON DUPLICATE KEY UPDATE account_num=?,dialog_id=?,dialog_type=?,abs_path=?,top_value=top_value+?,down_value=down_value+?,create_time=?")

	if stmt!= nil {
		defer stmt.Close()
		t := time.Now().Unix()
		_, err = stmt.Exec(account_num, dialogid, dialogtype, abs_path, topValue, downValue, t, account_num, dialogid, dialogtype, abs_path, topValue, downValue, t)
		if err != nil {
			log.Errorln(err)
			return err
		}
	}
	if err != nil {
		log.Errorln(err)
		return err
	}

	return nil
}

// 更改指定对白的顶贴值
func (this *Master) UpdateDialogTopstick(dialogid int, abs_path string, value int) (count int, e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()

	if dialogid <= 0 {
		return count, errors.New("paramter value error")
	}


	err := this.db.QueryRow("SELECT up_worth FROM tb_dialog_stat  WHERE abs_path =? AND dialog_id =? AND dialog_type = ?",
		abs_path,dialogid,1).Scan(&count)

	if err != nil {
		log.Errorln(err)
		return count, err
	}

	// 有则更新新无则插入(必须先在数据库建立dialog_id和abs_path字段的唯一索引)
	stmt, err := MasterDB.db.Prepare("UPDATE tb_dialog_stat SET up_worth=? WHERE abs_path =? AND dialog_id =? AND dialog_type = ?")

	if stmt != nil {
		defer stmt.Close()

		_, err := stmt.Exec(value+count,abs_path,dialogid,1)

		if err != nil {
			log.Errorln(err)
			return count, err
		}

	}

	if err != nil {
		log.Errorln(err)
		return count, err
	}

	return value+count, nil
}
