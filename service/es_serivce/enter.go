package es_serivce

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/service"
	"context"
	"encoding/json"
	"errors"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"strings"
)

type Option struct {
	models.PageInfo
	Fields []string
	Tag    string
}

func (o *Option) GetForm() int {
	if o.Page == 0 {
		o.Page = 1
	}
	if o.Limit == 0 {
		o.Limit = 10
	}
	return (o.Page - 1) * o.Limit
}

func CommonList(option Option) (list []models.ArticleModel, count int, err error) {
	boolSearch := elastic.NewBoolQuery()
	if option.Key != "" {
		boolSearch.Must(
			elastic.NewMultiMatchQuery(option.Key, option.Fields...),
		)
	}
	if option.Tag != "" {
		boolSearch.Must(
			elastic.NewMultiMatchQuery(option.Tag, "tags"),
		)
	}

	//排序查询
	type SortField struct {
		Field     string //以什么排序
		Ascending bool   //排序的类型
	}
	//默认按照什么排序
	sortField := SortField{
		Field:     "create_at",
		Ascending: false, // false为降序，true为升序
	}
	if option.Sort != "" {
		//进行截取，因为传过去的query内sort的值的输入基本都是create_at desc ←这种，字段+排序类型，所以要进行亲后截取
		_list := strings.Split(option.Sort, " ") //以空格作为分割线进行截取
		if len(_list) == 2 && (_list[1] == "desc" || _list[1] == "asc") {
			sortField.Field = _list[0]
			//根据对应排序设置好es的sort的bool值进行同样的排序
			if _list[1] == "desc" {
				sortField.Ascending = false //降序
			}
			if _list[1] == "asc" {
				sortField.Ascending = true //升序
			}
		}
	}

	res, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Highlight(elastic.NewHighlight().Field("title")).
		From(option.GetForm()).
		Sort(sortField.Field, sortField.Ascending).
		Size(option.Limit).
		Do(context.Background())
	if err != nil {
		return
	}

	count = int(res.Hits.TotalHits.Value) //搜索到结果总条数
	articleList := []models.ArticleModel{}
	diggInfo := service.Service.RedisService.GetDiggInfo()
	LookInfo := service.Service.RedisService.GetLookInfo()
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
		title, ok := hit.Highlight["title"]
		if ok {
			article.Title = title[0]
		}
		article.ID = hit.Id
		//设置一个文章redis内点赞数大于等于10自动同步es和redis数据的机制
		//for循环文章，当redis内该哈希表对应的id文章点赞数大于等于10，则触发同步数据机制，更新es表对应文章的点赞数的同时删除redis内对应id文章的点赞数据
		digg, ok := diggInfo[hit.Id]
		if ok && digg >= 10 {
			// 更新 ES 中的文章点赞数
			article.DiggCount = article.DiggCount + digg
			_, updateErr := global.ESClient.Update().
				Index(models.ArticleModel{}.Index()).
				Id(hit.Id).
				Doc(map[string]int{
					"digg_count": article.DiggCount,
				}).
				Do(context.Background())
			if updateErr != nil {
				logrus.Error(updateErr.Error())
				continue
			}
			logrus.Info(article.Title, "点赞数更新成功，新点赞数为:", article.DiggCount)
			// 从 Redis 中删除点赞数
			service.Service.RedisService.DiggClearByID(hit.Id)
		} else {
			article.DiggCount = article.DiggCount + digg
		}
		//设置一个文章redis内浏览量大于等于10自动同步es和redis数据的机制
		//for循环文章，当redis内该哈希表对应的id文章浏览量大于等于10，则触发同步数据机制，更新es表对应文章的浏览量的同时删除redis内对应id文章的浏览量数据
		look, ok := LookInfo[hit.Id]
		if ok && look >= 10 {
			// 更新 ES 中的文章点赞数
			article.LookCount = article.LookCount + look
			_, updateErr := global.ESClient.Update().
				Index(models.ArticleModel{}.Index()).
				Id(hit.Id).
				Doc(map[string]int{
					"look_count": article.LookCount,
				}).
				Do(context.Background())
			if updateErr != nil {
				logrus.Error(updateErr.Error())
				continue
			}
			logrus.Info(article.Title, "浏览数更新成功，新浏览数为:", article.LookCount)
			// 从 Redis 中删除点赞数
			service.Service.RedisService.LookClearByID(hit.Id)
		} else {
			article.LookCount = article.LookCount + look
		}
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
	//访问一次文章详细内容，文章浏览量+1
	article.LookCount = article.LookCount + service.Service.RedisService.GetLook(res.Id)
	return
}

func CommDetailByKeyword(key string) (model models.ArticleModel, err error) {
	res, err := global.ESClient.Search().
		Index(models.ArticleModel{}.Index()).
		Query(elastic.NewTermQuery("keyword", key)).
		Size(1).
		Do(context.Background())
	if err != nil {
		return
	}
	if res.Hits.TotalHits.Value == 0 {
		return model, errors.New("文章不存在")
	}
	hit := res.Hits.Hits[0]

	err = json.Unmarshal(hit.Source, &model)
	if err != nil {
		return
	}
	model.ID = hit.Id
	return
}
