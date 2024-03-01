package settings_api

import (
	"GoRoLingG/models/res"
	"github.com/gin-gonic/gin"
)

func (SettingsApi) SettingsInfoView(c *gin.Context) {
	//c.JSON(200, gin.H{"Msg": "settings测试"})

	//res.OK(map[string]string{}, "测试OK接口", c)

	//res.OKWithData(map[string]string{
	//	"id": "测试OKWithData",
	//}, c)

	res.FailWithCode(2, c)
}
