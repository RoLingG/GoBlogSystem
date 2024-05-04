package routers

import "GoRoLingG/api"

// 方法一
//func SettingsRouter(router *gin.Engine) {
//	settingsApi := api.ApiGroupApp.SettingsApi
//	router.GET("settings", settingsApi.SettingsInfoView)
//}

// 方法二
// 通过routers\enter里的RouterGroup进行两个文件之间的连接
//func (RouterGroup) SettingsRouter(router *gin.Engine) {
//	settingsApi := api.ApiGroupApp.SettingsApi
//	router.GET("settings", settingsApi.SettingsInfoView)
//}

// 方法三
// 通过routers\enter里的RouterGroup进行两个文件之间的连接
func (router RouterGroup) SettingsRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("settings/:name", settingsApi.SettingsInfoView)
	router.PUT("settings/:name", settingsApi.SettingsInfoUpdateView)
}
