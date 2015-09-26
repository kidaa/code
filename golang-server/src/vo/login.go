package vo

import (
	"encoding/json"
)

func (this *CtoS1000Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS1000Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS1000Data1 struct {
	T       int    `json:"t"`
	Session string `json:"session"`
}

func (this *StoC1000Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC1000Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC1000Data1 struct {
	T    int `json:"t"`
	Data int `json:"data"`
}

func (this *StoC1010Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC1010Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC1010Data1 struct {
	T int `json:"t"`
}

func (this *CtoS1111Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CtoS1111Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type CtoS1111Data1 struct {
	T int `json:"t"`
}

func (this *StoC1111Data1) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *StoC1111Data1) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}

type StoC1111Data1 struct {
	T int `json:"t"`
	E int `json:"e"`
}
