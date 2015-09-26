/**
 * Created by Michael on 2015/7/31.
 * garyburd redis 再次封裝
 * 实现基于redis的hashmap和sortedset
 */
package redis

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"strings"
	"time"
	"bytes"
	"encoding/gob"
)

var pool *redis.Pool


// 连接登陆redis 服务器
func CacheInit(server, password string) {
	pool = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func GetRedis() redis.Conn {
	return pool.Get()
}

type HashMap struct {
	Name string
}

func NewHashMap(name string) *HashMap {
	Do("PING")
	return &HashMap{name}
}

func Send(cmd string, args ...interface{}) error {
	red := GetRedis()
	defer red.Close()
	err := red.Send(cmd, args...)
	if err != nil {
		return err
	}
	return red.Flush()
}

func Do(cmd string, args ...interface{}) (interface{}, error) {
	red := GetRedis()
	defer red.Close()
	return red.Do(cmd, args...)
}

func titleCasedName(name string) string {
	newstr := make([]rune, 0)
	upNextChar := true

	for _, chr := range name {
		switch {
		case upNextChar:
			upNextChar = false
			chr -= ('a' - 'A')
		case chr == '_':
			upNextChar = true
			continue
		}

		newstr = append(newstr, chr)
	}

	return string(newstr)
}
func (this *HashMap) SetExpire(second int) error {
	return Send("EXPIRE", this.Name, second)
}

func (this *HashMap) PutObject(k string, v interface{}) error {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err:= enc.Encode(v)
	b:=buffer.Bytes()

	if err != nil {
		return err
	}
	return Send("HSET", this.Name, k, b)
}

func (this *HashMap) Put(k, v string) error {
	return Send("HSET", this.Name, k, []byte(v))
}

func (this *HashMap) GetObject(k string, clazz interface{}) error {
	b, err := redis.Bytes(Do("HGET", this.Name, k))
	if err != nil {
		return err
	}
	var buffer bytes.Buffer
	buffer.Write(b)
	dec := gob.NewDecoder(&buffer)
	return dec.Decode(clazz)
}

func (this *HashMap) Get(k string) (string, error) {
	b, err := redis.Bytes(Do("HGET", this.Name, k))
	if err != nil {
		return "", err
	}
	return string(b), nil
}

//func (orm *HashMap) ScanPK(output interface{}) *Model {
//  if reflect.TypeOf(reflect.Indirect(reflect.ValueOf(output)).Interface()).Kind() == reflect.Slice {
//      sliceValue := reflect.Indirect(reflect.ValueOf(output))
//      sliceElementType := sliceValue.Type().Elem()
//      for i := 0; i < sliceElementType.NumField(); i++ {
//          bb := reflect.ValueOf(sliceElementType.Field(i).Tag)
//          if bb.String() == "PK" {
//              orm.PrimaryKey = sliceElementType.Field(i).Name
//          }
//      }
//  } else {
//      tt := reflect.TypeOf(reflect.Indirect(reflect.ValueOf(output)).Interface())
//      for i := 0; i < tt.NumField(); i++ {
//          bb := reflect.ValueOf(tt.Field(i).Tag)
//          if bb.String() == "PK" {
//              orm.PrimaryKey = tt.Field(i).Name
//          }
//      }
//  }
//  return orm

//}

//func (this *HashMap) GetObjectList(k []string, objs []interface{}) error {
//  args := []interface{}{}
//  args = append(args, this.Name)
//  for _, v := range k {
//      args = append(args, v)
//  }
//  b, err := redis.MultiBulk(Do("HMGET", args...))
//  if err != nil {
//      return err
//  }
//  for i, v := range b {
//      bb, err := redis.Bytes(v, nil)
//      if err != nil {
//          break
//      }
//      err = json.Unmarshal(bb, objs[i])
//      if err != nil {
//          break
//      }
//  }
//  return err
//}

func (this *HashMap) PutString(k string, v string) error {
	return Send("HSET", this.Name, k, v)
}

func (this *HashMap) GetString(k string) (string, error) {
	str, err := redis.String(Do("HGET", this.Name, k))
	if err == nil {
		str = strings.Trim(str, "\"")
	}
	return str, err
}

func (this *HashMap) GetStringList(k []string) ([]string, error) {
	args := []interface{}{}
	args = append(args, this.Name)
	for _, v := range k {
		args = append(args, v)
	}
	reply, err := redis.MultiBulk(Do("HMGET", args...))
	if err != nil {
		return nil, err
	}
	var list = make([]string, 0)
	for _, v := range reply {
		s, err := redis.String(v, nil)
		if err != nil {
			break
		}
		s = strings.Trim(s, "\"")
		list = append(list, s)
	}
	return list, err
}

func (this *HashMap) MultiGet(k []string) ([]string, error) {
	args := []interface{}{}
	args = append(args, this.Name)
	for _, v := range k {
		args = append(args, v)
	}
	reply, err := redis.MultiBulk(Do("HMGET", args...))
	if err != nil {
		return nil, err
	}
	var list = make([]string, 0)
	for _, v := range reply {
		b, err := redis.Bytes(v, nil)
		if err != nil {
			break
		}
		list = append(list, string(b))
	}
	return list, err
}

func (this *HashMap) Size() (int, error) {
	return redis.Int(Do("HLEN", this.Name))
}

func (this *HashMap) Remove(k string) error {
	return Send("HDEL", this.Name, k)
}

func (this *HashMap) Exists(k string) bool {
	v, err := redis.Bool(Do("HEXISTS", this.Name, k))
	if err != nil {
		return false
	}
	return v
}

func (this *HashMap) Clear() error {
	return Send("DEL", this.Name)
}

type SortedSet struct {
	Name string
}

func NewSortedSet(name string) *SortedSet {
	Do("PING")
	return &SortedSet{name}
}

func (this *SortedSet) SetExpire(second int) error {
	return Send("EXPIRE", this.Name, second)
}

func (this *SortedSet) AddObject(score float64, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return Send("ZADD", this.Name, score, b)
}

func (this *SortedSet) Set(score float64, v string) error {
	return Send("ZADD", this.Name, score, []byte(v))
}

func (this *SortedSet) AddString(score float64, v string) error {
	return Send("ZADD", this.Name, score, v)
}

func (this *SortedSet) Size() int {
	b, err := redis.Int(Do("ZCARD", this.Name))
	if err != nil {
		return -1
	}
	return b
}

func (this *SortedSet) SizeByScore(min, max float64) int {
	b, err := redis.Int(Do("ZCOUNT", this.Name, min, max))
	if err != nil {
		return -1
	}
	return b
}

func (this *SortedSet) GetObject(index int, clazz interface{}) error {
	b, err := redis.Bytes(Do("ZRANGE", this.Name, index, index+1))
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, clazz)
	return err
}

