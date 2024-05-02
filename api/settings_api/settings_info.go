package settings_api

import (
	"GoRoLingG/global"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	res.OKWithData(global.Config.SiteInfo, c)
}
