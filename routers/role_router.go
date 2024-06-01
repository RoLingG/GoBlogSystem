package routers

import "GoRoLingG/api"

func (router RouterGroup) RoleRouter() {
	roleApi := api.ApiGroupApp.RoleApi
	router.GET("/roleIDList", roleApi.RoleIDListView)
}
