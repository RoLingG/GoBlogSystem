package comment_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

type CommentListRequest struct {
	ArticleID string `json:"article_id"`
}

func (CommentApi) CommentListView(c *gin.Context) {
	var cr CommentListRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//使用filter，过滤掉不必要显示的数据信息
	rootArticleComment := FindArticleCommmentList(cr.ArticleID)
	res.OKWithData(filter.Omit("comment", rootArticleComment), c)
	return
}

func FindArticleCommmentList(articleID string) (rootCommentList []*models.CommentModel) {
	//先把文章下的根评论查出来,同时查出该评论的发布人部分信息
	global.DB.Preload("User").Find(&rootCommentList, "article_id = ? and parent_comment_id is null", articleID)
	//遍历父评论，将对应的子评论给递归出来
	for _, model := range rootCommentList {
		var subCommentList []models.CommentModel
		FindSubComment(*model, &subCommentList)
		model.SubComments = subCommentList
		fmt.Println(model.Content, subCommentList)
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
