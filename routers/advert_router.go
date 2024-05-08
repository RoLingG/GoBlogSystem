package routers

import "GoRoLingG/api"

func (router RouterGroup) AdvertRouter() {
	advertApi := api.ApiGroupApp.AdertApi
	router.POST("advertUpload", advertApi.AdvertCreateView)
}
