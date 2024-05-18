package flag

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
)

func Makemigrations() {
	var err error
	//global.DB.SetupJoinTable(&models.UserModel{}, "CollectModels", &models.UserCollectModel{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "MenuImage", &models.MenuImageModel{})
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.ImageModel{},
		&models.TagModel{},
		&models.MessageModel{},
		&models.AdvertModel{},
		&models.UserModel{},
		&models.CommentModel{},
		//&models.ArticleModel{},
		&models.UserCollectModel{},
		&models.MenuModel{},
		&models.MenuImageModel{},
		&models.FeedBackModel{},
		&models.LoginDataModel{},
	)
	if err != nil {
		global.Log.Error("[ Error ] 生成数据库表结构失败")
		return
	}
	global.Log.Info("[ Success ] 生成数据库表结构成功")
}
