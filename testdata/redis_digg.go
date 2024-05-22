package main

import (
	"GoRoLingG/core"
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/service/redis_service"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func main() {
	// 读取配置文件
	core.InitConfig()
	// 初始化日志
	global.Log = core.InitLogger()

	global.Redis = core.ConnectRedis()
	global.ESClient = core.ConnectES()

	//手动同步数据
	//把所有文章都获取出来到result里头
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Size(10000).
		Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}

	diggInfo := redis_service.NewArticleDiggIndex().GetInfo()
	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		digg := diggInfo[hit.Id]
		newDigg := article.DiggCount + digg
		if article.DiggCount == newDigg {
			logrus.Info(article.Title, "点赞数无变化")
			continue
		}
		_, err := global.ESClient.
			Update().
			Index(models.ArticleModel{}.Index()).
			Id(hit.Id).
			Doc(map[string]int{
				"digg_count": newDigg,
			}).
			Do(context.Background())
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		logrus.Info(article.Title, "点赞数据同步成功， 点赞数", newDigg)
	}
	redis_service.NewArticleDiggIndex().Clear()
}
