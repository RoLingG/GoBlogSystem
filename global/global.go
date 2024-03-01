package global

import (
	"GoRoLingG/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// 部署全局变量，用于保存配置文件
var (
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
)
