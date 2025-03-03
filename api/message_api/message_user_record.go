package message_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/common"
	"github.com/gin-gonic/gin"
)

type MessageUserRecordRequest struct {
	models.PageInfo
	SendUserID uint `json:"send_user_id" form:"send_user_id" binding:"required"`
	RevUserID  uint `json:"rev_user_id" form:"rev_user_id" binding:"required"`
}

// MessageUserRecordView 两个用户之间的聊天记录
// @Tags 消息管理
// @Summary 两个用户之间的聊天记录
// @Description 两个用户之间的聊天记录
// @Router /api/messageUserRecord [get]
// @Param token header string  true  "token"
// @Param data query MessageUserRecordRequest   false  "查询参数"
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.MessageModel]}
func (MessageApi) MessageUserRecordView(c *gin.Context) {
	var cr MessageUserRecordRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithMsg("参数错误", c)
		return
	}

	//查询条件，就只查发送和接受两人之间的消息 二者互为消息发送/接收者
	sqlCondition := global.DB.Where("(send_user_id = ? and rev_user_id = ? ) or ( rev_user_id = ? and send_user_id = ? )",
		cr.SendUserID, cr.RevUserID, cr.SendUserID, cr.RevUserID)

	list, count, _ := common.CommonList(models.MessageModel{}, common.Option{
		PageInfo: cr.PageInfo,
		Where:    sqlCondition,
		RoleBool: false,
	})

	res.OKWithList(list, count, c)
}
