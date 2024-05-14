package middleware

import (
	"GoRoLingG/models/ctype"
	"GoRoLingG/res"
	"GoRoLingG/utils/jwt"
	"github.com/gin-gonic/gin"
)

// 普通用户权限中间件
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
		//登录的用户
		c.Set("claims", claims)
	}
}

// 管理员调用中间件
func JwtAdmin() gin.HandlerFunc {
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
		//登录的用户
		if claims.Role != int(ctype.PermissionAdmin) {
			res.FailWithMsg("权限错误，非管理员权限", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}
