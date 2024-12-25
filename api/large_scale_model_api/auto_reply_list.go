package large_scale_model_api

import (
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/common"
	"github.com/gin-gonic/gin"
)

// AutoReplyListView 自动回复列表
func (LargeScaleModelApi) AutoReplyListView(c *gin.Context) {
	var cr models.PageInfo
	c.ShouldBindQuery(&cr)

	list, count, _ := common.CommonList(models.AutoReplyModel{}, common.Option{
		PageInfo: cr,
		Likes:    []string{"rule_name"}, //模糊匹配
	})
	res.OKWithList(list, count, c)
	return
}
