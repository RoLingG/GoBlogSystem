package large_scale_model_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/utils/jwt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (LargeScaleModelApi) ModelSessionDeleteView(c *gin.Context) {
	var cr models.IDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}

	// 查找会话是否存在
	var session models.LargeScaleModelSessionModel
	err = global.DB.Preload("ChatList").Take(&session, cr.ID).Error
	if err != nil {
		res.FailWithMsg("大模型会话不存在", c)
		return
	}

	// 查找用户是否存在
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	if session.UserID != claims.UserID {
		res.FailWithMsg("用户信息错误", c)
		return
	}

	// 对话记录得先删除，然后才能删除会话
	if len(session.ChatList) > 0 {
		err = global.DB.Delete(&session.ChatList).Error
		if err != nil {
			logrus.Error(err)
		} else {
			logrus.Infof("删除关联大模型对话 %d 个", len(session.ChatList))
		}
	}

	// 删除对应的大模型会话
	err = global.DB.Delete(&session).Error
	if err != nil {
		logrus.Error(err)
		res.FailWithMsg("会话删除失败", c)
		return
	}

	res.OKWithMsg("大模型会话删除成功", c)
}
