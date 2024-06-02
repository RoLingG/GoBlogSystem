package feedback_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

// FeedbackRemoveView 反馈删除
// @Tags 反馈管理
// @Summary 反馈删除
// @Description	批量删除反馈
// @Param data body models.RemoveRequest true	"删除反馈的一些参数"
// @Router /api/feedbackRemove [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (FeedbackApi) FeedbackRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var feedbackList []models.FeedBackModel
	count := global.DB.Find(&feedbackList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("反馈不存在，删除失败", c)
		return
	}
	global.DB.Delete(&feedbackList)
	res.OKWithMsg(fmt.Sprintf("反馈删除成功，共删除了 %d 条反馈", count), c)
}
