package large_scale_model_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/utils/jwt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ModelUserChatDeleteView 大模型普通用户对话删除接口
func (LargeScaleModelApi) ModelUserChatDeleteView(c *gin.Context) {
	var cr models.IDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}

	// 查找用户的对话
	var chat models.LargeScaleModelChatModel
	err = global.DB.Take(&chat, cr.ID).Error
	if err != nil {
		res.FailWithMsg("对话不存在", c)
		return
	}

	// 用户鉴权
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	if chat.UserID != claims.UserID {
		res.FailWithMsg("对话鉴权失败", c)
		return
	}

	// 删除对话
	err = global.DB.Delete(&chat).Error
	if err != nil {
		logrus.Error(err)
		res.FailWithMsg("对话删除失败", c)
		return
	}
	res.OKWithMsg("对话删除成功", c)
}
