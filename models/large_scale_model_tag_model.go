package models

// LargeScaleModelTagModel 大模型角色标签表
type LargeScaleModelTagModel struct {
	Model
	RoleTitle string                     `gorm:"size:16" json:"role_title"`                                                                                                                             // 标签的名称
	Color     string                     `gorm:"size:16" json:"color"`                                                                                                                                  // 颜色
	Roles     []LargeScaleModelRoleModel `gorm:"many2many:large_scale_model_role_tag_models;joinForeignKey:large_scale_model_tag_model_id;JoinReferences:large_scale_model_role_model_id" json:"roles"` // 角色列表
}
