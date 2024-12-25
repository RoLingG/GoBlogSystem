package large_scale_model_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TagUpdateRequest struct {
	ID        uint   `json:"id"`                            // 更新使用
	RoleTitle string `json:"role_title" binding:"required"` // 名称
	Color     string `json:"color" binding:"required"`      // 颜色
}

// LargeScaleModelTagUpdateView 大模型角色标签新增和更新
func (LargeScaleModelApi) LargeScaleModelTagUpdateView(c *gin.Context) {
	var cr TagUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMsg("参数错误", c)
		return
	}
	// 不传ID/无ID则创建对应的大模型角色标签
	if cr.ID == 0 {
		//增加标签
		var tagModel models.LargeScaleModelTagModel
		err = global.DB.Take(&tagModel, "role_title = ?", cr.RoleTitle).Error
		if err == nil {
			res.FailWithMsg("标签名称不能相同", c)
			return
		}
		err = global.DB.Create(&models.LargeScaleModelTagModel{
			RoleTitle: cr.RoleTitle,
			Color:     cr.Color,
		}).Error
		if err != nil {
			logrus.Errorf("角色标签添加失败 err：%s, 角色标签数据内容 %#v", err, cr)
			res.FailWithMsg("角色标签添加失败", c)
			return
		}
		res.OKWithMsg("角色标签添加成功", c)
		return
	}
	// 角色标签记录是否存在
	var tagModel models.LargeScaleModelTagModel
	err = global.DB.Take(&tagModel, "id = ?", cr.ID).Error
	if err != nil {
		res.FailWithMsg("记录不存在", c)
		return
	}
	// 存在，角色标签名重复校验
	var tagModelExist models.LargeScaleModelTagModel
	err = global.DB.Take(&tagModelExist, "role_title = ? and id <> ?", cr.RoleTitle, cr.ID).Error
	if err != nil {
		res.FailWithMsg("和已有的角色标签名重复", c)
		return
	}
	err = global.DB.Model(&tagModel).Updates(map[string]any{
		"role_title": cr.RoleTitle,
		"color":      cr.Color,
	}).Error
	if err != nil {
		logrus.Errorf("角色标签数据更新失败 err：%s, 角色标签数据内容 %#v", err, cr)
		res.FailWithMsg("角色标签更新失败", c)
		return
	}
	res.OKWithMsg("大模型角色标签更新成功", c)
}
