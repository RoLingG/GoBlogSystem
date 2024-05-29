package log_v2_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/plugins/log_stash_v2"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
)

// LogReadView 日志读取
// @Tags 日志管理V2
// @Summary 日志读取
// @Description 日志读取
// @Description 1. 前端判断这个日志的读取状态，未读就去请求这个接口，让这个日志变成已读的
// @Description 2. 如果是已读状态，就不需要调这个接口了
// @Param token header string true "Token"
// @Param data query models.ESIDRequest true "参数"
// @Router /api/logV2Read [get]
// @Produce json
// @Success 200 {object} res.Response{}
func (LogV2Api) LogReadView(c *gin.Context) {
	var cr models.ESIDRequest //这里图方便不想写多一个结构体就直接用es的了
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithMsg("参数错误", c)
		return
	}
	var logStash log_stash_v2.LogStashModel
	err = global.DB.Take(&logStash, cr.ID).Error
	if err != nil {
		res.FailWithMsg("日志不存在", c)
		return
	}
	if logStash.ReadStatus {
		res.OKWithData("日志读取成功", c)
		return
	}
	//日志读取成功后，将该日志的状态变为true，视为"已读"
	global.DB.Model(&logStash).Update("readStatus", true)
	res.OKWithMsg("日志读取成功", c)
	return
}
