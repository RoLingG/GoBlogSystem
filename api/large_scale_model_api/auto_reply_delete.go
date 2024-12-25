package large_scale_model_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (LargeScaleModelApi) AutoReplyDeleteView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var list []models.AutoReplyModel
	count := global.DB.Find(&list, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("自动回复数据记录不存在", c)
		return
	}

	if len(list) > 0 {
		global.DB.Delete(&list)
		logrus.Infof("删除自动回复数据记录 %d 条", len(list))
	}
	res.OKWithMsg(fmt.Sprintf("共删除 %d 条自动回复记录", len(list)), c)
}
