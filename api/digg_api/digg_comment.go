package digg_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/redis_service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type CommentIDRequest struct {
	ID uint `json:"id" form:"id" uri:"id"`
}

// DiggCommentView 点赞评论
// @Tags 评论管理
// @Summary 点赞评论
// @Description	点赞评论
// @Param id path int true "需要点赞的评论ID"
// @Router /api/diggComment/{id} [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (DiggApi) DiggCommentView(c *gin.Context) {
	var cr CommentIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var commentModel models.CommentModel
	//获取评论，如果获取不到则评论不存在
	err = global.DB.Take(&commentModel, cr.ID).Error
	if err != nil {
		res.FailWithMsg("评论不存在", c)
		return
	}
	redis_service.NewArticleCommentDiggIndex().Set(fmt.Sprintf("%d", cr.ID))
	err = global.DB.Model(&commentModel).Update("digg_count", commentModel.DiggCount+1).Error
	if err != nil {
		res.FailWithMsg("评论点赞失败", c)
		return
	}
	res.OKWithMsg("评论点赞成功", c)
	return

}
