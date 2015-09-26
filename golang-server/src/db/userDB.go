/*
*
* 有关用户账户详细资料数据库存取操作
*
 */

package db

import (
	//	"strconv"
	//	"time"
	//	"vo"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/golang/glog"
	"time"
	. "vo"
	"csv"
	"vo"
	"strings"
	"strconv"
)

//  玩家绑定的手机号搜索当前用户的账号
func (this *Slave) GetUserIDByPhone(phone string) int {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
		}
	}()
	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
	}
	log.Infoln(phone)
	var result int
	err = this.db.QueryRow("SELECT account_num FROM tb_user_member WHERE phone_num=?",
		phone).Scan(&result)

	if err != nil {
		log.Errorln(err)
	}
	log.Infoln(result, phone)
	return result
}

//  玩家绑定的邮箱号搜索当前用户的账号
func (this *Slave) GetUserIDByEmail(email string) int {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
		}
	}()
	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
	}
	log.Infoln(email)
	var result int
	err = this.db.QueryRow("SELECT account_num FROM tb_user_member WHERE account_email=?",
		email).Scan(&result)

	if err != nil {
		log.Errorln(err)
	}
	log.Infoln(result, email)
	return result
}


//检测昵称是否存在
func (this *Slave) CheckNickname(nickname string) (num int, e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()
	var count int
	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
		return count, err
	}

	err = this.db.QueryRow("SELECT COUNT(*) FROM tb_user_member where nickname=?", nickname).Scan(&count)
	if err != nil {
		log.Errorln(err)
		return count, err
	}


	return count, nil
}

// 获取用户对白数量
func (this *Slave) UserDialogCount(account_num int) (count int, e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()

	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
		return count, err
	}

	err = this.db.QueryRow("SELECT COUNT(*) FROM tb_dialog WHERE account_num = ?",
		account_num).Scan(&count)

	if err != nil {
		log.Errorln(err)
		return count, err
	}

	return count, nil

}


//获取未注册的账号
func (this *Slave) GetUnregID() (id int, e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()

	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
		return 0, err
	}
	var userid int
	err = this.db.QueryRow("SELECT account_num FROM tb_vragon_account WHERE id >= (SELECT floor(rand()*(SELECT max(id) FROM tb_vragon_account WHERE flag = 0))) AND flag = 0").Scan(&userid)
	if err != nil {
		log.Errorln(err)
		return 0, err
	}
	return userid, nil
}

//数据库注册账户
func (this *Master) RegisterUser(userData *UserData) (ctime int64, e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln(err)
			e = err
		}
	}()

	err := this.db.Ping()
	if err != nil {
		log.Errorln(err)
		return 0, err
	}
	createTime := time.Now().Unix()


	user:= Tb_user_member{Account_num:userData.Userid}
	user.Account_email = userData.Email
	user.Phone_num = userData.Phone
	user.Terminal_type = userData.TerminalType
	user.Creater_ip = userData.CreaterIP
	user.Qq_uid = userData.Qq_uid
	user.Create_time = createTime
	user.Nickname = userData.Nickname
	user.Sex = userData.Sex
	err= user.Insert()

	if err != nil {
		log.Errorln(err)
		return 0, err
	}

	pwd:=Tb_user_passwd{}
	pwd.Account_num = userData.Userid
	pwd.Account_email = userData.Email
	pwd.Phone_num = userData.Phone
	pwd.Passwd = userData.Pwd
	pwd.Auth = userData.Auth
	pwd.Qq_uid = userData.Qq_uid

	err = pwd.Insert()
	if err != nil {
		log.Errorln(err)
		return 0, err
	}

	world:= Tb_world_currently{Account_num:userData.Userid}
	world.World_id = userData.WorldID
	world.Update_time = createTime

	err = world.Insert()
	if err != nil {
		log.Errorln(err)
		return 0, err
	}

	worldData,ok := csv.World.Hash[userData.WorldID]
	if ok && worldData.Status == 1 {
		woldVO := &vo.UserWorld{}
		woldVO.WorldID = worldData.Id
		woldVO.AccountNum = userData.Userid
		buildArr := strings.Split(worldData.Builds, ",")
		if (len(buildArr) > 0) {
			buildPartCloud := make([]*vo.BuildPartCloud, 0, 10)
			for _, value := range buildArr {
				buildid, _ := strconv.Atoi(value)
				data, ok := csv.Build.Hash[buildid]
				if ok {
					buildVO := &vo.BuildPartCloud{}
					buildVO.Buildid = data.Id
					buildVO.AccountNum = userData.Userid
					buildVO.Part = data.Part
					buildVO.Cloud = data.Cloud
					buildPartCloud = append(buildPartCloud, buildVO)
				}

			}
			err := this.OpenUserWorld(woldVO, buildPartCloud)
			if err != nil {
				return 0,err
			}
		}
	}

	stmt3, err := this.db.Prepare("UPDATE tb_vragon_account SET flag = 1 WHERE account_num = ?")
	if err != nil {
		log.Errorln(err)
		return 0, err
	}
	defer stmt3.Close()
	_, err = stmt3.Exec(userData.Userid)
	if err != nil {
		log.Errorln(err)
		return 0, err
	}

	tb:=Tb_user_active{Account_num:userData.Userid}
	tb.Login_ip = userData.CreaterIP
	tb.Login()
	return createTime, nil
}
