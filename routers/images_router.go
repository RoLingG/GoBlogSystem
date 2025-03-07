package routers

import (
	"GoRoLingG/api"
	"GoRoLingG/middleware"
)

func (router RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	router.POST("/imagesUpload", imagesApi.ImagesUploadView)
	router.POST("/imageUploadSingle", middleware.JwtAdmin(), imagesApi.ImageUploadSingleView)
	router.GET("/imagesList", imagesApi.ImagesListView)
	router.DELETE("/imagesRemove", imagesApi.ImagesRemoveView)
	router.PUT("/imagesUpdate", imagesApi.ImagesUpdateView)
	router.GET("/imagesNameList", imagesApi.ImageNameListView)
}
