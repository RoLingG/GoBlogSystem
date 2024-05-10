package routers

import "GoRoLingG/api"

func (router RouterGroup) AdvertRouter() {
	advertApi := api.ApiGroupApp.AdertApi
	router.POST("advertUpload", advertApi.AdvertCreateView)
	router.GET("advertList", advertApi.AdvertListView)
	router.PUT("advertUpdate/:id", advertApi.AdvertUpdateView)
	router.DELETE("advertRemove", advertApi.AdvertRemoveView)
}
