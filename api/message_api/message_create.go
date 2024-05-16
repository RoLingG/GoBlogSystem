package message_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
)

type MessageRequest struct {
	SendUserID uint   `json:"send_user_id" binding:"required"` // 发送人id
	RevUserID  uint   `json:"rev_user_id" binding:"required"`  // 接收人id
	Content    string `json:"content" binding:"required"`      // 消息内容
}

// MessageCreateView 发布消息
func (MessageApi) MessageCreateView(c *gin.Context) {
	// 当前用户发布消息
	// SendUserID 就是当前登录人的id
	var cr MessageRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var sendUser, revUser models.UserModel

	err = global.DB.Take(&sendUser, cr.SendUserID).Error
	if err != nil {
		res.FailWithMsg("发送人不存在", c)
		return
	}
	err = global.DB.Take(&revUser, cr.RevUserID).Error
	if err != nil {
		res.FailWithMsg("接收人不存在", c)
		return
	}

	err = global.DB.Create(&models.MessageModel{
		SendUserID:       cr.SendUserID,
		SendUserNickName: sendUser.NickName,
		SendUserAvatar:   sendUser.Avatar,
		RevUserID:        cr.RevUserID,
		RevUserNickName:  revUser.NickName,
		RevUserAvatar:    revUser.Avatar,
		IsRead:           false,
		Content:          cr.Content,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("消息发送失败", c)
		return
	}
	res.OKWithMsg("消息发送成功", c)
	return
}
