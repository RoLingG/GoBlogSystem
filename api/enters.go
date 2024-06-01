package api

import (
	"GoRoLingG/api/advert_api"
	"GoRoLingG/api/article_api"
	"GoRoLingG/api/chat_api"
	"GoRoLingG/api/comment_api"
	"GoRoLingG/api/data_api"
	"GoRoLingG/api/digg_api"
	"GoRoLingG/api/images_api"
	"GoRoLingG/api/log_api"
	"GoRoLingG/api/log_v2_api"
	"GoRoLingG/api/menu_api"
	"GoRoLingG/api/message_api"
	"GoRoLingG/api/news_api"
	"GoRoLingG/api/role_api"
	"GoRoLingG/api/settings_api"
	"GoRoLingG/api/tag_api"
	"GoRoLingG/api/user_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdvertApi   advert_api.AdvertApi
	MenuApi     menu_api.MenuApi
	UserApi     user_api.UserApi
	TagApi      tag_api.TagApi
	MessageApi  message_api.MessageApi
	ArticleApi  article_api.ArticleApi
	DiggApi     digg_api.DiggApi
	CommentApi  comment_api.CommentApi
	NewsApi     news_api.NewsApi
	ChatApi     chat_api.ChatApi
	LogApi      log_api.LogApi
	LogStash    log_v2_api.LogV2Api
	DataApi     data_api.DataApi
	RoleApi     role_api.RoleApi
}

var ApiGroupApp = new(ApiGroup)
