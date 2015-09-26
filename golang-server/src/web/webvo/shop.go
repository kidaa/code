package webvo

import (
	"encoding/json"
)

func (this *WebCtoS11000Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebCtoS11000Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebCtoS11000Data1 struct {
	T      int `json:"t"`
	Shopid int `json:"shopid"`
	Count  int `json:"count"`
	Userid int `json:"userid"`
}

func (this *WebStoC11000Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebStoC11000Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebStoC11000Data1 struct {
	T    int                `json:"t"`
	E    int                `json:"e"`
	Data *WebStoC11000Data2 `json:"data"`
}

type WebStoC11000Data2 struct {
	Coin   int                  `json:"coin"`
	Gold   int                  `json:"gold"`
	Widget []*WebStoC11000Data3 `json:"widget"`
}

type WebStoC11000Data3 struct {
	Widgetid int `json:"widgetid"`
	Count    int `json:"count"`
}
