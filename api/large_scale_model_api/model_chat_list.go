package large_scale_model_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/common"
	"GoRoLingG/utils/jwt"
	"github.com/gin-gonic/gin"
)

type ChatListRequest struct {
	SessionID uint `json:"session_id" form:"session_ID" binding:"require"`
	models.PageInfo
}

type ChatListResponse struct {
	models.Model
	UserContent string `json:"user_content"`
	UserAvatar  string `json:"user_avatar"`
	AIContent   string `json:"ai_content"`
	AIAvatar    string `json:"ai_avatar"`
	Status      bool   `json:"status"`
}

// ModelChatListView 大模型对话列表接口
func (LargeScaleModelApi) ModelChatListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(jwt.CustomClaims)
	var cr ChatListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	var session models.LargeScaleModelSessionModel
	err = global.DB.Take(&session, cr.SessionID).Error
	if err != nil {
		res.FailWithMsg("大模型会话ID错误", c)
		return
	}
	if claims.Role != models.AdminRole {
		// 验证这个会话是不是当前用户创建的
		if claims.UserID != session.UserID {
			res.FailWithMsg("会话鉴权失败", c)
			return
		}
	}
	_list, count, _ := common.CommonList(models.LargeScaleModelChatModel{SessionID: cr.SessionID}, common.Option{
		PageInfo: cr.PageInfo,
		Preload:  []string{"RoleModel", "UserModel"},
	})
	var list = make([]ChatListResponse, 0)
	for _, item := range _list {
		list = append(list, ChatListResponse{
			Model:       item.Model,
			UserContent: item.UserContent,
			UserAvatar:  item.UserModel.Avatar,
			AIContent:   item.AIContent,
			AIAvatar:    item.RoleModel.Icon,
			Status:      item.Status,
		})
	}
	res.OKWithList(list, count, c)
}
