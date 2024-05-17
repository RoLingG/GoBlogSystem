package main

import (
	"GoRoLingG/core"
	_ "GoRoLingG/docs"
	"GoRoLingG/flag"
	"GoRoLingG/global"
	"GoRoLingG/routers"
)

// @title GoRoLingG API文档
// @version	1.0
// @description GoRoLingG API文档
// @host 127.0.0.01:8080
// @BasePath /
func main() {
	//读取配置文件，main中调用InitConfig
	core.InitConfig()

	//初始化日志
	global.Log = core.InitLogger()

	//连接数据库
	global.DB = core.InitGorm()
	//连接Redis
	global.Redis = core.ConnectRedis()
	//链接ES
	global.ESClient = core.ConnectES()

	//flag迁移数据库肯定是在连接数据库之后，路由连接之前
	//命令行参数绑定
	option := flag.Parse()
	//如果需要web项目停止运行，则后面的操作都不能执行，立刻停止web项目
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	//路由连接
	router := routers.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Infof("gvb_Server运行在: %s", addr) //传输路由连接log
	err := router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
