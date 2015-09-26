package weather
import (
	"web/proxy"
	"web/webvo"
	"strconv"
)




func init() {
	proxy.Regist(21005, getWeather) //
}

//  获取天气数据
func getWeather(uid int,body *[]byte) *[]byte{

	stoc:=webvo.WebStoC21005Data1{T:21005}
	stoc.E = 15016
	var cityId = ""
/*

	userdata:= &vo.UserData{}
	db.SlaveDB.GetUserData(uid,uid,userdata)

	if region,ok:= CityHash[userdata.Province];ok{
		if city,ok:= region[userdata.City];ok{
			cityId = city
		}
	}
*/

//	var cityId = ""
	city:= IP2City("119.137.58.164")
	if city != nil{
		if city.Region != ""{
			if region,ok:= CityHash[city.Region];ok{
				if city,ok:= region[city.City];ok{
					cityId = city
				}
			}
		}
	}

	if cityId != ""{
		weather:= GetWeather(cityId)
		if weather != nil{
			data:= &webvo.WebStoC21005Data2{}
			data.City = city.City
			data.Province = city.Region
			tp,_:= strconv.Atoi(weather.Tp)
			data.Temps =tp -2
			data.Tempe =tp + 3
			data.Weather = weather.Wt
			stoc.Data =data

			stoc.E = 0
		}
	}

	b,_:=stoc.Encode()
	return b
}