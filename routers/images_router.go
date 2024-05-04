package routers

import "GoRoLingG/api"

func (router RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	router.POST("imagesUpload", imagesApi.ImagesUploadView)
}
