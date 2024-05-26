package message_api

import (
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/common"
	"github.com/gin-gonic/gin"
)

// MessageListAdminView 消息列表(管理员)
// @Tags 消息管理
// @Summary 消息列表
// @Description	查询所有消息的列表
// @Param token header string true "Authorization token"
// @Param data query models.PageInfo true	"查询消息列表的一些参数"
// @Router /api/messageAdminList [get]
// @Produce json
// @Success 200 {object} res.Response{}
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
