package tag_api

import (
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/common"
	"github.com/gin-gonic/gin"
)

// TagListView 标签列表
// @Tags 标签管理
// @Summary 标签列表
// @Description	标签列表，用于展示所有标签
// @Param data query models.PageInfo false	"查询标签列表的一些参数"
// @Router /api/tagList [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.TagModel]}
func (TagApi) TagListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, err := common.CommonList(models.TagModel{}, common.Option{
		PageInfo: cr,
	})

	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	//需要展示标签关联的文章数
	res.OKWithList(list, count, c)
	return
}
