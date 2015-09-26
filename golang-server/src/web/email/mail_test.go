package email


import (
	"testing"

)

//
func Test(t *testing.T) {
	code:=GenerateIdentifyCode()
	Send("444153452@qq.com",code,60229)
}