package settings_api

import (
	"GoRoLingG/config"
	"GoRoLingG/core"
	"GoRoLingG/global"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
)

/*
	之所以不会有mysql那些是因为那些要一开始就配置好，这些可以动态修改的配置才有意义。
*/

// 接口程序数量减少，但是代码增多(待优化)
// 某一项配置信息更新

func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	switch cr.Name {
	case "site":
		var updateInfo config.SiteInfo
		err := c.ShouldBindJSON(&updateInfo)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
		}
		global.Config.SiteInfo = updateInfo
		res.OKWithData(global.Config.SiteInfo, c)
	case "email":
		var updateInfo config.Email
		err := c.ShouldBindJSON(&updateInfo)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
		}
		global.Config.Email = updateInfo
		res.OKWithData(global.Config.Email, c)
	case "qq":
		var updateInfo config.QQ
		err := c.ShouldBindJSON(&updateInfo)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
		}
		global.Config.QQ = updateInfo
		res.OKWithData(global.Config.QQ, c)
	case "qiniu":
		var updateInfo config.QiNiu
		err := c.ShouldBindJSON(&updateInfo)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
		}
		global.Config.QiNiu = updateInfo
		res.OKWithData(global.Config.QiNiu, c)
	case "jwt":
		var jwtInfo config.JWT
		err := c.ShouldBindJSON(&jwtInfo)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
		}
		global.Config.JWT = jwtInfo
		res.OKWithData(global.Config.JWT, c)
	case "chat_group":
		var chatInfo config.ChatGroup
		err = c.ShouldBindJSON(&chatInfo)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.ChatGroup = chatInfo
		res.OKWithData(global.Config.ChatGroup, c)
	default:
		res.FailWithMsg("没有对应的配置信息", c)
		return
	}
	core.SetYaml()
}
