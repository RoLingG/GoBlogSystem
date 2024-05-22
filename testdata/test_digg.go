package main

import (
	"GoRoLingG/core"
	"GoRoLingG/global"
	"GoRoLingG/service/redis_service"
	"fmt"
)

func main() {
	// 读取配置文件
	core.InitConfig()
	// 初始化日志
	global.Log = core.InitLogger()
	global.Redis = core.ConnectRedis()
	digg := redis_service.NewArticleDiggIndex()
	digg.Set("iDA_lo8BgM_PmuvUtu50")
	//service.Service.RedisService.Digg("iDA_lo8BgM_PmuvUtu50")
	fmt.Println(digg.Get("iDA_lo8BgM_PmuvUtu50"))
}
