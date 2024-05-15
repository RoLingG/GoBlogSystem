package user_api

import (
	"GoRoLingG/global"
	"GoRoLingG/res"
	"GoRoLingG/utils/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func (UserApi) UserLogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")         //从jwt.auth中获取claims
	claims := _claims.(*jwt.CustomClaims) //断言

	token := c.Request.Header.Get("token") //从传过来的header获取token

	//需要计算距离当前时间的token还有多久过期
	exp := claims.ExpiresAt   //获取token过期时间
	now := time.Now()         //获取当前时间
	diff := exp.Time.Sub(now) //用token过期时间-当前时间就算出距离当前还有多久过期
	err := global.Redis.Set(fmt.Sprintf("logout_%s", token), "", diff).Err()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("注销失败", c)
		return
	}
	res.OKWithMsg("注销成功", c)
}
