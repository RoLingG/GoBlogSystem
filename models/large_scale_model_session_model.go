package models

// LargeScaleModelSessionModel 大模型会话表
type LargeScaleModelSessionModel struct {
	Model
	SessionName string                     `json:"session_name"`
	UserID      uint                       `json:"user_id"` // 用户id
	UserModel   UserModel                  `gorm:"foreignKey:UserID" json:"-"`
	RoleID      uint                       `json:"role_id"` // 角色id
	RoleModel   LargeScaleModelRoleModel   `gorm:"foreignKey:RoleID" json:"-"`
	ChatList    []LargeScaleModelChatModel `gorm:"foreignKey:SessionID" json:"-"` // 会话列表
}
