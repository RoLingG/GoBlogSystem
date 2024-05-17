package message_api

import (
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/common"
	"github.com/gin-gonic/gin"
)

// MessageListAdminView 管理员消息列表
func (MessageApi) MessageListAdminView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, err := common.CommonList(models.MessageModel{}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})

	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}

	res.OKWithList(list, count, c)
	return
}
