package advert_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

// AdvertUpdateView 广告更新
// @Tags 广告管理
// @Summary 广告更新
// @Description	广告更新
// @Param id path int true "需要更新的广告ID"
// @Param _ body AdvertRequest true	"广告更新的一些参数"
// @Router /api/advertUpdate/{id} [put]
// @Produce json
// @Success 200 {object} res.Response{data=string}
func (AdvertApi) AdvertUpdateView(c *gin.Context) {
	id := c.Param("id") //因为AdvertRequest结构体里面没有ID，但是修改要id，所以我们从前端拿id来进行修改
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	//创建对应表的ORM实例对象
	var advert models.AdvertModel
	//需要修改的对应ID广告是否存在判断
	err = global.DB.Debug().Take(&advert, id).Error
	//有err就说明在数据库中不存在该ID对应的广告
	if err != nil {
		res.FailWithMsg("该广告不存在，请重传", c)
		return
	}

	//重新赋值一遍，清空上一次操作，可以用session替代这种怪怪的做法，但我不会（
	advert = models.AdvertModel{}
	//标题是否重复判断
	err = global.DB.Debug().Take(&advert, "title = ?", cr.Title).Error
	//无err就说明在数据库中找到了
	if err == nil {
		res.FailWithMsg("该广告标题已存在，请修改标题重传", c)
		return
	}

	//判断传过来的Href和Image的url是否合法
	var isValid bool //创建一个bool来接收ValidateURL传过来的参数
	if cr.Href != "" {
		fmt.Println(cr.Href)
		isValid = utils.ValidateURL(cr.Href)
		if !isValid {
			res.FailWithMsg("链接非法，请输入合法的跳转链接", c)
			return
		}
	}
	//if cr.Images != "" {
	//	fmt.Println(cr.Images)
	//	isValid = utils.ValidateURL(cr.Images)
	//	if !isValid {
	//		res.FailWithMsg("图片链接非法，请输入合法的图片链接", c)
	//		return
	//	}
	//}
	fmt.Println("IsShow:", *cr.IsShow)
	//入库，这里updates的结构体实例如果属性过多的话，可以找个结构体转map的函数去进行快速转换，这样就方便很多，要用的话可以终端输入go get github.com/fatih/structs
	//注：这里用select进行指定字段更新的原因是因为bool类型在gorm框架里默认是false为零值，如果这样会直接updates更新多字段自动识别会略过bool的零值，导致字段更新问题出现。
	err = global.DB.Debug().Where(id).Select("Title", "Href", "Images", "IsShow").Updates(&models.AdvertModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: *cr.IsShow,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("修改广告失败", c)
		return
	}
	res.OKWithMsg("修改广告成功", c)
}
