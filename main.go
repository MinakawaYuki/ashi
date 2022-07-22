package main

import (
	"ashi/setting"
	"ashi/utils"
)

func main() {
	// 初始化mysql
	setting.InitMysql()
	// 初始化redis
	setting.InitRedis()
	// 初始化路由 服务
	setting.InitServer()
	// 初始化log
	var log utils.Log
	log.Init()
}
