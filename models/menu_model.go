package models

import (
	"GoRoLingG/models/ctype"
)

type MenuModel struct {
	Model
	MenuTitle    string       `gorm:"size:32" json:"menu_title"`
	MenuTitleEn  string       `gorm:"size:32" json:"menu_title_en"`
	Slogan       string       `gorm:"size:64" json:"slogan"`
	Abstract     ctype.Array  `gorm:"type:string" json:"abstract"`                                                               //简介
	AbstractTime int          `json:"abstract_time"`                                                                             //简介的切换时间
	MenuImage    []ImageModel `gorm:"many2many:menu_image_model;joinForeignKey:MenuID;joinReferences:ImageID" json:"menu_image"` //菜单的图片列表
	MenuTime     int          `json:"menu_time"`                                                                                 //菜单图片的切换时间，为0不切换
	Sort         int          `gorm:"size:10" json:"sort"`                                                                       //菜单的顺序
}
