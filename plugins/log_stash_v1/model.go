package log_stash_v1

type LogModel struct {
	ID       uint     `gorm:"primarykey" json:"id"` // 主键ID
	CreateAt string   `json:"create_at"`            // 创建时间
	IP       string   `gorm:"size:32" json:"ip"`
	Addr     string   `gorm:"size:64" json:"addr"`
	Level    LogLevel `gorm:"size:4" json:"level"`     // 日志的等级
	Content  string   `gorm:"size:128" json:"content"` // 日志消息内容
	UserID   uint     `json:"user_id"`                 // 登录用户的用户id，需要自己在查询的时候做关联查询
	NickName string   `json:"nick_name"`               // 登录用户的昵称
}
