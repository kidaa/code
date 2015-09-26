/**
 * Created by Michael on 2015/7/29.
 *
 * 将所有CSV数据表解析映射成Golang数据结构，并全局存储
 *	获取数据以表的唯一id作为索引
 *
 */
package csv

import (
	log "github.com/golang/glog"
	"io/ioutil"
)

func init() {
	Prop = &Widgetlist{Hash: make(map[int]*TableWidget)}
	Shop = &Shoplist{Hash: make(map[int]*TableShop)}
	Suck = &SuckList{Hash: make(map[int]*TableSuck)}
	Defsuck = &DefsuckList{Hash: make(map[int]*TableDefsuck)}
	World = &WorldList{Hash: make(map[int]*TableWorld)}

	ActiveDaylist = & ActiveDayStruct{Hash: make(map[int]*ActiveDay)}
	Lv2Explist = & Lv2ExpStruct{Hash: make(map[int]*Lv2Exp)}

	Build = &BuildList{Hash: make(map[int]*TableBuild)}

	BoxOdds = &BoxOddsList{Hash: make(map[int]*TableBoxOdds)}
	BoxNum = &BoxNumList{Hash: make(map[int]*TableBoxNum)}

	Shop.shoplistParse()
	Prop.widgetlistParse()
	Suck.suckListParse()
	Defsuck.defsuckListParse()
	World.worldListParse()

	Lv2Explist.Parse()
	ActiveDaylist.Parse()

	Build.buildListParse()
	BoxOdds.boxOddsListParse()
	BoxNum.boxNumListParse()
}

type Widgetlist struct {
	Hash map[int]*TableWidget
}

var Prop *Widgetlist

func (this *Widgetlist) widgetlistParse() {
	data, err := ioutil.ReadFile("./csv/widget.csv")

	if err != nil {
		log.Errorln(err)
	}
	var list []TableWidget
	err = Unmarshal(data, &list)
	for i := 0; i < len(list); i++ {
		data := list[i]
		this.Hash[data.Id] = &data
	}
}

type Shoplist struct {
	Hash map[int]*TableShop
}

var Shop *Shoplist
func (this *Shoplist) shoplistParse() {
	data, err := ioutil.ReadFile("./csv/shop.csv")

	if err != nil {
		log.Errorln(err)
	}
	var list []TableShop
	err = Unmarshal(data, &list)
	for i := 0; i < len(list); i++ {
		data := list[i]
		this.Hash[data.Id] = &data
	}
}

type SuckList struct {
	Hash map[int]*TableSuck
}

var Suck *SuckList
func (this *SuckList) suckListParse() {
	data, err := ioutil.ReadFile("./csv/suck.csv")

	if err != nil {
		log.Errorln(err)
	}
	var list []TableSuck
	err = Unmarshal(data, &list)
	for i := 0; i < len(list); i++ {
		data := list[i]
		this.Hash[data.Id] = &data
	}
}

type DefsuckList struct {
	Hash map[int]*TableDefsuck
}

var Defsuck *DefsuckList
func (this *DefsuckList) defsuckListParse() {
	data, err := ioutil.ReadFile("./csv/defsuck.csv")

	if err != nil {
		log.Errorln(err)
	}
	var list []TableDefsuck
	err = Unmarshal(data, &list)
	for i := 0; i < len(list); i++ {
		data := list[i]
		this.Hash[data.Id] = &data
	}
}


type WorldList struct {
	Hash map[int]*TableWorld
}

var World *WorldList
func (this *WorldList) worldListParse() {
	data, err := ioutil.ReadFile("./csv/world.csv")

	if err != nil {
		log.Errorln(err)
	}
	var list []TableWorld
	err = Unmarshal(data, &list)
	for i := 0; i < len(list); i++ {
		data := list[i]
		this.Hash[data.Id] = &data
	}
}


type BuildList struct {
	Hash map[int]*TableBuild
}

