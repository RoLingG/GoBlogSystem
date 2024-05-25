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
