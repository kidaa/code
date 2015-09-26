/**
 * Created by Michael on 2015/8/19.
 */
package redis

import (
	"strconv"
	. "vo"
	"errors"
)

type FriendsCache struct {
	HashMap
}

// 以用户ID设置Redis用户好友关系表
func (this *FriendsCache) SetFriendByID(userid int, hashmap *[]*Relationship) (e interface{}) {
	err := this.PutObject(strconv.Itoa(userid), hashmap)
	return err
}

// 以用户ID获取Redis用户好友关系表
func (this *FriendsCache) GetFriendByID(userid int, hashmap *[]*Relationship) (e interface{}) {
	err := this.GetObject(strconv.Itoa(userid), hashmap)
	return err
}

// 添加好友
func (this *FriendsCache) AddFriend(userid int, data *Relationship) (e interface{}) {

	if data == nil || data.FriendID == 0 {
		return errors.New("data is nil")
	}
	hashmap := make([]*Relationship, 0, 10)



	err := this.GetObject(strconv.Itoa(userid), &hashmap)

	if err != nil{
		return  errors.New("data exist")
	}
	for _, v := range hashmap {
		if v.FriendID == data.FriendID {
			return  errors.New("data exist")
		}
	}

	hashmap = append(hashmap, data)

	err = this.PutObject(strconv.Itoa(userid), &hashmap)

	return err
}

// 删除好友
func (this *FriendsCache) RemoveFriend(userid int, friendid int) (e interface{}) {

	hashmap := make([]*Relationship, 0, 10)
	err:= this.GetObject(strconv.Itoa(userid), &hashmap)

	if err != nil{
		return  errors.New("data exist")
	}

	for k, v := range hashmap {
		if v.FriendID == friendid {
			hashmap = append(hashmap[:k], hashmap[k+1:]...)
			this.PutObject(strconv.Itoa(userid), &hashmap)
			break
		}
	}

	return err
}

// 修改好友备注
func (this *FriendsCache) ModifyFriendRemark(userid int, friendid int, remark string) (e interface{}) {

	hashmap := make([]*Relationship, 0, 10)

	err := this.GetObject(strconv.Itoa(userid), &hashmap)
	if err != nil{
		return  errors.New("data exist")
	}

	for _, v := range hashmap {
		if v.FriendID == friendid {
			v.Remarks = remark

			break
		}
	}

	err = this.PutObject(strconv.Itoa(userid), &hashmap)

	return err
}
