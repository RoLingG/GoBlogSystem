package log_stash

import (
	"GoRoLingG/global"
	"GoRoLingG/utils/jwt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Log struct {
	ip     string `json:"ip"`
	addr   string `json:"addr"`
	userID uint   `json:"user_id"` //用户ID存在Token里
}

func New(ip string, token string) *Log {
	// 解析token
	claims, err := jwt.ParseToken(token)
	var userID uint
	if err == nil {
		userID = claims.UserID
	}
	// 拿到用户id
	return &Log{
		ip:     ip,
		addr:   "内网",
		userID: userID,
	}
}

func NewLogByGin(c *gin.Context) *Log {
	ip := c.ClientIP()
	token := c.Request.Header.Get("token")
	return New(ip, token)
}

// 入库
func (log Log) send(level LogLevel, content string) {
	err := global.DB.Create(&LogStashModel{
		IP:       log.ip,
		Addr:     log.addr,
		LogLevel: level,
		Content:  content,
		UserID:   log.userID,
	}).Error
	if err != nil {
		logrus.Error(err)
	}
}

// 重写logrus里面的方法

func (log Log) Debug(content string) {
	log.send(DebugLevel, content)
}
func (log Log) Info(content string) {
	log.send(InfoLevel, content)
}
func (log Log) Warning(content string) {
	log.send(WarningLevel, content)
}
func (log Log) Error(content string) {
	log.send(ErrorLevel, content)
}
