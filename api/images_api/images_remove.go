package images_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

// ImagesRemoveView 图片删除
// @Tags 图片管理
// @Summary 图片删除
// @Description	图片删除，用于批量删除数据库内的图片数据
// @Param data body models.RemoveRequest true	"要删除的图片id列表"
// @Router /api/imagesRemove [delete]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (ImagesApi) ImagesRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//批量删除
	var imagesList []models.ImageModel
	count := global.DB.Find(&imagesList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("图片不存在", c)
		return
	}
	global.DB.Delete(&imagesList)
	res.OKWithMsg(fmt.Sprintf("共删除 %d 张图片", count), c)
}
