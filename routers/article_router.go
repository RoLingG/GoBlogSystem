package routers

import (
	"GoRoLingG/api"
	"GoRoLingG/middleware"
)

func (router RouterGroup) ArticleRouter() {
	articleApi := api.ApiGroupApp.ArticleApi
	router.POST("/articleCreate", middleware.JwtAuth(), articleApi.ArticleCreateView)
	router.GET("/articleList", articleApi.ArticleListView)
}
