package router

import (
	"ashi/controller"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	user(r)
}

/**
用户模块router
*/
func user(r *gin.Engine) {
	var group = r.Group("/api/user")
	{
		c := &controller.UserController{}
		group.GET("/add", c.Add)
	}
}
