package routers

import "GoRoLingG/api"

func (router RouterGroup) DiggRouter() {
	diggApi := api.ApiGroupApp.DiggApi
	router.POST("/diggArticle/:id", diggApi.DiggArticleView)
	router.POST("/diggComment/:id", diggApi.DiggCommentView)
}
