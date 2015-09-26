package db

// 论坛对白属性表
type Tb_dialog_stat struct {
//	Stat_id        int `orm:"size(10)"` //	表唯一id
	Dialog_id      int `orm:"pk;size(10)"` //   对白唯一id
	Up_worth       int   `orm:"size(8)"`     // 置顶值(登天值), 默认值0
	Up_total_time  int   `orm:"size(8)"`     // 置顶总共已经停留时间, 默认值0
	Up_end_time    int `orm:"size(10)"`    //当处于置顶时的结束时间; 默认值0
	Up_remain_time int   `orm:"size(5)"`     //置顶剩余时间, 默认值0

	Up_flag    int `orm:"size(1)"` //置顶标示位; 默认值0,1:置顶
	Follow_num int `orm:"size(8)"` //跟帖数, 默认值0
	Hot_num    int `orm:"size(5)"` //热度数, 默认值0

	Dialog_type int    `orm:"size(1)"`  //对白类型；默认值1:用户对白， 2:原始对白
	Abs_path    string `orm:"size(15)"` // 对白位置:世界_栋_楼_房_情景; 其他的表别名: dialog_path
	Create_time int  `orm:"size(10)"` // 创建时间戳
}

// 读取一整行
func (this *Tb_dialog_stat) Read() error {
	return OrmerSlave.Read(this)
}

//写入一整行
func (this *Tb_dialog_stat) Insert() (int64,error) {
	return OrmerMaster.Insert(this)
}
