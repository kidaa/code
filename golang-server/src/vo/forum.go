package vo

import (
	"encoding/json"
)

func (this *CtoS8004Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS8004Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS8004Data1 struct {
	T       int    `json:"t"`
	Abspath string `json:"abspath"`
	Content string `json:"content"`
}

func (this *StoC8004Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC8004Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC8004Data1 struct {
	T int `json:"t"`
	E int `json:"e"`
}

func (this *CtoS8018Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS8018Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS8018Data1 struct {
	T          int    `json:"t"`
	Dialogid   int    `json:"dialogid"`
	Dialogtype int    `json:"dialogtype"`
	Abspath    string `json:"abspath"`
	Widgetid   int    `json:"widgetid"`
}

func (this *StoC8018Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC8018Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC8018Data1 struct {
	T    int            `json:"t"`
	E    int            `json:"e"`
	Data *StoC8018Data2 `json:"data"`
}

type StoC8018Data2 struct {
	Widgetid   int    `json:"widgetid"`
	Count      int    `json:"count"`
	Dialogid   int    `json:"dialogid"`
	Dialogtype int    `json:"dialogtype"`
	Abspath    string `json:"abspath"`
	Upworth    int    `json:"upworth"`
}

func (this *CtoS8019Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS8019Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS8019Data1 struct {
	T          int    `json:"t"`
	Dialogid   int    `json:"dialogid"`
	Dialogtype int    `json:"dialogtype"`
	Abspath    string `json:"abspath"`
	Widgetid   int    `json:"widgetid"`
}

func (this *StoC8019Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC8019Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC8019Data1 struct {
	T    int            `json:"t"`
	E    int            `json:"e"`
	Data *StoC8019Data2 `json:"data"`
}

type StoC8019Data2 struct {
	Widgetid   int    `json:"widgetid"`
	Count      int    `json:"count"`
	Dialogid   int    `json:"dialogid"`
	Dialogtype int    `json:"dialogtype"`
	Abspath    string `json:"abspath"`
	Upworth    int    `json:"upworth"`
}

func (this *CtoS8020Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS8020Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS8020Data1 struct {
	T          int    `json:"t"`
	Dialogid   int    `json:"dialogid"`
	Abspath    string `json:"abspath"`
	Dialogtype int    `json:"dialogtype"`
}

func (this *StoC8020Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC8020Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC8020Data1 struct {
	T    int            `json:"t"`
	E    int            `json:"e"`
	Data *StoC8020Data2 `json:"data"`
}

type StoC8020Data2 struct {
	Dialogid   int    `json:"dialogid"`
	Dialogtype int    `json:"dialogtype"`
	Abspath    string `json:"abspath"`
	Upworth    int    `json:"upworth"`
}

func (this *CtoS8021Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS8021Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS8021Data1 struct {
	T          int    `json:"t"`
	Dialogid   int    `json:"dialogid"`
	Abspath    string `json:"abspath"`
	Dialogtype int    `json:"dialogtype"`
}

func (this *StoC8021Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC8021Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC8021Data1 struct {
	T    int            `json:"t"`
	E    int            `json:"e"`
	Data *StoC8021Data2 `json:"data"`
}

type StoC8021Data2 struct {
	Dialogid   int    `json:"dialogid"`
	Dialogtype int    `json:"dialogtype"`
	Abspath    string `json:"abspath"`
	Upworth    int    `json:"upworth"`
}

func (this *CtoS8030Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS8030Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS8030Data1 struct {
	T       int    `json:"t"`
	Abspath string `json:"abspath"`
	Page    int    `json:"page"`
}

func (this *StoC8030Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC8030Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC8030Data1 struct {
	T    int              `json:"t"`
	E    int              `json:"e"`
	Data []*StoC8030Data2 `json:"data"`
}

type StoC8030Data2 struct {
	Dialogid   int    `json:"dialogid"`
	Dialogtype int    `json:"dialogtype"`
	Abspath    string `json:"abspath"`
	Upworth    int    `json:"upworth"`
	Content    string `json:"content"`
	Time       int    `json:"time"`
	Userid     int    `json:"userid"`
	Nickname   string `json:"nickname"`
	Provinceid int    `json:"provinceid"`
	Cityid     int    `json:"cityid"`
	Sex        int    `json:"sex"`
	Headpic    int    `json:"headpic"`
}

func (this *CtoS8031Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS8031Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS8031Data1 struct {
	T       int    `json:"t"`
	Abspath string `json:"abspath"`
}

func (this *StoC8031Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC8031Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC8031Data1 struct {
	T    int              `json:"t"`
	E    int              `json:"e"`
	Data []*StoC8031Data2 `json:"data"`
}

type StoC8031Data2 struct {
	Dialogid   int    `json:"dialogid"`
	Dialogtype int    `json:"dialogtype"`
	Abspath    string `json:"abspath"`
	Upworth    int    `json:"upworth"`
	Content    string `json:"content"`
	Time       int    `json:"time"`
	Userid     int    `json:"userid"`
	Nickname   string `json:"nickname"`
	Provinceid int    `json:"provinceid"`
	Cityid     int    `json:"cityid"`
	Sex        int    `json:"sex"`
	Headpic    int    `json:"headpic"`
}

func (this *CtoS8032Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS8032Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS8032Data1 struct {
	T    int `json:"t"`
	Page int `json:"page"`
}

func (this *StoC8032Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC8032Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC8032Data1 struct {
	T    int              `json:"t"`
	E    int              `json:"e"`
	Data []*StoC8032Data2 `json:"data"`
}

type StoC8032Data2 struct {
	Abspath string `json:"abspath"`
	Upworth int    `json:"upworth"`
	Content string `json:"content"`
	Time    int    `json:"time"`
}

func (this *CtoS8034Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS8034Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS8034Data1 struct {
	T          int    `json:"t"`
	Abspath    string `json:"abspath"`
	Dialogid   int    `json:"dialogid"`
	Dialogtype int    `json:"dialogtype"`
}

func (this *StoC8034Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC8034Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC8034Data1 struct {
	T    int            `json:"t"`
	E    int            `json:"e"`
	Data *StoC8034Data2 `json:"data"`
}

type StoC8034Data2 struct {
	Abspath    string `json:"abspath"`
	Dialogid   int    `json:"dialogid"`
	Dialogtype int    `json:"dialogtype"`
}

func (this *CtoS8036Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS8036Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS8036Data1 struct {
	T        int    `json:"t"`
	Dialogid int    `json:"dialogid"`
	Otherid  int    `json:"otherid"`
	Content  string `json:"content"`
	Anno     int    `json:"anno"`
}

func (this *StoC8036Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC8036Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC8036Data1 struct {
	T int `json:"t"`
	E int `json:"e"`
}

func (this *CtoS8038Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS8038Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS8038Data1 struct {
	T        int `json:"t"`
	Dialogid int `json:"dialogid"`
	Page     int `json:"page"`
}

func (this *StoC8038Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC8038Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC8038Data1 struct {
	T    int              `json:"t"`
	E    int              `json:"e"`
	Data []*StoC8038Data2 `json:"data"`
}

type StoC8038Data2 struct {
	Dialogid  int    `json:"dialogid"`
	Userid    int    `json:"userid"`
	Nickname  string `json:"nickname"`
	Headpic   int    `json:"headpic"`
	Otherid   int    `json:"otherid"`
	Othername string `json:"othername"`
	Otherpic  int    `json:"otherpic"`
	Content   string `json:"content"`
	Time      int    `json:"time"`
	Anno      int    `json:"anno"`
}
