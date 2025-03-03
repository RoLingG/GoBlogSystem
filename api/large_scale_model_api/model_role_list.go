package large_scale_model_api

import (
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/common"
	"github.com/gin-gonic/gin"
)

// ModelRoleListView 大模型角色列表
func (LargeScaleModelApi) ModelRoleListView(c *gin.Context) {
	var cr models.PageInfo
	c.ShouldBindQuery(&cr)

	list, count, _ := common.CommonList(models.LargeScaleModelRoleModel{}, common.Option{
		PageInfo: cr,
		Likes:    []string{"name"}, //模糊匹配
		Preload:  []string{"Tags"},
	})
	res.OKWithList(list, count, c)
	return
}
