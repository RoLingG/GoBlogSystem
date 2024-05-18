package routers

import "GoRoLingG/api"

func (router RouterGroup) DiggRouter() {
	diggApi := api.ApiGroupApp.DiggApi
	router.POST("/diggArticle", diggApi.DiggArticleView)
}
