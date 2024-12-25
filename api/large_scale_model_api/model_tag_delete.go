package large_scale_model_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LargeScaleModelTagDeleteView 大模型角色标签删除
func (LargeScaleModelApi) LargeScaleModelTagDeleteView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMsg("参数错误", c)
		return
	}
	var tagList []models.LargeScaleModelTagModel
	count := global.DB.Preload("Roles").Find(&tagList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("大模型角色标签不存在", c)
		return
	}

	if len(tagList) > 0 {
		// 删除角色标签记录
		for _, tag := range tagList {
			//删除角色标签关联的数据
			global.DB.Model(&tag).Association("Roles").Delete(tag.Roles)
		}
		//删除数据库内的角色标签数据
		err = global.DB.Delete(&tagList).Error
		if err != nil {
			logrus.Error(err)
			res.FailWithMsg("删除大模型角色标签失败", c)
			return
		}
		logrus.Infof("删除大模型角色标签 %d 个", len(tagList))
	}
	res.OKWithMsg("删除大模型角色标签成功", c)
}
