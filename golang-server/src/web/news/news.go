package news

import (
	"web/proxy"
	"web/webvo"
	"redis"
	"time"
)


type NewsItem struct  {
	Category       string    `xml:"category"`
	Description       string    `xml:"description"`
	Title       string    `xml:"title"`
	Link       string    `xml:"link"`
}

type ChannelVO struct  {
	Item []NewsItem     `xml:"item"`
}

type NewsChannel struct  {
	Channel ChannelVO   `xml:"channel"`
}

func init() {
	proxy.Regist(12001, getRandomNews)
	//	resp, err := http.Get("http://rss.sina.com.cn/news/society/wonder15.xml")
	//	http://rss.sina.com.cn/news/society/misc15.xml

	go fetchNews()
}


func getNewsNow() {
	go crawlers("http://rss.sina.com.cn/news/society/wonder15.xml","奇闻")
	go crawlers("http://rss.sina.com.cn/news/society/misc15.xml","万象")
	go crawlers("rss.sina.com.cn/news/allnews/eladies.xml","轶事")

}

func fetchNews() {
	<- time.After(time.Second * 10)
	getNewsNow()
	for{
		if time.Now().Hour() == 10{
			getNewsNow()
			<- time.After(time.Second * 10)
			go redis.NewsCache.ClearOldData()
		}

		<- time.After(50 * time.Minute)
	}
}

func getRandomNews(uid int,body *[]byte) (bit *[]byte) {
	stoc := webvo.WebStoC12001Data1{T: 12001}

	list,err:=redis.NewsCache.GetNews()
	if err != nil{
		stoc.E = 16003
	}else{
		for i := 0; i < len(list); i++ {
			data := &webvo.WebStoC12001Data2{}
			data.Category = list[i].Category
			data.Url = list[i].Url
			data.Title =  list[i].Title
			data.Source = list[i].Source
			stoc.Data = append(stoc.Data,data)
		}
	}

	b, _ := stoc.Encode()
	return b
}
