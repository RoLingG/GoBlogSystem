package log_stash

import "time"

type LogStashModel struct {
	ID       uint      `gorm:"primarykey" json:"id"`                          //日志ID
	CreateAt time.Time `gorm:"default:current_timestamp(3)" json:"create_at"` //日志创建时间
	IP       string    `gorm:"size:32" json:"ip"`                             //造成日志的IP
	Addr     string    `gorm:"size:64" json:"addr"`                           //造成日志的地址
	LogLevel LogLevel  `gorm:"size:4" json:"log_level"`                       //日志等级
	Content  string    `gorm:"size:128" json:"content"`                       //日志内容
	UserID   uint      `json:"user_id"`                                       //登录用户的用户id
}
