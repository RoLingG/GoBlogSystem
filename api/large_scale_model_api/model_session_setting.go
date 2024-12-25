package large_scale_model_api

import (
	"GoRoLingG/global"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
)

// ModelSessionSettingView 大模型会话配置信息获取
func (LargeScaleModelApi) ModelSessionSettingView(c *gin.Context) {
	res.OKWithData(global.Config.LargeScaleModel.ModelSessionSetting, c)
	return
}
