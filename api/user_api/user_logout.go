package user_api

import (
	"GoRoLingG/global"
	"GoRoLingG/res"
	"GoRoLingG/service"
	"GoRoLingG/utils/jwt"
	"github.com/gin-gonic/gin"
)

// UserLogoutView 用户注销账号
// @Tags 用户管理
// @Summary	用户注销账号
// @Description	用户注销账号
// @Param	token header string true "Authorization token"
// @Router /api/userLogout [post]
// @Produce	json
// @Success	200 {subject} res.Response{}
func (UserApi) UserLogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")         //从jwt.auth中获取claims
	claims := _claims.(*jwt.CustomClaims) //断言

	token := c.Request.Header.Get("token") //从传过来的header获取token

	//通过service内的方法去进行用户注销操作，涉及redis
	err := service.Service.UserService.UserLogoutService(claims, token)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("注销失败", c)
		return
	}
	res.OKWithMsg("注销成功", c)
}
