package large_scale_model_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
)

// LargeScaleModelTagOptionListView 大模型角色标签选项接口
func (LargeScaleModelApi) LargeScaleModelTagOptionListView(c *gin.Context) {
	var optionList []models.Options[uint]
	global.DB.Model(models.LargeScaleModelTagModel{}).Select("id as value", "role_title as label").Scan(&optionList)
	res.OKWithData(optionList, c)
}
