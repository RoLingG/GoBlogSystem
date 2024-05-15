package middleware

import (
	"GoRoLingG/global"
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
		//判断token是否在redis中，如果在，则用户注销过，需要重新登录
		keys := global.Redis.Keys("logout_*").Val() //普通的keys返回的是对应条件的所有键值的集合指针(?不知道这样说对不对)，要加上.result()或者.val()才能获取到keys集合
		for _, key := range keys {
			//真实开发环境别用这种方法，会出现阻塞Redis的情况，损耗redis性能
			if "logout_"+token == key {
				res.FailWithMsg("token已失效", c)
				c.Abort()
				return
			}
		}

		//登录的用户
		c.Set("claims", claims)
	}
}
