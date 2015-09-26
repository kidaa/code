/**
 * Created by Michael on 2015/8/6.
 */
package webworld
import (
	"web/proxy"
	"web/webvo"
	"db"
	"strings"
	"utils"
	"strconv"
	"csv"
	log"github.com/golang/glog"
	"vo"
)
func init() {
	proxy.Regist(21002, getWorldBuildShow)
	proxy.Regist(21003, updateBuildPart)
	proxy.Regist(21004, updateBuildCloud)
}


//世界大楼展示位置（包括大楼生长位置，云消散位置）
func getWorldBuildShow(uid int, body *[]byte) *[]byte {
	ctos := webvo.WebCtoS21002Data1{}
	ctos.Decode(body)
	stoc := webvo.WebStoC21002Data1{T:ctos.T}
	if uid<= 0 || ctos.Worldid <= 0 {
		stoc.E = 15019
	}else {
		data, ok := csv.World.Hash[ctos.Worldid]
		if !ok || data.Builds == "" {
			stoc.E = 15019
		}else {
			showData, err := db.SlaveDB.GetWorldBuildsShow(uid, strings.Split(data.Builds, ","))
			if err != nil {
				stoc.E = 15019
			}else {
				for _, value := range showData {
					log.Infoln(value.Buildid, value.Part, value.Cloud)
					data := &webvo.WebStoC21002Data2{}
					data.Buildid = value.Buildid
					data.Cloud = strings.Split(value.Cloud, ",")
					data.Part = strings.Split(value.Part, ",")
					stoc.Data = append(stoc.Data, data)
				}
			}
		}
	}
	b, _ := stoc.Encode()
	return b
}

//更新世界大楼part生长位置
func updateBuildPart(uid int, body *[]byte) *[]byte {
	ctos := webvo.WebCtoS21003Data1{}
	ctos.Decode(body)
	stoc := webvo.WebStoC21003Data1{T:ctos.T}

	if uid<= 0 || ctos.Buildid <= 0 || ctos.Partid <= 0 {
		stoc.E = 15020
	}else {
		part, err := db.SlaveDB.GetBuildPart(uid, ctos.Buildid)
		if (err != nil) {
			stoc.E = 15020
		}else {
			if (part == "") {
				stoc.E = 15021
			}else {
				partid := strconv.Itoa(ctos.Partid)
				partArr := strings.Split(part, ",");
				index := utils.SliceIndexOf(partArr, partid)
				if index == -1 {
					stoc.E = 15021
				}else {
					partArr = append(partArr[:index], partArr[index+1:]...)
					part = strings.Join(partArr, ",")
				}
			}
		}
		if (stoc.E == 0) {
			err := db.MasterDB.UpdateBuildPart(uid, ctos.Buildid, part)
			if err != nil {
				stoc.E = 15020
			}else {
				stoc.Data = ctos.Partid
			}
		}
	}
	b, _ := stoc.Encode()
	return b
}

//更新世界大楼cloud云消散位置
func updateBuildCloud(uid int, body *[]byte) *[]byte {
	ctos := webvo.WebCtoS21004Data1{}
	ctos.Decode(body)

	stoc := webvo.WebStoC21004Data1{T:ctos.T}
	if uid <= 0 || ctos.CloudArr == nil || len(ctos.CloudArr) <= 0 {
		stoc.E = 15022
	}else {
		buildArr := make([]string, 0, len(ctos.CloudArr))
		removeClouds := map[int][]string{}
		for _, value := range ctos.CloudArr {
			if value.Buildid > 0 {
				buildArr = append(buildArr, strconv.Itoa(value.Buildid))
				removeClouds[value.Buildid] = value.Cloud
			}
		}
		if len(buildArr) <= 0 {
			stoc.E = 15022
		}else {
			cloud, err := db.SlaveDB.GetBuildCloud(uid, buildArr)
			if (err != nil || len(cloud) <= 0) {
				stoc.E = 15023
			}else {
				cloudList := make([]*vo.BuildPartCloud, 0, len(cloud))
				for _, value := range cloud {
					removeArr := removeClouds[value.Buildid]
					if removeArr != nil {
						oriArr := strings.Split(value.Cloud, ",")
						endArr := utils.SliceRemoveFormSlice(oriArr, removeClouds[value.Buildid])
						if len(oriArr) != len(endArr) {
							cloudVO := &vo.BuildPartCloud{}
							cloudVO.Buildid = value.Buildid
							cloudVO.Cloud = strings.Join(endArr, ",")
							cloudList = append(cloudList, cloudVO)
						}
					}
				}
				if len(cloudList) > 0 {
					err := db.MasterDB.UpdateBuildCloud(uid, cloudList)
					if (err != nil) {
						log.Infoln(err)
						stoc.E = 15022
					}
				}else {
					stoc.E = 15023
				}
			}
		}

	}
	b, _ := stoc.Encode()
	return b
}