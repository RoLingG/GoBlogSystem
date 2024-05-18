package es_serivce

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
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
