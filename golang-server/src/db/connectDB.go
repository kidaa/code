/*
*
* 数据库实现读写分类，连接池
*
*/
package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/golang/glog"
	"regexp"
	"github.com/beedb/orm"
)

func init() {
	orm.RegisterModel(new(Tb_user_member),new(Tb_world_currently),new(Tb_user_passwd),
		new(Tb_user_active),new(Tb_score_stat),new(Tb_build_user),
		new(Tb_dialog),new(Tb_dialog_stat),new(Tb_sys_dialog_stat),
		new(Tb_subdialog),new(Tb_user_topstick),
	)
}

type Connector struct{
	db  *sql.DB
}

// 从数据库，负责查询
type Slave struct{
	Connector
}
// 主数据库，负责更新，插入，删除
type Master struct{
	Connector
}

var SlaveDB  = &Slave{}
var MasterDB  = &Master{}


var OrmerSlave orm.Ormer
var OrmerMaster orm.Ormer

//  连接数据库
func (this *Connector)Create(name string,ip string, port string, account string, pwd string, dbname string) orm.Ormer {
	var err error
	this.db, err = sql.Open("mysql", account+":"+pwd+"@tcp("+ip+":"+port+")/"+dbname+"?charset=utf8")

	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}
//	设置最大打开连接数，默认为0表示没有连接上限
//	this.db.SetMaxOpenConns(2000)

//	设置最大空闲连接数
	this.db.SetMaxIdleConns(10)
	//只有当需要使用时才会创建连接，如果想立即验证连接，需要用Ping()方法
	err = this.db.Ping()
	if err != nil {
		// do something here
		log.Fatalf("Open database error: %s\n", err)
	}

	orm.RegisterDataBase(name, "mysql",account+":"+pwd+"@tcp("+ip+":"+port+")/"+dbname+"?charset=utf8", 30)

	o :=  orm.NewOrm()

	o.Using(name)
	return o
}


// SQL字符安全检查
func (this *Connector)SecurityCheck(target string) bool {
	str := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`

	re, err := regexp.Compile(str)
	if err != nil {
		return true
	}

	result:= re.MatchString(target)
	if result == true{
		log.Errorln("There are non security characters")
	}

	return result
}


