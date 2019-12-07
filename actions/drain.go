package actions

import (
	"WeatheringWithYou_Golang/model"
	"WeatheringWithYou_Golang/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

type Drain struct {
}

func (d Drain) GetDrain() (func(ctx *gin.Context)) {
	return func(ctx *gin.Context) {
		util.Output(ctx, model.GetDrainData())
	}
}

func (d Drain) AnalyseDrain() (func(ctx *gin.Context)) {
	return func(ctx *gin.Context) {
		drainData := model.GetDrainData()
		var ch = make(chan model.DrainElevation, 10000)
		waitGroup := &sync.WaitGroup{}
		for k := range drainData {
			waitGroup.Add(1)
			go model.GetDrainElevation(waitGroup, drainData[k], ch)
		}
		waitGroup.Wait()
		close(ch)

		var drainRet []model.DrainElevation
		for data := range ch {
			drainRet = append(drainRet, data)
		}

		fmt.Println(drainRet)
		util.Output(ctx, drainRet)
	}
}