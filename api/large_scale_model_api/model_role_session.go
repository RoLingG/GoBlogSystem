package large_scale_model_api

import (
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/common"
	"GoRoLingG/utils/jwt"
	"github.com/gin-gonic/gin"
	"time"
)

type RoleSessionsRequest struct {
	models.PageInfo
	RoleID uint `json:"role_id" form:"role_id" binding:"required"`
}

type RoleSessionResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"create_at"`
	SessionName string    `json:"session_name"`
}

// ModelRoleSessionsView 角色会话列表
func (LargeScaleModelApi) ModelRoleSessionsView(c *gin.Context) {
	var cr RoleSessionsRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	_list, count, _ := common.CommonList(models.LargeScaleModelSessionModel{UserID: claims.UserID, RoleID: cr.RoleID}, common.Option{
		PageInfo: cr.PageInfo,
		Likes:    []string{"session_name"},
	})
	var list = make([]RoleSessionResponse, 0)
	for _, model := range _list {
		list = append(list, RoleSessionResponse{
			ID:          model.ID,
			CreatedAt:   model.CreateAt,
			SessionName: model.SessionName,
		})
	}
	res.OKWithList(list, count, c)
}
