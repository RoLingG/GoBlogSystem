package models

import "time"

type Model struct {
	ID       uint      `gorm:"primarykey" json:"ID"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}
