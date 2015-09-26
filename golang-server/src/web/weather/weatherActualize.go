package weather
import (
	log"github.com/golang/glog"
	"encoding/json"
	"net/http"
	"io/ioutil"
)

func GetWeather(cityId string) * WD {
	defer func() {
		if e := recover(); e != nil {
			log.Errorln(e)
		}
	}()
	h3 := H3VO{}
	if cityId != "" {
		weatherURL := "http://weather.gtimg.cn/city/" + cityId + ".js"
		resp, _ := http.Get(weatherURL)
		if resp != nil{
			defer resp.Body.Close()

			content, _ := ioutil.ReadAll(resp.Body)
			content = content[20:len(content)-1]
			err := json.Unmarshal(content, &h3)
			if err != nil {
				log.Errorln(err,"天气 json 无法解析")
				return nil
			}
		}else{
			return nil
		}
	}


	if len(h3.H3.VO) > 0{
		return  h3.H3.VO[0]
	}
	return nil
}