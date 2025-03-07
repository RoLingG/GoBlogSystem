package routers

import (
	"GoRoLingG/api"
	"GoRoLingG/middleware"
)

func (router RouterGroup) ArticleRouter() {
	articleApi := api.ApiGroupApp.ArticleApi
	//这里所有都要加用户登录中间件去检测，只是为了方便开发所以没加
	router.POST("/articleCreate", middleware.JwtAuth(), articleApi.ArticleCreateView)
	router.GET("/articleList", articleApi.ArticleListView)
	router.GET("/articleCalendar", articleApi.ArticleCalendarView)
	router.GET("/articleTagsList", articleApi.ArticleTagsListView)
	router.GET("/articleDetail/:id", articleApi.ArticleDetailView)
	router.PUT("/articleUpdate", middleware.JwtAuth(), articleApi.ArticleUpdateView)
	router.DELETE("/articleRemove", middleware.JwtAuth(), articleApi.ArticleRemoveView)
	router.POST("/articleCollect/:id", middleware.JwtAuth(), articleApi.ArticleUserCollectView)
	router.GET("/articleCollectList", middleware.JwtAuth(), articleApi.ArticleUserCollectListView)
	router.DELETE("/articleCollectRemove", middleware.JwtAuth(), articleApi.ArticleUserCollectRemoveView)
	router.GET("/articleFullTextSearch", articleApi.FullTextSearchView) //全文搜索
	router.GET("/articleCategoryList", articleApi.ArticleCategoryListView)
	router.GET("/articleContent/:id", articleApi.ArticleContentByIDView)
}
