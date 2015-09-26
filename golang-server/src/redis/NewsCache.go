/**
 * Created by Michael on 2015/7/31.
 * 新闻资讯存取
 *
 */

package redis

import (
	"bytes"
	"encoding/gob"
	"github.com/garyburd/redigo/redis"
	"vo"
	"time"
	"github.com/golang/glog"
)

type News struct{}

const (
	KeyName = "news_set_cache"
	NEWS_INCR = "NEWS_INCR"
	PAGE_MAX = 50
)

var NewsCache News

// 随机获取新闻,每页30条
func (this *News) GetNews() ([]*vo.NewsData,error) {
	c, err := Do("ZCARD", KeyName)
	if err != nil {
		return nil, err
	}
	count, err := redis.Int(c, err)

	if err != nil {
		return nil, err
	}
	if count == 0{
		return nil, err
	}

	limit:= 0
	if count <PAGE_MAX{
		limit = count
	}else {
		limit = PAGE_MAX
	}

	a, err := Do("ZREVRANGE", KeyName, 0, limit-1)
	if err != nil {
		return nil, err
	}
	b, err := redis.MultiBulk(a, err)
	if err != nil {
		return nil, err
	}
	list:= make([]*vo.NewsData,0,PAGE_MAX)
	for _, v := range b {
		s, err := redis.Bytes(v, nil)
		var buffer bytes.Buffer
		buffer.Write(s)
		data:= &vo.NewsData{}
		dec := gob.NewDecoder(&buffer)
		err= dec.Decode(data)
		if err != nil {
			continue
		}
		list = append(list, data)
	}
	return list, err
}

// 存储新闻
func (this *News) SetNews(v *vo.NewsData) error {
		Send("INCR","NEWS_INCR")
		incr, err := redis.Int(Do("GET", NEWS_INCR))
		if err != nil {
			return err
		}

		v.CreateTime = time.Now().Unix()
		v.ID = incr

		var buffer bytes.Buffer
		enc := gob.NewEncoder(&buffer)
		err = enc.Encode(v)
		if err != nil {
			return err
		}

		b := buffer.Bytes()


		err = Send("ZADD", KeyName, v.ID,b)
		if err != nil {
			return err
		}
		return nil
}



//
func (this *News) ClearOldData() error {

	a, err := Do("ZRANGE", KeyName, 0, -1)
	if err != nil {
		return err
	}
	b, err := redis.MultiBulk(a, err)
	if err != nil {
		return err
	}
//	list:= make([]*vo.NewsData,0,PAGE_MAX)
	for k, v := range b {
		s, err := redis.Bytes(v, nil)
		var buffer bytes.Buffer
		buffer.Write(s)
		data:= &vo.NewsData{}
		dec := gob.NewDecoder(&buffer)
		err= dec.Decode(data)
		if err != nil {
			continue
		}

		if time.Now().Unix() - data.CreateTime  > 60 * 60 *24 * 7{
//		if time.Now().Unix() - data.CreateTime  > 7{
			err:=Send("ZREMRANGEBYRANK",KeyName,k,k)
			glog.Info(time.Now().Unix() , data.CreateTime,err)
		}
//		list = append(list, data)
	}
	return nil
}
