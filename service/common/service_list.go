package common

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"fmt"
	"gorm.io/gorm"
)

type Option struct {
	models.PageInfo          // 分页查询
	Likes           []string // 需要模糊匹配的字段列表
	Debug           bool     // 是否打印sql
	Where           *gorm.DB // 额外的查询
	Preload         []string // 预加载的字段列表
	RoleBool        bool     // 是否开启 role 查询条件
}

func CommonList[T any](model T, option Option) (list []T, count int64, err error) {

	// 查model中非空字段
	query := global.DB.Where(model)
	if option.Debug {
		query = query.Debug()
	}

	// 默认按照时间往后排
	if option.Sort == "" {
		option.Sort = "create_at desc"
	}

	// 默认一页显示10条
	if option.Limit == 0 {
		option.Limit = 10
	}

	// 如果有高级查询就加上
	if option.Where != nil {
		query.Where(option.Where)
	}

	//默认所有角色全部选上
	if option.RoleBool {
		if len(option.Role) == 0 {
			option.Role = []int{1, 2, 3}
		}
		query.Where("role IN (?)", option.Role)
	}

	// 模糊匹配
	if option.Key != "" {
		likeQuery := global.DB.Where("")
		for index, column := range option.Likes {
			// 第一个模糊匹配和前面的关系是and关系，后面的和前面的模糊匹配是or的关系
			if index == 0 {
				likeQuery.Where(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key)) //%%实际为输出一个%，所以这实际上是% key %
			} else {
				likeQuery.Or(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
			}
		}
		// 整个模糊匹配它是一个整体
		query = query.Where(likeQuery) //query := global.DB.Where(model)
	}

	// 查列表，获取总数
	count = query.Find(&list).RowsAffected

	// 预加载
	for _, preload := range option.Preload {
		query = query.Preload(preload)
	}

	// 计算偏移
	offset := (option.Page - 1) * option.Limit
	err = query.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error

	return
}
