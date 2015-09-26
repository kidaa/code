/**
 * Created by Michael on 2015/8/10.
 *	道具商城系统
 *
 */
package shop

import (
	"csv"
	"db"
	log "github.com/golang/glog"
	"web/proxy"
	"web/webvo"
)

func init() {
	proxy.Regist(11000, buy) //
}

//  购买商品
func buy(uid int,body *[]byte) *[]byte{
	webctos := webvo.WebCtoS11000Data1{}
	webctos.Decode(body)
	stoc := webvo.WebStoC11000Data1{}
	stoc.T = webctos.T
	userid := webctos.Userid

	if webctos.Count > 0 {
		tableShop, ok := csv.Shop.Hash[webctos.Shopid]
		if ok {

			currentcy:= db.Tb_score_stat{Account_num:userid}
			currentcy.Read()
			//一级货币购买
			if tableShop.Currency == 1 {
				if currentcy.Currency_total < tableShop.Gold*int64(webctos.Count) {
					stoc.E = 13202
				} else {
					curGold := currentcy.Currency_total - tableShop.Gold*int64(webctos.Count)
					currentcy.Currency_total = curGold
					e:=currentcy.UpdateGold()
					if e != nil {
						log.Errorln(e)
					} else {
						stoc.Data = &webvo.WebStoC11000Data2{}
						stoc.Data.Gold = int(curGold)
						stoc.Data.Coin = int(currentcy.Score_total)
						widgetID := tableShop.WidgetID
						num, _ := db.MasterDB.AddUserWidget(userid, widgetID, webctos.Count)
						d := &webvo.WebStoC11000Data3{}
						d.Count = num
						d.Widgetid = widgetID
						stoc.Data.Widget = append(stoc.Data.Widget, d)
					}
				}

			} else {// 二级货币购买

				if currentcy.Score_total < tableShop.Coin*int64(webctos.Count) {
					stoc.E = 13201
				} else {
					curCoin := currentcy.Score_total - tableShop.Coin*int64(webctos.Count)
					currentcy.Score_total = curCoin
					e:=currentcy.UpdateCoin()
					if e != nil {
						log.Errorln(e)
					} else {
						stoc.Data = &webvo.WebStoC11000Data2{}
						stoc.Data.Coin = int(curCoin)
						stoc.Data.Gold = int(currentcy.Currency_total)
						widgetID := tableShop.WidgetID
						num, _ := db.MasterDB.AddUserWidget(userid, widgetID, webctos.Count)
						d := &webvo.WebStoC11000Data3{}
						d.Count = num
						d.Widgetid = widgetID
						stoc.Data.Widget = append(stoc.Data.Widget, d)
					}
				}
			}
		} else {
			log.Errorln("没有该商品")
			stoc.E = 13201 // 没有该商品
		}
	} else {
		stoc.E = 13204 //  购买数量为0
	}

	b, _ := stoc.Encode()
	return  b
}
