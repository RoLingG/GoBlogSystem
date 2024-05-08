package advert_api

import (
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/common"
	"github.com/gin-gonic/gin"
	"strings"
)

// 广告列表
func (AdvertApi) AdvertListView(c *gin.Context) {
	var cr models.PageModel
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
	}
	//判断Referer 是否包含admin，如果包含则全部返回；如果不包含，就返回is_show=true
	//isShow默认为true，当走不进if时，就说明不包含admin，也就只能看为IsShow为true的内容
	isShow := true
	referer := c.GetHeader("Referer")
	if strings.Contains(referer, "admin") {
		//Referer包含admin，则可以看IsShow两种情况的内容
		isShow = false
	}
	list, count, err := common.CommonList(models.AdverModel{IsShow: isShow}, common.Option{
		PageModel: cr,
		Debug:     true,
	})

	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}

	res.OKWithList(list, count, c)
	return
}
