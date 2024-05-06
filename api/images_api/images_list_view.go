package images_api

import (
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/common"
	"github.com/gin-gonic/gin"
)

// 图片列表
func (ImagesApi) ImagesListView(c *gin.Context) {
	var cr models.PageModel
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
	}
	//将获取列表集成成一个function，封装起来便于广泛利用
	list, count, err := common.CommonList(models.ImageModel{}, common.Option{
		PageModel: cr,
		Debug:     true,
	})
	res.OKWithList(list, count, c)
	return
}
