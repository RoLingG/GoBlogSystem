package routers

import "GoRoLingG/api"

func (router RouterGroup) UserRouter() {
	userApi := api.ApiGroupApp.UserApi
	router.POST("/emailLogin", userApi.EmailLoginView)
}
