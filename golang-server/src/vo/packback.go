package vo

import (
	"encoding/json"
)

func (this *CtoS2000Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS2000Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS2000Data1 struct {
	T int `json:"t"`
}

func (this *StoC2000Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC2000Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC2000Data1 struct {
	T    int              `json:"t"`
	Data []*StoC2000Data2 `json:"data"`
	E    int              `json:"e"`
}

type StoC2000Data2 struct {
	Id    int `json:"id"`
	Count int `json:"count"`
	Level int `json:"level"`
}

func (this *CtoS2001Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS2001Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS2001Data1 struct {
	T     int `json:"t"`
	Id    int `json:"id"`
	Count int `json:"count"`
}

func (this *StoC2001Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC2001Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC2001Data1 struct {
	T int `json:"t"`
	E int `json:"e"`
}

func (this *CtoS2002Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS2002Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS2002Data1 struct {
	T     int `json:"t"`
	Id    int `json:"id"`
	Count int `json:"count"`
}

func (this *StoC2002Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC2002Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC2002Data1 struct {
	T int `json:"t"`
	E int `json:"e"`
}

func (this *CtoS2003Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS2003Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS2003Data1 struct {
	T     int `json:"t"`
	Id    int `json:"id"`
	Count int `json:"count"`
}

func (this *StoC2003Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC2003Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC2003Data1 struct {
	T int `json:"t"`
	E int `json:"e"`
}

func (this *CtoS2004Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS2004Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS2004Data1 struct {
	T int `json:"t"`
}

func (this *StoC2004Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC2004Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC2004Data1 struct {
	T    int `json:"t"`
	E    int `json:"e"`
	Data int `json:"data"`
}

func (this *StoC2005Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC2005Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC2005Data1 struct {
	T    int            `json:"t"`
	Data *StoC2005Data2 `json:"data"`
	E    int            `json:"e"`
}

type StoC2005Data2 struct {
	Id    int `json:"id"`
	Count int `json:"count"`
}
