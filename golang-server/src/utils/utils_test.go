/**
 * Created by Michael on 2015/8/4.
 */
package utils
import (
	"testing"
	"encoding/json"
	"bytes"
	"encoding/gob"
)

func TestHeader(t *testing.T) {
	d := []string{"15","17", "16", "13", "14", "7", "2", "3", "4", "1", "5", "8", "9", "10", "11", "12"}
	b := []string{"15","17", "16", "13"}
	e := SliceRemoveFormSlice(d,b)
	t.Log(e)
	t.Log(len(d),len(e), len(d) == len(e))
//	SliceRemoveFormSlice([]string("15", "17", "16", "13", "14", "7", "2", "3", "4", "1", "5", "8", "9", "10", "11", "12"),[]string("15", "17", "16", "13"))
//	t.Log((InetToaton("192.168.1.190")))
//	t.Log((InetTobton(net.IPv4(192,168,1,190))))
//	t.Log((InetTontoa(3232235966).String()))
//	t.Log((GetStrLen("323223596")))

//	var aa AA
//	aa.A =1323
//

//	var bb BB= aa

//	b,_:=bb.Encode()


//	var cc AA
//	v := reflect.ValueOf(aa)
//	k	:= reflect.ValueOf(cc)
//	t.Log(aa.(BB))
	var s = []string{"ca"}
	index:= SliceIndexOf(s,"ca")
	t.Log(index)
	s = append(s[:index], s[index+1:]...)
	t.Log(s)
	type P struct {
		X, Y, Z int
		Name    string
	}

	type Q struct {
		X, Y *int32
		Name string
	}


	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)
	// Encode (send) the value.
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		t.Log("encode error:", err)
	}
	var q Q
	// Decode (receive) the value.
	err = dec.Decode(&q)
	if err != nil {
		t.Log( err)
	}
	t.Log(q)
	t.Log(q.Name, *q.X, *q.Y)


//	v := reflect.ValueOf(s)

//	t.Log(v)



//	ti:=time.Unix(1436788251,0)
	t.Log(TimeToHeadphpoto(1432603887,60121))
}

type AA struct {
	CC
	A int `json:"a"`
}

type BB interface{
	Decode(b *[]byte) error
	Encode() (*[]byte, error)
}

type CC struct {}
func (this *CC) Decode(b *[]byte) error {
	return json.Unmarshal(*b, this)
}

func (this *CC) Encode() (*[]byte, error) {
	data, err := json.Marshal(this)
	return &data, err
}