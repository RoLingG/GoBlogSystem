package article_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/redis_service"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

// ArticleContentByIDView 获取文章正文
// @Tags 文章管理
// @Summary 获取文章正文
// @Description 获取文章正文
// @Param id path string  true  "id"
// @Router /api/articleContent/{id} [get]
// @Produce json
// @Success 200 {object} res.Response{}
func (ArticleApi) ArticleContentByIDView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	redis_service.NewArticleLookIndex().Set(cr.ID)

	//将在es中对应ID的文章的所有数据获取出来
	result, err := global.ESClient.Get().
		Index(models.ArticleModel{}.Index()).
		Id(cr.ID).
		Do(context.Background())
	if err != nil {
		res.FailWithMsg("查询失败", c)
		return
	}
	var article models.ArticleModel
	//将对应ID文章的数据反序列化，用类型为models.ArticleModel的article实例去接收
	err = json.Unmarshal(result.Source, &article)
	if err != nil {
		return
	}
	//只需要返回对应ID文章的正文
	res.OKWithData(article.Content, c)
}
