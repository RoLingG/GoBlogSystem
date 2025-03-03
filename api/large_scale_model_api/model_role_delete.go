package large_scale_model_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (LargeScaleModelApi) ModelRoleDeleteView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMsg("参数错误", c)
		return
	}
	var roleList []models.LargeScaleModelRoleModel
	count := global.DB.Preload("Tags").Find(&roleList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("大模型角色不存在", c)
		return
	}

	if len(roleList) > 0 {
		for _, role := range roleList {
			//删除角色关联的数据
			global.DB.Model(&role).Association("Tags").Delete(role.Tags)
		}
		//删除数据库内的角色数据
		err = global.DB.Delete(&roleList).Error
		if err != nil {
			logrus.Error(err)
			res.FailWithMsg("删除大模型角色失败", c)
			return
		}
		logrus.Infof("删除大模型角色 %d 个", len(roleList))
	}
	res.OKWithMsg("删除大模型角色成功", c)
}
