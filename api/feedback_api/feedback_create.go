package feedback_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
)

type FeedbackCreateRequest struct {
	Email   string `json:"email"`
	Content string `json:"content"`
}

// FeedbackCreateView 创建反馈
// @Tags 反馈管理
// @Summary 创建反馈
// @Description	创建反馈
// @Param data body FeedbackCreateRequest true	"创建反馈的一些参数"
// @Router /api/feedbackCreate [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (FeedbackApi) FeedbackCreateView(c *gin.Context) {
	var cr FeedbackCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	err = global.DB.Create(&models.FeedBackModel{
		Email:   cr.Email,
		Content: cr.Content,
	}).Error
	if err != nil {
		res.FailWithMsg("反馈失败", c)
		return
	}

	res.OKWithMsg("反馈成功", c)
}
