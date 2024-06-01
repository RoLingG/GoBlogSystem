package role_api

import (
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
)

type OptionResponse struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}

// RoleIDListView 用户权限列表
// @Tags 权限管理
// @Summary 查看所有用户权限的列表
// @Description 查看所有用户权限的列表
// @Router /api/roleIDList [get]
// @Produce json
// @Success 200 {object} res.Response{data=OptionResponse}
func (RoleApi) RoleIDListView(c *gin.Context) {
	res.OKWithData([]OptionResponse{
		{"管理员", 1},
		{"普通用户", 2},
		{"游客", 3},
	}, c)
}
