package feedback_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/utils"
	"github.com/fatih/structs"
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

	//检测邮箱是否合法
	if !utils.IsValidEmail(cr.Email) {
		res.FailWithMsg("邮箱输入非法，请重新输入", c)
		return
	}

	var feedback = models.FeedBackModel{
		Email:   cr.Email,
		Content: cr.Content,
	}
	maps := structs.Map(feedback)
	err = global.DB.Create(&feedback).Error
	if err != nil {
		res.FailWithMsg("反馈失败", c)
		return
	}

	res.OKWithDataAndMsg(maps, "反馈成功", c)
}
