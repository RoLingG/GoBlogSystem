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

// InitRouter 路由初始化,建议每一个模块的路由都新建一个文件
func InitRouter() *gin.Engine {

	//Router := gin.New()
	//PublicGroup := Router.Group("")
	//if global.Config.System.Env == "dev" {
	//	PublicGroup.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	//}

	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	//设置swagger路由，让网页也能访问swagger
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	//router.GET("/login", user_api.UserApi{}.QQLoginView)

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
	//图片配置(images)api的路由
	routerGroup.ImagesRouter()
	//广告配置(advert)api的路由
	routerGroup.AdvertRouter()
	//菜单配置(menu)api的路由
	routerGroup.MenuRouter()
	//用户配置(user)api的路由
	routerGroup.UserRouter()
	//文章tag配置(tag)api的路由
	routerGroup.TagRouter()
	//消息配置(message)api的路由
	routerGroup.MessageRouter()
	//文章配置(article)api的路由
	routerGroup.ArticleRouter()
	//点赞配置(digg)api的路由
	routerGroup.DiggRouter()
	//评论配置(comment)api的路由
	routerGroup.CommentRouter()
	//新闻配置(news)api的路由
	routerGroup.NewsRouter()
	//聊天室配置(chat)api的路由
	routerGroup.ChatRouter()
	//日志配置(log)api的路由
	routerGroup.LogRouter()
	//按日统计数据配置(date)api的路由
	routerGroup.DataRouter()
	//用户权限(role)api的路由
	routerGroup.RoleRouter()
	//用户反馈(feedback)api的路由
	routerGroup.FeedBackRouter()

	return router
}
