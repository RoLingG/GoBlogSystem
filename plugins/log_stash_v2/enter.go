package log_stash_v2

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
	userID uint   `json:"user_id"` //用户ID存在Token里
}

func New(ip string, token string) *Log {
	var userID uint
	// 解析token
	if token != "" {
		claims, err := jwt.ParseToken(token)
		if err == nil {
			userID = claims.UserID
		}
	}

	addr := utils.GetAddr(ip)
	// 拿到用户id
	return &Log{
		ip:     ip,
		addr:   addr,
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
