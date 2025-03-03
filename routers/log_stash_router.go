package routers

import "GoRoLingG/api"

func (router RouterGroup) LogStashRouter() {
	logStashApi := api.ApiGroupApp.LogStash
	router.GET("/log_v2", logStashApi.LogListView)
	router.GET("/log_v2_read", logStashApi.LogReadView)
	router.DELETE("/log_v2_remove", logStashApi.LogRemoveView)
}
