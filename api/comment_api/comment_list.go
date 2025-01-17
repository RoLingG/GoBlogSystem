package comment_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/redis_service"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

type CommentListRequest struct {
	ArticleID string `json:"article_id"`
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
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//使用filter，过滤掉不必要显示的数据信息
	rootArticleComment := FindArticleCommentList(cr.ArticleID)
	data := filter.Select("comment", rootArticleComment)
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

func FindArticleCommentList(articleID string) (rootCommentList []*models.CommentModel) {
	//先把文章下的根评论查出来,同时查出该评论的发布人部分信息
	global.DB.Preload("User").Find(&rootCommentList, "article_id = ? and parent_comment_id is null", articleID)
	//遍历父评论，将对应的子评论给递归出来
	diggInfo := redis_service.NewArticleCommentDiggIndex().GetInfo()
	for _, comment := range rootCommentList {
		//modelDigg := diggInfo[fmt.Sprintf("%d", comment.ID)]
		commentDigg := diggInfo[string(comment.ID)]
		comment.DiggCount = comment.DiggCount + commentDigg
		models.GetCommentTree(comment)
	}
	return
}

// FindSubComment 递归查评论下的子评论
func FindSubComment(model models.CommentModel, subCommentList *[]models.CommentModel) {
	//根评论的子评论查出来,同时查出该评论的发布人部分信息
	global.DB.Preload("SubComments.User").Take(&model)
	for _, sub := range model.SubComments {
		*subCommentList = append(*subCommentList, sub)
		FindSubComment(sub, subCommentList) //子评论的子评论也是在根评论的下一级，属于同级关系，不会出现一直套娃的现象
	}
	return
}

func FindSubCommentCount(model models.CommentModel) (subCommentList []models.CommentModel) {
	findSubCommentList(model, &subCommentList)
	return subCommentList
}

func findSubCommentList(model models.CommentModel, subCommentList *[]models.CommentModel) {
	//根评论的子评论查出来,同时查出该评论的发布人部分信息
	global.DB.Preload("SubComments").Take(&model)
	for _, sub := range model.SubComments {
		*subCommentList = append(*subCommentList, sub)
		FindSubComment(sub, subCommentList) //子评论的子评论也是在根评论的下一级，属于同级关系，不会出现一直套娃的现象
	}
	return
}
