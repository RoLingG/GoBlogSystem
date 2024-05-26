package article_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/es_service"
	"GoRoLingG/service/redis_service"
	"github.com/gin-gonic/gin"
)

// ArticleDetailView 文章详情列表
// @Tags 文章管理
// @Summary 文章详情列表
// @Description	查询文章详情的列表
// @Param id path string true "需要查询详情的文章ID"
// @Produce json
// @Router /api/articleDetail/{id} [get]
// @Success 200 {object} res.Response{data=models.ArticleModel}
func (ArticleApi) ArticleDetailView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindUri(&cr) //通过uri去进行获取文章id
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	redis_service.NewArticleLookIndex().Set(cr.ID)
	article, err := es_service.CommonDetail(cr.ID)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("文章详情查询出错", c)
		return
	}
	res.OKWithData(article, c)
}
