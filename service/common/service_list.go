package common

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"fmt"
	"gorm.io/gorm"
)

//type Option struct {
//	models.PageInfo
//	Debug bool
//}

// 封装统一用列表查询的方法，包括分页
// 这里的model即便函数内没有显式使用也不能删掉，因为它能让GORM实际的定义到所要操作的model，删了就不知道要对哪个model进行操作了
//func CommonList[T any](model T, option Option) (list []T, count int64, err error) {
//	DB := global.DB //默认无日志
//	if option.Debug {
//		//如果有debug就说明有日志，开启显示日志模式
//		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
//	}
//
//	if option.Sort == "" {
//		option.Sort = "create_at desc" //默认按照时间往前排,desc(降序),asc(升序)
//	}
//
//	//count = DB.Select("id").Find(&list).RowsAffected
//	query := DB.Where(model)
//	//加上select可不用select *，优化sql语句
//	count = query.Select("id").Find(&list).RowsAffected
//	//query因为会受上一次query的查询影响，所以这里要再赋值重置一遍
//	query = DB.Where(model)
//
//	//页码有效检测
//	totalPages := int(count) / option.Limit
//	if int(count)%option.Limit > 0 {
//		totalPages++
//	}
//
//	//偏移量设置
//	//这样page默认就是1了，而不是未设置的0，从1开始分页更直观
//	offset := (option.Page - 1) * option.Limit
//	if offset < 0 {
//		offset = 0
//	}
//
//	currentPage := option.Page
//	if currentPage < 1 || currentPage > totalPages {
//		// 如果页码无效，返回错误
//		return list, 0, fmt.Errorf("无效的页码: %d, 总页数: %d", currentPage, totalPages)
//	}
//
//	err = query.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error
//	//err = DB.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error
//	return list, count, err
//}

type Option struct {
	models.PageInfo          // 分页查询
	Likes           []string // 需要模糊匹配的字段列表
	Debug           bool     // 是否打印sql
	Where           *gorm.DB // 额外的查询
	Preload         []string // 预加载的字段列表
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

	// 模糊匹配
	if option.Key != "" {
		likeQuery := global.DB.Where("")
		for index, column := range option.Likes {
			// 第一个模糊匹配和前面的关系是and关系，后面的和前面的模糊匹配是or的关系
			if index == 0 {
				likeQuery.Where(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
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
