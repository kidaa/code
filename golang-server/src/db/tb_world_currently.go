package db
import "time"


// 用户当前所在世界
type Tb_world_currently struct {
										  //	Id          int64 `orm:"size(10)"`
	World_id    int   `orm:"size(10)"`    //所在世界id
	Account_num int   `orm:"pk;size(10)"` //用户账号id
	Update_time int64 `orm:"size(10)"`    //
}

// 获取用户所在地图
func (this *Tb_world_currently) GetWorldid() error {
	return OrmerSlave.QueryTable("Tb_world_currently").Filter("Account_num", this.Account_num).One(this, "World_id")
}

// 切换用户所在地图
func (this *Tb_world_currently) UpdateWorldID() error {
	t := time.Now().Unix()
	_, err := OrmerMaster.Raw("INSERT INTO tb_world_currently(world_id,account_num,update_time) VALUES(?,?,?)"+
	"ON DUPLICATE KEY UPDATE world_id=?,update_time=?", this.World_id, this.Account_num, t, this.World_id, t).Exec()
	return err
}

func (this *Tb_world_currently) Read() error {
	return OrmerSlave.Read(this)
}

func (this *Tb_world_currently) Insert() error {
	_,err:= OrmerMaster.Insert(this)
	return err
}
