package message_api

import (
	"GoRoLingG/utils/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
)

// MessageUserListByMyself 我与其他用户的聊天列表
// @Tags 消息管理
// @Summary 我与其他用户的聊天列表
// @Description 我与其他用户的聊天列表
// @Router /api/messageUserListByMyself [get]
// @Param token header string  true  "Token"
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[MessageUserListResponse]}
func (message MessageApi) MessageUserListByMyself(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	//获取自己的token内携带的id，用于MessageUserListByUser内获取聊天信息
	c.Request.URL.RawQuery = fmt.Sprintf("user_id=%d", claims.UserID)
	message.MessageUserListByUser(c)
}
