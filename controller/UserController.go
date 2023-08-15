package controller

import (
	"ashi/service"
	"ashi/utils"
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

func (u *UserController) Login(c *gin.Context) {
	params, _ := c.Get("params")
	param := utils.AnyMapToStringMap(params.(map[string]interface{}))
	data, loginErr := service.Login(param["username"], param["password"])
	if loginErr != nil {
		ApiError(loginErr.Error(), map[string]interface{}{}, c)
		return
	}

	ApiSuccess("登录成功", data, c)
}
