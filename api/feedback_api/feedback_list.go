package feedback_api

import (
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/common"
	"github.com/gin-gonic/gin"
)

// FeedbackListView 反馈列表
// @Tags 反馈管理
// @Summary 反馈列表
// @Description	查询所有反馈的列表
// @Param data body FeedbackCreateRequest true	"查询反馈的一些参数"
// @Router /api/feedbackList [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.FeedBackModel]}
func (FeedbackApi) FeedbackListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	list, count, err := common.CommonList(models.FeedBackModel{}, common.Option{
		PageInfo: cr,
	})

	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}

	res.OKWithList(list, count, c)
	return
}
