package message_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/utils/jwt"
	"github.com/gin-gonic/gin"
)

type MessageRecordRequest struct {
	UserID uint `json:"user_id" binding:"required" msg:"请输入查询的用户id"`
}

// MessageRecordView 消息记录列表
// @Tags 消息管理
// @Summary 消息记录列表
// @Description	查询当前用户所有的消息记录
// @Param token header string true "Authorization token"
// @Param data body MessageRecordRequest true "查询与ID用户的消息记录"
// @Router /api/messageRecord [get]
// @Produce json
// @Success 200 {object} res.Response{data=models.MessageModel}
func (MessageApi) MessageRecordView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var cr MessageRecordRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var _messageList []models.MessageModel
	var messageList = make([]models.MessageModel, 0)
	global.DB.Order("create_at asc").Find(&_messageList, "send_user_id = ? or rev_user_id = ?", claims.UserID, claims.UserID)
	for _, message := range _messageList {
		if message.RevUserID == cr.UserID || message.SendUserID == cr.UserID {
			messageList = append(messageList, message)
		}
	}
	// 点开消息，里面的每一条消息，都从未读变成已读

	res.OKWithData(messageList, c)
}
