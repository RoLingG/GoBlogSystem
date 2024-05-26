package menu_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
)

type MenuNameResponse struct {
	ID        uint   `json:"id"`
	MenuTitle string `json:"menu_title"`
	MenuPath  string `json:"menu_path"`
}

// MenuNameListView 菜单项名列表
// @Tags 菜单管理
// @Summary 菜单项名列表
// @Description	查看菜单项名字的列表
// @Produce json
// @Router /api/menusNameList [get]
// @Success 200 {object} res.Response{data=MenuNameResponse}
func (MenuApi) MenuNameListView(c *gin.Context) {
	var menuNameList []MenuNameResponse
	global.DB.Debug().Model(models.MenuModel{}).Select("id", "menu_title", "menu_path").Scan(&menuNameList)
	res.OKWithData(menuNameList, c)
}
