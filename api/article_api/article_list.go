package article_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service"
	"GoRoLingG/service/es_service"
	"GoRoLingG/utils/jwt"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
	"github.com/olivere/elastic/v7"
)

type ArticleSearchRequest struct {
	models.PageInfo
	Tag    string `json:"tag" form:"tag"`
	IsUser bool   `json:"is_user" form:"is_user"` //是否是当前用户发布的文章
}

// ArticleListView 文章列表
// @Tags 文章管理
// @Summary 文章列表
// @Description	查询所有文章的列表
// @Param token header string true "Token"
// @Param data query ArticleSearchRequest true	"查询文章的一些参数"
// @Produce json
// @Router /api/articleList [get]
// @Success 200 {object} res.Response{data=res.ListResponse[models.AdvertModel]}
func (ArticleApi) ArticleListView(c *gin.Context) {
	var cr ArticleSearchRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	boolSearch := elastic.NewBoolQuery() //创建一个bool查询，
	if cr.IsUser {
		token := c.GetHeader("token")
		claims, err := jwt.ParseToken(token)
		if err == nil && service.Service.RedisService.CheckLogout(token) {
			//bool查询的Must()同等于用And逻辑联系，bool查询用于组成复杂的查询逻辑关系
			boolSearch.Must(elastic.NewTermQuery("user_id", claims.UserID))
		}
	}

	list, count, err := es_service.CommonList(es_service.Option{
		PageInfo: cr.PageInfo,
		Fields:   []string{"title", "content", "category"},
		Tag:      cr.Tag,
		Query:    boolSearch, //查询条件
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("分页查询失败", c)
		return
	}
	//一般来说列表不需要看到文章内容，所以给content加上omit(list)标签，在list场景下过滤掉content
	//别忘了加 github.com/liu-cn/json-filter/filter ←这个第三方库
	data := filter.Omit("list", list)
	_list, _ := data.(filter.Filter)
	//判断当list为空时，该怎么让它传过去的样子从空json{}转换成空集合[]，解决json-filter空值问题
	if string(_list.MustMarshalJSON()) == "{}" {
		list = make([]models.ArticleModel, 0) //去除零值，返回正常空集合[]
		res.OKWithList(list, int64(count), c)
		return
	}
	res.OKWithList(_list, int64(count), c) //content上有omit(list)标签，当用filter过滤掉之后，返回给前端的json就不会有content了，这样写比较灵活
}
