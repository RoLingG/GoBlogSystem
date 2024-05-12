package routers

import "GoRoLingG/api"

func (router RouterGroup) ImagesRouter() {
	imagesApi := api.ApiGroupApp.ImagesApi
	//这种增删改查都是对应一个功能模块的路由都可以写成同relativePath，只要后面的api对应的函数不同就行
	router.POST("imagesUpload", imagesApi.ImagesUploadView)
	router.GET("imagesList", imagesApi.ImagesListView)
	router.DELETE("imagesRemove", imagesApi.ImagesRemoveView)
	router.PUT("imagesUpdate", imagesApi.ImagesUpdateView)
	router.GET("imagesNameList", imagesApi.ImageNameListView)
}
