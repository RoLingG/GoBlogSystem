package routers

import "GoRoLingG/api"

func (router RouterGroup) LogStashRouter() {
	logStashApi := api.ApiGroupApp.LogStash
	router.GET("/logV2List", logStashApi.LogListView)
	router.GET("/logV2Read", logStashApi.LogReadView)
	router.DELETE("/logV2Remove", logStashApi.LogRemoveView)
}
