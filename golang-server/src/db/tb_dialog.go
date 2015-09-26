package db

// 论坛对白表
type Tb_dialog struct {
	Dialog_id     int  `orm:"pk;size(10)"` //   评论唯一id
	World_id      int    `orm:"size(3)"`     // 世界id
	Build_id      int    `orm:"size(5)"`     //栋id
	Terminal_type int    `orm:"size(5)"`     // 终端类型; PC:1, 手机:2
	Dialog_type   int    `orm:"size(1)"`     //对白类型；默认值1:用户对白， 2:原始对白
	Account_num   int  `orm:"size(10)"`    // 发布对白的用户id
	Dialog_title  string `orm:"size(200)"`   // 评论的内容
	Abs_path      string `orm:"size(15)"`    // 对白位置:世界_栋_楼_房_情景; 其他的表别名: dialog_path
	Is_anno       int   `orm:"size(1)"`     //是否匿名; 默认值0:实名，1:匿名
	Status        int   `orm:"size(1)"`     //默认值0:正常, 1:屏蔽
	Create_time   int  `orm:"size(10)"`    // 创建时间戳
}

// 读取整张表
func (this *Tb_dialog) Read() error {
	return OrmerSlave.Read(this)
}


//写入一整行
func (this *Tb_dialog) Insert() (int64,error) {
	return OrmerMaster.Insert(this)
}




