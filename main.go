package main

import (
	"GoRoLingG/core"
	_ "GoRoLingG/docs"
	"GoRoLingG/flag"
	"GoRoLingG/global"
	"GoRoLingG/routers"
	"GoRoLingG/utils"
	"strings"
)

// @title GoRoLingG API文档
// @version	1.0
// @description GoRoLingG API文档
// @host 127.0.0.01:8080
// @BasePath /
func main() {
	//读取配置文件，main中调用InitConfig
	core.InitConfig()
	//连接IP地址数据库
	core.InitAddrDB()
	defer global.AddrDB.Close() //主程序结束，则关闭IP地址数据库

	//初始化日志
	global.Log = core.InitLogger()

	//连接数据库
	global.DB = core.InitGorm()
	//连接Redis
	global.Redis = core.ConnectRedis()
	//连接ES
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
	host := global.Config.System.Host
	port := global.Config.System.Port
	addr := global.Config.System.Addr()

	if host == "0.0.0.0" {
		ipList := utils.GetIPList()
		for _, ip := range ipList {
			parts := strings.Split(ip, ".")
			if parts[0] == "169" { //过滤掉访问不了的内网ip
				continue
			}
			global.Log.Infof("gvb_Server 运行在: http://%s:%d/api", ip, port) //传输路由连接log
			global.Log.Infof("gvb_Server api文档 运行在: http://%s:%d/swagger/index.html#", ip, port)
		}
	} else {
		global.Log.Infof("gvb_Server 运行在: http://%s:%d/api", host, port) //传输路由连接log
		global.Log.Infof("gvb_Server api文档 运行在: http://%s:%d/swagger/index.html#", host, port)
	}
	err := router.Run(addr) //将路由运行到该地址下
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
