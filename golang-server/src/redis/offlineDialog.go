/**
 * Created by Michael on 2015/8/19.
 *	Redis缓存，离线对白
 *	系统在Redis 建立存储离线对白的有序列表
 *	1.根据用户发表对白的时间戳排序
 *	2.列表每个元素设置过去时间为7天
 *	3.根据用户离线时间计算用户离线对白内容
 *
 *
 */
package redis

import (
	"strconv"
	"time"
	"github.com/garyburd/redigo/redis"
	log"github.com/golang/glog"
	"strings"
)

const (
	OFFLINE_DIALOG = "Offline_dialog"
	OFFLINE_Expires = 7* 24 * 60 * 60

)

type OfflineDialog struct  {}

// 移除过期的离线对白
func (this *OfflineDialog)removeOFFLINEExpires() {
	defer func() {
		if err := recover(); err!= nil {
			log.Errorln(err)
		}
	}()

	for{
		Do("ZREMRANGEBYSCORE", OFFLINE_DIALOG, time.Now().Unix() - OFFLINE_Expires, time.Now().Unix() + OFFLINE_Expires*9)
		<- time.After(time.Hour*24)
	}
}

// 添加离线对白
func  (this *OfflineDialog) AddOfflineDialg(dialogid int,abs_path string) error {
	return Send("ZADD", OFFLINE_DIALOG, time.Now().Unix(), strconv.Itoa(dialogid) +":"+abs_path)
}

// 获取离线对白,unix用户离线时间戳,只获取位置对白，过来掉重复位置的对白
func  (this *OfflineDialog)GetOfflineDialg(unix int64) ([]string,error) {

	var l = make([]string, 0, 10)

	a, err := Do("ZREVRANGEBYSCORE", OFFLINE_DIALOG, time.Now().Unix(),unix)
	if err != nil {
		return  l,err
	}
	b, err := redis.MultiBulk(a, err)
	if err != nil {
		return  l,err
	}

	hash := make(map[string]bool)
	for _, v := range b {
		s, err := redis.String(v, nil)
		if err != nil {
			continue
		}

		slist := strings.Split(s,":")
		path:=  slist[1]
		if _,ok:= hash[path];!ok{
			l = append(l,path)
			hash[path] = true
		}
	}

	return l,nil
}

