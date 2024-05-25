package chat_api

import (
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/common"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

func (ChatApi) ChatRecordView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	cr.Sort = "create_at desc"
	list, count, _ := common.CommonList(models.ChatModel{IsGroup: true}, common.Option{
		PageInfo: cr,
	})

	data := filter.Omit("list", list)
	_list, _ := data.(filter.Filter)
	//判断当list为空时，该怎么让它传过去的样子从空json{}转换成空集合[]，解决json-filter空值问题
	if string(_list.MustMarshalJSON()) == "{}" {
		list = make([]models.ChatModel, 0) //去除零值，返回正常空集合[]
		res.OKWithList(list, int64(count), c)
		return
	}
	res.OKWithList(data, count, c)
}
