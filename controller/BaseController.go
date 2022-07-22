package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiSuccess(msg string, data map[string]interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  msg,
		"data": data,
	})
}

func ApiError(msg string, data map[string]interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  msg,
		"data": data,
	})
}
