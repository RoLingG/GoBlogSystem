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
		var Update_info config.SiteInfo
		err := c.ShouldBindJSON(&Update_info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
		}
		global.Config.SiteInfo = Update_info
		res.OKWithData(global.Config.SiteInfo, c)
	case "email":
		var Update_info config.Email
		err := c.ShouldBindJSON(&Update_info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
		}
		global.Config.Email = Update_info
		res.OKWithData(global.Config.Email, c)
	case "qq":
		var Update_info config.QQ
		err := c.ShouldBindJSON(&Update_info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
		}
		global.Config.QQ = Update_info
		res.OKWithData(global.Config.QQ, c)
	case "qiniu":
		var Update_info config.QiNiu
		err := c.ShouldBindJSON(&Update_info)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
		}
		global.Config.QiNiu = Update_info
		res.OKWithData(global.Config.QiNiu, c)
	case "jwt":
		var updateInfo config.JWT
		err := c.ShouldBindJSON(&updateInfo)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
		}
		global.Config.JWT = updateInfo
		res.OKWithData(global.Config.JWT, c)
	case "chat_group":
		var chatGroupInfo config.ChatGroup
		err = c.ShouldBindJSON(&chatGroupInfo)
		if err != nil {
			res.FailWithCode(res.ArgumentError, c)
			return
		}
		global.Config.ChatGroup = chatGroupInfo
		res.OKWithData(global.Config.ChatGroup, c)
	default:
		res.FailWithMsg("没有对应的配置信息", c)
		return
	}
	core.SetYaml()
}
