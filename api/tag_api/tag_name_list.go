package tag_api

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

type TagResponse struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// TagNameListView 标签名称列表
// @Tags 标签管理
// @Summary 标签名称列表
// @Description 标签名称列表
// @Router /api/tagNameList [get]
// @Produce json
// @Success 200 {object} res.Response{data=[]TagResponse}
func (TagApi) TagNameListView(c *gin.Context) {
	type T struct {
		DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
		SumOtherDocCount        int `json:"sum_other_doc_count"`
		Buckets                 []struct {
			Key      string `json:"key"`
			DocCount int    `json:"doc_count"`
		} `json:"buckets"`
	}
	query := elastic.NewBoolQuery()
	//将领域为tags的每项进行聚合，将文章现有的所有tag都聚合出来
	aggregation := elastic.NewTermsAggregation().Field("tags")
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("tags", aggregation). //将聚合的结果都归类到tags下，以k-v的形式保存
		Size(0).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	byteData := result.Aggregations["tags"] //获取文章中tags的聚合结果
	var tagType T
	_ = json.Unmarshal(byteData, &tagType) //将搜索的结果json解析到类型T的tagType内

	var tagList = make([]TagResponse, 0)
	for _, bucket := range tagType.Buckets {
		tagList = append(tagList, TagResponse{
			Label: bucket.Key, //文章标签的名字
			Value: bucket.Key, //文章标签的名字值(其实就是key对应的value，只不过刚好相等，但是又要符合tagType.Buckets返回的对应type类型结构，只能这样存)
		})
	}

	res.OKWithData(tagList, c)
}
