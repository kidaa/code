package news

import (
	"encoding/xml"
	log "github.com/golang/glog"
	"io/ioutil"
	"net/http"
	"redis"
	"time"
	"vo"
)

func crawlers(source string, category string) {
	defer func() {
		if e := recover(); e != nil {
			log.Errorln(e)
		}
	}()
	<-time.After(time.Second * 3)
	resp, err := http.Get(source)

	if err != nil {
		return
	}
	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	//	t.Log(string(content))

	m := &NewsChannel{}

	err = xml.Unmarshal(content, m)
	//	t.Log(len(m.Channel))
	for _, v := range m.Channel.Item {
		data := &vo.NewsData{}
		data.Title = v.Title
		data.Url = v.Link
		data.Category = category
		go redis.NewsCache.SetNews(data)
	}

}
