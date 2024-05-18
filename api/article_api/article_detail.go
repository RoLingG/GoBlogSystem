package article_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service"
	"GoRoLingG/service/es_serivce"
	"github.com/gin-gonic/gin"
)

func (ArticleApi) ArticleDetailView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindUri(&cr) //通过uri去进行获取文章id
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	service.Service.RedisService.Look(cr.ID)
	article, err := es_serivce.CommonDetail(cr.ID)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("文章详情查询出错", c)
		return
	}
	res.OKWithData(article, c)
}
