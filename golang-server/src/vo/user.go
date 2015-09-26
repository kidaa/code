package vo

import (
	"encoding/json"
)

func (this *CtoS3001Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS3001Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS3001Data1 struct {
	T      int `json:"t"`
	Userid int `json:"userid"`
}

func (this *StoC3001Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC3001Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC3001Data1 struct {
	T    int            `json:"t"`
	E    int            `json:"e"`
	Data *StoC3001Data2 `json:"data"`
}

type StoC3001Data2 struct {
	Isfriend        int    `json:"isfriend"`
	Remark          string `json:"remark"`
	Account_email   string `json:"account_email"`
	Account_num     int    `json:"account_num"`
	Address         string `json:"address"`
	Birth           int    `json:"birth"`
	City_id         int    `json:"city_id"`
	Creater_ip      int    `json:"creater_ip"`
	Currency_total  int    `json:"currency_total"`
	Ens             string `json:"ens"`
	Grade           int    `json:"grade"`
	Grade_exp       int    `json:"grade_exp"`
	Ispassword      int    `json:"ispassword"`
	Headpic         int    `json:"headpic"`
	E               int    `json:"e"`
	Nickname        string `json:"nickname"`
	Phone_num       string `json:"phone_num"`
	Province_id     int    `json:"province_id"`
	Score_total     int    `json:"score_total"`
	Sex             int    `json:"sex"`
	Sign            string `json:"sign"`
	Star            int    `json:"star"`
	Terminal_type   int    `json:"terminal_type"`
	User_id         int    `json:"user_id"`
	User_type       int    `json:"user_type"`
	Validate_status int    `json:"validate_status"`
	World_id        int    `json:"world_id"`
}

func (this *CtoS3002Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS3002Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS3002Data1 struct {
	T      int `json:"t"`
	Userid int `json:"userid"`
}

func (this *StoC3002Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC3002Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC3002Data1 struct {
	T    int `json:"t"`
	E    int `json:"e"`
	Data int `json:"data"`
}
