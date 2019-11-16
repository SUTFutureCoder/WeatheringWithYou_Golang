package model

import (
	"WeatheringWithYou_Golang/constant"
	"WeatheringWithYou_Golang/third_party/denverdino/aliyungo/opensearch"
	"WeatheringWithYou_Golang/util"
	"fmt"
	"github.com/bitly/go-simplejson"
	"strconv"
	"strings"
	"sync"
)
type Points struct {
	Point []Point
}

type Point struct {
	Lng float64
	Lat float64
	Elevation int
}

var KeyConf map[string]string

func AnalysePoints(waitGroup *sync.WaitGroup, ch chan []Point, minLng, minLat, maxLng, maxLat float64, distCount, distTimes int) {
	defer waitGroup.Done()
	if KeyConf == nil {
		KeyConf, _ = util.GetConfig("key", "opensearch")
	}
	client := opensearch.NewClient(constant.OpenSearchNetworkType ,constant.OpenSearchReigon, KeyConf["accessKeyId"], KeyConf["accessKeySecret"])
	query := "query=loc:'rectangle(%f %f,%f %f)'&&distinct=dist_key:groupid,dist_count:%d,dist_times:%d,reserved:false&&config=start:0,hit:500"
	searchArgs := opensearch.SearchArgs{
		Query: fmt.Sprintf(query, minLng, minLat, maxLng, maxLat, distCount, distTimes),
		Index_name: constant.OpenSearchAppId,
	}
	resp, err := client.Search(searchArgs)

	var	analyse []Point
	js, err := simplejson.NewJson(resp)
	if err != nil {
		ch <- analyse
		return
	}

	status, _ := js.Get("status").String()
	if status != "OK" {
		ch <- analyse
		return
	}

	arr, _ := js.Get("result").Get("items").Array()
	for i := 0; i < len(arr); i++ {
		if dataMap, ok := (arr[i]).(map[string]interface {}); ok {
			intElevation, _ := strconv.Atoi(dataMap["elevation"].(string))
			locStrList := strings.Split(dataMap["loc"].(string), " ")
			float64Lng, _ := strconv.ParseFloat(locStrList[0], 64)
			float64Lat, _ := strconv.ParseFloat(locStrList[1], 64)
			point := Point{
				Lng: float64Lng,
				Lat: float64Lat,
				Elevation: intElevation,
			}
			analyse = append(analyse, point)
		}
	}
	ch <- analyse
	return
}
