package webworld
import (
	"testing"
//	"web/webvo"
	"github.com/Unknwon/goconfig"
	"db"
)


func Test(t *testing.T) {
	config, err := goconfig.LoadConfigFile("../config.ini")
	if err != nil {
		t.Error("config.GetValue socket port: ", err)
	}
	//连接从数据库
	data, err := config.GetSection("mysql_slave")
	db.OrmerSlave = db.SlaveDB.Create("default", data["ip"], data["port"], data["account"], data["pwd"], data["dbname"])


	//连接主数据库
	data, err = config.GetSection("mysql_master")
	db.OrmerMaster= db.MasterDB.Create("master", data["ip"], data["port"], data["account"], data["pwd"], data["dbname"])
	if err != nil {
		t.Error("config.GetValue socket port: ", err)
	}


	jsonstr:=[]byte(`{"t":21002,"worldid":1}`)

//	vo:=webvo.WebCtoS21002Data1{}

	result:= getWorldBuildShow(60110,&jsonstr)
	t.Log(string(*result))
	t.Log(result)
}