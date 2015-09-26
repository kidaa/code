package db

import (
	"github.com/Unknwon/goconfig"
	"testing"
	"redis"
)


// 测试数据库
func Test(t *testing.T) {
	// 加载配置文件
	config, err := goconfig.LoadConfigFile("../config.ini")
	if err != nil {
		t.Error("config.GetValue socket port: ", err)
	}
	//连接从数据库
	data, err := config.GetSection("mysql_slave")
	OrmerSlave = SlaveDB.Create("default", data["ip"], data["port"], data["account"], data["pwd"], data["dbname"])


	//连接主数据库
	data, err = config.GetSection("mysql_master")
	OrmerMaster = MasterDB.Create("master", data["ip"], data["port"], data["account"], data["pwd"], data["dbname"])
	if err != nil {
		t.Error("config.GetValue socket port: ", err)
	}


	section, err := config.GetSection("redis")
	addr := section["addr"]
	pwd := section["pwd"]

	redis.CacheInit(addr, pwd)
	redis.CreateCache()

	user := Tb_user_member{Account_num:60434}
	user.ReadNameAndHeadPic()

	t.Log(user)
	/*args := []interface{}{}
	args = append(args, this.Name)
	for _, v := range k {
		args = append(args, v)
	}
	reply, err := redis.MultiBulk(Do("HMGET", args...))
	if err != nil {
		return nil, err
	}
	var list = make([]string, 0)
	for _, v := range reply {
		s, err := redis.String(v, nil)
		if err != nil {
			break
		}
		s = strings.Trim(s, "\"")
		list = append(list, s)
	}*/



	//	DELETE FROM 表B WHERE C not in (SELECT C FROM A)
	//	MasterDB.db.Query("DELETE FROM tb_dialog WHERE dialog_id not in (SELECT dialog_id FROM tb_dialog_stat)")

	//	d:= &vo.UserData{}
	//	SlaveDB.GetUserData(60229,60434,d)
	//	t.Log(*d)


	//	MasterDB.ChangeWorld(60229,99)

	//	t.Log(SlaveDB.GetUserDialogs(60229,0))


	//	t.Log((time.Now().Unix() - 1439883707 )/(24 * 3600),time.Now().Unix())
	//
	//	ids:= []int{60229,60434,60175}
	//	d,e:=SlaveDB.GetUsersData(ids)
	//	t.Log(d,e)
	//	d:=
	//	d,_:=MasterDB.GetCoin(60110)
	//	t.Log(MasterDB.UpdateDialogTopstick(60229,"1_1_2_1_7",-88))
	//	t.Log(MasterDB.AddKarma(60229,0,"hello",1,0,0,0))
	//	t.Log(MasterDB.AddKarma(60229,0,"hello",2,0,0,0))
	//	t.Log(MasterDB.AddKarma(60434,0,"hello",2,0,0,0))


	//	stmt, err := MasterDB.db.Prepare(`DELETE FROM tb_friend WHERE account_num=? AND friend_id= ?`)
	//	defer stmt.Close()
	//
	//	res, err := stmt.Exec(60229, 0)
	//
	//	num, err := res.RowsAffected()
	//
	//	t.Log(num,err)
	//	userData:=SlaveDB.SearchUserByIDOrNickname(60229,60434,"")

	//	t.Log(QueryFriendAndUserTable(60229,0,1))
	//	t.Log(QueryFriendAndUserTable(60229,0,1))

	//		t.Log(SlaveDB.FriendsCount(60229))
	//		MasterDB.UpdateRequest(60334,60172)
	//	t.Log(NoReadMsgCount(60229),"asdad")

	//	b,_:=json.Marshal(QueryFriendAndUserTable(60229))
	//	err=InsertFriendRequst(60334,60172)
	//	if err != nil {
	//		t.Log("已经添加好友",err)
	//	}
	//	InsertAddFriendRequst(60334,60172) //	t.Log()
	//	addRequest(60229,60172)
	//	t.Log(MsgList(60434))
	//
	//		MasterDB.Changepwd(60135,"6d6a6d585e550ee364f542f7142c2378")
	//		t.Log(len(d))

	//
	//	str := `(?:')|(?:--)|(/\\*(?:.|[\\n\\r])*?\\*/)|(\b(select|update|and|or|delete|insert|trancate|char|chr|into|substr|ascii|declare|exec|count|master|into|drop|execute)\b)`
	//
	//
	//	re, err := regexp.Compile(str)
	//	if err != nil {
	//		t.Log(err.Error())
	//		return
	//	}
	//	aa := "There are non security characters"
	//	t.Log(re.MatchString(aa))  //打印出false。
}

func Test1(t *testing.T) {

}
