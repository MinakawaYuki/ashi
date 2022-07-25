package controller

import (
	"ashi/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

type login struct {
	UserName string `json:"username,omitempty"`
	Password string `json:"password"`
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
	var login login
	err := c.ShouldBindJSON(&login)
	if err != nil {
		ApiError("用户名为空", map[string]interface{}{}, c)
		return
	}
	if err != nil {
		ApiError("密码为空", map[string]interface{}{}, c)
		return
	}
	data, loginErr := service.Login(login.UserName, login.Password)
	if loginErr != nil {
		ApiError(loginErr.Error(), map[string]interface{}{}, c)
		return
	}

	ApiSuccess("登录成功", data, c)
}
