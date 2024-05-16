package tag_api

import (
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/common"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagListView(c *gin.Context) {
	var cr models.PageModel
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, err := common.CommonList(models.TagModel{}, common.Option{
		PageModel: cr,
	})

	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	//需要展示标签关联的文章数
	res.OKWithList(list, count, c)
	return
}
