package routers

import (
	"GoRoLingG/api"
	"GoRoLingG/middleware"
)

func (router RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	router.POST("/emailLogin", userApi.EmailLoginView)
	router.GET("/userList", middleware.JwtAuth(), userApi.UserListView)
	router.PUT("/userUpdateAdmin", middleware.JwtAdmin(), userApi.UserUpdateAdminView)
	router.PUT("/userUpdatePassword", middleware.JwtAuth(), userApi.UserUpdatePasswordView)
	router.POST("/userLogout", middleware.JwtAuth(), userApi.UserLogoutView)
	router.DELETE("/userRemove", middleware.JwtAdmin(), userApi.UserRemove)
	router.POST("/userBindEmail", middleware.JwtAuth(), userApi.UserBindEmail)
}
