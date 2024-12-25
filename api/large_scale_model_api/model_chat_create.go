package large_scale_model_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/large_scale_model_service"
	"GoRoLingG/utils/jwt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
)

type ChatCreateRequest struct {
	SessionID uint   `form:"session_id" json:"session_id" binding:"required"` // 角色id
	Content   string `form:"content" json:"content" binding:"required"`       //聊天内容
}

func (LargeScaleModelApi) ModelChatCreateView(c *gin.Context) {
	// 认证
	token := c.Query("token")
	claims, err := jwt.ParseToken(token)
	if err != nil {
		res.FailWithMsgSSE("认证失败", c)
		return
	}

	var cr ChatCreateRequest
	err = c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithMsgSSE("参数错误", c)
		return
	}
	// 查找用户选中会话是否存在
	var session models.LargeScaleModelSessionModel
	err = global.DB.Take(&session, cr.SessionID).Error
	if err != nil {
		res.FailWithMsgSSE("会话不存在", c)
		return
	}

	// 用户是否存在，以及其是否可以创建对话
	var user models.UserModel
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMsgSSE("用户信息错误", c)
		return
	}
	scope := global.Config.LargeScaleModel.ModelSessionSetting.ChatScope
	if user.Scope-scope < 0 {
		res.FailWithMsgSSE("角色积分不足，无法创建对话", c)
		return
	}

	msgChan, err := large_scale_model_service.Send(cr.SessionID, cr.Content)
	if err != nil {
		res.FailWithMsgSSE(err.Error(), c)
		return
	}
	//没问题，则流式传输内容
	var aiContent string
	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-msgChan; ok {
			res.OKWithDataSSE(msg, c)
			aiContent += msg
			return true
		}
		return false
	})
	chat := models.LargeScaleModelChatModel{
		SessionID:   cr.SessionID,
		Status:      true,
		UserContent: cr.Content,
		AIContent:   aiContent,
		RoleID:      session.RoleID,
		UserID:      session.UserID,
	}
	// 创建对话
	err = global.DB.Create(&chat).Error
	if err != nil {
		res.FailWithMsgSSE("对话失败", c)
		return
	}

	//创建对话成功，扣除用户对应的积分
	global.DB.Model(&user).Update("scope", gorm.Expr("scope - ?", scope))
	res.OKWithDataAndMsg(chat.ID, "OK", c)
}
