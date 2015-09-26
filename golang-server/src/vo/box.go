package vo

import (
	"encoding/json"
)

func (this *CtoS4000Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS4000Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS4000Data1 struct {
	T     int `json:"t"`
	Boxid int `json:"boxid"`
}

func (this *StoC4000Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC4000Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC4000Data1 struct {
	T    int `json:"t"`
	E    int `json:"e"`
	Data int `json:"data"`
}

func (this *CtoS4001Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS4001Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS4001Data1 struct {
	T     int `json:"t"`
	Boxid int `json:"boxid"`
}

func (this *StoC4001Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC4001Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC4001Data1 struct {
	T    int            `json:"t"`
	E    int            `json:"e"`
	Data *StoC4001Data2 `json:"data"`
}

type StoC4001Data2 struct {
	Id       int `json:"id"`
	Num      int `json:"num"`
	Genre    int `json:"genre"`
	Count    int `json:"count"`
	Boxtotal int `json:"boxtotal"`
}
