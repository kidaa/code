package db
import (
)

//  用户经济
type Tb_build_user struct {
	Build_id    int64 	`orm:"size(10)"`    //大楼ID
	Part    	string	`orm:"size(150)"`   //大楼未激活部分
	Cloud		string	`orm:"size(300)"`   //云未点开部分
	Account_num int 	`orm:"pk;size(10)"` 	//用户账号id
//	Id int "pk"
}

/*
func (this *Tb_build_user) Read()error {
	return  OrmerSlave.Read(this)
}*/







