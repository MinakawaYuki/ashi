package controller

import (
	"ashi/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func (u *UserController) Add(c *gin.Context) {
	service.AddUser(c)
}
