package routers

import "GoRoLingG/api"

func (router RouterGroup) MenuRouter() {
	menusApi := api.ApiGroupApp.MenuApi
	router.POST("/menusUpload", menusApi.MenuCreateView)
	router.GET("/menusList", menusApi.MenuListView)
	router.GET("/menusNameList", menusApi.MenuNameListView)
}
