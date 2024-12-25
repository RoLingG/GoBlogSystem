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
	//之所以有下面这个结构体是因为聚合之后的结果byteData反序列化解构之后对应的结构体就长这样（
	type T struct {
		DocCountErrorUpperBound int        `json:"doc_count_error_upper_bound"`
		SumOtherDocCount        int        `json:"sum_other_doc_count"`
		Buckets                 []struct { //这里面存的就是聚合之后的结果了
			Key      string `json:"key"`
			DocCount int    `json:"doc_count"`
		} `json:"buckets"`
	}
	query := elastic.NewBoolQuery()
	//创建一个名为tags的聚合
	aggregation := elastic.NewTermsAggregation().Field("tags")
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("tags", aggregation). //将es索引内字段为tags聚合的结果都归类到aggregation聚合下，以k-v的形式保存
		Size(0).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	byteData := result.Aggregations["tags"] //获取文章中tags字段的聚合结果
	var tagType T
	_ = json.Unmarshal(byteData, &tagType) //将搜索的结果反序列化到类型T的tagType内

	var tagList = make([]TagResponse, 0)
	for _, bucket := range tagType.Buckets {
		tagList = append(tagList, TagResponse{
			Label: bucket.Key, //文章标签的名字
			Value: bucket.Key, //文章标签的名字值(其实就是key对应的value，只不过刚好相等，但是又要符合tagType.Buckets返回的对应type类型结构，只能这样存)
		})
	}

	res.OKWithData(tagList, c)
}
