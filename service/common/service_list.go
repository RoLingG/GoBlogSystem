package common

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"gorm.io/gorm"
)

type Option struct {
	models.PageModel
	Debug bool
}

func CommonList[T any](model T, option Option) (list T, count int64, err error) {
	DB := global.DB //默认无日志
	if option.Debug {
		//如果有debug就说明有日志，开启显示日志模式
		DB = global.DB.Session(&gorm.Session{
			Logger: global.MysqlLog,
		})
	}

	//加上select可不用select *，优化sql语句
	count = DB.Select("id").Find(&list).RowsAffected
	//偏移量设置
	//这样page默认就是1了，而不是未设置的0，从1开始分页更直观
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	err = DB.Limit(option.Limit).Offset(offset).Find(&list).Error
	return list, count, err
}
