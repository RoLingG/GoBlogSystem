package core

import (
	"GoRoLingG/global"
	geoip2db "github.com/cc14514/go-geoip2-db"
	"github.com/sirupsen/logrus"
)

// InitAddrDB 实例化IP寻找国省市地址数据库
func InitAddrDB() {
	db, err := geoip2db.NewGeoipDbByStatik()
	if err != nil {
		logrus.Errorf("ip地址数据库加载失败")
	}
	global.AddrDB = db
}
