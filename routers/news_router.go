package routers

import "GoRoLingG/api"

func (router RouterGroup) NewsRouter() {
	newsApi := api.ApiGroupApp.NewsApi
	router.POST("/newsList", newsApi.NewsListView)
}
