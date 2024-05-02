package routers

import (
	"GoRoLingG/global"
	"github.com/gin-gonic/gin"
)

//方法一：无RouterGroup这个struct

// 方法二
//type RouterGroup struct {
//}

// 方法三
type RouterGroup struct {
	//*gin.Engine
	*gin.RouterGroup
}

// 路由初始化,建议每一个模块的路由都新建一个文件
func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	//router.GET("", func(c *gin.Context) {
	//	c.String(200, "测试用连接")
	//})

	//每一个模块都是go的一个文件，明了
	//方法一
	//SettingsRouter(router)

	//方法二
	//routerGroup := RouterGroup{}
	////系统配置(settings)api
	//routerGroup.SettingsRouter(router)

	//方法三
	//如果要分组,则要router.Group()
	apiRouterGroup := router.Group("/api")
	//xxxGroup := router.Group("/xxx")

	routerGroup := RouterGroup{apiRouterGroup}
	//系统配置(settings)api
	routerGroup.SettingsRouter()

	return router
}
