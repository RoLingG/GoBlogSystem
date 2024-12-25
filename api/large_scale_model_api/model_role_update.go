package large_scale_model_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RoleUpdateRequest struct {
	ID        uint   `json:"id"`
	Name      string `binding:"required" json:"name"`     // 角色名称
	Enable    bool   `json:"enable"`                      // 是否启用
	Icon      string `json:"icon"`                        // 可以选择系统默认的一些，也可以图片上传
	Abstract  string `binding:"required" json:"abstract"` // 简介
	Scope     int    `json:"scope"`                       // 消耗的积分
	Prologue  string `binding:"required" json:"prologue"` // 开场白
	Prompt    string `binding:"required" json:"prompt"`   // 设定词
	AutoReply bool   `json:"auto_reply"`                  // 自动回复
	TagList   []uint `json:"tagList"`                     // 标签的id列表
}

// ModelRoleUpdateView 大模型角色新增和更新
func (LargeScaleModelApi) ModelRoleUpdateView(c *gin.Context) {
	var cr RoleUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	// 先在数据库找这些标签，只看都对得上和都没有
	var tagList []models.LargeScaleModelTagModel
	if len(cr.TagList) == 0 {
		tagList = make([]models.LargeScaleModelTagModel, 0)
	} else {
		global.DB.Find(&tagList, cr.TagList)
		if len(cr.TagList) != len(tagList) {
			res.FailWithMsg("标签选择不一致", c)
			return
		}
	}

	if cr.ID == 0 {
		//新增标签
		var roleModel models.LargeScaleModelRoleModel
		roleModel = models.LargeScaleModelRoleModel{
			Name:      cr.Name,
			Enable:    cr.Enable,
			Icon:      cr.Icon,
			Abstract:  cr.Abstract,
			Tags:      tagList,
			Scope:     cr.Scope,
			Prologue:  cr.Prologue,
			Prompt:    cr.Prompt,
			AutoReply: cr.AutoReply,
		}
		err = global.DB.Create(&roleModel).Error
		if err != nil {
			logrus.Errorf("大模型角色添加失败 err：%s, 大模型角色原始数据内容 %#v", err, cr)
			res.FailWithMsg("大模型角色添加失败", c)
			return
		}
		res.OKWithMsg("大模型角色添加成功", c)
		return
	}
	// 更新
	var roleModelExist models.LargeScaleModelRoleModel
	err = global.DB.Preload("Tags").Take(&roleModelExist, cr.ID).Error
	if err != nil {
		res.FailWithMsg("大模型角色记录不存在", c)
		return
	}
	var roleModelNameExist models.LargeScaleModelRoleModel
	err = global.DB.Take(&roleModelNameExist, "name = ? and id <> ?", cr.Name, cr.ID).Error
	if err == nil {
		res.FailWithMsg("和已有的大模型角色名称重复", c)
		return
	}
	err = global.DB.Model(&roleModelExist).Updates(map[string]any{
		"name":       cr.Name,
		"enable":     cr.Enable,
		"icon":       cr.Icon,
		"abstract":   cr.Abstract,
		"scope":      cr.Scope,
		"prologue":   cr.Prologue,
		"prompt":     cr.Prompt,
		"auto_reply": cr.AutoReply,
	}).Error
	if err != nil {
		logrus.Errorf("大模型角色数据更新失败 err：%s, 大模型角色原始数据内容 %#v", err, cr)
		res.FailWithMsg("大模型角色更新失败", c)
		return
	}
	// 把之前的大模型角色标签替换掉
	err = global.DB.Model(&roleModelExist).Association("Tags").Replace(tagList)
	if err != nil {
		res.FailWithMsg("大模型角色标签替换失败", c)
		return
	}
	res.OKWithMsg("大模型角色更新成功", c)
}
