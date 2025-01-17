package global

import (
	"GoRoLingG/config"
	"github.com/cc14514/go-geoip2"
	"github.com/go-redis/redis"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 部署全局变量，用于保存配置文件
var (
	Config   *config.Config
	DB       *gorm.DB
	Log      *logrus.Logger
	MysqlLog logger.Interface
	Redis    *redis.Client
	ESClient *elastic.Client
	AddrDB   *geoip2.DBReader
)
