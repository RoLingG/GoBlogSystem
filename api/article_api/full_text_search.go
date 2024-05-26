package article_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

// FullTextSearchView 全文搜索列表
// @Tags 标签管理
// @Summary 全文搜索列表
// @Description	全文搜索列表
// @Param data query models.PageInfo true	"全文搜索列表的一些参数"
// @Router /api/articleFullTextSearch [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.FullTextSearchModel]}
func (ArticleApi) FullTextSearchView(c *gin.Context) {
	var cr models.PageInfo
	_ = c.ShouldBindQuery(&cr)

	boolQuery := elastic.NewBoolQuery()
	if cr.Key != "" {
		boolQuery.Must(elastic.NewMultiMatchQuery(cr.Key, "title", "body"))
	}

	result, err := global.ESClient.
		Search(models.FullTextSearchModel{}.Index()).
		Query(boolQuery).
		Highlight(elastic.NewHighlight().Field("body")). //因为全文搜索搜的是文章内容，所以高亮的是body对应搜索的内容
		Size(100).
		Do(context.Background())
	if err != nil {
		return
	}

	count := result.Hits.TotalHits.Value //搜索到结果总条数
	fmt.Println(count)
	fullTextList := make([]models.FullTextSearchModel, 0)
	for _, hit := range result.Hits.Hits {
		var fullText models.FullTextSearchModel
		json.Unmarshal(hit.Source, &fullText)
		body, ok := hit.Highlight["body"]
		if ok {
			fullText.Body = body[0]
		}

		fullTextList = append(fullTextList, fullText)
	}

	res.OKWithList(fullTextList, count, c)
}
