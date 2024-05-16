package routers

import (
	"GoRoLingG/api"
	"GoRoLingG/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

/*
	会话中间件的作用
	当在路由组中使用 sessions.Sessions("sessions", store) 中间件时，对于该路由组下的所有路由：
	对于每个HTTP请求，中间件都会尝试从请求的cookie中恢复会话数据。
	如果请求中没有会话数据，中间件将创建一个新的会话，并将其存储在响应的cookie中，以便在后续请求中使用。
	可以在处理函数中通过 sessions.Default(ctx) 获取到会话实例，然后使用它来设置或获取会话值
*/

// 初始化了一个 cookie 存储，它将被中间件用来存储和检索会话数据
var store = cookie.NewStore([]byte("GFH2RG3S2DFG4D6FG1D32SF"))

func (router RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	//在Gin路由组中启用了session会话中间件
	router.Use(sessions.Sessions("sessions", store))
	router.POST("/emailLogin", userApi.EmailLoginView)
	router.GET("/userList", middleware.JwtAuth(), userApi.UserListView)
	router.PUT("/userUpdateAdmin", middleware.JwtAdmin(), userApi.UserUpdateAdminView)
	router.PUT("/userUpdatePassword", middleware.JwtAuth(), userApi.UserUpdatePasswordView)
	router.POST("/userLogout", middleware.JwtAuth(), userApi.UserLogoutView)
	router.DELETE("/userRemove", middleware.JwtAdmin(), userApi.UserRemove)
	router.POST("/userBindEmail", middleware.JwtAuth(), userApi.UserBindEmail)
}
