/**
 * Created by Michael on 2015/7/31.
 */
package redis

import (
	"github.com/Unknwon/goconfig"
	"testing"
	"vo"
	"time"
)

func Test(t *testing.T) {
	// 加载配置文件
	config, err := goconfig.LoadConfigFile("../config.ini")
	if err != nil {
		t.Error("config.GetValue redis_test: ", err)
	}
	section, err := config.GetSection("redis")
	addr := section["addr"]
	pwd := section["pwd"]

	CacheInit(addr, pwd)

	msg:= &vo.MsgData{}
	msg.Userid = 60033
	msg.Msg = "hello2"
	msg.Sendtime = (time.Now().Unix())

//	data:= &vo.MsgData{}
//	list:= make([]*vo.MsgData,0,5)
	list:= make([]*vo.MsgData,0,10)
	ChatMsg.GetAllOfflineChatMsg(60229,&list)

	for _,v:= range list{
		t.Log(*v)
	}

//	<-time.After(time.Second*3)
//	t.Log(ChatMsg.GetOfflineChatMsgByFriendid(60434,60229,0,&list))
//	err=ChatMsg.DelOfflineChatMsgByFriendid(60434,60033)
//	t.Log(list,err)

	//	cookie:= &http.Cookie{}
	//			session.Sessions[cookie.Value] =sessonData

	//			expire,_:= strconv.Atoi(cookie.Expires)

	//	userLoginSessionStruct:= &UserLoginSessionStruct{HashMap{"user_login_session_cache"}}
	//	userLoginSessionStruct.SetSessionByID(60229,cookie)

	//	newsData:= &vo.NewsData{}
	//	newsData.ID = 123456
	//	newsData.Title = "朋友订婚宴上，朋友的母亲端起一杯酒走到亲家母位子边，恭恭敬敬说了一大套客气话，最后说：这闺女模样好，嘴巴甜，又聪明又乖巧，我和全家都打心眼里喜欢，俗话说真是破窑出好货呀……全桌顿时凝固 "
	//	NewsCache.SetNews(newsData)
//
//	for i := 0; i < 10; i++ {
//		newsd := &vo.NewsData{}
//		newsd.Title = "michael" + strconv.Itoa(i)
//		newsd.Source = "sina"
//		newsd.Url = "http://baidu.com"
//		newsd.Category = "资讯"
//		err:=NewsCache.SetNews(newsd)
//		t.Log(err)
//	}

//	d, err:= NewsCache.GetNews()
//	for i := 0; i < len(d); i++ {
//		t.Log(*(d[i]))
//	}


//	list,_:=NewsCache.GetNews()
//	t.Log(len(list))
	//	b:=GetUserForgotPWDSession("cuJzcQEja04vRxigMWTwMGa2Nxe5NVRmRh4vCYJqBX1zRxigNW35NvAsNxaoMWywNva5OWTtMFJ9",sessonData)

	//	t.Log(b)

	//	byteArr,e:= redis.Bytes(Do("GET","UserSession:"+strconv.Itoa(60229)))
	//	t.Log(string(byteArr),e)
	//	hash := NewHashMap("user_data_cache")

	//	data := vo.UserData{}
	//	data.Userid = 60434
	//	data.Email = "dolotech@163.com"
	//	data.Phone = "13435333722"
	//	hash.PutObject(strconv.Itoa(data.Userid), data)
	//
	//	data = vo.UserData{}
	//	data.Userid = 60229
	//	data.Email = "xuzhenhui0425@163.com"
	//	data.Phone = "13435333722"
	//	hash.PutObject(strconv.Itoa(data.Userid), data)

	//	var result vo.UserData
	//	hash.GetObject("60229", &result)
	//	t.Log(result)

	//	hash.GetObject("60434", &result)
	//	t.Log(result)

}
