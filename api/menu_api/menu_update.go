package menu_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

// MenuUpdateView 菜单项更新
// @Tags 菜单管理
// @Summary 菜单项更新
// @Description	菜单内更新存在的菜单项
// @Param id path int true "需要更新的菜单项ID"
// @Param data body MenuRequest true	"删除菜单项的一些参数"
// @Produce json
// @Router /api/menusUpdate/{id} [put]
// @Success 200 {object} res.Response{}
func (MenuApi) MenuUpdateView(c *gin.Context) {
	//cr为post传入过来的数据载体
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	id := c.Param("id")
	// 先把之前的image清空
	var menuModel models.MenuModel
	err = global.DB.Debug().Take(&menuModel, id).Error
	if err != nil {
		res.FailWithMsg("菜单项不存在，操作失败", c)
		return
	}
	global.DB.Debug().Model(&menuModel).Association("MenuImage").Clear() //一对多操作，因为外键想关联，所以清楚menu表也会清空第三张表姑关联着的数据
	//清空后如果选择了image，就进行添加
	if len(cr.ImageSortList) > 0 {
		//操作第三张表
		var imageList []models.MenuImageModel
		for _, sort := range cr.ImageSortList {
			imageList = append(imageList, models.MenuImageModel{
				MenuID:  menuModel.ID,
				ImageID: sort.ImageID,
				Sort:    sort.Sort,
			})
		}
		err = global.DB.Debug().Create(&imageList).Error
		if err != nil {
			res.FailWithMsg("更新菜单项图片失败", c)
			return
		}
	}
	//有需要添加图片，则普通更新
	//这里更新如果用广告的update那种方法，会出现sort零值更新问题
	//因为sort不好设置默认值，情况复杂，所以这里用了第三方依赖structs，可以将对应结构体转化成map，这样就可以将sort更新成0了，前提是要在type的参数设置好structs标签
	maps := structs.Map(&cr)
	err = global.DB.Debug().Model(&menuModel).Updates(maps).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("修改菜单失败", c)
		return
	}
	res.OKWithMsg("修改菜单成功", c)
}
