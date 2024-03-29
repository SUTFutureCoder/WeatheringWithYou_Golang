package main

import (
	"WeatheringWithYou_Golang/actions"
	"WeatheringWithYou_Golang/util"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.Use(util.CORSMiddleware())

	analyse := new(actions.Analyse)
	drain := new(actions.Drain)
	r.POST("/analyse", analyse.AnalysePoint())
	r.POST("/getdrain", drain.GetDrain())
	r.POST("/analysedrain", drain.AnalyseDrain())

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":17565") // あまの ひな
	//r.Run(":3001")
}
