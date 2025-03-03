package es_service

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/service/redis_service"
	"context"
	"encoding/json"
	"errors"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"strings"
)

type Option struct {
	models.PageInfo
	Fields          []string
	Tag             string
	Date            string
	ArticleCategory string
	Query           *elastic.BoolQuery
}

func (o *Option) GetForm() int {
	page := o.Page
	limit := o.Limit
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	return (page - 1) * limit
}

// CommonList 列表查询文章分页
func CommonList(option Option) (list []models.ArticleModel, count int, err error) {
	boolSearch := option.Query
	if boolSearch == nil {
		boolSearch = elastic.NewBoolQuery()
	}
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
	if option.Date != "" {
		oldDate := option.Date + " " + "00:00:00"
		newDate := option.Date + " " + "23:59:59"
		boolSearch.Must(
			elastic.NewRangeQuery("create_at").Gte(oldDate).Lte(newDate),
		)
	}

	if option.ArticleCategory != "" {
		boolSearch.Must(
			elastic.NewMultiMatchQuery(option.ArticleCategory, "category"),
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
	diggInfo := redis_service.NewArticleDiggIndex().GetInfo()
	lookInfo := redis_service.NewArticleLookIndex().GetInfo()
	commentInfo := redis_service.NewArticleCommentIndex().GetInfo()
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
		digg, ok := diggInfo[hit.Id]
		article.DiggCount = article.DiggCount + digg
		look, ok := lookInfo[hit.Id]
		article.LookCount = article.LookCount + look
		comment, ok := commentInfo[hit.Id]
		article.CommentCount = article.CommentCount + comment

		//将当前文章加入文章列表中，便于显示给前端
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
	//使用文章详情接口，访问一次文章详细内容，则文章浏览量+1
	article.LookCount = article.LookCount + redis_service.NewArticleLookIndex().Get(res.Id)
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

func ArticleUpdate(id string, data map[string]any) error {
	_, err := global.ESClient.
		Update().
		Index(models.ArticleModel{}.Index()).
		Id(id).
		Doc(data).
		Do(context.Background())
	return err
}
