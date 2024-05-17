package article_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/es_serivce"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

func (ArticleApi) ArticleListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	list, count, err := es_serivce.CommonList(cr.Key, cr.Limit, cr.Page)
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("分页查询失败", c)
		return
	}
	//一般来说列表不需要看到文章内容，所以给content加上omit(list)标签，在list场景下过滤掉content
	//别忘了加 github.com/liu-cn/json-filter/filter ←这个第三方库
	filter.Omit("list", list)
	res.OKWithList(filter.Omit("list", list), int64(count), c) //content上有omit(list)标签，当用filter过滤掉之后，返回给前端的json就不会有content了，这样写比较灵活
}
