package models

import (
	"GoRoLingG/models/ctype"
)

type UserModel struct {
	Model `json:",omit(comment)"`
	//gorm.Model                     //gorm的Model自带逻辑删除，如果需要用就用gorm自带的
	NickName   string           `gorm:"size:36" json:"nick_name,select(comment),select(info)"`
	UserName   string           `gorm:"size:36" json:"user_name,select(info),omit(comment)"`
	Password   string           `gorm:"size:128" json:"-"`
	Avatar     string           `gorm:"size:256" json:"avatar,select(comment),select(info)"`
	Email      string           `gorm:"size:128" json:"email,select(info),omit(comment)"`
	Telephone  string           `gorm:"size:18" json:"telephone,select(info),omit(comment)"`
	Address    string           `gorm:"size:64" json:"address,select(comment),select(info)"`
	Token      string           `gorm:"size:512" json:"token,omit(comment)"`
	IP         string           `gorm:"size:20" json:"ip,select(comment),select(info)"`
	Role       ctype.Role       `gorm:"size:4;default:1" json:"role,select(info),omit(comment)"`        //角色权限：1为管理员，2为用户，3为游客
	SignStatus ctype.SignStatus `gorm:"type=smallint(6)" json:"sign_status,select(info),omit(comment)"` //注册来源
	Signature  string           `gorm:"size:48" json:"signature,select(info),omit(comment)"`            //用户签名
	Scope      int              `gorm:"default:0" json:"scope,select(info),omit(comment)"`              //用户积分
	Blog       string           `gorm:"size:36" json:"blog,select(info),omit(comment)"`                 //用户博客
}
