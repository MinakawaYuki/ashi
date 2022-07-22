package main

import (
	"ashi/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	gin.SetMode("debug")
	router.Register(r)
	r.Run(":10490")
}