func (this *SortedSet) Get(index int) (string, error) {
	b, err := redis.Bytes(Do("ZRANGE", this.Name, index, index+1))
	if err != nil {
		return "", err
	}
	return string(b), err
}

//func (this *SortedSet) GetObjects(clazz []interface{}, start, limit int) error {
//  b, err := redis.MultiBulk(Do("ZRANGE", this.Name, start, start+limit))
//  if err != nil {
//      return err
//  }
//  for i, v := range b {
//      bb, err := redis.Bytes(v, nil)
//      if err != nil {
//          break
//      }
//      err = json.Unmarshal(bb, &clazz[i])
//      if err != nil {
//          break
//      }
//  }
//  return err
//}

func (this *SortedSet) RemoveObject(v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return Send("ZREM", this.Name, b)
}

func (this *SortedSet) Remove(v string) error {
	return Send("ZREM", this.Name, []byte(v))
}

func (this *SortedSet) GetString(index int) (string, error) {
	str, err := redis.String(Do("ZRANGE", this.Name, index, index+1))
	if err == nil {
		str = strings.Trim(str, "\"")
	}
	return str, err
}

func (this *SortedSet) GetAllStrings() ([]string, error) {
	return this.GetStrings(0, -1)
}

func (this *SortedSet) FindAll() ([]string, error) {
	return this.Find(0, -1)
}

func (this *SortedSet) FindAllRev() ([]string, error) {
	return this.FindRev(0, -1)
}

func (this *SortedSet) GetStrings(start, limit int) ([]string, error) {
	a, err := Do("ZRANGE", this.Name, start, start+limit-1)
	if err != nil {
		return nil, err
	}
	b, err := redis.MultiBulk(a, err)
	if err != nil {
		return nil, err
	}

	var list = make([]string, 0)
	for _, v := range b {
		s, err := redis.String(v, nil)
		if err != nil {
			break
		}
		s = strings.Trim(s, "\"")
		list = append(list, s)
	}
	return list, err
}

func (this *SortedSet) Find(start, limit int) ([]string, error) {
	a, err := Do("ZRANGE", this.Name, start, start+limit-1)
	if err != nil {
		return nil, err
	}
	b, err := redis.MultiBulk(a, err)
	if err != nil {
		return nil, err
	}

	var list = make([]string, 0)
	for _, v := range b {
		b, err := redis.Bytes(v, nil)
		if err != nil {
			break
		}
		list = append(list, string(b))
	}
	return list, err
}

func (this *SortedSet) GetStringsRev(start, limit int) ([]string, error) {
	a, err := Do("ZREVRANGE", this.Name, start, start+limit-1)
	if err != nil {
		return nil, err
	}
	b, err := redis.MultiBulk(a, err)
	if err != nil {
		return nil, err
	}

	var list = make([]string, 0)
	for _, v := range b {
		s, err := redis.String(v, nil)
		if err != nil {
			break
		}
		s = strings.Trim(s, "\"")
		list = append(list, s)
	}
	return list, err
}

func (this *SortedSet) FindRev(start, limit int) ([]string, error) {
	a, err := Do("ZREVRANGE", this.Name, start, start+limit-1)
	if err != nil {
		return nil, err
	}
	b, err := redis.MultiBulk(a, err)
	if err != nil {
		return nil, err
	}

	var list = make([]string, 0)
	for _, v := range b {
		b, err := redis.Bytes(v, nil)
		if err != nil {
			break
		}
		list = append(list, string(b))
	}
	return list, err
}

func (this *SortedSet) RemoveString(v string) error {
	return Send("ZREM", this.Name, v)
}

func (this *SortedSet) RemoveRange(start, limit int) error {
	return Send("ZREMRANGEBYRANK", this.Name, start, start+limit-1)
}

func (this *SortedSet) RemoveIndex(index int) error {
	return Send("ZREMRANGEBYRANK", this.Name, index, index+1)
}

func (this *SortedSet) ObjectScore(v interface{}) int {
	b, err := json.Marshal(v)
	if err != nil {
		return -1
	}
	r, err := redis.Int(Do("ZINCRBY", this.Name, b))
	if err != nil {
		return -1
	}
	return r
}

func (this *SortedSet) StringScore(v string) int {
	r, err := redis.Int(Do("ZINCRBY", this.Name, v))
	if err != nil {
		return -1
	}
	return r
}

func (this *SortedSet) Score(v string) int {
	r, err := redis.Int(Do("ZINCRBY", this.Name, []byte(v)))
	if err != nil {
		return -1
	}
	return r
}

func (this *SortedSet) Clear() error {
	return Send("DEL", this.Name)
}
