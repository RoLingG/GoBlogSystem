package api

import (
	"GoRoLingG/api/advert_api"
	"GoRoLingG/api/images_api"
	"GoRoLingG/api/menu_api"
	"GoRoLingG/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdvertApi   advert_api.AdvertApi
	MenuApi     menu_api.MenuApi
}

var ApiGroupApp = new(ApiGroup)
