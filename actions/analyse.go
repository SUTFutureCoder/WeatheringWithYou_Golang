package actions

import (
	"WeatheringWithYou_Golang/model"
	"WeatheringWithYou_Golang/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Analyse struct {
	SwLng float64 `json:"sw_lng"`
	SwLat float64 `json:"sw_lat"`
	NeLng float64 `json:"ne_lng"`
	NeLat float64 `json:"ne_lat"`
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

		nwX, nwY := util.Latlng2Tile(reqInfo.NeLat, reqInfo.SwLng, 15)
		seX, seY := util.Latlng2Tile(reqInfo.SwLat, reqInfo.NeLng, 15)

		// STEP2 计算每块数据量
		// 东京tile总量为864 - 109，点总量为49,414,144，因为平均分布，可以估算出数据量。
		analysePointsSum := int(float64((seY - nwY) * (seX - nwX)) / (864 - 109) * 49414144)
		// 假设前端处理能力为10w点
		analysePointsSlice := analysePointsSum / 100000

		// 如结果为30，则范围内，每30点聚合一个

		// 估算返回值总量除以聚合量，按照500返回进行多协程请求
		routineNum := analysePointsSum / analysePointsSlice / 500


		fmt.Println(analysePointsSum)
		fmt.Println(analysePointsSlice)
		fmt.Println(routineNum)

		var ch = make(chan []model.Point, 100)

		// STEP3

		// STEP4 请求opensearch获取
		// 对区块进行分割，进行随机处理
		go model.AnalysePoints(ch, reqInfo.SwLng, reqInfo.SwLat, reqInfo.NeLng, reqInfo.NeLat, 5000,1)

		var pointsRet model.Points
		point := <- ch
		pointsRet.Point = append(pointsRet.Point, point...)
		util.Output(ctx, pointsRet.Point)
	}
}