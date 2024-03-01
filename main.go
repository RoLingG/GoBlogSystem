package main

import (
	"GoRoLingG/core"
	"GoRoLingG/global"
	"GoRoLingG/routers"
)

func main() {
	//读取配置文件，main中调用InitConfig
	//fmt.Println(global.Config)
	core.InitConfig()

	//初始化日志
	global.Log = core.InitLogger()
	//global.Log.Warn("warn")
	//global.Log.Error("error")
	//global.Log.Info("info")
	//global.Log.Debug("debug")

	//连接数据库
	global.DB = core.InitGorm()
	//fmt.Println(global.DB)

	//路由连接
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("gvb_Server运行在: %s", addr) //传输路由连接log
	router.Run(addr)
}
