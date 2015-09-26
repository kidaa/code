package news
import (

	"testing"
	"net/http"
	"io/ioutil"
	"crypto/tls"
	"encoding/json"
)

func TestProtoFriendListData(t *testing.T) {

	ip:= "https://119.147.19.43/v3/user/is_login?"
	appid:= "1104749408"
	pf:= "desktop_m_qq-10000144-android-2002-"
	opendid:= "3677E0F673EB88250B8FF73C7AA21C20"
	openkey:= "0569B069BF7EDD2EF874DC4103EB21CB"
	url:= ip+"openid="+opendid+"&openkey=" + openkey + "&appid="+appid+"&pf=" + pf

	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err:= client.Get(url)

	if resp != nil{
		defer resp.Body.Close()
		content, _ := ioutil.ReadAll(resp.Body)
//		{"ret":-5,"msg":"signature verification failed"}
		type Ret struct  {
			Ret int
			Msg string
		}
		m:= Ret{}
		json.Unmarshal(content,&m)
		t.Log(err,m.Ret==0)
	}
}








