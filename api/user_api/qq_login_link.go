package user_api

import (
	"GoRoLingG/global"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
)

// QQLoginLinkView 获取qq登录的跳转链接
// @Tags 用户管理
// @Summary 获取qq登录的跳转链接
// @Description 获取qq登录的跳转链接,data就是qq的跳转地址
// @Router /api/qqLoginLink [get]
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) QQLoginLinkView(c *gin.Context) {
	path := global.Config.QQ.GetPath()
	if path == "" {
		res.FailWithMsg("未配置qq登录地址", c)
		return
	}
	res.OKWithData(path, c)
	return
}
