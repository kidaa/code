/*
*
*Web服务器程序入口
*
 */

package main

import (
	_"csv"
	"db"
	"flag"
	"github.com/Unknwon/goconfig"
	log "github.com/golang/glog"
	"io/ioutil"
	"math/rand"
	"net/http"
	"redis"
	"regexp"
	"rpc"
	"strconv"
	"time"
	"web/email"
	_ "web/login"
	"web/proxy"
	_ "web/register"
	_ "web/shop"
	_ "web/webaccount"
	_"web/news"
	_"web/weather"
	_"web/webworld"
	_"web/message"
)

//   获取图片验证码
func pic(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if e := recover(); e != nil {
			log.Errorln(w, r, e)
			b := []byte("bad request!")
			w.Write(b)
		}
	}()

	validCode := email.ValidationCode{}
	strCode := validCode.NewCdoe(4)

	cookie, _ := r.Cookie("pngverify")
	if cookie == nil || cookie.Value == "" {
		// set cookie一定要在内容返回之前设置，才有效
		cookie = redis.SetCookie(w, "pngverify", int(rand.Int())) // 先设置cookie客户端才能收到
	}

	se := &redis.SessionData{PngCode: strCode}
	go se.Encode(cookie.Value)

	w.Header().Set("Content-Type", "image/png")
	validCode.DrawToImg(strCode, w)
}

func service(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if e := recover(); e != nil {
			log.Errorln(w, r, e)
			w.Write([]byte("request error"))
		}
	}()
	defer r.Body.Close()

	if "POST" != r.Method {
		w.Write([]byte("request error"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("request paramater error"))
		return
	}

	// 获取协议号
	reg := regexp.MustCompile(`"t":[\d]+`)
	s := reg.Find(body)
	reg = regexp.MustCompile(`[\d]+`)
	in, err := strconv.Atoi(string(reg.Find(s)))

	if in > 0 {
		proxy.Run(in,&body,r,w)
	}else{
		w.Write([]byte("proto code error"))
	}
}

func main() {
	defer func() {
		if e := recover(); e != nil {
			log.Errorln(e)
			log.Flush()

		}
	}()

	flag.Parse()

	// 加载配置文件
	config, err1 := goconfig.LoadConfigFile("./config.ini")

	checkErr(err1)

	webPort, er := config.GetValue("php", "port")
	checkErr(er)



	//连接从数据库
	data, dberr := config.GetSection("mysql_slave")
	checkErr(dberr)
	db.OrmerSlave = db.SlaveDB.Create("default",data["ip"], data["port"], data["account"], data["pwd"], data["dbname"])

	//连接主数据库
	data, dberr = config.GetSection("mysql_master")
	checkErr(dberr)
	db.OrmerMaster = db.MasterDB.Create("master",data["ip"], data["port"], data["account"], data["pwd"], data["dbname"])


	rpcConfig, _ := config.GetSection("rpc")
	rpc.CreateClient(rpcConfig["ip"], rpcConfig["port"])

	section, err := config.GetSection("redis")
	checkErr(err)
	addr := section["addr"]
	pwd := section["pwd"]

	redis.CacheInit(addr, pwd)
	redis.CreateCache()

	http.HandleFunc("/uc/show_img", pic)
	http.HandleFunc("/", service)
	s := &http.Server{
		Addr:           ":" + webPort,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Infoln("Listen on: ", webPort)
	s.ListenAndServe()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal("get config file err: ", err)
	}
}
