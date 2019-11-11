package actions

import (
	"WeatheringWithYou_Golang/util"
	"github.com/gin-gonic/gin"
)

type OpenSearch struct {

}

func (o *OpenSearch) AnalysePoint() (func(ctx *gin.Context)) {
	return func(ctx *gin.Context) {
		// STEP1 对输入值获取tile

		// STEP2 对tile进行分块

		// STEP3 计算每块数据量

		// STEP4 请求opensearch获取


		util.Output(ctx, "{\"goodjob\"}")
	}
}