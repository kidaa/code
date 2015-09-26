package vo

import (
	"encoding/json"
)

func (this *CtoS1009Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS1009Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS1009Data1 struct {
	T    int `json:"t"`
	Page int `json:"page"`
}

func (this *StoC1009Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC1009Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC1009Data1 struct {
	T    int              `json:"t"`
	E    int              `json:"e"`
	Data []*StoC1009Data2 `json:"data"`
}

type StoC1009Data2 struct {
	Userid   int    `json:"userid"`
	Nickname string `json:"nickname"`
	Remark   string `json:"remark"`
	Sex      int    `json:"sex"`
	Headpic  int    `json:"headpic"`
}

func (this *CtoS1001Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS1001Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS1001Data1 struct {
	T      int `json:"t"`
	Userid int `json:"userid"`
}

func (this *StoC1001Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC1001Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC1001Data1 struct {
	T    int              `json:"t"`
	E    int              `json:"e"`
	Data []*StoC1001Data2 `json:"data"`
}

type StoC1001Data2 struct {
	Userid   int    `json:"userid"`
	Msg      string `json:"msg"`
	Sendtime int    `json:"sendtime"`
}

func (this *CtoS1002Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS1002Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS1002Data1 struct {
	T    int `json:"t"`
	Page int `json:"page"`
}

func (this *StoC1002Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC1002Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC1002Data1 struct {
	T    int              `json:"t"`
	E    int              `json:"e"`
	Data []*StoC1002Data2 `json:"data"`
}

type StoC1002Data2 struct {
	Userid   int    `json:"userid"`
	Nickname string `json:"nickname"`
	Msg      string `json:"msg"`
	Sex      int    `json:"sex"`
	Headpic  int    `json:"headpic"`
	Status   int    `json:"status"`
}

func (this *CtoS1003Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS1003Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS1003Data1 struct {
	T    int    `json:"t"`
	Key  string `json:"key"`
	Page int    `json:"page"`
}

func (this *StoC1003Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC1003Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC1003Data1 struct {
	T    int              `json:"t"`
	E    int              `json:"e"`
	Data []*StoC1003Data2 `json:"data"`
}

type StoC1003Data2 struct {
	Userid     int    `json:"userid"`
	Nickname   string `json:"nickname"`
	Sex        int    `json:"sex"`
	Headpic    int    `json:"headpic"`
	Provinceid int    `json:"provinceid"`
	Cityid     int    `json:"cityid"`
	Sign       string `json:"sign"`
	Isfriend   int    `json:"isfriend"`
}

func (this *CtoS1004Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS1004Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS1004Data1 struct {
	T      int    `json:"t"`
	Userid int    `json:"userid"`
	Msg    string `json:"msg"`
}

func (this *StoC1004Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC1004Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC1004Data1 struct {
	T int `json:"t"`
	E int `json:"e"`
}

func (this *StoC1005Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC1005Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC1005Data1 struct {
	T    int            `json:"t"`
	E    int            `json:"e"`
	Data *StoC1005Data2 `json:"data"`
}

type StoC1005Data2 struct {
	Userid   int    `json:"userid"`
	Nickname string `json:"nickname"`
	Sex      int    `json:"sex"`
	Msg      string `json:"msg"`
	Headpic  int    `json:"headpic"`
}

func (this *StoC1011Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC1011Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC1011Data1 struct {
	T    int            `json:"t"`
	E    int            `json:"e"`
	Data *StoC1011Data2 `json:"data"`
}

type StoC1011Data2 struct {
	Userid   int    `json:"userid"`
	Nickname string `json:"nickname"`
	Sex      int    `json:"sex"`
	Msg      string `json:"msg"`
	Headpic  int    `json:"headpic"`
}

func (this *CtoS1006Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS1006Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS1006Data1 struct {
	T      int `json:"t"`
	Userid int `json:"userid"`
}

func (this *StoC1006Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC1006Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC1006Data1 struct {
	T    int            `json:"t"`
	E    int            `json:"e"`
	Data *StoC1006Data2 `json:"data"`
}

type StoC1006Data2 struct {
	Userid   int    `json:"userid"`
	Nickname string `json:"nickname"`
	Sex      int    `json:"sex"`
	Headpic  int    `json:"headpic"`
}

func (this *CtoS1007Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS1007Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS1007Data1 struct {
	T      int    `json:"t"`
	Userid int    `json:"userid"`
	Msg    string `json:"msg"`
}

func (this *StoC1007Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC1007Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC1007Data1 struct {
	T    int            `json:"t"`
	E    int            `json:"e"`
	Data *StoC1007Data2 `json:"data"`
}

type StoC1007Data2 struct {
	Msg      string `json:"msg"`
	Sendtime int    `json:"sendtime"`
}

func (this *StoC1008Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC1008Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC1008Data1 struct {
	T    int            `json:"t"`
	E    int            `json:"e"`
	Data *StoC1008Data2 `json:"data"`
}

type StoC1008Data2 struct {
	Userid   int    `json:"userid"`
	Msg      string `json:"msg"`
	Sendtime int    `json:"sendtime"`
}

func (this *CtoS1012Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS1012Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS1012Data1 struct {
	T int `json:"t"`
}

func (this *StoC1012Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC1012Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC1012Data1 struct {
	T    int `json:"t"`
	E    int `json:"e"`
	Data int `json:"data"`
}

func (this *CtoS1013Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS1013Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS1013Data1 struct {
	T int `json:"t"`
}

func (this *StoC1013Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC1013Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC1013Data1 struct {
	T    int `json:"t"`
	E    int `json:"e"`
	Data int `json:"data"`
}

func (this *CtoS1014Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS1014Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS1014Data1 struct {
	T int `json:"t"`
}

func (this *StoC1014Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC1014Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC1014Data1 struct {
	T    int `json:"t"`
	Data int `json:"data"`
	E    int `json:"e"`
}

func (this *CtoS1016Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS1016Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS1016Data1 struct {
	T        int    `json:"t"`
	Friendid int    `json:"friendid"`
	Remarks  string `json:"remarks"`
}

func (this *StoC1016Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC1016Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC1016Data1 struct {
	T int `json:"t"`
	E int `json:"e"`
}

func (this *CtoS1018Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS1018Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS1018Data1 struct {
	T        int `json:"t"`
	Friendid int `json:"friendid"`
}

func (this *StoC1018Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC1018Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC1018Data1 struct {
	T    int `json:"t"`
	E    int `json:"e"`
	Data int `json:"data"`
}

func (this *CtoS1020Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS1020Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS1020Data1 struct {
	T int `json:"t"`
}

func (this *StoC1020Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC1020Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC1020Data1 struct {
	T    int              `json:"t"`
	E    int              `json:"e"`
	Data []*StoC1020Data2 `json:"data"`
}

type StoC1020Data2 struct {
	Userid   int    `json:"userid"`
	Msgnum   int    `json:"msgnum"`
	Msg      string `json:"msg"`
	Sendtime int    `json:"sendtime"`
}
