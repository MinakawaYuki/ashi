package router

import (
	"ashi/controller"
	"ashi/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(r *gin.Engine) {
	r.Use(gin.Recovery())
	r.StaticFS("/upload", http.Dir("./runtime/upload"))
	user(r)
	upload(r)
	character(r)
}

/**
用户模块router
*/
func user(r *gin.Engine) {
	var group = r.Group("/api/user")
	group.Use(middleware.JWTAuth())
	{
		c := &controller.UserController{}
		group.POST("/add", c.Add)     //注册
		group.POST("/login", c.Login) //登录
	}
}

/**
上传
*/
func upload(r *gin.Engine) {
	var group = r.Group("/api/upload")
	group.Use(middleware.JWTAuth())
	{
		c := &controller.UploadController{}
		group.POST("/file", c.UploadFile)   //单文件上传
		group.POST("/files", c.UploadFiles) //多文件上传
	}
}

/**
角色
*/
func character(r *gin.Engine) {
	var group = r.Group("/api/character")
	group.Use(middleware.JWTAuth())
	{
		c := &controller.CharacterController{}
		group.GET("/getById", c.GetCharacterById)
		group.POST("/getInfo", c.GetCharacter)
		group.POST("/getList", c.GetCharacterList)
	}
}
