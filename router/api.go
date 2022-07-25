package router

import (
	"ashi/controller"
	"ashi/middleware"
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
	group.Use(middleware.JWTAuth())
	{
		c := &controller.UserController{}
		group.POST("/add", c.Add)
		group.POST("/login", c.Login)
	}
}
