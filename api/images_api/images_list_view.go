package images_api

import (
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/common"
	"github.com/gin-gonic/gin"
)

// 图片列表

// ImagesListView 图片列表
// @Tags 图片管理
// @Summary 图片列表
// @Description	图片列表，用于显示所有的图片
// @Param data query models.PageModel false "查询参数"
// @Produce json
// @Router /api/imagesList [get]
// @Success 200 {object} res.Response{data=res.ListResponse[models.ImageModel]}
func (ImagesApi) ImagesListView(c *gin.Context) {
	var cr models.PageModel
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
	}
	//将获取列表集成为一个function，封装起来便于广泛利用
	list, count, err := common.CommonList(models.ImageModel{}, common.Option{
		PageModel: cr,
		Debug:     true,
	})

	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}

	res.OKWithList(list, count, c)
	return
}
