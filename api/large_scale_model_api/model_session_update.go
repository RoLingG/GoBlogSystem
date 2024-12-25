package large_scale_model_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/utils/jwt"
	"github.com/gin-gonic/gin"
)

type SessionUpdateRequest struct {
	SessionID   uint   `json:"session_id"`
	SessionName string `json:"session_name"`
}

func (LargeScaleModelApi) ModelSessionUpdateView(c *gin.Context) {
	var cr SessionUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}

	// 查找会话是否存在
	var session models.LargeScaleModelSessionModel
	err = global.DB.Take(&session, cr.SessionID).Error
	if err != nil {
		res.FailWithMsg("大模型会话不存在", c)
		return
	}

	// 查找用户是否存在
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	if session.UserID != claims.UserID {
		res.FailWithMsg("用户信息错误", c)
		return
	}

	// 修改会话名称
	err = global.DB.Model(&session).Updates(SessionUpdateRequest{
		SessionID:   cr.SessionID,
		SessionName: cr.SessionName,
	}).Error
	if err != nil {
		res.FailWithMsg("大模型会话名称修改失败", c)
		return
	}

	res.OKWithMsg("大模型会话名称更新成功", c)
}
