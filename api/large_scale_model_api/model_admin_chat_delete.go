package large_scale_model_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ModelAdminChatDeleteView 大模型管理员对话删除接口
func (LargeScaleModelApi) ModelAdminChatDeleteView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}

	var list []models.LargeScaleModelChatModel
	count := global.DB.Find(&list, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("记录不存在", c)
		return
	}

	if len(list) > 0 {
		// 先删除表外键关联数据
		err = global.DB.Delete(&list).Error
		if err != nil {
			logrus.Error(err)
			res.FailWithMsg("删除对话失败", c)
			return
		}
		logrus.Infof("删除对话 %d 个", len(list))
	}
	res.OKWithMsg(fmt.Sprintf("共删除 %d 个对话", count), c)
}
