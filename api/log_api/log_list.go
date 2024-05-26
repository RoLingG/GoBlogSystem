package log_api

import (
	"GoRoLingG/models"
	"GoRoLingG/plugins/log_stash"
	"GoRoLingG/res"
	"GoRoLingG/service/common"
	"github.com/gin-gonic/gin"
)

type LogRequest struct {
	models.PageInfo
	Level log_stash.LogLevel `form:"level"`
}

// LogListView 日志列表
// @Tags 日志管理
// @Summary 日志列表
// @Description	查询日志列表
// @Param data query LogRequest true	"查询日志列表的一些参数"
// @Router /api/logList [get]
// @Produce json
// @Success 200 {object} res.Response{data=log_stash.LogStashModel}
func (LogApi) LogListView(c *gin.Context) {
	var cr LogRequest
	c.ShouldBindQuery(&cr)
	list, count, _ := common.CommonList(log_stash.LogStashModel{LogLevel: cr.Level}, common.Option{
		PageInfo: cr.PageInfo,
		Debug:    true,
		Likes:    []string{"ip", "addr"},
	})
	res.OKWithList(list, count, c)
	return
}
