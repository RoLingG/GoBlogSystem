package article_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type TagsType struct {
	DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int `json:"sum_other_doc_count"`
	Buckets                 []struct {
		Key      string `json:"key"`
		DocCount int    `json:"doc_count"`
		Articles struct {
			DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
			SumOtherDocCount        int `json:"sum_other_doc_count"`
			Buckets                 []struct {
				Key      string `json:"key"`
				DocCount int    `json:"doc_count"`
			} `json:"buckets"`
		} `json:"articles"`
	} `json:"buckets"`
}

type TagsResponse struct {
	Tag           string   `json:"tag"`
	Count         int      `json:"count"`
	ArticleIDList []string `json:"article_id_list"`
	CreateAt      string   `json:"create_at"`
}

// ArticleTagsListView 文章标签列表
// @Tags 文章管理
// @Summary 文章标签列表
// @Description	查询文章标签列表
// @Param data body models.PageInfo true	"查询文章标签的一些参数"
// @Produce json
// @Router /api/articleTagsList [get]
// @Success 200 {object} res.Response{data=res.ListResponse[TagsResponse]}
func (ArticleApi) ArticleTagsListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//分页的参数设置
	if cr.Limit == 0 {
		cr.Limit = 10
	}
	offset := (cr.Page - 1) * cr.Limit
	if offset < 0 {
		offset = 0
	}

	//获取tag总数的查询
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		//NewCardinalityAggregation()   会对指标进行去重
		//NewValueCountAggregation()    不会进行去重操作
		Aggregation("tags", elastic.NewCardinalityAggregation().Field("tags")).
		Size(0).
		Do(context.Background())
	totalTags, _ := result.Aggregations.Cardinality("tags") //totalTags内是这样的→ &{map[value:[50]] 0xc00026b340 map[]}
	//直接打印totalTags发现其是指向值的地址
	tagCount := int64(*totalTags.Value) //tag的总数

	agg := elastic.NewTermsAggregation().Field("tags")
	query := elastic.NewBoolQuery()
	agg.SubAggregation("articles", elastic.NewTermsAggregation().Field("keyword"))             //通过索引结构中的keyword去查找tag底下的文章有哪些
	agg.SubAggregation("page", elastic.NewBucketSortAggregation().From(offset).Size(cr.Limit)) //分别对应mysql中的offset和find，用于实现分页效果

	//查询通用写法
	result, err = global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("tags", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("对应tag的文章查询失败", c)
		return
	}
	//底下这些bucket的数据基本都可以在fmt.Println(string(result.Aggregations["tags"]))里看到什么字段对应什么参数
	//然后用对应的结构体接过来以后实例化再for循环append进去就可以得到想要的实例类型组了
	var tagType TagsType
	var tagList = make([]*TagsResponse, 0)
	_ = json.Unmarshal(result.Aggregations["tags"], &tagType)
	var tagStrList []string
	for _, bucket := range tagType.Buckets {
		var articleList []string
		for _, s := range bucket.Articles.Buckets {
			articleList = append(articleList, s.Key)
		}
		tagList = append(tagList, &TagsResponse{
			Tag:           bucket.Key,
			Count:         bucket.DocCount,
			ArticleIDList: articleList,
		})
		tagStrList = append(tagStrList, bucket.Key) //获取tag的title
	}
	var tagModelList []models.TagModel
	global.DB.Debug().Find(&tagModelList, "title in ?", tagStrList) //在mysql中将对应title的tag数据获取出来
	var tagDate = map[string]string{}                               //获取tag创建的时间
	//从tagModelList获取tag的创建时间，tagModelList相当于将mysql中tag表内的数据都获取了出来
	for _, tagModel := range tagModelList {
		tagDate[tagModel.Title] = tagModel.CreateAt.Format("2006-01-02 15:04:05")
	}
	//将传给前端的tagList实例赋值好后丢给前端
	for _, response := range tagList {
		response.CreateAt = tagDate[response.Tag]
	}
	//这里难的主要是es与mysql同步数据，es内有tag，从es内获取tag名字，到mysql中找对应的tag，找到对应的tag就去拿它的创建时间，然后返回给前端
	//做了这么多操作其实就是为了从mysql中获取tag创建的时间
	res.OKWithList(tagList, tagCount, c)
}
