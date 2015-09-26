/**
 * Created by Michael on 2015/8/19.
 *	Redis缓存离线聊天消息,聊天离线消息只存储在Redis里，数据库没有备份
 *	系统在Redis 建立存储离线聊天消息的有序列表
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
	OFFLINE_CHAT_MSG = "OfflineChatMsg:"
)

type OfflineChatMsg struct{}

//	添加跟指定好友的离线聊天记录
func (this *OfflineChatMsg) AddOfflineChatMsg(receiveid int, sendid int, msg *vo.MsgData) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return Send("ZADD", OFFLINE_CHAT_MSG+strconv.Itoa(receiveid), sendid, data)
}

//	获取跟指定好友的离线聊天记录数量
func (this *OfflineChatMsg) GetOfflineChatMsgCountByFriendid(selfid int, friendid int) (int, error) {
	count, err := Do("ZCOUNT", OFFLINE_CHAT_MSG+strconv.Itoa(selfid), friendid, friendid)
	return redis.Int(count, err)
}

//	获取用户跟所有好友的离线聊天记录数量
func (this *OfflineChatMsg) GetOfflineChatMsgCount(selfid int) (int, error) {
	count, err := Do("ZCOUNT", OFFLINE_CHAT_MSG+strconv.Itoa(selfid), "-inf", "+inf")

	return redis.Int(count, err)
}

//	获取用户和所有好友的离线聊天记录
func (this *OfflineChatMsg) GetAllOfflineChatMsg(selfid int,msg *[]*vo.MsgData) error {
	a, err := Do("ZRANGEBYSCORE", OFFLINE_CHAT_MSG+strconv.Itoa(selfid), 0, "+inf")
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
func (this *OfflineChatMsg) GetOfflineChatMsgByFriendid(selfid int, friendid int, page int, msg *[]*vo.MsgData) error {
	a, err := Do("ZREVRANGEBYSCORE", OFFLINE_CHAT_MSG+strconv.Itoa(selfid), friendid, friendid, "LIMIT", page*PAGE_MAX, PAGE_MAX-1)
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
		Send("ZREM", OFFLINE_CHAT_MSG+strconv.Itoa(selfid),s)
	}

	return nil
}
