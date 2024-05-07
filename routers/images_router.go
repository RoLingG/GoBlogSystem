package routers

import "GoRoLingG/api"

func (router RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	router.POST("imagesUpload", imagesApi.ImagesUploadView)
	router.GET("imagesList", imagesApi.ImagesListView)
	router.DELETE("imagesRemove", imagesApi.ImagesRemoveView)
	router.PUT("imagesUpdate", imagesApi.ImagesUpdateView)
}
