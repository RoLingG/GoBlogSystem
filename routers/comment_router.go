package routers

import (
	"GoRoLingG/api"
	"GoRoLingG/middleware"
)

func (router RouterGroup) CommentRouter() {
	commentRouter := api.ApiGroupApp.CommentApi
	router.POST("/commentCreate", middleware.JwtAuth(), commentRouter.CommentCreateView)
	router.GET("/commentList/:article_id", commentRouter.CommentListView)
	router.DELETE("/commentRemove/:id", middleware.JwtAuth(), commentRouter.CommentRemoveView)
	router.GET("/commentByArticle", middleware.JwtAuth(), commentRouter.CommentByArticleListView)
}
