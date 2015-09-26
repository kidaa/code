package vo

import (
	"encoding/json"
)

func (this *StoC5002Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC5002Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC5002Data1 struct {
	T    int            `json:"t"`
	Data *StoC5002Data2 `json:"data"`
	E    int            `json:"e"`
}

type StoC5002Data2 struct {
	Id string `json:"id"`
}

func (this *StoC5003Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC5003Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC5003Data1 struct {
	T    int            `json:"t"`
	Data *StoC5003Data2 `json:"data"`
	E    int            `json:"e"`
}

type StoC5003Data2 struct {
	Id     int `json:"id"`
	Hotnum int `json:"hotnum"`
}

func (this *StoC5004Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC5004Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC5004Data1 struct {
	T    int            `json:"t"`
	Data *StoC5004Data2 `json:"data"`
	E    int            `json:"e"`
}

type StoC5004Data2 struct {
	Dialogid int `json:"dialogid"`
}

func (this *CtoS5005Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS5005Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS5005Data1 struct {
	T       int `json:"t"`
	Worldid int `json:"worldid"`
}

func (this *StoC5005Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC5005Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC5005Data1 struct {
	T    int `json:"t"`
	Data int `json:"data"`
	E    int `json:"e"`
}

func (this *CtoS5006Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS5006Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS5006Data1 struct {
	T       int `json:"t"`
	Worldid int `json:"worldid"`
}

func (this *StoC5006Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC5006Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC5006Data1 struct {
	T    int `json:"t"`
	Data int `json:"data"`
	E    int `json:"e"`
}

func (this *CtoS5007Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS5007Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS5007Data1 struct {
	T int `json:"t"`
}

func (this *StoC5007Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC5007Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC5007Data1 struct {
	T    int              `json:"t"`
	Data []*StoC5007Data2 `json:"data"`
	E    int              `json:"e"`
}

type StoC5007Data2 struct {
	Worldid   int `json:"worldid"`
	Currently int `json:"currently"`
}

func (this *CtoS5010Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS5010Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS5010Data1 struct {
	T int `json:"t"`
}

func (this *StoC5010Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC5010Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC5010Data1 struct {
	T    int      `json:"t"`
	Data []string `json:"data"`
	E    int      `json:"e"`
}
