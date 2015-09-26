package vo

import (
	"encoding/json"
)

func (this *CtoS6001Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS6001Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS6001Data1 struct {
	T   int    `json:"t"`
	Id  int    `json:"id"`
	Msg string `json:"msg"`
}

func (this *StoC6001Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC6001Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC6001Data1 struct {
	T int `json:"t"`
	E int `json:"e"`
}

func (this *StoC6002Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC6002Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC6002Data1 struct {
	T    int            `json:"t"`
	E    int            `json:"e"`
	Data *StoC6002Data2 `json:"data"`
}

type StoC6002Data2 struct {
	Msg      string `json:"msg"`
	Nickname string `json:"nickname"`
	Id       int    `json:"id"`
}

func (this *StoC6004Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC6004Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC6004Data1 struct {
	T int `json:"t"`
	E int `json:"e"`
}

func (this *CtoS6010Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS6010Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS6010Data1 struct {
	T int `json:"t"`
}

func (this *StoC6010Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC6010Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC6010Data1 struct {
	T    int              `json:"t"`
	E    int              `json:"e"`
	Data []*StoC6010Data2 `json:"data"`
}

type StoC6010Data2 struct {
	Widgetid int `json:"widgetid"`
	Endtime  int `json:"endtime"`
}

func (this *CtoS6011Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS6011Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS6011Data1 struct {
	T        int `json:"t"`
	Widgetid int `json:"widgetid"`
	Userid   int `json:"userid"`
}

func (this *StoC6011Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC6011Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC6011Data1 struct {
	T    int            `json:"t"`
	E    int            `json:"e"`
	Data *StoC6011Data2 `json:"data"`
}

type StoC6011Data2 struct {
	Coin       int `json:"coin"`
	Userid     int `json:"userid"`
	Widgetid   int `json:"widgetid"`
	Success    int `json:"success"`
	Defensedid int `json:"defensedid"`
}

func (this *CtoS6012Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS6012Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS6012Data1 struct {
	T        int `json:"t"`
	Widgetid int `json:"widgetid"`
}

func (this *StoC6012Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC6012Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC6012Data1 struct {
	T    int            `json:"t"`
	E    int            `json:"e"`
	Data *StoC6012Data2 `json:"data"`
}

type StoC6012Data2 struct {
	Widgetid int `json:"widgetid"`
	Endtime  int `json:"endtime"`
}

func (this *StoC6013Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC6013Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC6013Data1 struct {
	T    int            `json:"t"`
	E    int            `json:"e"`
	Data *StoC6013Data2 `json:"data"`
}

type StoC6013Data2 struct {
	Userid   int    `json:"userid"`
	Nickname string `json:"nickname"`
	Widgetid int    `json:"widgetid"`
}

func (this *StoC6014Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC6014Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC6014Data1 struct {
	T    int            `json:"t"`
	E    int            `json:"e"`
	Data *StoC6014Data2 `json:"data"`
}

type StoC6014Data2 struct {
	Userid   int    `json:"userid"`
	Nickname string `json:"nickname"`
	Widgetid int    `json:"widgetid"`
	Coin     int    `json:"coin"`
}

func (this *StoC6015Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC6015Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC6015Data1 struct {
	T    int `json:"t"`
	E    int `json:"e"`
	Data int `json:"data"`
}

func (this *CtoS6016Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS6016Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS6016Data1 struct {
	T int `json:"t"`
}

func (this *StoC6016Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC6016Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC6016Data1 struct {
	T    int `json:"t"`
	E    int `json:"e"`
	Data int `json:"data"`
}

func (this *CtoS6017Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS6017Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS6017Data1 struct {
	T int `json:"t"`
}

func (this *StoC6017Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC6017Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC6017Data1 struct {
	T    int              `json:"t"`
	E    int              `json:"e"`
	Data []*StoC6017Data2 `json:"data"`
}

type StoC6017Data2 struct {
	Attackuser   int `json:"attackuser"`
	Attackprop   int `json:"attackprop"`
	Defensedprop int `json:"defensedprop"`
	Changenum    int `json:"changenum"`
	Changetype   int `json:"changetype"`
	Status       int `json:"status"`
	Createtime   int `json:"createtime"`
}

func (this *CtoS6023Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS6023Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS6023Data1 struct {
	T        int    `json:"t"`
	Userid   int    `json:"userid"`
	Widgetid int    `json:"widgetid"`
	Message  string `json:"message"`
}

func (this *StoC6023Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC6023Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC6023Data1 struct {
	T    int            `json:"t"`
	E    int            `json:"e"`
	Data *StoC6023Data2 `json:"data"`
}

type StoC6023Data2 struct {
	Userid   int    `json:"userid"`
	Widgetid int    `json:"widgetid"`
	Sex      int    `json:"sex"`
	Nickname string `json:"nickname"`
	Message  string `json:"message"`
	Headpic  int    `json:"headpic"`
}

func (this *StoC6024Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC6024Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC6024Data1 struct {
	T    int            `json:"t"`
	E    int            `json:"e"`
	Data *StoC6024Data2 `json:"data"`
}

type StoC6024Data2 struct {
	Userid   int    `json:"userid"`
	Widgetid int    `json:"widgetid"`
	Sex      int    `json:"sex"`
	Nickname string `json:"nickname"`
	Message  string `json:"message"`
	Headpic  int    `json:"headpic"`
}
