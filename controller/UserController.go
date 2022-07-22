package controller

import (
	"ashi/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (u *UserController) Add(c *gin.Context) {
	err := service.AddUser(c)
	if err != nil {
		ApiError("用户创建失败", map[string]interface{}{
			"err": err.Error(),
		}, c)
		return
	}
	ApiSuccess("用户创建成功", map[string]interface{}{}, c)
}
