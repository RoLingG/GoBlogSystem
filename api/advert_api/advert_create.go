package advert_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/utils"
	"github.com/gin-gonic/gin"
)

type AdvertRequest struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题"`
	Href   string `json:"href" binding:"required,url" msg:"请输入广告跳转链接"`
	Images string `json:"images" binding:"required,url" msg:"请输入广告图片链接"`
	IsShow bool   `gorm:"default:false" json:"is_show" msg:"选择是否显示,默认为false"`
}

func (AdvertApi) AdvertCreateView(c *gin.Context) {
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	//广告重复判断
	var advert models.AdverModel
	err = global.DB.Take(&advert, "title = ?", cr.Title).Error
	//无err就说明在数据库中找到了
	if err == nil {
		res.FailWithMsg("该广告重复存在，请重传", c)
		return
	}

	//判断传过来的Href和Image的url是否合法
	isValid := utils.ValidateURL(cr.Href)
	if !isValid {
		res.FailWithMsg("链接非法，请输入合法的跳转链接", c)
		return
	}
	isValid = utils.ValidateURL(cr.Images)
	if !isValid {
		res.FailWithMsg("图片链接非法，请输入合法的图片链接", c)
		return
	}

	//入库
	err = global.DB.Create(&models.AdverModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: cr.IsShow,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("添加广告失败", c)
		return
	}
	res.OKWithMsg("添加广告成功", c)
}
