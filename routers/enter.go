package routers

import (
	"GoRoLingG/global"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
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
	//设置swagger路由，让网页也能访问swagger
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

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
	//系统配置(settings)api的路由
	routerGroup.SettingsRouter()
	//图片上传配置(images)api的路由
	routerGroup.ImagesRouter()
	//广告上传配置(advert)api的路由
	routerGroup.AdvertRouter()
	//菜单上传配置(menu)api的路由
	routerGroup.MenuRouter()

	return router
}
