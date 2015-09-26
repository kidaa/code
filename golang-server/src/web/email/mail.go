/*
*
*邮件验证更改密码
*
 */

package email

import (
	"github.com/Unknwon/goconfig"
	log "github.com/golang/glog"
	"math/rand"
	"net/smtp"
	"regexp"
	"strconv"
	"strings"
	"time"
    "utils"
)

/*
 *  user : example@example.com login smtp server user
 *  password: xxxxx login smtp server password
 *  host: smtp.example.com:port   smtp.163.com:25
 *  to: example@example.com;example1@163.com;example2@sina.com.cn;...
 *  subject:The subject of mail
 *  body: The content of mail
 *  mailtyoe: mail type html or text
 */

var config *goconfig.ConfigFile

func sendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

// 产生6位数验证码
func GenerateIdentifyCode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return strconv.Itoa(int(r.Float32()*899999 + 100000))
}



// 替换邮件内容的验证码和玩家id,(1888888)里面的是玩家ID，<font color="red">223568</font>,“><”里面的是验证码，不能有空格
func replaceContent(content string, code string, id int) string {
	//    text:=`<html><body><b>亲爱的游区用户:</b><br> 您好！<br> <b>您正在为(1888888)</b> 申请密码找回，请将下方验证码填写到密码找回页面的验证码输入框中进入下一步：<br> <h3 align="center" > 您本次申诉的验证码为： <font color="red">223568</font> </h3>该验证码2个小时内有效，验证成功后立即失效。如非您本人操作请忽略。<br><br> 游区客服团队</body></html>`
	reg := regexp.MustCompile(`\(\d+\)`)
	content = reg.ReplaceAllString(content, "("+strconv.Itoa(id)+")")

	reg = regexp.MustCompile(`\>\d+\<`)
	content = reg.ReplaceAllString(content, ">"+code+"<")

	return content
}

func Send(to string, code string, id int) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorln("Send email error: ", err)
		}
	}()
	if utils.EmailRegexp(to) == false {
		log.Errorln("email addr is wrong")
		return
	}

	if id == 0 {
		log.Errorln("email mode id is 0")
		return
	}
	user, err := config.GetValue("mail", "user")
	if err != nil {
		log.Errorln("get config file err: ", err)
	}

	password, err := config.GetValue("mail", "password")
	host, err := config.GetValue("mail", "host")

	subject, err := config.GetValue("content", "subject")

	body, err := config.GetValue("content", "body")

	body = replaceContent(body, code, id)

	log.Infoln("send email")
	err1 := sendMail(user, password, host, to, subject, body, "html")
	if err1 != nil {
		log.Errorln("send mail error!", err1)
	} else {
		log.Infoln("send mail success!")
	}
}

func init() {
	// 加载配置文件
	var err error
	config, err = goconfig.LoadConfigFile("./mail.ini")
	if err != nil {
		log.Fatal("get config file err: ", err)
	}
}
