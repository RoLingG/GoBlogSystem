package data_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type DataCollectResponse struct {
	UserCount       int64 `json:"user_count"`
	ArticleCount    int64 `json:"article_count"`
	MessageCount    int64 `json:"message_count"`
	ChatGroupCount  int64 `json:"chat_group_count"`
	TodayLoginCount int64 `json:"today_login_count"`
	TodaySignCount  int64 `json:"today_sign_count"`
}

func (DataApi) DataCollectView(c *gin.Context) {
	var cr DataCollectResponse
	global.DB.Model(&models.UserModel{}).Select("count(id)").Scan(&cr.UserCount)
	global.DB.Model(&models.MessageModel{}).Select("count(id)").Scan(&cr.MessageCount)
	global.DB.Model(&models.ChatModel{IsGroup: true}).Select("count(id)").Scan(&cr.ChatGroupCount)
	global.DB.Model(&models.LoginDataModel{}).Where("to_days(create_at)=to_days(now())").Select("count(id)").Scan(&cr.TodayLoginCount)
	global.DB.Model(&models.UserModel{}).Where("to_days(create_at)=to_days(now())").Select("count(id)").Scan(&cr.TodaySignCount)

	result, _ := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(elastic.NewMatchAllQuery()).
		Do(context.Background())
	cr.ArticleCount = result.Hits.TotalHits.Value

	res.OKWithData(cr, c)
}
