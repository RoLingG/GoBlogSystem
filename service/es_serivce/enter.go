package es_serivce

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func CommonList(key string, limit, page int) (list []models.ArticleModel, count int, err error) {
	boolSearch := elastic.NewBoolQuery()
	from := page
	if key != "" {
		boolSearch.Must(
			elastic.NewMatchQuery("title", key),
		)
	}
	if limit == 0 {
		limit = 10
	}
	if from == 0 {
		from = 1
	}

	res, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		From((from - 1) * limit).
		Size(limit).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return
	}

	count = int(res.Hits.TotalHits.Value) //搜索到结果总条数
	articleList := []models.ArticleModel{}
	for _, hit := range res.Hits.Hits {
		var article models.ArticleModel
		data, err := hit.Source.MarshalJSON()
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		err = json.Unmarshal(data, &article)
		if err != nil {
			logrus.Error(err)
			continue
		}
		article.ID = hit.Id
		articleList = append(articleList, article)
	}
	return articleList, count, err
}

func CommonDetail(id string) (article models.ArticleModel, err error) {
	//get获取es内对应index(也就是article_index)中_id为id的文章数据
	res, err := global.ESClient.Get().
		Index(models.ArticleModel{}.Index()).
		Id(id).
		Do(context.Background())
	if err != nil {
		return
	}
	err = json.Unmarshal(res.Source, &article)
	if err != nil {
		return
	}
	article.ID = res.Id
	return
}
