package routers

import "GoRoLingG/api"

func (router RouterGroup) TagRouter() {
	tagApi := api.ApiGroupApp.TagApi
	router.POST("tagUpload", tagApi.TagCreateView)
	router.GET("tagList", tagApi.TagListView)
	router.PUT("tagUpdate/:id", tagApi.TagUpdateView)
	router.DELETE("tagRemove", tagApi.TagRemoveView)
}
