package models

type FeedBackModel struct {
	Model
	Email        string `gorm:"size:64" json:"email"`         //问题反馈人邮箱
	Content      string `gorm:"size:128" json:"content"`      //问题反馈内容
	ApplyContent string `gorm:"size:128" json:"applyContent"` //回复内容
	IsApply      bool   `json:"is_apply"`                     //是否回复
}
