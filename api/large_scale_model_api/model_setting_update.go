package large_scale_model_api

import (
	"GoRoLingG/config"
	"GoRoLingG/core"
	"GoRoLingG/global"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
)

// ModelSettingUpdateView 更新大模型配置
func (LargeScaleModelApi) ModelSettingUpdateView(c *gin.Context) {
	var cr config.ModelSetting
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var ok bool
	for _, option := range global.Config.LargeScaleModel.ModelOption {
		if option.Value == cr.Name {
			ok = true
			break
		}
	}
	if !ok {
		res.FailWithMsg("大模型名称错误", c)
		return
	}

	global.Config.LargeScaleModel.ModelSetting = cr
	core.SetYaml()
	res.OKWithMsg("大模型配置修改成功", c)
	return
}
