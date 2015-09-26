package vo

import (
	"encoding/json"
	"testing"
)

//
func TestProtoFriendListData(t *testing.T) {

	type ProtoMsgDataToc struct {
		T    int `json:"t"`
		Data []struct {
			Userid int    `json:"userid"`
			Msg    string `json:"msg"`
		} `json:"data"`
	}

	vo := ProtoMsgDataToc{}

	var d struct {
		Userid int    `json:"userid"`
		Msg    string `json:"msg"`
	}

	d.Userid = 132132
	d.Msg = "asd"

	data := make([]struct {
		Userid int    `json:"userid"`
		Msg    string `json:"msg"`
	}, 10)
	data = data[:0]
	vo.Data = append(data, d)

	b, _ := json.Marshal(vo)

	t.Log(string(b))
}
