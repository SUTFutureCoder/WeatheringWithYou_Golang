package actions

import (
	"WeatheringWithYou_Golang/model"
	"WeatheringWithYou_Golang/util"
	"github.com/gin-gonic/gin"
)

type Drain struct {
}

func (d Drain) GetDrain() (func(ctx *gin.Context)) {
	return func(ctx *gin.Context) {
		util.Output(ctx, model.GetDrainData())
	}
}