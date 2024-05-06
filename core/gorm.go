package core

import (
	"GoRoLingG/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		//global.Log.Warn("未配置mysql,取消gorm的连接")
		global.Log.Warn("未配置mysql,取消gorm的连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()

	//sql日志
	var mysqlLogger logger.Interface

	if global.Config.System.Env == "dev" {
		//开发环境显示所有的sql
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}

	//sql日志配置，可以不用debug()去看了
	global.MysqlLog = logger.Default.LogMode(logger.Info)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		//global.Log.Error(fmt.Sprintf("[%s] mysql连接失败", dsn))
		global.Log.Fatalf(fmt.Sprintf("[%s] mysql连接失败", dsn))
		//panic(err)
	} else {
		global.Log.Info(fmt.Sprintf("[%s] mysql连接成功", dsn))
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               //最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              //最多可容纳数
	sqlDB.SetConnMaxLifetime(time.Hour * 4) //连接最大复用时间，不能超过mysql的wait_timeout
	return db
}
