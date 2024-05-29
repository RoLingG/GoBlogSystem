package log_stash_v2

import (
	"GoRoLingG/global"
	"GoRoLingG/utils"
	"github.com/gin-gonic/gin"
)

// NewSuccessLogin 登录成功的日志
func NewSuccessLogin(c *gin.Context) {
	//从token中获取登陆成功的用户信息
	token := c.Request.Header.Get("token")
	jwyPayLoad := parseToken(token)
	saveLoginLog("登录成功", "——", jwyPayLoad.UserID, jwyPayLoad.UserName, true, c)
}

// NewFailLogin 登录失败的日志
func NewFailLogin(title, userName, pwd string, c *gin.Context) {
	saveLoginLog(title, pwd, 0, userName, false, c)
}

func saveLoginLog(title string, content string, userID uint, userName string, status bool, c *gin.Context) {
	ip := c.ClientIP()
	addr := utils.GetAddr(ip)
	global.DB.Create(&LogStashModel{
		IP:       ip,
		Addr:     addr,
		Title:    title,
		Content:  content,
		UserID:   userID,
		UserName: userName,
		Status:   status,
		Type:     LoginType,
	})
}
