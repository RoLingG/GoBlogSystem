package models

import "time"

// UserCollectModel 自定义第三张表，记录用户什么时候收藏什么文章
type UserCollectModel struct {
	UserID    uint      `gorm:"primaryKey"`
	UserModel UserModel `gorm:"foreignKey:UserID"`
	ArticleID string    `gorm:"size:32;primaryKey"`
	CreateAt  time.Time `gorm:"default:current_timestamp(3)"`
}
