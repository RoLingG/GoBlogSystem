package settings_api

import (
	"GoRoLingG/global"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	switch cr.Name {
	case "site":
		res.OKWithData(global.Config.SiteInfo, c)
	case "email":
		emailInfo := global.Config.Email
		emailInfo.Password = "******"
		res.OKWithData(global.Config.Email, c)
	case "qq":
		qqInfo := global.Config.QQ
		qqInfo.Key = "******"
		res.OKWithData(global.Config.QQ, c)
	case "qiniu":
		qiNiuInfo := global.Config.QiNiu
		qiNiuInfo.SecretKey = "******"
		res.OKWithData(global.Config.QiNiu, c)
	case "jwt":
		jwt := global.Config.JWT
		jwt.Secret = "******"
		res.OKWithData(global.Config.JWT, c)
	case "chat_group":
		res.OKWithData(global.Config.ChatGroup, c)
	case "large_scale_model":
		res.OKWithData(global.Config.LargeScaleModel.ModelSetting, c)
	default:
		res.FailWithMsg("没有对应的配置信息", c)
	}
}
