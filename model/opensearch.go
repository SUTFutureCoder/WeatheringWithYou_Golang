package model

import (
	"WeatheringWithYou_Golang/constant"
	"WeatheringWithYou_Golang/third_party/denverdino/aliyungo/opensearch"
	"WeatheringWithYou_Golang/util"
	"fmt"
	"github.com/bitly/go-simplejson"
	"strconv"
)
type Points struct {
	Point []Point
}

type Point struct {
	Loc string
	Elevation int
}

func AnalysePoints(ch chan Points, minLng, minLat, maxLng, maxLat float64, distCount, distTimes int) {
	keyConf, _ := util.GetConfig("key", "opensearch")
	client := opensearch.NewClient(constant.OpenSearchNetworkType ,constant.OpenSearchReigon, keyConf["accessKeyId"], keyConf["accessKeySecret"])
	query := "query=loc:'rectangle(%f %f,%f %f)'&&distinct=dist_key:groupid,dist_count:%d,dist_times:%d,reserved:false"
	searchArgs := opensearch.SearchArgs{
		Query: fmt.Sprintf(query, minLng, minLat, maxLng, maxLat, distCount, distTimes),
		Index_name: constant.OpenSearchAppId,
	}
	resp, _ := client.Search(searchArgs)

	analyse := Points{Point:nil}
	js, err := simplejson.NewJson(resp)
	if err != nil {
		fmt.Println(err)
		ch <- analyse
	}
	status, _ := js.Get("status").String()
	if status != "OK" {
		fmt.Println(err)
		ch <- analyse
	}

	arr, _ := js.Get("result").Get("items").Array()
	for i := 0; i < len(arr); i++ {
		if dataMap, ok := (arr[i]).(map[string]interface {}); ok {
			intElevation, _ := strconv.Atoi(dataMap["elevation"].(string))
			point := Point{
				Loc:       dataMap["loc"].(string),
				Elevation: intElevation,
			}
			analyse.Point = append(analyse.Point, point)
		}
	}

	ch <- analyse
}
