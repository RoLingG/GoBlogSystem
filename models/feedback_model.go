package models

type FeedBackModel struct {
	Model        `structs:"-"`
	Email        string `gorm:"size:64" json:"email" structs:"email"`      //问题反馈人邮箱
	Content      string `gorm:"size:128" json:"content" structs:"content"` //问题反馈内容
	ApplyContent string `gorm:"size:128" json:"applyContent" structs:"-"`  //回复内容
	IsApply      bool   `json:"is_apply" structs:"-"`                      //是否回复
}
