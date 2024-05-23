package routers

import "GoRoLingG/api"

func (router RouterGroup) ChatRouter() {
	chatApi := api.ApiGroupApp.ChatApi
	router.GET("/chatGroup", chatApi.ChatGroupView)
}
