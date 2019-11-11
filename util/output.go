package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Output(c *gin.Context, resultData ...interface{}){
	if resultData != nil {
		if len(resultData) == 1 {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"error_no": 0,
				"error_msg": "",
				"result": resultData[0],
			})
		} else {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"error_no": 0,
				"error_msg": "",
				"result": resultData,
			})
		}
	} else {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"error_no": 0,
			"error_msg": "",
			"result": true,
		})
	}
}

func Exception(c *gin.Context, errNo int, errMessage ...string){
	if errMessage[0] != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_no": errNo,
			"error_msg": errMessage[0],
		})
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_no": errNo,
			"error_msg": GetErrorMessage(errNo),
		})
	}
}