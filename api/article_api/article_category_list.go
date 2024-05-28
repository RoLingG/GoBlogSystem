package article_api

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

type CategoryResponse struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

//写法基本和标签名称列表一样

// ArticleCategoryListView 文章分类列表
// @Tags 文章管理
// @Summary 文章分类列表
// @Description 文章分类列表
// @Router /api/articleCategoryList [get]
// @Produce json
// @Success 200 {object} res.Response{data=[]CategoryResponse}
func (ArticleApi) ArticleCategoryListView(c *gin.Context) {
	type T struct {
		DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
		SumOtherDocCount        int `json:"sum_other_doc_count"`
		Buckets                 []struct {
			Key      string `json:"key"`
			DocCount int    `json:"doc_count"`
		} `json:"buckets"`
	}

	aggregation := elastic.NewTermsAggregation().Field("category") //聚合es内文章索引内的所有文章分类
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewBoolQuery()).
		Aggregation("category", aggregation).
		Size(0).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	byteData := result.Aggregations["category"]
	var categoryType T
	_ = json.Unmarshal(byteData, &categoryType)
	var categoryList = make([]CategoryResponse, 0)
	for _, bucket := range categoryType.Buckets {
		categoryList = append(categoryList, CategoryResponse{
			Label: bucket.Key,
			Value: bucket.Key,
		})
	}
	res.OKWithData(categoryList, c)
}
