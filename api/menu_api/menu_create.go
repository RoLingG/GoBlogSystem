package menu_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/models/ctype"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

// 前端传过来的数据
type MenuRequest struct {
	MenuTitle     string      `json:"menu_title" binding:"required" msg:"请完善菜单项名称"`
	MenuPath      string      `json:"menu_path" binding:"required" msg:"请完善菜单项路径"`
	Slogan        string      `json:"slogan"`
	Abstract      ctype.Array `json:"abstract"`
	AbstractTime  int         `json:"abstract_time"`                         //简介的切换时间，单位为秒
	MenuTime      int         `json:"menu_time"`                             //图片的切换时间，单位为秒
	Sort          int         `json:"sort" binding:"required" msg:"请输入菜单序号"` //菜单的序号
	ImageSortList []ImageSort `json:"image_sort_list"`                       //具体图片的顺序，要单独给ImageSortList创建一个类型是因为如果用[]imageModel要传的参数太多了，实际上我们只需要对应的ID和序号就行
}

func (MenuApi) MenuCreateView(c *gin.Context) {
	//cr为post传入过来的数据载体
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	//重复值判断
	var titleCheck models.MenuModel
	err = global.DB.Take(&titleCheck, "menu_title = ?", cr.MenuTitle).Error
	//无err就说明在数据库中找到了
	if err == nil {
		res.FailWithMsg("已存在相同菜单标题，请重新创建", c)
		return
	}
	var pathCheck models.MenuModel
	err = global.DB.Take(&pathCheck, "menu_path = ?", cr.MenuPath).Error
	//无err就说明在数据库中找到了
	if err == nil {
		res.FailWithMsg("已存在相同菜单路径，请重新创建", c)
		return
	}
	var sortCheck models.MenuModel
	err = global.DB.Take(&sortCheck, "sort = ?", cr.Sort).Error
	//无err就说明在数据库中找到了
	if err == nil {
		res.FailWithMsg("已存在相同菜单序号，请重新创建", c)
		return
	}

	//多表操作，操作第三张表menu_image_model
	//菜单数据入库，create入库之后，因为MenuModel继承enter里面的Model，所以已入库相应的数据就会自动生成对应的ID
	menuModel := &models.MenuModel{
		MenuTitle:    cr.MenuTitle,
		MenuPath:     cr.MenuPath,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		MenuTime:     cr.MenuTime,
		Sort:         cr.Sort,
	}
	err = global.DB.Debug().Create(&menuModel).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("菜单添加失败", c)
		return
	}

	//前面说到自动生成ID，因为生成了，所以存的时候就能有对应的MenuID
	var menuImageList []models.MenuImageModel
	//post过来没有图片数据时，判定为无图菜单添加，添加成功
	if len(cr.ImageSortList) == 0 {
		res.OKWithMsg("菜单添加成功", c)
		return
	}
	//创建一个图片表实例，去用于判断post过来的图片id在表内是否存在数据
	for _, sort := range cr.ImageSortList {
		//判断ImageID是否存在，imageModel不能放在外面，不然每次for循环都不会刷新操作，每次都在一个查询里不断添加语句
		var imageModel models.ImageModel
		err = global.DB.Debug().Where("id = ?", sort.ImageID).Take(&imageModel).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMsg("图片不存在，菜单添加失败", c)
			return
		}
		menuImageList = append(menuImageList, models.MenuImageModel{
			MenuID:  menuModel.ID,
			ImageID: sort.ImageID,
			Sort:    sort.Sort,
		})
	}
	//第三张表入库
	err = global.DB.Debug().Create(&menuImageList).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("菜单数据入库失败", c)
		return
	}
	res.OKWithMsg("菜单添加成功", c)
}
