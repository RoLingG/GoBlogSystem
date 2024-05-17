package api

import (
	"GoRoLingG/api/advert_api"
	"GoRoLingG/api/article_api"
	"GoRoLingG/api/images_api"
	"GoRoLingG/api/menu_api"
	"GoRoLingG/api/message_api"
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
}

var ApiGroupApp = new(ApiGroup)
