package tag_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
)

// TagUpdateView 更新标签
// @Tags 标签管理
// @Summary 更新标签
// @Description	更新现有的标签
// @Param id path int true "需要更新的标签ID"
// @Param data body TagRequest true	"更新现有标签的一些参数"
// @Router /api/tagUpdate/{id} [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (TagApi) TagUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	var tag models.AdvertModel
	err = global.DB.Debug().Take(&tag, id).Error
	if err != nil {
		res.FailWithMsg("该标签不存在，请重传", c)
		return
	}

	tag = models.AdvertModel{}
	//标题是否重复判断
	err = global.DB.Debug().Take(&tag, "title = ?", cr.Title).Error
	//无err就说明在数据库中找到了
	if err == nil {
		res.FailWithMsg("该标签标题已存在，请修改标题重传", c)
		return
	}

	err = global.DB.Debug().Where(id).Updates(&models.TagModel{
		Title: cr.Title,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("修改标签失败", c)
		return
	}
	res.OKWithMsg("修改标签成功", c)
}
