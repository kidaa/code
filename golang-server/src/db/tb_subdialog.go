package db

// 子对白结构（对白评论）
type Tb_subdialog struct {
	Dialog_id    int  `orm:"pk;size(10)"` //   评论唯一id
	Parent_id    int  `orm:"size(10)"`    //父对话id
	Account_num  int  `orm:"size(10)"`    // 发布对白的用户id
	Other_account int  `orm:"size(10)"`    // 发布对白的用户id
	Dialog_title string `orm:"size(200)"`   // 评论的内容
	Is_anno      int   `orm:"size(1)"`     //是否匿名; 默认值0:实名，1:匿名
	Status       int   `orm:"size(1)"`     //默认值0:正常, 1:屏蔽
	Create_time  int  `orm:"size(10)"`    // 创建时间戳
}


// 读取一整行
func (this *Tb_subdialog) Read() error {
	return OrmerSlave.Read(this)
}

//写入一整行
func (this *Tb_subdialog) Insert() (int64,error) {
	return OrmerMaster.Insert(this)
}
