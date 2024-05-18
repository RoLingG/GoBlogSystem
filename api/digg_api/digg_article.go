package digg_api

import (
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service"
	"github.com/gin-gonic/gin"
)

func (DiggApi) DiggArticleView(c *gin.Context) {
	var cr models.ESIDRequest
	err := c.ShouldBindJSON(&cr) //通过uri去进行获取文章id
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//对长度校验
	if len(cr.ID) != 20 {
		res.FailWithMsg("文章id非法，点赞失败", c)
		return
	}
	service.Service.RedisService.Digg(cr.ID)
	res.OKWithMsg("文章点赞成功", c)
}
