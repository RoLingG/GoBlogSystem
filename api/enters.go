package api

import (
	"GoRoLingG/api/advert_api"
	"GoRoLingG/api/images_api"
	"GoRoLingG/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdertApi    advert_api.AdvertApi
}

var ApiGroupApp = new(ApiGroup)
