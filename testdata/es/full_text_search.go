package main

import (
	"GoRoLingG/core"
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/service/es_service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func main() {
	core.InitConfig()
	core.InitLogger()
	global.ESClient = core.ConnectES()

	boolSearch := elastic.NewMatchAllQuery()
	articleRes, _ := global.ESClient.Search(models.ArticleModel{}.Index()).Query(boolSearch).Size(1000).Do(context.Background())
	for _, hit := range articleRes.Hits.Hits {
		var article models.ArticleModel
		_ = json.Unmarshal(hit.Source, &article)
		indexList := es_service.GetSearchIndexDataByContent(hit.Id, article.Title, article.Content)
		bulk := global.ESClient.Bulk()
		for _, indexData := range indexList {
			request := elastic.NewBulkIndexRequest().Index(models.FullTextSearchModel{}.Index()).Doc(indexData)
			bulk.Add(request)
		}
		fullTextSearchRes, err := bulk.Do(context.Background())
		if err != nil {
			logrus.Error(err)
			continue
		}
		fmt.Println(article.Title, "添加成功", "共", len(fullTextSearchRes.Succeeded()), "条")
	}
}
