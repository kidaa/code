/**
 * Created by Michael on 2015/7/27.
* 华信短信平台,向用户发送修改密码验证码信息
*
 */
package sms

import (
	"net/http"
	"io/ioutil"
	"net/url"
	"strings"
)

type SMSStruct struct {
	Userid string `json:"userid"`		//企业ID
	Account string `json:"account"`		//用户账号
	Password string `json:"password"`	//用户密码
	Mobile string `json:"mobile"`		//发送到的目标手机号码,多个用逗号隔开
	Content string `json:"content"`		//短信内容
	Action string `json:"action"`		// 发送任务命令,设置为固定的:send
	Extno string `json:"extno"`			// 扩展子号
}

//
func SendMSM(mobile string,code string) error{
	v := url.Values{}
	v.Set("mobile", mobile)
	v.Set("account", "szzd0059")
	v.Set("password", "wwawo2015")
	v.Set("action", "send")
	// 信息内容要在华信短信后台设置模板，核审通过才能发出去
	v.Set("content", "【游区】尊敬的用户:您本次手机绑定的验证码"+code+"。30分钟内有效，请勿泄露。")

	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://sz.ipyy.com/sms.aspx", body)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value") //这个一定要加，不加form的值post不过去，被坑了两小时

	resp, err := client.Do(req) //发送
	defer resp.Body.Close()     //一定要关闭resp.Body
//	data, _ := ioutil.ReadAll(resp.Body)
	return err
}