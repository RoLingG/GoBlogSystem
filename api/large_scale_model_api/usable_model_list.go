package large_scale_model_api

import (
	"GoRoLingG/global"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
)

// UsableModelListView 可用大模型列表接口
func (LargeScaleModelApi) UsableModelListView(c *gin.Context) {
	res.OKWithData(global.Config.LargeScaleModel.ModelOption, c)
	return
}
