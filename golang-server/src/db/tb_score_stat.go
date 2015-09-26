package db

//  用户经济
type Tb_score_stat struct {
	Currency_total    int64 `orm:"size(10)"`    // 元宝一级货币
	Score_total    int64 `orm:"size(10)"`    //金币二级货币
	Account_num int `orm:"pk;size(10)"` //用户账号id
}

// 读取整张表
func (this *Tb_score_stat) Read()error {
	return  OrmerSlave.Read(this)
}

// 获取一级货币
func (this *Tb_score_stat) GetGold()error {
	return  OrmerSlave.QueryTable("Tb_score_stat").Filter("Account_num",this.Account_num).One(this, "Currency_total")
}

func (this *Tb_score_stat) UpdateGold()error {
	_,err:=OrmerMaster.Update(this,"Currency_total")
	return  err
}


// 获取二级货币
func (this *Tb_score_stat) GetCoin()error {
	return  OrmerSlave.QueryTable("Tb_score_stat").Filter("Account_num",this.Account_num).One(this, "Score_total")
}


func (this *Tb_score_stat) UpdateCoin()error {
	_,err:=OrmerMaster.Update(this,"Score_total")
	return  err
}




