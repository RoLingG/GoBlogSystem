package routers

import "GoRoLingG/api"

func (router RouterGroup) DataRouter() {
	dataApi := api.ApiGroupApp.DataApi
	router.GET("/dateLogin", dataApi.SevenLogin)
	router.GET("/dataCollect", dataApi.DataCollectView)
}
