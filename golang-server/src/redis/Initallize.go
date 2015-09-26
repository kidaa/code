/**
 * Created by Michael on 2015/8/19.
 */
package redis
import (
	log"github.com/golang/glog"
	"time"
	"strconv"
	"errors"
	"github.com/garyburd/redigo/redis"
)


var UserFriendCache *FriendsCache			// 好友关系


var ChatMsg OfflineChatMsg			// 离线聊天记录
var OfflineDlg OfflineDialog
//
func CreateCache() {
	_,err:=Do("PING")
	if err != nil {
		log.Errorln(err)
	}
	UserFriendCache = &FriendsCache{HashMap{"user_friend_cache"}}

	go OfflineDlg.removeOFFLINEExpires()
}



const TmpExpires  = "TmpExpires:"
func SetTmpExpires(s string,expires int) error {
	if s != ""{
		v:= strconv.Itoa(int(time.Now().Unix()))
		err := Send("SET", TmpExpires+s, v)
		if err != nil {
			return err
		}

		err = Send("EXPIRE", TmpExpires+s, expires)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("session nill")
}

func GetTmpExpires(s string) error {
	if s != ""{
		_,err:=redis.String(Do("GET", TmpExpires+s))
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("session nill")
}


func SetTmpExpiresNum(s string,num int,expires int64) error {
	if s != ""{
		err := Send("SET", TmpExpires+s, num)
		if err != nil {
			return err
		}

		err = Send("EXPIREAT", TmpExpires+s, expires)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("session nill")
}

func GetTmpExpiresNum(s string) (num int, e error) {
	if s != ""{
		total,err:=redis.Int(Do("GET", TmpExpires+s))
		if err != nil {
			return 0,err
		}
		return total,nil
	}
	return 0,errors.New("session nill")
}


