/**
* Created by Michael on 2015/8/7.
*	用户数据列表
*
 */
package vo

import (
	"errors"
	"sync"
)

var Users *UserList

func init() {
	Users = NewUsers()
}

func NewUsers() (u *UserList) {
	u = new(UserList)
	u.L = map[int]*UserData{}
	return
}

type UserList struct {
	sync.RWMutex
	L map[int]*UserData
}

func (self *UserList) Add(key int, val *UserData) (err error) {
	self.Lock()
	defer self.Unlock()

	if _, ok := self.L[key]; ok {
		err = errors.New("Key exists!")
	} else {
		self.L[key] = val
	}

	return
}

func (self *UserList) Set(key int, val *UserData) (err error) {
	self.Lock()
	defer self.Unlock()

	self.L[key] = val

	return
}

func (self *UserList) Del(key int) (err error) {
	self.Lock()
	defer self.Unlock()

	if _, ok := self.L[key]; ok {
		delete(self.L, key)
	} else {
		err = errors.New("Key not exists!")
	}

	return
}

func (self *UserList) Get(key int) (info *UserData, err error) {
	self.RLock()
	defer self.RUnlock()
	var ok bool
	if info, ok = self.L[key]; ok {
	} else {
		err = errors.New("Key not exists!")
	}

	return
}
