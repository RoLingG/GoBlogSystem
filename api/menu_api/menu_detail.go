package menu_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
)

// MenuDetailView 菜单项详情列表
// @Tags 菜单管理
// @Summary 菜单详情列表
// @Description	查看菜单项详情的列表
// @Param id path int true "需要查询的菜单项ID"
// @Produce json
// @Router /api/menuDetailList/{id} [get]
// @Success 200 {object} res.Response{}
func (MenuApi) MenuDetailView(c *gin.Context) {
	id := c.Param("id")
	var menuModel models.MenuModel
	err := global.DB.Debug().Take(&menuModel, id).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("菜单ID获取失败，请检查是否有菜单内菜单项", c)
		return
	}
	//查连接表
	var menuImages []models.MenuImageModel
	err = global.DB.Debug().Preload("ImageModel").Order("sort desc").Find(&menuImages).Select("menu_id = ?", id).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("菜单项图片数据获取失败，请检查是否菜单项是否有图片", c)
		return
	}
	var images = make([]Image, 0)
	for _, image := range menuImages {
		if menuModel.ID != image.MenuID {
			continue
		}
		images = append(images, Image{
			ID:   image.ImageID,
			Path: image.ImageModel.Path,
		})
	}
	menuResponse := MenuResponse{
		MenuModel: menuModel,
		MenuImage: images,
	}
	res.OKWithData(menuResponse, c)
	return
}
