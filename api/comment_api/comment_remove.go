package comment_api

import (
	"GoRoLingG/api/digg_api"
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/redis_service"
	"GoRoLingG/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CommentRemoveView 删除评论
// @Tags 评论管理
// @Summary 删除评论
// @Description	删除现有的评论
// @Param id path int true "需要删除的评论ID"
// @Router /api/commentRemove/{id} [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (CommentApi) CommentRemoveView(c *gin.Context) {
	var cr digg_api.CommentIDRequest
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
	//统计要删除评论下的子评论数，总删除量要把评论本身算上
	subCommentList := FindSubCommentCount(commentModel)
	CommentCount := len(subCommentList) + 1
	redis_service.NewArticleCommentIndex().SetCount(commentModel.ArticleID, -CommentCount)

	//判断是否是子评论
	if commentModel.ParentCommentID != nil {
		//要删的评论是子评论，则要找父评论，减掉对应删除的评论数量
		global.DB.Model(&models.CommentModel{}).
			Where("id = ?", *commentModel.ParentCommentID).
			Update("comment_count", gorm.Expr("comment_count - ?", CommentCount))
	}

	//删除当前评论以及其子评论
	var deleteCommentIDList []uint
	for _, model := range subCommentList {
		//将当前评论下的所有子评论id加进要删除的评论id列表中
		deleteCommentIDList = append(deleteCommentIDList, model.ID)
	}
	// 反转后逐个删除
	utils.Reverse(deleteCommentIDList)
	//记得把评论本身也删了，别光删子评论了
	deleteCommentIDList = append(deleteCommentIDList, commentModel.ID)
	for _, id := range deleteCommentIDList {
		global.DB.Model(models.CommentModel{}).Delete("id = ?", id)
	}

	res.OKWithMsg(fmt.Sprintf("共删除 %d 条评论", len(deleteCommentIDList)), c)
	return
}
