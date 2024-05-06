package images_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

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
