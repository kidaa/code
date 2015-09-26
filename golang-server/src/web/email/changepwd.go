/*
*
*邮件验证更改密码
*
*/

package email

/*
type CodeStruct struct {
	time int64
	code int
}
*/

//var hash = make(map[string]CodeStruct)



/*func SaveCode(email string,code int){
	d:= CodeStruct{}
	d.time = time.Now().Unix()
	d.code = code
	hash[email] = d
	log.Infoln(email,d)
//	go delCode(email)
}*/
// 2小时后删除验证码
/*
func delCode(email string) {
	select {
		case <- time.After(time.Hour * 2):
		delete(hash,email)
	}
}
*/
/*

func CheckCode(email string,code int)int{
	codeStr,ok:= hash[email]
	log.Infoln(email,codeStr,ok,code)
	rel:= 0
	if ok {
		if time.Now().Unix() -  codeStr.time > 2* 60* 60{
			rel =   13005
		}else{
			if codeStr.code == code{
				rel =   0
			}else{
				rel =   13006
			}
		}
	}else{
		rel =   13010
	}

	return  rel
}
*/
