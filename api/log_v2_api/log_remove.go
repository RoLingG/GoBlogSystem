package log_v2_api

import (
	"GoRoLingG/global"
	"GoRoLingG/plugins/log_stash_v2"
	"GoRoLingG/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type LogRemoveRequest struct {
	IDList    []uint `json:"id_list"`   // 可以传id列表删除
	StartTime string `json:"startTime"` // 年月日格式的开始时间
	EndTime   string `json:"endTime"`   // 年月日格式的结束时间
	UserID    uint   `json:"userID"`    // 根据用户删除日志
	IP        string `json:"ip"`        // 根据用户ip删除
}

// LogRemoveView 删除日志
// @Tags 日志管理V2
// @Summary 删除日志
// @Description 删除日志
// @Param token header string true "Token"
// @Param data body LogRemoveRequest true "参数"
// @Router /api/logV2Remove [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (LogV2Api) LogRemoveView(c *gin.Context) {
	var cr LogRemoveRequest
	logStash := log_stash_v2.NewAction(c)
	logStash.SetRequest(c)
	logStash.SetResponse(c)

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var logs []log_stash_v2.LogStashModel
	//多重判断删除，如果id列表找得到就用id列表删，如果找不到就用对应给的userID删，如果还找不到就找对应的IP删，如果还还还找不到，就按创建时间与结束时间删
	if len(cr.IDList) > 0 {
		logStash.SetItemInfo("IDList", cr.IDList)
		global.DB.Find(&logs, cr.IDList).Delete(&logs)
	} else if cr.UserID != 0 {
		global.DB.Find(&logs, "user_id = ?", cr.UserID).Delete(&logs)
	} else if cr.IP != "" {
		global.DB.Find(&logs, "ip = ?", cr.IP).Delete(&logs)
	} else if cr.StartTime != "" && cr.EndTime != "" {
		_, startTimeErr := time.Parse("2006-01-02", cr.StartTime)
		_, endTimeErr := time.Parse("2006-01-02", cr.EndTime)
		if startTimeErr != nil {
			res.FailWithMsg("开始时间格式错误", c)
			return
		}
		if endTimeErr != nil {
			res.FailWithMsg("结束时间格式错误", c)
			return
		}
		global.DB.Find(&logs, "create_at > date(?) and create_at < date(?)", cr.StartTime, cr.EndTime).Delete(&logs)
	}

	logStash.SetItemInfo("共删除日志", len(logs))
	logStash.Info("删除日志成功")

	res.OKWithMsg(fmt.Sprintf("共删除 %d 条日志", len(logs)), c)
}
