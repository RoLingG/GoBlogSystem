package log_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/plugins/log_stash"
	"GoRoLingG/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (LogApi) LogRemoveListView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var logList []log_stash.LogStashModel
	count := global.DB.Find(&logList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("日志不存在", c)
		return
	}
	global.DB.Delete(&logList)
	res.OKWithMsg(fmt.Sprintf("共删除 %d 个日志", count), c)
}
