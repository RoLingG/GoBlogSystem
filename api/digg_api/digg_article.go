package digg_api

import (
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/es_service"
	"GoRoLingG/service/redis_service"
	"github.com/gin-gonic/gin"
)

func (DiggApi) DiggArticleView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindUri(&cr) //通过uri去进行获取文章id
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//对长度校验
	if len(cr.ID) != 20 {
		res.FailWithMsg("文章id非法，点赞失败", c)
		return
	}
	redis_service.NewArticleDiggIndex().Set(cr.ID)
	article, err := es_service.CommonDetail(cr.ID)
	err = es_service.ArticleUpdate(cr.ID, map[string]any{
		"digg_count": article.CollectCount + 1,
	})
	res.OKWithMsg("文章点赞成功", c)
}
