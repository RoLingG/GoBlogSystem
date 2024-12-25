package large_scale_model_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/utils/jwt"
	"github.com/gin-gonic/gin"
)

// ModelHistoricalSessionListView 大模型历史会话列表
func (LargeScaleModelApi) ModelHistoricalSessionListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	var roleIDList []uint
	global.DB.Model(models.LargeScaleModelSessionModel{}).Where("user_id = ?", claims.UserID).Group("role_id").Select("role_id").Scan(&roleIDList)
	var roleList []models.LargeScaleModelRoleModel
	global.DB.Order("create_at").Find(&roleList, "id in ?", roleIDList) //这里加"id in ?"条件是为了让用户没有角色聊天时，不查全部人的信息，和gorm有关

	var list = make([]RoleItem, 0)
	for _, item := range roleList {
		list = append(list, RoleItem{
			ID:       item.ID,
			RoleName: item.Name,
			Abstract: item.Abstract,
			Icon:     item.Icon,
		})
	}
	res.OKWithDataAndMsg(list, "大模型历史会话查询成功", c)
	return
}
