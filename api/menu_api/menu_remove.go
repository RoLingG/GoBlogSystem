package menu_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// MenuRemoveView 菜单项删除
// @Tags 菜单管理
// @Summary 菜单项删除
// @Description	菜单内删除存在的菜单项
// @Param data body models.RemoveRequest true	"删除菜单项的一些参数"
// @Produce json
// @Router /api/menusRemove [delete]
// @Success 200 {object} res.Response{}
func (MenuApi) MenuRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//批量删除
	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("所要删除的菜单项不存在", c)
		return
	}
	//批量删除菜单项事务(成功就一起成功，失败就一起失败)
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		//因为连表操作了，所以也要删除第三张表的数据
		err = global.DB.Debug().Model(&menuList).Association("MenuImage").Clear()
		if err != nil {
			global.Log.Error(err)
			return err
		}
		err = global.DB.Delete(&menuList).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("删除菜单项事务失败", c)
		return
	}

	res.OKWithMsg(fmt.Sprintf("共删除 %d 个菜单项", count), c)
}
