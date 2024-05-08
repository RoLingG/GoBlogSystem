package common

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"fmt"
	"gorm.io/gorm"
)

type Option struct {
	models.PageModel
	Debug bool
}

// 封装统一用列表查询的方法，包括分页
// 这里的model即便函数内没有显式使用也不能删掉，因为它能让GORM实际的定义到所要操作的model，删了就不知道要对哪个model进行操作了
func CommonList[T any](model T, option Option) (list []T, count int64, err error) {
	DB := global.DB //默认无日志
	if option.Debug {
		//如果有debug就说明有日志，开启显示日志模式
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}

	if option.Sort == "" {
		option.Sort = "create_at desc" //默认按照时间往前排,desc(降序),asc(升序)
	}

	//count = DB.Select("id").Find(&list).RowsAffected
	query := DB.Where(model)
	//加上select可不用select *，优化sql语句
	count = query.Select("id").Find(&list).RowsAffected
	//query因为会受上一次query的查询影响，所以这里要再赋值重置一遍
	query = DB.Where(model)

	//页码有效检测
	totalPages := int(count) / option.Limit
	if int(count)%option.Limit > 0 {
		totalPages++
	}

	//偏移量设置
	//这样page默认就是1了，而不是未设置的0，从1开始分页更直观
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}

	currentPage := option.Page
	if currentPage < 1 || currentPage > totalPages {
		// 如果页码无效，返回错误
		return list, 0, fmt.Errorf("无效的页码: %d, 总页数: %d", currentPage, totalPages)
	}

	err = query.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error
	//err = DB.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error
	return list, count, err
}
