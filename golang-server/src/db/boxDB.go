package db

import (
	_ "github.com/go-sql-driver/mysql"
	"redis"
	"strconv"
	"csv"
	"time"
)

/**
* 获取用户世界宝箱数量
* @param int account_num 用户ID
* @param int box_id 宝箱ID
* @param int world_id 世界ID
* @return int num 数量
*/
func (this *Slave) GetWorldBoxNum(account_num int, world_id int, box_id int) (num int, err error) {
	total, err := redis.GetTmpExpiresNum("WorldBoxNum" + strconv.Itoa(account_num) + strconv.Itoa(world_id))
	if err == nil {
		return total, err
	}else {
		member := &Tb_user_member{Account_num: account_num}
		err := member.Read()
		if err != nil {
			return 0, err
		}
		numVO, ok := csv.BoxNum.Hash[box_id]
		if ok {
			total = numVO.BaseNum + int(member.Grade/numVO.EachLevel) * numVO.EachNum
			nowTime := time.Now().Unix()
			nowTime  += int64(60*60*24)
			tm := time.Unix(nowTime, 0)
			renewTime := time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, time.UTC)
			err := redis.SetTmpExpiresNum("WorldBoxNum" + strconv.Itoa(account_num) + strconv.Itoa(world_id), total, renewTime.Unix())
			if err != nil {
				return total, err
			}
		}
	}
	return total, nil
}

/**
*更新用户世界宝箱数量
* @param int account_num 用户ID
* @param int box_id 宝箱ID
* @param int world_id 世界ID
* @return int num 数量
*/
func (this *Master) UpdateWorldBoxNum(account_num int, world_id int, box_id int, count_num int) (num int, err error) {
	nowTime := time.Now().Unix()
	nowTime  += int64(60*60*24)
	tm := time.Unix(nowTime, 0)
	renewTime := time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, time.UTC)
	err = redis.SetTmpExpiresNum("WorldBoxNum" + strconv.Itoa(account_num) + strconv.Itoa(world_id), count_num, renewTime.Unix())
	return count_num, err
}

