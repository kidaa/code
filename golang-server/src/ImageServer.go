/**
 * Created by Michael on 2015/8/10.
 *	用户头像和相册存储服务器
 *
 */
package main
import (
	"flag"
	"net/http"
	"time"
	log "github.com/golang/glog"
	"os"
	"io/ioutil"
	"db"
	"utils"
	"strings"
	"redis"
	"github.com/Unknwon/goconfig"
)

func main() {
	flag.Parse()

	defer func() {
		if e := recover(); e != nil {
			log.Errorln(e)
		}
	}()


	// 加载配置文件
	config, _ := goconfig.LoadConfigFile("./config.ini")


	section, err := config.GetSection("redis")
	if err!= nil{
		log.Fatal(err)
	}
	addr := section["addr"]
	pwd := section["pwd"]

	redis.CacheInit(addr, pwd)
	redis.CreateCache()


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

	imagePort, _ := config.GetValue("image", "port")


	http.HandleFunc("/upload", uploadImageHdr)
	http.HandleFunc("/", loadImageHdr)
//
//	http.Handle("/", http.FileServer(http.Dir("/headpic/")))
//	http.ListenAndServe(":8080", nil)


	s := &http.Server{
		Addr:           ":" + imagePort,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Infoln("Listen on: ", imagePort)



	s.ListenAndServe()
}

var realPath string = "./"

func loadImageHdr(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if e := recover(); e != nil {
			log.Errorln(w, r, e)
			w.WriteHeader(404)
		}
	}()

	path := r.URL.Path
	if path != ""{
		index:=strings.LastIndex(path, ".")
		if index != -1{
			request_type := path[index:]
			if request_type == ".jpg" {
				w.Header().Set("Content-Type", "image/jpeg")
				fin, _ := os.Open(realPath + path)
				if fin != nil{
					defer fin.Close()
					fd, _ := ioutil.ReadAll(fin)
					w.Write(fd)
					return
				}
			}
		}

	}
	w.WriteHeader(404)
}


func uploadImageHdr(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if e := recover(); e != nil {
//			log.Errorln(w, r, e)
			w.Write([]byte("post error"))
		}
	}()
	defer r.Body.Close()

	if "POST" != r.Method {
		return
	}

	var err interface{}
	cookie, _ := r.Cookie("login")
	if cookie == nil || cookie.Value == "" {
		w.Write([]byte("no login error"))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorln(err)
		w.Write([]byte("post error"))
		return
	}


	if len(body) == 0 {
		log.Errorln(err)
		w.Write([]byte("file enpty"))
		return
	}

	userid ,err:= redis.GetUseridByCookie(cookie.Value)

	if err != nil {
		w.Write([]byte("cookie error"))
		return
	}

	if userid <= 0{
		w.Write([]byte("userid error"))
		return
	}

	if redis.IsExitsUserSession(cookie.Value){
		userData := &db.Tb_user_member{Account_num:userid}
		err = userData.GetCreateTime()

		if err != nil {
			w.Write([]byte("no this user"))
			return
		}

		path,name:= utils.TimeToHeadphpoto(userData.Create_time,userid)
		f,err:=os.Open(path)
		if f!=nil{
			defer f.Close()
		}else{
			err=os.MkdirAll(path,0777)
			log.Errorln(err)
			if err != nil {
				log.Errorln(err)
				w.Write([]byte("error"))
				return
			}
		}

		err = ioutil.WriteFile(path+name, body, 0777)  //写入文件(字节数组)
		if err != nil {
			log.Errorln(err)
			w.Write([]byte("error"))
		}else{
			f.Write(body)
			w.Write([]byte("ok"))
		}
	}else{
		w.Write([]byte("session error"))
	}
}
