package models

// UserLargeScaleModelRoleModel 用户选择的大模型角色表
type UserLargeScaleModelRoleModel struct {
	Model
	UserID    uint                     `json:"user_id"`
	RoleID    uint                     `json:"role_id"`
	RoleModel LargeScaleModelRoleModel `gorm:"foreignKey:RoleID“ json:"-"`
}
