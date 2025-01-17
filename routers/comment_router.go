package routers

import (
	"GoRoLingG/api"
	"GoRoLingG/middleware"
)

func (router RouterGroup) CommentRouter() {
	commentRouter := api.ApiGroupApp.CommentApi
	router.POST("/commentCreate", middleware.JwtAuth(), commentRouter.CommentCreateView)
	router.GET("/commentList", commentRouter.CommentListView)
	router.DELETE("/commentRemove/:id", commentRouter.CommentRemoveView)
}
