package large_scale_model_api

import (
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/common"
	"github.com/gin-gonic/gin"
)

type TagListResponse struct {
	models.Model
	RoleTitle string `json:"title"`     // 角色名称
	Color     string `json:"color"`     // 颜色
	RoleCount int    `json:"roleCount"` // 角色个数
}

// LargeScaleModelTagListView 大模型角色标签列表
func (LargeScaleModelApi) LargeScaleModelTagListView(c *gin.Context) {
	var cr models.PageInfo
	c.ShouldBindQuery(&cr)
	_list, count, _ := common.CommonList(models.LargeScaleModelTagModel{}, common.Option{
		Likes:   []string{"role_title"},
		Preload: []string{"Roles"},
	})
	var list = make([]TagListResponse, 0)
	for _, model := range _list {
		list = append(list, TagListResponse{
			Model:     model.Model,
			RoleTitle: model.RoleTitle,
			Color:     model.Color,
			RoleCount: len(model.Roles),
		})
	}
	res.OKWithList(list, count, c)
}
