package images_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
)

type ImageResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

// 图片列表精简版

// ImageNameListView 图片列表精简版
// @Tags 图片管理
// @Summary 图片列表精简查询
// @Description	图片列表，用于精简显示所有的图片数据
// @Produce json
// @Router /api/imagesNameList [get]
// @Success 200 {object} res.Response{data=[]ImageResponse}
func (ImagesApi) ImageNameListView(c *gin.Context) {
	var imageList []ImageResponse
	global.DB.Debug().Model(models.ImageModel{}).Select("id", "name", "path").Scan(&imageList)
	res.OKWithData(imageList, c)
	return
}
