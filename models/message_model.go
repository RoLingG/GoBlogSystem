package models

type MessageModel struct {
	Model
	SendUserID       uint      `gorm:"primaryKey" json:"send_user_id"` //发送者ID
	SendUserModel    UserModel `gorm:"foreignKey:SendUserID" json:"-"`
	SendUserNickName string    `gorm:"size:42" json:"send_user_nick_name"` //发送者名称
	SendUserAvatar   string    `json:"send_user_avatar"`                   //发送者头像

	RevUserID       uint      `gorm:"primaryKey" json:"rev_user_id"` //接收者ID
	RevUserModel    UserModel `gorm:"foreignKey:RevUserID" json:"-"`
	RevUserNickName string    `gorm:"size:42" json:"rev_user_nick_name"` //接收者名称
	RevUserAvatar   string    `json:"rev_user_avatar"`                   //接收者头像
	IsRead          bool      `gorm:"default:false" json:"is_read"`      //接收方是否收到消息
	Content         string    `json:"content"`                           //消息内容
}
