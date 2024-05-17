package article_api

import (
	"GoRoLingG/global"
	"GoRoLingG/res"
	"GoRoLingG/service/es_serivce"
	"github.com/gin-gonic/gin"
)

type ESIDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}

func (ArticleApi) ArticleDetailView(c *gin.Context) {
	var cr ESIDRequest
	err := c.ShouldBindUri(&cr) //通过uri去进行获取文章id
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	article, err := es_serivce.CommonDetail(cr.ID)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("文章详情查询出错", c)
		return
	}
	res.OKWithData(article, c)
}
