package db

// 记录用户对某条对白的顶踩贴
type Tb_user_topstick struct {
	Id          int  `orm:"pk;size(10)"` //   评论唯一id
	Dialog_id   int  `orm:"size(10)"`    //   评论唯一id
	Dialog_type int    `orm:"size(1)"`     //对白类型；默认值1:用户对白， 2:原始对白
	Account_num int  `orm:"size(10)"`    // 发布对白的用户id
	Abs_path    string `orm:"size(15)"`    // 对白位置:世界_栋_楼_房_情景; 其他的表别名: dialog_path
	Top_value   int   `orm:"size(10)"`
	Down_value  int   `orm:"size(10)"`
	Create_time int  `orm:"size(10)"` // 创建时间戳
}

// 读取整张表
func (this *Tb_user_topstick) Read() error {
	return OrmerSlave.Read(this)
}
