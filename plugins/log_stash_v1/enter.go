package log_stash_v1

import (
	"GoRoLingG/global"
	"GoRoLingG/utils"
	"GoRoLingG/utils/jwt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Log struct {
	ip     string `json:"ip"`
	addr   string `json:"addr"`
	userId uint   `json:"user_id"`
}

func New(ip string, token string) *Log {
	// 解析token
	claims, err := jwt.ParseToken(token)
	var userID uint
	if err == nil {
		userID = claims.UserID
	}
	addr := utils.GetAddr(ip)

	// 拿到用户id
	return &Log{
		ip:     ip,
		addr:   addr,
		userId: userID,
	}
}

func NewLogByGin(c *gin.Context) *Log {
	ip := c.ClientIP()
	token := c.Request.Header.Get("token")
	// 检查 token 是否为空
	if token == "" {
		// 可以选择返回一个错误或者一个默认的 Log 实例
		logrus.Warn("No token found in request header")
		return &Log{ip: ip, addr: "未知", userId: 0}
	}
	return New(ip, token)
}

func (l Log) Debug(content string) {
	l.send(DebugLevel, content)
}
func (l Log) Info(content string) {
	l.send(InfoLevel, content)
}
func (l Log) Warning(content string) {
	l.send(WarningLevel, content)
}
func (l Log) Error(content string) {
	l.send(ErrorLevel, content)
}

func (l Log) send(level LogLevel, content string) {
	err := global.DB.Create(&LogModel{
		IP:      l.ip,
		Addr:    l.addr,
		Level:   level,
		Content: content,
		UserID:  l.userId,
	}).Error
	if err != nil {
		logrus.Error(err)
	}
}
