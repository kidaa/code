package weather

import (
	"encoding/json"
	log "github.com/golang/glog"
	"io/ioutil"
	"net/http"
)

type jsonVO struct {
	Code int    `json:"code"`
	Data CityVO `json:"data"`
}

type CityVO struct {
	Country   string `json:"country"` // 国家
	Cuntry_id string `json:"country_id,omitempty"`

	Area    string `json:"area"` // 地区如：华南区
	Area_id string `json:"area_id,omitempty"`

	Region    string `json:"region"` //  省份如：广东省
	Region_id string `json:"region_id,omitempty"`

	City    string `json:"city"` // 城市如：深圳市
	City_id string `json:"city_id,omitempty"`

	County    string `json:"county"` // 如：宝安区
	County_id string `json:"county_id,omitempty"`

	Isp    string `json:"isp"` // 网络运用商如：电信
	Isp_id string `json:"isp_id,omitempty"`
	Ip     string `json:"ip"`
}

func IP2City(ip string) *CityVO {

	defer func() {
		if e := recover(); e != nil {
			log.Errorln(e)
		}
	}()

	//	ip:= "220.173.122.113" //广西壮族自治区
	//		ip:= "119.137.58.164" //深圳
	//	ip:= "61.128.126.23" //新疆维吾尔自治区

	resp, _ := http.Get("http://ip.taobao.com/service/getIpInfo.php?ip=" + ip)
	m := &jsonVO{}
	if resp != nil{
		defer resp.Body.Close()
		content, _ := ioutil.ReadAll(resp.Body)
		//		t.Log(string(content))
		err := json.Unmarshal(content, m)
		if err != nil {
			log.Errorln("ip 地址无法解析")
			return nil
		}
	}else{
		return nil
	}

	//	t.Log("Area: ",m.Data.Area)
	//	t.Log("City: ",m.Data.City)
	//	t.Log("Country: ",m.Data.Country)
	//	t.Log("County: ", m.Data.County)
	//	t.Log("Region: ",m.Data.Region)
	//	t.Log("Isp: ",m.Data.Isp)

//	var cityId = ""
//	if m.Data.Region != "" {
//		if region, ok := CityHash[m.Data.Region]; ok {
//			if city, ok := region[m.Data.City]; ok {
//				cityId = city
//			}
//		}
//	}

	return &m.Data
}
