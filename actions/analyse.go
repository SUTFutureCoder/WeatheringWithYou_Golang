package actions

import (
	"WeatheringWithYou_Golang/model"
	"WeatheringWithYou_Golang/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

type Analyse struct {
	SwLng float64 `json:"sw_lng"`
	SwLat float64 `json:"sw_lat"`
	NeLng float64 `json:"ne_lng"`
	NeLat float64 `json:"ne_lat"`
	WSlice float64 `json:"wslice"`
	HSlice float64 `json:"hslice"`
}

func (o *Analyse) AnalysePoint() (func(ctx *gin.Context)) {
	return func(ctx *gin.Context) {
		// STEP1 对输入值获取tile
		var reqInfo Analyse
		err := ctx.BindJSON(&reqInfo)
		if err != nil {
			fmt.Println(err)
		}
fmt.Println(reqInfo)
		// STEP2 计算每块数据量
		// 12:9对区块进行切割 假定前端最多处理5w点，每个各自返回500随机点数据
		var ch = make(chan []model.Point, 1000)

		// STEP3 请求opensearch获取
		// 对区块进行分割，进行随机处理
		lngSlice := (reqInfo.NeLng - reqInfo.SwLng) / reqInfo.WSlice
		latSlice := (reqInfo.NeLat - reqInfo.SwLat) / reqInfo.HSlice

		fmt.Println(lngSlice)
		fmt.Println(latSlice)

		waitGroup := &sync.WaitGroup{}
		for currLng := reqInfo.SwLng; currLng < reqInfo.NeLng; currLng += lngSlice {
			for currLat := reqInfo.SwLat; currLat < reqInfo.NeLat; currLat += latSlice {
				waitGroup.Add(1)
				go model.AnalysePoints(waitGroup, ch, currLng, currLat, currLng + lngSlice, currLat + latSlice, 500,1)
			}
		}
		fmt.Println(waitGroup)
		waitGroup.Wait()
		close(ch)

		var pointsRet model.Points
		for data := range ch{
			pointsRet.Point = append(pointsRet.Point, data...)
		}

		fmt.Println(len(pointsRet.Point))
		util.Output(ctx, pointsRet.Point)
	}
}