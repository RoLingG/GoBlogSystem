package log_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/plugins/log_stash_v1"
	"GoRoLingG/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

// LogRemoveListView 日志记录删除
// @Tags 日志管理
// @Summary 日志记录删除
// @Description	将日志记录从日志列表删除
// @Param data body models.RemoveRequest true	"删除日志记录的一些参数"
// @Router /api/logRemove [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (LogApi) LogRemoveListView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var logList []log_stash_v1.LogModel
	count := global.DB.Find(&logList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("日志不存在", c)
		return
	}
	global.DB.Delete(&logList)
	res.OKWithMsg(fmt.Sprintf("共删除 %d 个日志", count), c)
}
