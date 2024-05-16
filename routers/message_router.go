package routers

import (
	"GoRoLingG/api"
	"GoRoLingG/middleware"
)

func (router RouterGroup) MessageRouter() {
	messageApi := api.ApiGroupApp.MessageApi
	router.POST("/messageCreate", middleware.JwtAuth(), messageApi.MessageCreateView)
	router.GET("/messageAdminList", middleware.JwtAdmin(), messageApi.MessageListAdminView)
	router.GET("/messageUserList", middleware.JwtAuth(), messageApi.MessageListUserView)
	router.GET("/messageRecord", middleware.JwtAuth(), messageApi.MessageRecordView)
}
