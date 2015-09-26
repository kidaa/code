package user
import (

	"testing"
	"crypto/md5"

"encoding/hex"
)

//
func TestProtoFriendListData(t *testing.T) {

	h := md5.New()

	h.Write([]byte("123456")) // 需要加密的字符串为 123456

	t.Log( hex.EncodeToString(h.Sum(nil))) // 输出加密结果
}