package log_v2_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/plugins/log_stash_v2"
	"GoRoLingG/res"
	"GoRoLingG/service/common"
	"github.com/gin-gonic/gin"
	"time"
)

type LogListRequest struct {
	models.PageInfo
	Level    log_stash_v2.LogLevel `json:"level" form:"level"`       // 日志查询的等级
	Type     log_stash_v2.LogType  `json:"type" form:"type"`         // 日志的类型   1 登录日志  2 操作日志  3 运行日志
	IP       string                `json:"ip" form:"ip"`             // 根据ip查询
	UserID   uint                  `json:"userID" form:"userID"`     // 根据用户id查询
	Addr     string                `json:"addr" form:"addr"`         // 感觉地址查询
	Date     string                `json:"date" form:"date"`         // 查某一天的，格式是年月日
	Status   *bool                 `json:"status" form:"status"`     // 登录状态查询  true  成功  false 失败
	UserName string                `json:"userName" form:"userName"` // 查用户名
}

// LogListView 日志列表
func (LogV2Api) LogListView(c *gin.Context) {
	var cr LogListRequest
	c.ShouldBindQuery(&cr)

	var query = global.DB.Where("")
	if cr.Date != "" {
		_, dateTimeErr := time.Parse("2006-01-02", cr.Date)
		if dateTimeErr != nil {
			res.FailWithMsg("时间格式错误", c)
			return
		}
		query.Where("date(created_at) = ?", cr.Date)
	}
	if cr.Status != nil {
		query.Where("status = ?", cr.Status)
	}

	_list, count, _ := common.CommonList(log_stash_v2.LogStashModel{
		Type:     cr.Type,
		LogLevel: cr.Level,
		IP:       cr.IP,
		Addr:     cr.Addr,
		UserID:   cr.UserID,
		UserName: cr.UserName,
	}, common.Option{
		PageInfo: cr.PageInfo,
		Where:    query,
		Likes:    []string{"title", "user_name"},
	})
	res.OKWithList(_list, count, c)
}
