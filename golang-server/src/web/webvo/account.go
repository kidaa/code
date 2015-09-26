package webvo

import (
	"encoding/json"
)

func (this *WebCtoS10000Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebCtoS10000Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebCtoS10000Data1 struct {
	T       int    `json:"t"`
	Account string `json:"account"`
	Pwd     string `json:"pwd"`
}

func (this *WebStoC10000Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebStoC10000Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebStoC10000Data1 struct {
	T    int                `json:"t"`
	E    int                `json:"e"`
	Data *WebStoC10000Data2 `json:"data"`
}

type WebStoC10000Data2 struct {
	Userid     int    `json:"userid"`
	Servertime int    `json:"servertime"`
	Ens        string `json:"ens"`
}

func (this *WebCtoS10001Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebCtoS10001Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebCtoS10001Data1 struct {
	T           int    `json:"t"`
	PhoneORmail string `json:"phoneORmail"`
	Pngcode     string `json:"pngcode"`
}

func (this *WebStoC10001Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebStoC10001Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebStoC10001Data1 struct {
	T int `json:"t"`
	E int `json:"e"`
}

func (this *WebCtoS10004Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebCtoS10004Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebCtoS10004Data1 struct {
	T    int    `json:"t"`
	Code string `json:"code"`
}

func (this *WebStoC10004Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebStoC10004Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebStoC10004Data1 struct {
	T int `json:"t"`
	E int `json:"e"`
}

func (this *WebCtoS10009Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebCtoS10009Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebCtoS10009Data1 struct {
	T      int    `json:"t"`
	Newpwd string `json:"newpwd"`
}

func (this *WebStoC10009Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebStoC10009Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebStoC10009Data1 struct {
	T int `json:"t"`
	E int `json:"e"`
}

func (this *WebCtoS10010Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebCtoS10010Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebCtoS10010Data1 struct {
	T       int `json:"t"`
	Userid  int `json:"userid"`
	Otherid int `json:"otherid"`
}

func (this *WebStoC10010Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebStoC10010Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebStoC10010Data1 struct {
	T    int                `json:"t"`
	E    int                `json:"e"`
	Data *WebStoC10010Data2 `json:"data"`
}

type WebStoC10010Data2 struct {
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

func (this *WebCtoS10011Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebCtoS10011Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebCtoS10011Data1 struct {
	T        int    `json:"t"`
	Nickname string `json:"nickname"`
	Userid   int    `json:"userid"`
}

func (this *WebStoC10011Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebStoC10011Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebStoC10011Data1 struct {
	T int `json:"t"`
	E int `json:"e"`
}

func (this *WebCtoS10012Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebCtoS10012Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebCtoS10012Data1 struct {
	T      int `json:"t"`
	Sex    int `json:"sex"`
	Userid int `json:"userid"`
}

func (this *WebStoC10012Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebStoC10012Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebStoC10012Data1 struct {
	T int `json:"t"`
	E int `json:"e"`
}

func (this *WebCtoS10013Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebCtoS10013Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebCtoS10013Data1 struct {
	T          int `json:"t"`
	Cityid     int `json:"cityid"`
	Provinceid int `json:"provinceid"`
	Userid     int `json:"userid"`
}

func (this *WebStoC10013Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebStoC10013Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebStoC10013Data1 struct {
	T int `json:"t"`
	E int `json:"e"`
}

func (this *WebCtoS10014Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebCtoS10014Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebCtoS10014Data1 struct {
	T      int    `json:"t"`
	Sign   string `json:"sign"`
	Userid int    `json:"userid"`
}

func (this *WebStoC10014Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebStoC10014Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebStoC10014Data1 struct {
	T int `json:"t"`
	E int `json:"e"`
}

func (this *WebCtoS10015Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebCtoS10015Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebCtoS10015Data1 struct {
	T      int    `json:"t"`
	Pwd    string `json:"pwd"`
	Newpwd string `json:"newpwd"`
	Userid int    `json:"userid"`
}

func (this *WebStoC10015Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebStoC10015Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebStoC10015Data1 struct {
	T int `json:"t"`
	E int `json:"e"`
}

func (this *WebCtoS10018Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebCtoS10018Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebCtoS10018Data1 struct {
	T      int `json:"t"`
	Userid int `json:"userid"`
}

func (this *WebStoC10018Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebStoC10018Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebStoC10018Data1 struct {
	T    int `json:"t"`
	E    int `json:"e"`
	Data int `json:"data"`
}

func (this *WebCtoS10020Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebCtoS10020Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebCtoS10020Data1 struct {
	T       int   `json:"t"`
	Userids []int `json:"userids"`
}

func (this *WebStoC10020Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebStoC10020Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebStoC10020Data1 struct {
	T    int                  `json:"t"`
	E    int                  `json:"e"`
	Data []*WebStoC10020Data2 `json:"data"`
}

type WebStoC10020Data2 struct {
	Userid     int    `json:"userid"`
	Nickname   string `json:"nickname"`
	Provinceid int    `json:"provinceid"`
	Cityid     int    `json:"cityid"`
	Sex        int    `json:"sex"`
	Headpic    int    `json:"headpic"`
}

func (this *WebCtoS10100Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebCtoS10100Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebCtoS10100Data1 struct {
	T            int    `json:"t"`
	Account      string `json:"account"`
	Pwd          string `json:"pwd"`
	Authcode     string `json:"authcode"`
	Terminaltype int    `json:"terminaltype"`
}

func (this *WebStoC10100Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebStoC10100Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebStoC10100Data1 struct {
	T    int                `json:"t"`
	E    int                `json:"e"`
	Data *WebStoC10100Data2 `json:"data"`
}

type WebStoC10100Data2 struct {
	Userid     int `json:"userid"`
	Servertime int `json:"servertime"`
}

func (this *WebCtoS10021Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebCtoS10021Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebCtoS10021Data1 struct {
	T      int `json:"t"`
	Userid int `json:"userid"`
}

func (this *WebStoC10021Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebStoC10021Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebStoC10021Data1 struct {
	T    int                `json:"t"`
	E    int                `json:"e"`
	Data *WebStoC10021Data2 `json:"data"`
}

type WebStoC10021Data2 struct {
	Active int `json:"active"`
	Day    int `json:"day"`
	Time   int `json:"time"`
}

func (this *WebCtoS10023Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebCtoS10023Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebCtoS10023Data1 struct {
	T  int    `json:"t"`
	Id string `json:"id"`
}

func (this *WebStoC10023Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebStoC10023Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebStoC10023Data1 struct {
	T int `json:"t"`
	E int `json:"e"`
}

func (this *WebCtoS10022Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebCtoS10022Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebCtoS10022Data1 struct {
	T   int    `json:"t"`
	New string `json:"new"`
}

func (this *WebStoC10022Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebStoC10022Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebStoC10022Data1 struct {
	T int `json:"t"`
	E int `json:"e"`
}

func (this *WebCtoS10024Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebCtoS10024Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebCtoS10024Data1 struct {
	T    int    `json:"t"`
	New  string `json:"new"`
	Code string `json:"code"`
}

func (this *WebStoC10024Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebStoC10024Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebStoC10024Data1 struct {
	T int `json:"t"`
	E int `json:"e"`
}

func (this *WebCtoS10026Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebCtoS10026Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebCtoS10026Data1 struct {
	T   int    `json:"t"`
	Old string `json:"old"`
}

func (this *WebStoC10026Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebStoC10026Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebStoC10026Data1 struct {
	T int `json:"t"`
	E int `json:"e"`
}

func (this *WebCtoS10028Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebCtoS10028Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebCtoS10028Data1 struct {
	T        int    `json:"t"`
	Uid      string `json:"uid"`
	Openid   string `json:"openid"`
	Openkey  string `json:"openkey"`
	Pf       string `json:"pf"`
	Appid    string `json:"appid"`
	Nickname string `json:"nickname"`
	City     string `json:"city"`
	Province string `json:"province"`
	Sex      int    `json:"sex"`
	Headpic  string `json:"headpic"`
}

func (this *WebStoC10028Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *WebStoC10028Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type WebStoC10028Data1 struct {
	T    int `json:"t"`
	E    int `json:"e"`
	Data int `json:"data"`
}
