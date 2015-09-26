/*
*
* 通讯协议处理，主要处理封包和解包的过程
*
*/
package socket

import (
	"bytes"
	"encoding/binary"
	log"github.com/golang/glog"
	"errors"
)

const (
	Header    = "wwawo"     // 包头
	HeaderLen = len(Header) //　包头长度
	DataLen   = 4           //包信息数据长度占位长度
	HANDDLen  = HeaderLen + DataLen
)

//封包
func Packet(message []byte) []byte {
	return append(append([]byte(Header), IntToBytes(len(message))...), message...)
}

//解包
func Unpack(buffer []byte, readerChannel chan []byte) (data []byte ,e interface{}){
	defer func() {
		if err:=recover();err!= nil{
			log.Errorln(err)
			e = err
		}
	}()

	length := len(buffer)
	var i int
	for i = 0; i < length; {
		// 包头都不足
		if length < i+HANDDLen {
			break
		}
		if string(buffer[i:i+HeaderLen]) == Header {
			// 读取信息数据长度
			messageLength := BytesToInt(buffer[i+HeaderLen : i+HANDDLen])
			// 只有包头，数据不足一包
			if length < i+HANDDLen+messageLength {
				break
			}
			// 读取整包信息数据
			data := buffer[i+HANDDLen : i+HANDDLen+messageLength]
			readerChannel <- data

			i += HANDDLen + messageLength
		}else{
			return []byte{},errors.New("header error")
		}
	}
	// 刚好整包，包括2,３,4...个包
	if i == length {
		return make([]byte, 0, HeaderLen+1),nil
	}
	// 不足一包，或者残包余数
	return buffer[i:],nil
}

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}
