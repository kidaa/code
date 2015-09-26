/**
 * Created by Michael on 2015/7/29.
 */
package csv
import (

"testing"
)

//
func TestProtoFriendListData(t *testing.T) {
/*
	data, err := ioutil.ReadFile("widget.csv")

	if err != nil {
		t.Log(err)
	}
	pp := []TableWidget{}
	ppp:= make(map[int]*TableWidget )

	err = Unmarshal(data, &pp)

	for i:=0;i<len(pp);i++{
		data:= pp[i]
		ppp[data.Id] = &data
	}

	if err != nil {
		t.Log(err)
	}
*/
//
//	Prop.widgetlistParse()
//	t.Log(Prop.Hash[1100560097].Id)
//
	Lv2Explist.Parse()
	ActiveDaylist.Parse()


	t.Log(*Lv2Explist.Hash[60],*ActiveDaylist.Hash[1])


/*	data, err := ioutil.ReadFile("shop.csv")

	if err != nil {
		t.Log(err)
	}
	var list []TableShop
	err = Unmarshal(data, &list)
	for i := 0; i < len(list); i++ {
		data := list[i]
//		this.Hash[data.Id] = &data
		t.Log(data.Id,data.WidgetID)
	}*/

}