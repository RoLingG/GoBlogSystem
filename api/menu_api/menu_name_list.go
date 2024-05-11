package menu_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
)

type MenuNameResponse struct {
	ID    uint   `json:"id"`
	title string `json:"title"`
	path  string `json:"path"`
}

func (MenuApi) MenuNameListView(c *gin.Context) {
	var menuNameList []MenuNameResponse
	global.DB.Debug().Model(models.MenuModel{}).Select("id", "menu_title", "menu_path").Scan(&menuNameList)
	res.OKWithData(menuNameList, c)
}
