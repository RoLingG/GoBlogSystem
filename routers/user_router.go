package routers

import (
	"GoRoLingG/api"
	"GoRoLingG/middleware"
)

func (router RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	router.POST("/emailLogin", userApi.EmailLoginView)
	router.GET("/userList", middleware.JwtAuth(), userApi.UserListView)
}
