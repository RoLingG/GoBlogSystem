package images_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
)

type ImageUpdateRequest struct {
	ID   uint   `json:"id" binding:"required" msg:"请输入图片ID"`
	Name string `json:"name" binding:"required" msg:"请输入图片名称"`
}

// ImagesUpdateView 图片更新
// @Tags 图片管理
// @Summary 图片更新
// @Description 图片更新，用于更新图片名字
// @param _ body ImageUpdateRequest true "要更新的图片id和更改后的name"
// @Router /api/imagesUpdate [put]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (ImagesApi) ImagesUpdateView(c *gin.Context) {
	var cr ImageUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var imageModel models.ImageModel
	err = global.DB.Take(&imageModel, cr.ID).Error
	if err != nil {
		res.FailWithMsg("图片不存在", c)
		return
	}
	err = global.DB.Model(&imageModel).Update("name", cr.Name).Error
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.OKWithMsg("图片名称修改成功", c)
	return
}
