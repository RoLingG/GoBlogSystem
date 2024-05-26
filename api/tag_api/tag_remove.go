package tag_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

// TagRemoveView 删除标签
// @Tags 标签管理
// @Summary 删除标签
// @Description	删除现有的标签
// @Param data body TagRequest true	"删除现有标签的一些参数"
// @Router /api/tagRemove [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (TagApi) TagRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//批量删除
	var tagList []models.TagModel
	count := global.DB.Find(&tagList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("所要删除的标签不存在", c)
		return
	}
	//如果标签下有关联的文章怎么办？
	global.DB.Delete(&tagList)
	res.OKWithMsg(fmt.Sprintf("共删除 %d 个标签", count), c)
}
