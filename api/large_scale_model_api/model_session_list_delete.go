package large_scale_model_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ModelSessionListDeleteView 大模型会话批量删除接口
func (LargeScaleModelApi) ModelSessionListDeleteView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}

	var list []models.LargeScaleModelSessionModel
	count := global.DB.Preload("ChatList").Find(&list, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("记录不存在", c)
		return
	}

	if len(list) > 0 {
		for _, item := range list {
			err = global.DB.Delete(&item.ChatList).Error
			if err != nil {
				res.FailWithMsg("级联删除大模型对话失败", c)
				return
			}
			logrus.Infof("级联删除大模型对话 %d 条", len(item.ChatList))
		}
		err = global.DB.Delete(&list).Error
		if err != nil {
			logrus.Error(err)
			res.FailWithMsg("删除大模型会话失败", c)
			return
		}
		logrus.Infof("删除大模型会话 %d 条", len(list))
	}

	res.OKWithMsg(fmt.Sprintf("共删除 %d 条大模型会话", count), c)
}
