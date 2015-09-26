/**
 * Created by Michael on 2015/7/27.
 */
package sms
import (
	"testing"
)

//
func Test(t *testing.T) {
//	s :=SMSStruct{}
//	s.Account = "szzd0059"
//	s.Password = "wwawo2015"
//	s.Mobile ="13435333722"
//	s.Content ="尊敬的用户：您本次手机绑定的验证码： "+  "53132" +  " 。30分钟内有效，请勿泄露。【游区】"
//	s.Action = "send"
//	b, err := json.Marshal(s)
//	if err != nil {
//		t.Log("json err:", err)
//	}
//	t.Log("Marshal ", string(b))
//	body := bytes.NewBuffer([]byte(b))

//	res,err := http.Post("http://sz.ipyy.com/sms.aspx", "application/json;charset=utf-8", body)

	t.Log(SendMSM("13435333722","65655"))
}