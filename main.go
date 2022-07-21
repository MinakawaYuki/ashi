package main

import (
	"ashi/setting"
)

func main() {
	// 初始化mysql
	setting.InitMysql()
	// 初始化redis

	// 初始化路由 服务
	setting.InitServer()
}
