package settings_api

import (
	"GoRoLingG/config"
	"GoRoLingG/core"
	"GoRoLingG/global"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoUpdate(c *gin.Context) {
	var cr config.SiteInfo
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	global.Config.SiteInfo = cr
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.OKWithoutData(c)
}
