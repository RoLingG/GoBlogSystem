package models

import (
	"GoRoLingG/models/ctype"
)

// 菜单的路径可以是/path，也可以是路由别名
type MenuModel struct {
	Model
	MenuTitle    string       `gorm:"size:32" json:"menu_title"`                                                                 //标题
	MenuPath     string       `gorm:"size:32" json:"menu_path"`                                                                  //路径
	Slogan       string       `gorm:"size:64" json:"slogan"`                                                                     //标语
	Abstract     ctype.Array  `gorm:"type:string" json:"abstract"`                                                               //简介
	AbstractTime int          `json:"abstract_time"`                                                                             //简介的切换时间
	MenuImage    []ImageModel `gorm:"many2many:menu_image_models;joinForeignKey:MenuID;joinReferences:ImageID" json:"menu_image"` //菜单的图片列表
	MenuTime     int          `json:"menu_time"`                                                                                 //菜单图片的切换时间，为0不切换
	Sort         int          `gorm:"size:10" json:"sort"`                                                                       //菜单的顺序
}
