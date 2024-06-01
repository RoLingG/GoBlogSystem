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
	router.DELETE("/messageRecordRemove", middleware.JwtAuth(), messageApi.MessageRecordRemoveView)
	router.GET("/messageUserReceiveList", middleware.JwtAdmin(), messageApi.MessageUserReceiveList)
	router.GET("/messageUserListByMyself", middleware.JwtAuth(), messageApi.MessageUserListByMyself)
	router.GET("/messageUserListByUser", middleware.JwtAuth(), messageApi.MessageUserListByUser)
	router.GET("/messageUserRecord", middleware.JwtAuth(), messageApi.MessageUserRecordView)
	router.GET("/messageUserRecordByMyself", middleware.JwtAuth(), messageApi.MessageUserRecordByMyselfView)
}
