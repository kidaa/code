package db
import (
	"crypto/md5"
	"encoding/hex"
)



func (this *Tb_user_passwd) GetByQQUid() error {
	return OrmerSlave.QueryTable("Tb_user_passwd").Filter("Qq_uid", this.Qq_uid).One(this)
}


func (this *Tb_user_passwd) UpateEmail() error {
	_, err := OrmerMaster.Update(this, "Account_email")
	return err
}

func (this *Tb_user_passwd) UpatePhone() error {
	_, err := OrmerMaster.Update(this, "Phone_num")
	return err
}

func (this *Tb_user_passwd) Read() error {
	return OrmerSlave.Read(this)
}

func (this *Tb_user_passwd) Insert() error {
	_,err:= OrmerMaster.Insert(this)
	return err
}


func (this *Tb_user_passwd) UpdatePWD() error {
	OrmerSlave.Read(this)
	h := md5.New()
	h.Write([]byte(this.Passwd + this.Auth))     // 需要加密的字符串为 123456
	this.Passwd = hex.EncodeToString(h.Sum(nil)) // 输出加密结果
	_, err := OrmerMaster.Update(this, "Passwd")
	return err
}

//  用户登陆密码验证
func (this *Tb_user_passwd) PWDIsOK(pwd string) bool {
	err := OrmerSlave.Read(this)
	if err == nil {
		h := md5.New()
		h.Write([]byte(pwd + this.Auth)) // 需要加密的字符串为 123456
		//  密码正确
		if hex.EncodeToString(h.Sum(nil)) == this.Passwd {
			return true
		}
	}
	return false
}

// 用户密码表
type Tb_user_passwd struct {
	Account_num   int    `orm:"pk;size(10)"` //用户账号id
	Qq_uid string `orm:"size(32)"`
	Account_email string `orm:"size(70)"`    //用户邮箱
	Phone_num     string `orm:"size(11)"`    //手机号码
	Passwd        string `orm:"size(32)"`    //男1 女2 非男非女3
	Auth          string `orm:"size(6)"`     //密码验证码
}