package message_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/utils/jwt"
	"github.com/gin-gonic/gin"
	"time"
)

type Message struct {
	SendUserID       uint      `json:"send_user_id"` // 发送人id
	SendUserNickName string    `json:"send_user_nick_name"`
	SendUserAvatar   string    `json:"send_user_avatar"`
	RevUserID        uint      `json:"rev_user_id"` // 接收人id
	RevUserNickName  string    `json:"rev_user_nick_name"`
	RevUserAvatar    string    `json:"rev_user_avatar"`
	Content          string    `json:"content"`       // 消息内容
	CreateAt         time.Time `json:"create_at"`     // 最新的消息时间
	MessageCount     int       `json:"message_count"` // 消息条数
}

type MessageGroup map[uint]*Message //消息组

func (MessageApi) MessageListUserView(c *gin.Context) {
	//根据jwt的token获取当前用户信息
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	var messageGroup = MessageGroup{}
	var messageList []models.MessageModel
	var messages []Message

	global.DB.Order("create_at asc").Find(&messageList, "send_user_id = ? or rev_user_id = ?", claims.UserID, claims.UserID)
	//因为send_user_id和rev_user_id都被锁定成是当前用户，所以该用户的id+其他用户的id不会重复，所以采用id和的机制进行分组
	for _, messageModel := range messageList {
		// 判断是一个组的条件
		// send_user_id 和 rev_user_id 其中一个
		// 1 2  2 1
		// 1 3  3 1 是一组
		message := Message{
			SendUserID:       messageModel.SendUserID,
			SendUserNickName: messageModel.SendUserNickName,
			SendUserAvatar:   messageModel.SendUserAvatar,
			RevUserID:        messageModel.RevUserID,
			RevUserNickName:  messageModel.RevUserNickName,
			RevUserAvatar:    messageModel.RevUserAvatar,
			Content:          messageModel.Content,
			CreateAt:         messageModel.CreateAt,
			MessageCount:     1,
		}
		idNum := message.SendUserID + message.RevUserID //采用id和的方式进行分组
		val, ok := messageGroup[idNum]                  //取分组内消息发送人和接收人之间的最新消息
		if !ok {
			// 如果当前消息组消息不存在，则将当前消息存进去，并继续循环
			messageGroup[idNum] = &message //存该消息发送人和接收人之间的最新消息
			continue
		}
		//如果当前消息组消息存在，则消息计数+1，并将当前消息组的消息刷新
		message.MessageCount = val.MessageCount + 1 //同发送人和接收人之间的消息计数
		messageGroup[idNum] = &message              //存该消息发送人和接收人之间的最新消息
	}
	for _, message := range messageGroup {
		messages = append(messages, *message) //进行对应形式分组，便于反馈给前端
	}

	res.OKWithData(messages, c)
	return
}
