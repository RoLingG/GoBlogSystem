package models

import (
	"GoRoLingG/models/ctype"
)

type UserModel struct {
	Model //自己写的Model
	//gorm.Model                     //gorm的Model自带逻辑删除，如果需要用就用gorm自带的
	NickName      string           `gorm:"size:36" json:"nick_name"`
	UserName      string           `gorm:"size:36" json:"user_name"`
	Password      string           `gorm:"size:128" json:"password"`
	Avatar        string           `gorm:"size:256" json:"avatar"`
	Email         string           `gorm:"size:128" json:"email"`
	Telephone     string           `gorm:"size:18" json:"telephone"`
	Address       string           `gorm:"size:64" json:"address"`
	Token         string           `gorm:"size:64" json:"token"`
	IP            string           `gorm:"size:20" json:"ip"`
	Role          ctype.Role       `gorm:"size:4;default:1" json:"role"` //角色权限：1为管理员，2为用户，3为游客
	SignStatus    ctype.SignStatus `gorm:"type=smallint(6)" json:"sign_status"`
	ArticleModels []ArticleModel   `gorm:"foreignKey:UserID" json:"-"` //发布的文章列表
	CollectModels []ArticleModel   `gorm:"many2many:user_collect_models;joinForeignKey:UserID;JoinReferences:ArticleID" json:"-"`
}