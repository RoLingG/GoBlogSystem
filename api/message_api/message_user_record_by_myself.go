package message_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/common"
	"GoRoLingG/utils/jwt"
	"github.com/gin-gonic/gin"
)

type MessageUserRecordByMeRequest struct {
	models.PageInfo
	UserID uint `json:"user_id" form:"user_id" binding:"required"`
}

// MessageUserRecordByMyselfView 当前用户与某个用户的聊天列表
// @Tags 消息管理
// @Summary 当前用户与某个用户的聊天列表
// @Description 当前用户与某个用户的聊天列表
// @Router /api/messageUserRecordByMyself [get]
// @Param token header string  true  "Token"
// @Param data query MessageUserRecordByMeRequest  true  "参数"
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.MessageModel]}
func (message MessageApi) MessageUserRecordByMyselfView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	//获取聊天对象的ID
	var cr MessageUserRecordByMeRequest
	c.ShouldBindQuery(&cr)

	cr.Sort = "create_at asc"
	sqlCondition := global.DB.Where("(send_user_id = ? and rev_user_id = ? ) or ( rev_user_id = ? and send_user_id = ? )", claims.UserID, cr.UserID, claims.UserID, cr.UserID)
	list, count, _ := common.CommonList(models.MessageModel{}, common.Option{
		PageInfo: cr.PageInfo,
		Where:    sqlCondition,
	})

	res.OKWithList(list, count, c)
}
