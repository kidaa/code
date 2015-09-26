package vo

import (
	"encoding/json"
)

func (this *CtoS6105Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS6105Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS6105Data1 struct {
	T   int    `json:"t"`
	Id  int    `json:"id"`
	Msg string `json:"msg"`
}

func (this *StoC6105Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC6105Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC6105Data1 struct {
	T int `json:"t"`
	E int `json:"e"`
}

func (this *CtoS6106Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS6106Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS6106Data1 struct {
	T int `json:"t"`
}

func (this *StoC6106Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC6106Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC6106Data1 struct {
	T    int            `json:"t"`
	E    int            `json:"e"`
	Data *StoC6106Data2 `json:"data"`
}

type StoC6106Data2 struct {
	Userid     int    `json:"userid"`
	Id         int    `json:"id"`
	Birth      int    `json:"birth"`
	Sex        int    `json:"sex"`
	Provinceid int    `json:"provinceid"`
	Cityid     int    `json:"cityid"`
	Msg        string `json:"msg"`
}

func (this *CtoS6107Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS6107Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS6107Data1 struct {
	T      int    `json:"t"`
	Msg    string `json:"msg"`
	Id     int    `json:"id"`
	Userid int    `json:"userid"`
}

func (this *StoC6107Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC6107Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC6107Data1 struct {
	T int `json:"t"`
	E int `json:"e"`
}

func (this *StoC6108Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC6108Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC6108Data1 struct {
	T    int            `json:"t"`
	Data *StoC6108Data2 `json:"data"`
}

type StoC6108Data2 struct {
	Userid   int    `json:"userid"`
	Msg      string `json:"msg"`
	Sendtime int    `json:"sendtime"`
}

func (this *CtoS6109Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS6109Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS6109Data1 struct {
	T      int    `json:"t"`
	Userid int    `json:"userid"`
	Msg    string `json:"msg"`
}

func (this *StoC6109Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC6109Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC6109Data1 struct {
	T    int            `json:"t"`
	E    int            `json:"e"`
	Data *StoC6109Data2 `json:"data"`
}

type StoC6109Data2 struct {
	Msg      string `json:"msg"`
	Sendtime int    `json:"sendtime"`
}

func (this *CtoS6110Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS6110Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS6110Data1 struct {
	T      int `json:"t"`
	Userid int `json:"userid"`
}

func (this *StoC6110Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC6110Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC6110Data1 struct {
	T int `json:"t"`
	E int `json:"e"`
}

func (this *CtoS6113Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS6113Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS6113Data1 struct {
	T    int `json:"t"`
	Page int `json:"page"`
}

func (this *StoC6113Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC6113Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC6113Data1 struct {
	T    int              `json:"t"`
	E    int              `json:"e"`
	Data []*StoC6113Data2 `json:"data"`
}

type StoC6113Data2 struct {
	Userid     int `json:"userid"`
	Birth      int `json:"birth"`
	Sex        int `json:"sex"`
	Provinceid int `json:"provinceid"`
	Cityid     int `json:"cityid"`
}

func (this *StoC6114Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC6114Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC6114Data1 struct {
	T    int            `json:"t"`
	E    int            `json:"e"`
	Data *StoC6114Data2 `json:"data"`
}

type StoC6114Data2 struct {
	Userid     int `json:"userid"`
	Birth      int `json:"birth"`
	Sex        int `json:"sex"`
	Provinceid int `json:"provinceid"`
	Cityid     int `json:"cityid"`
}

func (this *StoC6122Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC6122Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC6122Data1 struct {
	T    int              `json:"t"`
	E    int              `json:"e"`
	Data []*StoC6122Data2 `json:"data"`
}

type StoC6122Data2 struct {
	Send     int    `json:"send"`
	Receive  int    `json:"receive"`
	Msg      string `json:"msg"`
	Sendtime int    `json:"sendtime"`
}

func (this *CtoS6118Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS6118Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS6118Data1 struct {
	T int `json:"t"`
}

func (this *StoC6118Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC6118Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC6118Data1 struct {
	T    int `json:"t"`
	E    int `json:"e"`
	Data int `json:"data"`
}

func (this *CtoS6119Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS6119Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS6119Data1 struct {
	T int `json:"t"`
}

func (this *StoC6119Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC6119Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC6119Data1 struct {
	T    int              `json:"t"`
	E    int              `json:"e"`
	Data []*StoC6119Data2 `json:"data"`
}

type StoC6119Data2 struct {
	Userid   int    `json:"userid"`
	Msgnum   int    `json:"msgnum"`
	Msg      string `json:"msg"`
	Sendtime int    `json:"sendtime"`
}

func (this *CtoS6120Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS6120Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS6120Data1 struct {
	T      int `json:"t"`
	Userid int `json:"userid"`
}

func (this *StoC6120Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC6120Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC6120Data1 struct {
	T    int              `json:"t"`
	E    int              `json:"e"`
	Data []*StoC6120Data2 `json:"data"`
}

type StoC6120Data2 struct {
	Userid   int    `json:"userid"`
	Msg      string `json:"msg"`
	Sendtime int    `json:"sendtime"`
}
