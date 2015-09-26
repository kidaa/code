/*
*
*socket服务器程序入口
*
 */

package main

import (
	_ "backpack"
	_ "chat"
	"db"
	"flag"
	"github.com/Unknwon/goconfig"
	log "github.com/golang/glog"
	_ "karma"
	"net/http"
	"policy"
	_ "proxy"
	"redis"
	"rpc"
	"socket"
	_ "user"
	_ "widget"
	_ "world"
	_"socketlogin"
	_"github.com/beedb/orm"
	_"box"
	_"forum"
)
import (
	_ "net/http/pprof"
)

func main() {
	defer func() {
		if e := recover(); e != nil {
			log.Errorln(e)
			log.Flush()
		}
	}()

	flag.Parse()
	//	go tool pprof http://localhost:3001/debug/pprof/profile

	// 加载配置文件
	_, err := goconfig.LoadConfigFile("./mail.ini")

	config, err := goconfig.LoadConfigFile("./config.ini")

	if err != nil {
		log.Fatal("get config file err: ", err)
	}
	policyPort, err := config.GetValue("policy", "port")
	if err != nil {
		log.Fatal("get config file err: ", err)
	}
	// Flash 策略文件
	go policy.Create(policyPort)

	section, err := config.GetSection("redis")
	addr := section["addr"]
	pwd := section["pwd"]

	redis.CacheInit(addr, pwd)
	redis.CreateCache()

	rpcConfig, err := config.GetSection("rpc")

	rpc.CreateServer(rpcConfig["port"])

	//连接从数据库
	data, err := config.GetSection("mysql_slave")
	if err != nil {
		log.Fatal("get config file err: ", err)
	}
	db.OrmerSlave = db.SlaveDB.Create("default",data["ip"], data["port"], data["account"], data["pwd"], data["dbname"])


	//连接主数据库
	data, err = config.GetSection("mysql_master")
	if err != nil {
		log.Fatal("get config file err: ", err)
	}
	db.OrmerMaster = db.MasterDB.Create("master",data["ip"], data["port"], data["account"], data["pwd"], data["dbname"])

	// 全局事件监听
	go socket.Hub.Run()
	http.HandleFunc("/", socket.WSHandler)
	socketPort, err := config.GetValue("socket", "port")
	log.Infoln("socket Listen on: ", socketPort, "policy Listen on: ", policyPort)
	if err != nil {
		log.Fatal("get config file err: ", err)
	}

	err = http.ListenAndServe(":"+socketPort, nil)
	if err != nil {
		log.Fatal("get config file err: ", err)
	}

	// 退出时调用，确保日志写入文件中
	log.Flush()
}
