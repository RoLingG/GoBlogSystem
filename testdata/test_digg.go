package main

import (
	"GoRoLingG/core"
	"GoRoLingG/global"
	"GoRoLingG/service"
)

func main() {
	// 读取配置文件
	core.InitConfig()
	// 初始化日志
	global.Log = core.InitLogger()
	global.Redis = core.ConnectRedis()

	service.Service.RedisService.Digg("c_LxiY8Bd9paV12IAy_B")
	service.Service.RedisService.GetDiggInfo()
}
