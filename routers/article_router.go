package routers

import (
	"GoRoLingG/api"
	"GoRoLingG/middleware"
)

func (router RouterGroup) ArticleRouter() {
	articleApi := api.ApiGroupApp.ArticleApi
	router.POST("/articleCreate", middleware.JwtAuth(), articleApi.ArticleCreateView)
	router.GET("/articleList", articleApi.ArticleListView)
	router.GET("/articleCalendar", articleApi.ArticleCalendarView)
	router.GET("/articleTagsList", articleApi.ArticleTagsListView)
	router.GET("/articleDetail/:id", articleApi.ArticleDetailView)
}