var Build *BuildList
func (this *BuildList) buildListParse() {
	data, err := ioutil.ReadFile("./csv/buildPartCloud.csv")

	if err != nil {
		log.Errorln(err)
	}
	var list []TableBuild
	err = Unmarshal(data, &list)
	for i := 0; i < len(list); i++ {
		data := list[i]
		this.Hash[data.Id] = &data
	}
}

//宝箱中物品概率表
type BoxOddsList struct {
	Hash map[int]*TableBoxOdds
}
var BoxOdds *BoxOddsList
func (this *BoxOddsList) boxOddsListParse() {
	data, err := ioutil.ReadFile("./csv/boxOdds.csv")

	if err != nil {
		log.Errorln(err)
	}
	var list []TableBoxOdds
	err = Unmarshal(data, &list)
	for i := 0; i < len(list); i++ {
		data := list[i]
		this.Hash[data.Id] = &data
	}
}

//宝箱在世界中的原始数量表
type BoxNumList struct {
	Hash map[int]*TableBoxNum
}
var BoxNum *BoxNumList
func (this *BoxNumList) boxNumListParse() {
	data, err := ioutil.ReadFile("./csv/boxNum.csv")

	if err != nil {
		log.Errorln(err)
	}
	var list []TableBoxNum
	err = Unmarshal(data, &list)
	for i := 0; i < len(list); i++ {
		data := list[i]
		this.Hash[data.Id] = &data
	}
}

// 道具表结构
type TableWidget struct {
	Id       int
	Category int
	Sort     int
	Name     string
	Icon     string
	MaxCount int
	Desc     string
	Tag      int
	Avatar   int
	Use      int
	PropType int
}

// 商城表结构
type TableShop struct {
	Id       int
	WidgetID int
	Category int
	Sort     int
	Name     string
	Icon     string
	Desc     string
	Discount int
	Limited  int
	Online   int
	New      int
	Hot      int
	Gold     int64
	Coin     int64
	Currency int
}

// 吸分道具表结构
type TableSuck struct {
	Id       	int
	Grade		int
	Lowerlimit	int
	Toplimit	int
}

// 护体道具表结构
type TableDefsuck struct {
	Id       int
	Grade	 int
	Def		 int
	Duration int
}

// 世界相关数据表结构
type TableWorld struct {
	Id       		int
	Name	 		string
	BuildCount		int
	Builds			string
	Sort 			int
	PCName			string
	PhoneName		string
	Icon			string
	Status			int
	Default			int
}

// 大楼部位、云数据表结构
type TableBuild struct {
	Id       		int
	Part	 		string
	Cloud			string
}


var Lv2Explist *Lv2ExpStruct
type Lv2ExpStruct struct {
	Hash map[int]*Lv2Exp
}
//经验等级映射表
type Lv2Exp struct {
	Id 		int
	Lv      int
	Exp	 	float32
}

// 活跃天数获取经验兑换表
type ActiveDay struct {
	Id int
	Day       int
	Exp	 float32
}

//宝箱中物品概率表
type TableBoxOdds struct  {
	Id			int
	Widget		string
	WidgetEach	string
	WidgetOdds	int
	Goods		string
	GoodsEach	string
	GoddsOdds	int
}

//宝箱在世界中的原始数量表
type TableBoxNum struct  {
	Id			int
	BaseNum		int
	EachLevel	int
	EachNum		int
}

var ActiveDaylist *ActiveDayStruct
type ActiveDayStruct struct {
	Hash map[int]*ActiveDay
}

func (this *Lv2ExpStruct) Parse() {
	data, err := ioutil.ReadFile("./csv/exp2lv.csv")

	if err != nil {
		log.Errorln(err)
	}
	var list []Lv2Exp
	err = Unmarshal(data, &list)
	for i := 0; i < len(list); i++ {
		data := list[i]
		this.Hash[data.Lv] = &data
	}
}

func (this *ActiveDayStruct) Parse() {
	data, err := ioutil.ReadFile("./csv/activeday.csv")

	if err != nil {
		log.Errorln(err)
	}
	var list []ActiveDay
	err = Unmarshal(data, &list)
	for i := 0; i < len(list); i++ {
		data := list[i]
		this.Hash[data.Day] = &data
	}
}


