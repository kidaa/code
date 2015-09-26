package webvo

import (
	"encoding/json"
)

func (this *WebCtoS21002Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebCtoS21002Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebCtoS21002Data1 struct {
	T       int `json:"t"`
	Worldid int `json:"worldid"`
}

func (this *WebStoC21002Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebStoC21002Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebStoC21002Data1 struct {
	T    int                  `json:"t"`
	Data []*WebStoC21002Data2 `json:"data"`
	E    int                  `json:"e"`
}

type WebStoC21002Data2 struct {
	Buildid int      `json:"buildid"`
	Part    []string `json:"part"`
	Cloud   []string `json:"cloud"`
}

func (this *WebCtoS21003Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebCtoS21003Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebCtoS21003Data1 struct {
	T       int `json:"t"`
	Buildid int `json:"buildid"`
	Partid  int `json:"partid"`
}

func (this *WebStoC21003Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebStoC21003Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebStoC21003Data1 struct {
	T    int `json:"t"`
	Data int `json:"data"`
	E    int `json:"e"`
}

func (this *WebCtoS21004Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebCtoS21004Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebCtoS21004Data1 struct {
	T        int                  `json:"t"`
	Buildid  int                  `json:"buildid"`
	CloudArr []*WebCtoS21004Data2 `json:"cloudArr"`
}

type WebCtoS21004Data2 struct {
	Buildid int      `json:"buildid"`
	Cloud   []string `json:"cloud"`
}

func (this *WebStoC21004Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebStoC21004Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebStoC21004Data1 struct {
	T int `json:"t"`
	E int `json:"e"`
}

func (this *WebCtoS21005Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebCtoS21005Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebCtoS21005Data1 struct {
	T int `json:"t"`
}

func (this *WebStoC21005Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebStoC21005Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebStoC21005Data1 struct {
	T    int                `json:"t"`
	Data *WebStoC21005Data2 `json:"data"`
	E    int                `json:"e"`
}

type WebStoC21005Data2 struct {
	Province string `json:"province"`
	City     string `json:"city"`
	Temps    int    `json:"temps"`
	Tempe    int    `json:"tempe"`
	Weather  string `json:"weather"`
}
