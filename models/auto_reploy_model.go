package models

import (
	"GoRoLingG/global"
	"regexp"
	"strings"
)

// AutoReplyModel 自动回复表
type AutoReplyModel struct {
	Model
	RuleName     string `gorm:"size:32" json:"rule_name"`       //规则名称
	RuleMode     int    `json:"rule_mode"`                      //规则类型 1 精确匹配 2 模糊匹配 3 前缀匹配 4 正则匹配
	Rule         string `gorm:"size:64" json:"rule"`            //匹配规则
	ReplyContent string `gorm:"size:1024" json:"reply_content"` //回复内容
}

// AutoReplyValidView 是否命中自动回复
func (AutoReplyModel) AutoReplyValidView(content string) *AutoReplyModel {
	var list []AutoReplyModel
	global.DB.Find(&list)
	for _, model := range list {
		switch model.RuleMode {
		case 1:
			// 精确匹配
			if model.Rule == content {
				return &model
			}
		case 2:
			// 模糊匹配
			if strings.Contains(content, model.Rule) {
				return &model
			}
		case 3:
			// 前缀匹配
			if strings.HasPrefix(content, model.Rule) {
				return &model
			}
		case 4:
			// 正则匹配
			regex, _ := regexp.Compile(model.Rule)
			if regex.MatchString(content) {
				return &model
			}
		}
	}
	return nil
}
