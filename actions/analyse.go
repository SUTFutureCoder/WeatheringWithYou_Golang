package actions

import (
	"WeatheringWithYou_Golang/model"
	"WeatheringWithYou_Golang/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type Analyse struct {
}

func (o *Analyse) AnalysePoint() (func(ctx *gin.Context)) {
	return func(ctx *gin.Context) {
		// STEP1 对输入值获取tile
		var swLng = ctx.Query("sw_lng")
		var swLat = ctx.Query("sw_lat")
		var neLng = ctx.Query("ne_lng")
		var neLat = ctx.Query("ne_lat")
		var distance = ctx.Param("distance")

		floatSwLng, _ := strconv.ParseFloat(swLng, 64)
		floatSwLat, _ := strconv.ParseFloat(swLat, 64)
		floatNeLng, _ := strconv.ParseFloat(neLng, 64)
		floatNeLat, _ := strconv.ParseFloat(neLat, 64)

		nwX, nwY := util.Latlng2Tile(floatNeLat, floatSwLng, 15)
		seX, seY := util.Latlng2Tile(floatSwLat, floatNeLng, 15)

		// STEP2 计算每块数据量
		// 东京tile总量为864 - 109，点总量为49,414,144，因为平均分布，可以估算出数据量。
		analysePointsSum := int(float64((seY - nwY) * (seX - nwX)) / (864 - 109) * 49414144)
		// 假设前端处理能力为10w点
		analysePointsSlice := analysePointsSum / 100000

		// 如结果为30，则范围内，每30点聚合一个

		// 估算返回值总量除以聚合量，按照5k返回进行多协程请求
		routineNum := analysePointsSum / analysePointsSlice / 5000


		fmt.Println(analysePointsSlice)
		fmt.Println(routineNum)
		fmt.Println(distance)

		var ch = make(chan model.Points, 100)

		// STEP3

		// STEP4 请求opensearch获取
		// 对区块进行分割，进行随机处理
		go model.AnalysePoints(ch, floatNeLng, floatSwLat, floatSwLng, floatNeLat, 2,2)

		fmt.Println(<- ch)
		util.Output(ctx, "{\"goodjob\"}")
	}
}