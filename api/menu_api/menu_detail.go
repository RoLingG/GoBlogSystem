package menu_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

type menuDetailRequest struct {
	Path string `form:"path"`
}

// MenuDetailView 菜单项详情列表
// @Tags 菜单管理
// @Summary 菜单详情列表
// @Description	查看菜单项详情的列表
// @Param id path int true "需要查询的菜单项ID"
// @Produce json
// @Router /api/menuDetailList/{id} [get]
// @Success 200 {object} res.Response{}
func (MenuApi) MenuDetailView(c *gin.Context) {
	var cr menuDetailRequest
	c.ShouldBindQuery(&cr)
	var menuModel models.MenuModel
	fmt.Println(cr.Path)
	err := global.DB.Debug().Where("menu_path = ?", cr.Path).Take(&menuModel).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("菜单获取失败，请检查是否有菜单内菜单项", c)
		return
	}
	//查连接表
	var menuImages []models.MenuImageModel
	err = global.DB.Debug().Preload("ImageModel").Order("sort desc").Find(&menuImages).Select("menu_path = ?", cr.Path).Error
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
