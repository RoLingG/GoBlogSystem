package models

import "GoRoLingG/models/ctype"

type ChatModel struct {
	Model    `json:",select(list)"`
	NickName string        `gorm:"size:15" json:"nick_name"`
	Avatar   string        `gorm:"size:128" json:"avatar"`
	Content  string        `gorm:"size:1024" json:"content"`
	IP       string        `gorm:"size:32" json:"ip,omit(list)"`
	Addr     string        `gorm:"size:64" json:"addr,omit(list)"`
	IsGroup  bool          `json:"is_group"` // 是否是群组消息
	MsgType  ctype.MsgType `gorm:"size:4" json:"msg_type"`
}
