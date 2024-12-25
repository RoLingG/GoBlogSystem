package large_scale_model_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"regexp"
)

type AutoReplyUpdateRequest struct {
	ID           uint   `json:"id"`
	RuleName     string `json:"rule_name" binding:"required"`               // 规则名称
	RuleMode     int    `json:"rule_mode" binding:"required,oneof=1 2 3 4"` // 匹配模式 1 精确匹配，2 模糊匹配，3 前缀匹配，4 正则匹配
	Rule         string `json:"rule" binding:"required"`                    // 匹配规则
	ReplyContent string `json:"reply_content" binding:"required"`           // 回复内容
}

func (LargeScaleModelApi) AutoReplyUpdateView(c *gin.Context) {
	var cr AutoReplyUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMsg("参数错误", c)
		return
	}
	// 校验正则是否写错，这里检查了，之后就不用担心这个问题了
	if cr.RuleMode == 4 {
		_, err := regexp.Compile(cr.Rule)
		if err != nil {
			res.FailWithMsg(fmt.Sprintf("正则表达式错误 %s", err.Error()), c)
			return
		}
	}
	if cr.ID == 0 {
		//新增规则
		var autoReplyModel models.AutoReplyModel
		err = global.DB.Take(&autoReplyModel, "rule_name = ?", cr.RuleName).Error
		if err == nil {
			res.FailWithMsg("规则名称不能相同", c)
			return
		}
		err = global.DB.Create(&models.AutoReplyModel{
			RuleName:     cr.RuleName,
			RuleMode:     cr.RuleMode,
			Rule:         cr.Rule,
			ReplyContent: cr.ReplyContent,
		}).Error
		if err != nil {
			res.FailWithMsg("自动回复规则添加失败", c)
			return
		}
		res.OKWithMsg("自动回复规则添加成功", c)
		return
	}
	//自动回复需要更新的规则是否存在
	var autoReplyModelExist models.AutoReplyModel
	err = global.DB.Take(&autoReplyModelExist, cr.ID).Error
	if err != nil {
		res.FailWithMsg("记录不存在", c)
		return
	}
	// 自动回复更新规则名重复校验
	var autoReplyModelNameCheck models.AutoReplyModel
	err = global.DB.Take(&autoReplyModelNameCheck, "rule_name = ? and id <> ?", cr.RuleName, cr.ID).Error
	if err == nil {
		res.FailWithMsg("和已有的规则名称重复", c)
		return
	}
	err = global.DB.Model(&autoReplyModelExist).Updates(map[string]any{
		"rule_name":     cr.RuleName,
		"rule_mode":     cr.RuleMode,
		"rule":          cr.Rule,
		"reply_content": cr.ReplyContent,
	}).Error
	if err != nil {
		logrus.Errorf("自动回复更新失败 err：%s, 自动回复数据内容 %#v", err, cr)
		res.FailWithMsg("自动回复更新失败", c)
		return
	}
	res.OKWithMsg("自动回复更新成功", c)
	return
}
