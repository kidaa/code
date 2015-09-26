package webvo

import (
	"encoding/json"
)

func (this *WebCtoS13000Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebCtoS13000Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebCtoS13000Data1 struct {
	T int `json:"t"`
}

func (this *WebStoC13000Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebStoC13000Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebStoC13000Data4 struct {
	Obtain int `json:"obtain"`
}

type WebStoC13000Data5 struct {
	Request int `json:"request"`
	Chat    int `json:"chat"`
}

type WebStoC13000Data6 struct {
	Dialog int `json:"dialog"`
	Attack int `json:"attack"`
}

type WebStoC13000Data1 struct {
	T    int                `json:"t"`
	E    int                `json:"e"`
	Data *WebStoC13000Data2 `json:"data"`
}

type WebStoC13000Data2 struct {
	Total  int                `json:"total"`
	Shop   *WebStoC13000Data3 `json:"shop"`
	Prop   *WebStoC13000Data4 `json:"prop"`
	Friend *WebStoC13000Data5 `json:"friend"`
	Sys    *WebStoC13000Data6 `json:"sys"`
}

type WebStoC13000Data3 struct {
	New   int `json:"new"`
	Sales int `json:"sales"`
}
