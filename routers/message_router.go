package routers

import (
	"GoRoLingG/api"
	"GoRoLingG/middleware"
)

func (router RouterGroup) MessageRouter() {
	messageApi := api.ApiGroupApp.MessageApi
	router.POST("/messageCreate", middleware.JwtAuth(), messageApi.MessageCreateView)                        // 添加消息记录
	router.GET("/messageAdminList", middleware.JwtAdmin(), messageApi.MessageListAdminView)                  // 管理员用的消息列表
	router.GET("/messageUserList", middleware.JwtAuth(), messageApi.MessageListUserView)                     // 用户用的消息列表
	router.GET("/messageRecord", middleware.JwtAuth(), messageApi.MessageRecordView)                         // 消息记录列表
	router.DELETE("/messageRecordRemove", middleware.JwtAuth(), messageApi.MessageRecordRemoveView)          // 删除用户的消息记录
	router.GET("/messageUserReceiveList", middleware.JwtAdmin(), messageApi.MessageUserReceiveList)          // 有消息的用户列表
	router.GET("/messageUserListByMyself", middleware.JwtAuth(), messageApi.MessageUserListByMyself)         // 当前用户与其他用户的聊天列表
	router.GET("/messageUserListByUser", middleware.JwtAuth(), messageApi.MessageUserListByUser)             // 某个用户的聊天列表
	router.GET("/messageUserRecord", middleware.JwtAuth(), messageApi.MessageUserRecordView)                 // 两个用户之间的聊天记录
	router.GET("/messageUserRecordByMyself", middleware.JwtAuth(), messageApi.MessageUserRecordByMyselfView) // 当前用户与某个用户的聊天列表
}
