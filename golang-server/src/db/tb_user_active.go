/**
 * Created by Michael on 2015/8/11.
 *	用户在线时长、连续登陆、活跃天数、成长等级
 *
 *
 */
package db

import (
	"csv"
	log "github.com/golang/glog"
	"time"
)

// 用户在线时长表
type Tb_user_active struct {
	Account_num int     `orm:"pk;size(10)"` //用户账号id
	Last_time int64 `orm:"size(10)"`		// 上一次请求时间
	Login_ip int64 `orm:"size(10)"`			// 登陆IP
	Is_online int8    `orm:"size(1)"`		//  是否在线
	Times       int8    `orm:"size(1)"`     //当天获得的活跃天数次数
	Day_num     int     `orm:"size(10)"`    //连续登陆天数
	Active_num  float32 `orm:"size(10)"`    // 累积在线天数
	Login_time  int64   `orm:"size(10)"`    // 登入时间戳
	Logout_time int64   `orm:"size(10)"`    // 登出时间戳
	Online_time int     `orm:"size(10)"`    //每天在线时长（单位秒）
}

func (this *Tb_user_active) UpdateIP() error{
	_,err:=OrmerMaster.Update(this,"Login_ip")
	return  err
}

func (this *Tb_user_active) GetIP() error{
	return  OrmerSlave.QueryTable("Tb_user_active").Filter("Account_num",this.Account_num).One(this, "Login_ip")
}

// 记录登陆时间,没有该玩家数据则插入
func (this *Tb_user_active) Login() {
	t := time.Now().Unix()
	/*	OrmerMaster.Raw("INSERT INTO tb_user_active(account_num,login_time) VALUES(?,?)" +
		"ON DUPLICATE KEY UPDATE login_time=?",this.Account_num,t,t).Exec()*/
	err := OrmerSlave.Read(this)
	if err != nil {
		this.Login_time = t
		OrmerMaster.Insert(this)
	} else {

		this.Continuation(this.Logout_time)
		this.Login_time = t
		this.Last_time = t
		OrmerMaster.Update(this)
	}
}

// 活跃天数计算
func (this *Tb_user_active) ActiveDay() {
	currrent := time.Now()
	gap := int(currrent.Unix() - this.Last_time)
	if gap > 0 && this.Times < 2{

		d := this.Day_num
		if d == 0 {
			d = 1
		}else if d >7{
			d = 7
		}
		this.Online_time = this.Online_time + gap


		if this.Online_time >= 60*60 && this.Times < 2 {
			if this.Times == 0 {
				this.Active_num = this.Active_num + csv.ActiveDaylist.Hash[d].Exp
			}
			this.Active_num = this.Active_num + (csv.ActiveDaylist.Hash[d].Exp)
			this.Times = 2
//			this.Online_time = this.Online_time + gap
//			log.Infoln("在线60分钟")
		} else if this.Online_time >= 20 * 60 && this.Times == 0 {
			this.Active_num = this.Active_num + csv.ActiveDaylist.Hash[d].Exp

			this.Times = 1
//			log.Infoln("在线20分钟")
		}else{
//			log.Infoln("在线<20分钟")
		}
	}
}

// 用户请求
func (this *Tb_user_active) Request() error {
	err := OrmerSlave.Read(this)
	if err != nil {
		return err
	} else {


		this.ActiveDay()				// 一定要先计算活跃天数再计算是否连续登陆
		this.Continuation(this.Last_time)
		this.Last_time = time.Now().Unix()
		_, err := OrmerMaster.Update(this)

		if err != nil {
			return err
		}


	}
	return nil
}

// 连续登陆计算
func (this *Tb_user_active) Continuation(oldTime int64) {
	currrent := time.Now()
	if oldTime < currrent.Unix() && this.Online_time > 0{

		t1 := time.Unix(oldTime, 0)

		// 日期不同一天，清除当日在线时长
		if t1.Day() != currrent.Day() || t1.Year() != currrent.Year() || t1.Month() != currrent.Month() {
			// 判断连续登陆，如果前一天的在线时长达到20分钟，则昨天算作连续登陆+1天
			if time.Unix(oldTime+24*60*60, 0).Day() == currrent.Day() {
				if this.Online_time >= 20 * 60{
					if this.Day_num < 1 {
						this.Day_num = 1
					}
					this.Day_num = this.Day_num + 1
//					log.Infoln("连续登陆")
				}
			}
			this.Times = 0
			this.Online_time = 0

//			log.Infoln("清除昨天数据")
		}
	}
}

// 记录退出时间，并累积在线时长
func (this *Tb_user_active) Logout() {
	err := OrmerSlave.Read(this)
	if err != nil {
		log.Errorln(err)
	} else {
		currrent := time.Now()
		// 一定要先计算活跃天数再计算是否连续登陆

		this.ActiveDay()
		this.Continuation(this.Last_time)
		this.Logout_time = currrent.Unix()

		OrmerMaster.Update(this)
	}
}


// 获取一级货币
func (this *Tb_user_active) GetOfflineTime()(error ){
	return  OrmerSlave.QueryTable("Tb_user_active").Filter("Account_num",this.Account_num).One(this, "Logout_time")
}
