package flag

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/plugins/log_stash_v1"
	"GoRoLingG/plugins/log_stash_v2"
)

func MakeMigrations() {
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
		&models.ChatModel{},
		&models.UserScopeModel{},
		&models.AutoReplyModel{},
		&models.LargeScaleModelRoleModel{},
		&models.LargeScaleModelTagModel{},
		&models.LargeScaleModelChatModel{},
		&models.LargeScaleModelSessionModel{},
		&log_stash_v1.LogModel{},
		&log_stash_v2.LogStashModel{},
	)
	if err != nil {
		global.Log.Error("[ Error ] 生成数据库表结构失败")
		return
	}
	global.Log.Info("[ Success ] 生成数据库表结构成功")
}
