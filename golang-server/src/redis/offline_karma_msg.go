/**
 * Created by Michael on 2015/8/19.
 *	Redis缓存离线缘分消息,聊天缘分消息只存储在Redis里，数据库没有备份
 *	系统在Redis 建立存储离线缘分消息的有序列表
 *	可以由客户端读取完成主动清除
 *	列表每个消息元素设置过去时间为7天
 *
 */
package redis

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"vo"
)

const (
	OFFLINE_KARMA_MSG = "OfflineKarmaMsg:"
)

type OfflineKarmaMsg struct{}

//	添加跟指定好友的离线聊天记录
func (this *OfflineKarmaMsg) AddMsg(receiveid int, sendid int, msg *vo.MsgData) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return Send("ZADD", OFFLINE_KARMA_MSG+strconv.Itoa(receiveid), sendid, data)
}

//	获取跟指定好友的离线聊天记录数量
func (this *OfflineKarmaMsg) GetMsgCountByFriend(selfid int, friendid int) (int, error) {
	count, err := Do("ZCOUNT", OFFLINE_KARMA_MSG+strconv.Itoa(selfid), friendid, friendid)
	return redis.Int(count, err)
}

//	获取用户跟所有好友的离线聊天记录数量
func (this *OfflineKarmaMsg) GeAllMsgCount(selfid int) (int, error) {
	count, err := Do("ZCOUNT", OFFLINE_KARMA_MSG+strconv.Itoa(selfid), "-inf", "+inf")

	return redis.Int(count, err)
}

//	获取用户和所有好友的离线聊天记录
func (this *OfflineKarmaMsg) GetAllMsg(selfid int,msg *[]*vo.MsgData) error {
	a, err := Do("ZRANGEBYSCORE", OFFLINE_KARMA_MSG+strconv.Itoa(selfid), 0, "+inf")
	if err != nil {
		return err
	}
	b, err := redis.MultiBulk(a, err)
	if err != nil {
		return err
	}
	for _, v := range b {
		s, err := redis.Bytes(v, nil)
		if err != nil {
			continue
		}
		msgData := &vo.MsgData{}
		json.Unmarshal(s, msgData)
		*msg = append(*msg, msgData)
	}
	return nil
}

//	获取跟指定好友的离线聊天记录
func (this *OfflineKarmaMsg) GetMsgByFriend(selfid int, friendid int, page int, msg *[]*vo.MsgData) error {
	a, err := Do("ZREVRANGEBYSCORE", OFFLINE_KARMA_MSG+strconv.Itoa(selfid), friendid, friendid, "LIMIT", page*PAGE_MAX, PAGE_MAX-1)
	if err != nil {
		return err
	}
	b, err := redis.MultiBulk(a, err)
	if err != nil {
		return err
	}

	for _, v := range b {
		s, err := redis.Bytes(v, nil)
		if err != nil {
			continue
		}
		msgData := &vo.MsgData{}
		json.Unmarshal(s, msgData)
		*msg = append(*msg, msgData)
		// 读取完删除
		Send("ZREM", OFFLINE_KARMA_MSG+strconv.Itoa(selfid),s)
	}

	return nil
}

