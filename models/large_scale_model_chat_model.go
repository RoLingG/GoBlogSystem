package models

// LargeScaleModelChatModel 大模型对话表
type LargeScaleModelChatModel struct {
	Model
	SessionID    uint                        `json:"session_id"`                    // 会话id
	SessionModel LargeScaleModelSessionModel `gorm:"foreignKey:SessionID" json:"-"` // 会话
	Status       bool                        `json:"status"`                        // 状态，ai有没有正常的回复用户
	UserContent  string                      `json:"user_content"`                  // 用户的聊天内容
	AIContent    string                      `json:"ai_content"`                    // ai的回复内容
	RoleID       uint                        `json:"role_id"`                       // 是哪一个角色
	RoleModel    LargeScaleModelRoleModel    `gorm:"foreignKey:RoleID" json:"-"`    // 角色
	UserID       uint                        `json:"user_id"`                       // 用户id
	UserModel    UserModel                   `gorm:"foreignKey:UserID" json:"-"`    // 用户
}
