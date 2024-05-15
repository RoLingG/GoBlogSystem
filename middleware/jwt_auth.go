package middleware

import (
	"GoRoLingG/res"
	"GoRoLingG/service"
	"GoRoLingG/utils/jwt"
	"github.com/gin-gonic/gin"
)

// JwtAuth 普通用户权限中间件
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMsg("token为空，未登录", c)
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			res.FailWithMsg("token解析错误", c)
			c.Abort()
			return
		}
		//判断token是否在redis中，如果在，则用户注销过，需要重新登录
		if service.Service.RedisService.CheckLogout(token) {
			res.FailWithMsg("token已失效，请重新登录", c)
			c.Abort()
			return
		}

		//登录的用户
		c.Set("claims", claims)
	}
}
