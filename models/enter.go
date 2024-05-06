package models

import "time"

type Model struct {
	ID       uint      `gorm:"primarykey" json:"id"`
	CreateAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"create_at"`
	UpdateAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
	//CreateAt time.Time `json:"create_at"`
	//UpdateAt time.Time `json:"-"`
}

// APIPOST的参数，不设置默认值都是0
// 因为Page会被高频使用(是分页就会用到)，所以可以封装起来
type PageModel struct {
	Page  int    `form:"page"`  //页数
	Key   string `form:"key"`   //模糊匹配的关键字
	Limit int    `form:"limit"` //每页限制显示量
	Sort  string `form:"sort"`  //排序
}
