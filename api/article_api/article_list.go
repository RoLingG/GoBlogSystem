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
	"strconv"
)

type ArticleSearchRequest struct {
	models.PageInfo
	Tag             string `json:"tag" form:"tag"`
	IsUser          bool   `json:"is_user" form:"is_user"` //是否是当前用户发布的文章
	Date            string `json:"date" form:"date"`
	ArticleCategory string `json:"article_category" form:"article_category"`
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
	if err := c.ShouldBindQuery(&cr); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	boolSearch := elastic.NewBoolQuery()

	if cr.IsUser {
		token := c.GetHeader("token")
		claims, err := jwt.ParseToken(token)
		if err == nil && !service.Service.RedisService.CheckLogout(token) {
			userID, _ := strconv.Atoi(strconv.Itoa(int(claims.UserID)))
			boolSearch.Must(elastic.NewTermQuery("user_id", userID))
		}
	}

	list, count, err := es_service.CommonList(es_service.Option{
		PageInfo:        cr.PageInfo,
		Fields:          []string{"title", "content", "category"},
		Tag:             cr.Tag,
		Date:            cr.Date,
		ArticleCategory: cr.ArticleCategory,
		Query:           boolSearch,
	})

	if err != nil {
		global.Log.Error(err)
		res.OKWithMsg("查询失败", c)
		return
	}

	// json-filter空值问题
	data := filter.Omit("list", list)
	_list, _ := data.(filter.Filter)
	if string(_list.MustMarshalJSON()) == "{}" {
		list = make([]models.ArticleModel, 0)
		res.OKWithList(list, int64(count), c)
		return
	}
	res.OKWithList(data, int64(count), c)
}
