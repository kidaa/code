/**
 * Created by Michael on 2015/8/3.
 */
package utils
import (
	"regexp"
	"net"
	"strings"
	"strconv"
	"bytes"
	"encoding/binary"
	"time"
)
// 验证是否邮箱
func EmailRegexp(mail string) bool {
	b := false
	if mail != "" {
		reg := regexp.MustCompile(`^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(\.[a-zA-Z0-9_-]+)$`)
		b = reg.FindString(mail) != ""
	}
	return b
}

// 验证是否手机
func PhoneRegexp(phone string) bool {
	b := false
	if phone != "" {
		reg := regexp.MustCompile(`^(0)?1[3|4|5|7|8][0-9]\d{8}$`)

		b = reg.FindString(phone) != ""
	}
	return b
}

// eg: t.Log((InetTontoa(3232235966).String()))
func InetTontoa(ipnr int64) net.IP {
	var bytes [4]byte
	bytes[0] = byte(ipnr & 0xFF)
	bytes[1] = byte((ipnr >> 8) & 0xFF)
	bytes[2] = byte((ipnr >> 16) & 0xFF)
	bytes[3] = byte((ipnr >> 24) & 0xFF)
	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}


// t.Log((InetTobton(net.IPv4(192,168,1,190))))
func InetTobton(ipnr net.IP) int64 {
	bits := strings.Split(ipnr.String(), ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}
// eg: t.Log((InetToaton("192.168.1.190")))
func InetToaton(ipnr string) int64 {
	bits := strings.Split(ipnr, ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum int64

	sum += int64(b0) << 24
	sum += int64(b1) << 16
	sum += int64(b2) << 8
	sum += int64(b3)

	return sum
}

/**
 * 获取字符串长度并去除首尾空格（两个英文或数字算一个字符，一个中文算一个字符）
 * @param string str
 * @return int 长度
 */
func GetTrimStrLen(str string) (string, int) {
	str = strings.Trim(str, " ")
	num := len([]rune(str)) + len([]byte(str))
	result := float64(num)/4.0
	sum := int(result + 0.99)
	return str, sum
}

/**
 * 截取字符串
 * @param string str
 * @param begin int
 * @param length int
 * @return int 长度
 */
func SubStr(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(rs[begin:end])
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


//整形转换成字节
func Int64ToBytes(n int64) []byte {
	x := int64(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt64(b []byte) int64 {
	bytesBuffer := bytes.NewBuffer(b)

	var x int64
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int64(x)
}

//切片中字符串第一个位置
func SliceIndexOf(arr []string, str string) int {
	var index int = -1
	arrlen := len(arr);
	for i := 0; i < arrlen; i++ {
		if arr[i] == str {
			index = i
			break;
		}
	}
	return index
}

//字节转换成整形
func SliceLastIndexOf(arr []string, str string) int {
	var index int = -1
	for arrlen := len(arr) - 1; arrlen > -1; arrlen-- {
		if arr[arrlen] == str {
			index = arrlen
			break;
		}
	}
	return index
}

//字节转换成整形
func SliceRemoveFormSlice(oriArr []string, removeArr []string) []string {
	//	endArr := make([]string, 0, 10)
	//	for i:=0; i < arrlen; i++ {
	//		index := SliceIndexOf(removeArr, oriArr[i])
	//		if index == -1{
	//			endArr = append(endArr,  oriArr[i])
	//		}
	//	}
	//	return endArr

	endArr := oriArr[:]
	for _,value:= range removeArr{
		index := SliceIndexOf(endArr, value)
		if(index != -1){
			endArr = append(endArr[:index],endArr[index + 1:]...)
		}
	}
	return endArr

}

func TimeToHeadphpoto(t int64, userid int) (string, string) {

	var str string
	ti := time.Unix(t, 0)
	str = ti.Format("2006/01/02/15")
	//	http://image.wa.com/headpic/2015/05/26/01/60434/130_60434.jpg?1436788251

	str = "./headpic/" + str +  "/" + strconv.Itoa(userid)
	name := "/130_" +  strconv.Itoa(userid) + ".jpg"
	return str, name

}