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

// MenuNameListView 这个接口用来传给前端主页之类要重要数据菜单列表数据的
func (MenuApi) MenuNameListView(c *gin.Context) {
	var menuNameList []MenuNameResponse
	global.DB.Debug().Model(models.MenuModel{}).Select("id", "menu_title", "menu_path").Scan(&menuNameList)
	res.OKWithData(menuNameList, c)
}
