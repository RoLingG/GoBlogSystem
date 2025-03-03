package comment_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

type CommentByArticleListRequest struct {
	models.PageInfo
	Title string `json:"title" form:"title"`
}

type CommentByArticleListResponse struct {
	Title string `json:"title"`
	ID    string `json:"id"`
	Count int    `json:"count"`
}

// CommentByArticleListView 有评论的文章列表
// @Tags 评论管理
// @Summary 有评论的文章列表
// @Description 有评论的文章列表
// @Param id path string  true  "ID"
// @Param data query CommentByArticleListRequest  true  "参数"
// @Router /api/commentsByArticle [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[CommentByArticleListResponse]}
func (CommentApi) CommentByArticleListView(c *gin.Context) {
	var cr CommentByArticleListRequest
	c.ShouldBindQuery(&cr)

	var count int64

	global.DB.Model(models.CommentModel{}).Group("article_id").Count(&count)

	type T struct {
		ArticleID string
		Count     int
	}

	if cr.Limit == 0 {
		cr.Limit = 10
	}

	offset := (cr.Page - 1) * cr.Limit

	var _list []T
	global.DB.Model(models.CommentModel{}).
		Group("article_id").Order("count desc").Limit(cr.Limit).Offset(offset).Select("article_id", "count(id) as count").Scan(&_list)

	var articleIDCountMap = map[string]int{}
	var articleIDList []interface{}
	for _, t := range _list {
		articleIDCountMap[t.ArticleID] = t.Count
		articleIDList = append(articleIDList, t.ArticleID)
	}

	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewTermsQuery("_id", articleIDList...)).
		Size(10000).
		Do(context.Background())
	if err != nil {
		res.FailWithMsg("es查询错误", c)
		return
	}

	var list = make([]CommentByArticleListResponse, 0)
	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			logrus.Error(err)
			continue
		}

		article.ID = hit.Id

		list = append(list, CommentByArticleListResponse{
			Title: article.Title,
			ID:    hit.Id,
			Count: articleIDCountMap[hit.Id],
		})
	}

	res.OKWithList(list, count, c)
	return
}
