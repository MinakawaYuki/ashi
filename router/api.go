package router

import (
	"ashi/controller"
	"ashi/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(r *gin.Engine) {
	r.Use(gin.Recovery())
	r.Use(middleware.GetParams())
	r.StaticFS("/upload", http.Dir("./runtime/upload"))
	user(r)
	upload(r)
	character(r)
	job(r)
	weapon(r)
}

/**
用户模块router
*/
func user(r *gin.Engine) {
	var group = r.Group("/api/user")
	c := &controller.UserController{}
	group.POST("/login", c.Login) //登录
	group.Use(middleware.JWTAuth())
	{
		group.POST("/add", c.Add) //注册
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

/**
职业
*/
func job(r *gin.Engine) {
	var group = r.Group("/api/job")
	group.Use(middleware.JWTAuth())
	{
		c := &controller.JobController{}
		group.POST("/getList", c.GetList)
	}
}

/**
武器
*/
func weapon(r *gin.Engine) {
	var group = r.Group("/api/weapon")
	group.Use(middleware.JWTAuth())
	{
		c := &controller.WeaponController{}
		group.POST("/getList", c.GetList)
	}
}
