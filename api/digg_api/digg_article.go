package digg_api

import (
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/es_service"
	"GoRoLingG/service/redis_service"
	"github.com/gin-gonic/gin"
)

// DiggArticleView 文章点赞
// @Tags 文章管理
// @Summary 文章点赞
// @Description	点赞文章
// @Param id path string true	"需要点赞的文章ID"
// @Param data body models.ESIDRequest true	"点赞文章的一些参数"
// @Produce json
// @Router /api/diggArticle/{id} [post]
// @Success 200 {object} res.Response{}
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
	_, err = es_service.CommonDetail(cr.ID)
	res.OKWithMsg("文章点赞成功", c)
}
