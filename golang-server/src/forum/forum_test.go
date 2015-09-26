package forum
import (
	"testing"
	"github.com/Unknwon/goconfig"
	"db"
	"redis"
)


func Test(t *testing.T) {
	// 加载配置文件
	config, err := goconfig.LoadConfigFile("../config.ini")
	if err != nil {
		t.Error("config.GetValue socket port: ", err)
	}
	//连接从数据库
	data, err := config.GetSection("mysql_slave")
	db.OrmerSlave = db.SlaveDB.Create("default", data["ip"], data["port"], data["account"], data["pwd"], data["dbname"])


	//连接主数据库
	data, err = config.GetSection("mysql_master")
	db.OrmerMaster =  db.MasterDB.Create("master", data["ip"], data["port"], data["account"], data["pwd"], data["dbname"])
	if err != nil {
		t.Error("config.GetValue socket port: ", err)
	}


	section, err := config.GetSection("redis")
	addr := section["addr"]
	pwd := section["pwd"]

	redis.CacheInit(addr, pwd)
	redis.CreateCache()



	msg:=`{"t":8020,"dialogid":19,"abspath":"1_1_2_1_7","dialogtype":1}`
	b:= []byte(msg)

	topstickFirst(&b,nil)

}
