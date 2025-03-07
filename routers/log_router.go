package routers

import "GoRoLingG/api"

func (router RouterGroup) LogRouter() {
	logApi := api.ApiGroupApp.LogApi
	router.GET("/logList", logApi.LogListView)
	router.GET("/logLevelList", logApi.LogLevelListView)
	router.DELETE("/logRemove", logApi.LogRemoveListView)
}
