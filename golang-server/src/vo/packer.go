/**
 * Created by Michael on 2015/8/5.
 */
package vo
import (
	"encoding/json"
	log"github.com/golang/glog"
)


type Packer  struct{


}
func (this *Packer)Decode(b *[]byte) error {
	err:= json.Unmarshal(*b,this)
	if err != nil {
		log.Errorln(err)
	}
	return err
}

func (this *Packer)Encode()(*[]byte, error){
	data,err:=json.Marshal(this)
	if err != nil {
		log.Errorln(err)
	}
	return &data,err
}
