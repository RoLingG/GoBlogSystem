package large_scale_model_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/utils/jwt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SessionCreateRequest struct {
	RoleID uint `json:"role_id" binding:"required"` // 角色id
}

func (LargeScaleModelApi) ModelSessionCreateView(c *gin.Context) {
	var cr SessionCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	// 查找用户选中大模型角色是否存在
	var role models.LargeScaleModelRoleModel
	err = global.DB.Take(&role, cr.RoleID).Error
	if err != nil {
		res.FailWithMsg("大模型角色不存在", c)
		return
	}

	// 用户是否存在，以及其是否可以创建会话
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var user models.UserModel
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMsg("用户信息错误", c)
		return
	}
	scope := global.Config.LargeScaleModel.ModelSessionSetting.SessionScope
	if user.Scope-scope < 0 {
		res.FailWithMsg("角色积分不足，无法创建会话", c)
		return
	}

	// 可以创建会话
	// 名字默认就是新的会话
	// 如果用户创建了一个新的会话，但是没有聊天，那就不能创建
	// 找这个用户相关的ai角色，有没有空的对话记录 > 1
	var sessionList []models.LargeScaleModelSessionModel
	global.DB.Preload("ChatList").Find(&sessionList, "user_id = ? and role_id = ?", user.ID, cr.RoleID)
	var empty bool
	for _, session := range sessionList {
		if len(session.ChatList) <= 1 {
			empty = true
		}
	}
	if empty {
		res.FailWithMsg("已存在新的未对话会话", c)
		return
	}

	// 创建会话
	var session = models.LargeScaleModelSessionModel{
		SessionName: "新建会话",
		UserID:      user.ID,
		RoleID:      cr.RoleID,
	}
	global.DB.Create(&session)
	//创建会话成功，扣除用户对应的积分
	global.DB.Model(&user).Update("scope", gorm.Expr("scope - ?", scope))

	res.OKWithDataAndMsg(session.ID, "大模型会话创建成功", c)
}
