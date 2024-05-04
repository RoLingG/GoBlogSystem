package api

import (
	"GoRoLingG/api/images_api"
	"GoRoLingG/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   images_api.ImagesApi
}

var ApiGroupApp = new(ApiGroup)
