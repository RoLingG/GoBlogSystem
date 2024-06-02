package routers

import "GoRoLingG/api"

func (router RouterGroup) FeedBackRouter() {
	feedbackApi := api.ApiGroupApp.FeedbackApi
	router.POST("/feedbackCreate", feedbackApi.FeedbackCreateView)
	router.GET("/feedbackList", feedbackApi.FeedbackListView)
	router.DELETE("/feedbackRemove", feedbackApi.FeedbackRemoveView)
}
