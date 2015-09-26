/**
 * Created by Michael on 2015/8/11.
 *	用户数据表,负责所以用户数据的存储
 *
 *
 */
package db

type Tb_user_member struct {
	Nickname      string  `orm:"size(30)"`    //用户昵称
	Account_num   int     `orm:"pk;size(10)"` //用户账号id
	Account_email string  `orm:"size(70)"`    //用户邮箱
	Qq_uid string `orm:"size(32)"`
	Phone_num     string  `orm:"size(11)"`    //手机号码
	Sex           int     `orm:"size(1)"`     //男1 女2 非男非女3
	Province_id   int     `orm:"size(2)"`     //省id
	City_id       int     `orm:"size(5)"`     //市id
	Birth         int64   `orm:"size(10)"`    //用户出生日期
	Sign          string  `orm:"size(51)"`    //签名
	Status        int     `orm:"size(1)"`     //正常1  锁定2  黑名单3
	User_type     int     `orm:"size(1)"`     //普通用户1 付费用户2  商家3
	Terminal_type int     `orm:"size(1)"`     //终端类型 1:PC，2:手机
	Grade         int     `orm:"size(5)"`     //等级
	Grade_exp     float32 `orm:"size(8)"`     //等级经验值
	Create_time   int64   `orm:"size(10)"`    //注册时间
	Creater_ip    int64   `orm:"size(10)"`    // 注册ip
}


func (this *Tb_user_member) Insert() error {
	_,err:= OrmerMaster.Insert(this)
	return err
}


//  读取全部数据
func (this *Tb_user_member) Read() error {
	return OrmerSlave.Read(this)
}

//  读取名字和头像
func (this *Tb_user_member) ReadNameAndHeadPic() error {
	return OrmerSlave.QueryTable("Tb_user_member").Filter("Account_num", this.Account_num).One(this, "Create_time","Nickname")
}


//  获取用户注册时间戳
func (this *Tb_user_member) GetCreateTime() error {
	return OrmerSlave.QueryTable("Tb_user_member").Filter("Account_num", this.Account_num).One(this, "Create_time")
}

func (this *Tb_user_member) UpateNickname() error {
	_, err := OrmerMaster.Update(this, "Nickname")
	return err
}
func (this *Tb_user_member) UpateSex() error {
	_, err := OrmerMaster.Update(this, "Sex")
	return err
}

func (this *Tb_user_member) UpateAddress() error {
	_, err := OrmerMaster.Update(this, "Province_id", "City_id")
	return err
}

func (this *Tb_user_member) UpateSign() error {
	_, err := OrmerMaster.Update(this, "Sign")
	return err
}

func (this *Tb_user_member) UpateEmail() error {
	_, err := OrmerMaster.Update(this, "Account_email")
	return err
}

func (this *Tb_user_member) UpatePhone() error {
	_, err := OrmerMaster.Update(this, "Phone_num")
	return err
}

