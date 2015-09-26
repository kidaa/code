package webvo

import (
	"encoding/json"
)

func (this *WebCtoS12001Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebCtoS12001Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebCtoS12001Data1 struct {
	T int `json:"t"`
}

func (this *WebStoC12001Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebStoC12001Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebStoC12001Data1 struct {
	T    int                  `json:"t"`
	E    int                  `json:"e"`
	Data []*WebStoC12001Data2 `json:"data"`
}

type WebStoC12001Data2 struct {
	Category string `json:"category"`
	Title    string `json:"title"`
	Url      string `json:"url"`
	Source   string `json:"source"`
}
