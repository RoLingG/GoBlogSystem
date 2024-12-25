package large_scale_model_api

import (
	"GoRoLingG/config"
	"GoRoLingG/core"
	"GoRoLingG/global"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
)

// ModelSessionSettingUpdateView 更新大模型会话配置
func (LargeScaleModelApi) ModelSessionSettingUpdateView(c *gin.Context) {
	var cr config.ModelSessionSetting
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	global.Config.LargeScaleModel.ModelSessionSetting = cr
	core.SetYaml()
	res.OKWithMsg("大模型会话配置修改成功", c)
	return
}
