package comment_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/redis_service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

type CommentListRequest struct {
	ArticleID string `json:"article_id" form:"article_id" uri:"article_id"`
}

// CommentListView 文章评论列表
// @Tags 评论管理
// @Summary 文章评论列表
// @Description	查询文章评论列表
// @Param data body CommentListRequest true	"查询评论列表的一些参数"
// @Router /api/commentList [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.CommentModel]}
func (CommentApi) CommentListView(c *gin.Context) {
	var cr CommentListRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//使用filter，过滤掉不必要显示的数据信息
	rootArticleComment := FindArticleCommentList(cr.ArticleID)
	data := filter.Omit("comment", rootArticleComment)
	_list, _ := data.(filter.Filter)
	//如果_list为空，则不让它按json为空显示，而是string的""显示
	if string(_list.MustMarshalJSON()) == "{}" {
		list := make([]models.CommentModel, 0)
		res.OKWithList(list, 0, c)
		return
	}
	res.OKWithList(data, int64(len(rootArticleComment)), c)
	return
}

//func FindArticleCommentList(articleID string) (rootCommentList []*models.CommentModel) {
//	//先把文章下的根评论查出来,同时查出该评论的发布人部分信息
//	global.DB.Preload("User").Find(&rootCommentList, "article_id = ? and parent_comment_id is null", articleID)
//	//遍历父评论，将对应的子评论给递归出来
//	diggInfo := redis_service.NewArticleCommentDiggIndex().GetInfo()
//	for _, comment := range rootCommentList {
//		//modelDigg := diggInfo[fmt.Sprintf("%d", comment.ID)]
//		commentDigg := diggInfo[string(comment.ID)]
//		comment.DiggCount = comment.DiggCount + commentDigg
//		models.GetCommentTree(comment)
//	}
//	return
//}

func FindArticleCommentList(articleID string) (rootCommentList []*models.CommentModel) {
	// 先把文章下的根评论查出来
	global.DB.Preload("User").Find(&rootCommentList, "article_id = ? and parent_comment_id is null", articleID)
	// 遍历根评论，递归查根评论下的所有子评论
	diggInfo := redis_service.NewArticleCommentDiggIndex().GetInfo()
	for _, model := range rootCommentList {
		modelDigg := diggInfo[fmt.Sprintf("%d", model.ID)]
		model.DiggCount = model.DiggCount + modelDigg
		models.GetCommentTree(model)
	}
	return
}
