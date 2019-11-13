package model

import (
	"WeatheringWithYou_Golang/constant"
	"WeatheringWithYou_Golang/third_party/denverdino/aliyungo/opensearch"
	"WeatheringWithYou_Golang/util"
	"fmt"
)

func AnalysePoints() {
	keyConf, _ := util.GetConfig("key", "opensearch")
	client := opensearch.NewClient(constant.OpenSearchNetworkType ,constant.OpenSearchReigon, keyConf["accessKeyId"], keyConf["accessKeySecret"])
	searchArgs := opensearch.SearchArgs{
		Query:"query=loc:'circle(139.710900 35.729446,1000)'&&aggregate=group_key:elevation,agg_fun:max(elevation)",
		Index_name: constant.OpenSearchAppId,
	}
	resp, _ := client.Search(searchArgs)
	fmt.Println(string(resp))
}
