package advert_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

// AdvertRemoveView 广告删除
// @Tags 广告管理
// @Summary 广告删除
// @Description	广告删除，用于批量删除广告
// @Param data body models.RemoveRequest true	"广告id列表"
// @Router /api/advertRemove [delete]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (AdvertApi) AdvertRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//批量删除
	var advertsList []models.AdvertModel
	count := global.DB.Find(&advertsList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("广告不存在", c)
		return
	}
	global.DB.Delete(&advertsList)
	res.OKWithMsg(fmt.Sprintf("共删除 %d 个广告", count), c)
}
