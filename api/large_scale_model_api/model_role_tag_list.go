package large_scale_model_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
)

type RoleItem struct {
	ID       uint   `json:"id"`        //角色ID
	RoleName string `json:"role_name"` //角色名称
	Abstract string `json:"abstract"`  //角色简介
	Icon     string `json:"icon"`      //角色icon
}

type RoleTagListResponse struct {
	ID       uint       `json:"id"`        //标签ID
	Title    string     `json:"title"`     // 名称
	RoleList []RoleItem `json:"role_list"` //角色列表
}

// LargeScaleModelRoleTagListView 大模型角色广场
func (LargeScaleModelApi) LargeScaleModelRoleTagListView(c *gin.Context) {
	var tagList []models.LargeScaleModelTagModel
	global.DB.Preload("Roles").Find(&tagList)
	var roleTagList = make([]RoleTagListResponse, 0)
	for _, tag := range tagList {
		roleList := make([]RoleItem, 0)
		for _, role := range tag.Roles {
			roleList = append(roleList, RoleItem{
				ID:       role.ID,
				RoleName: role.Name,
				Abstract: role.Abstract,
				Icon:     role.Icon,
			})
		}
		roleTagList = append(roleTagList, RoleTagListResponse{
			ID:       tag.ID,
			Title:    tag.RoleTitle,
			RoleList: roleList,
		})
	}

	res.OKWithData(roleTagList, c)
}
