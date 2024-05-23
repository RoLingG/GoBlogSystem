package redis_service

import (
	"GoRoLingG/global"
	"encoding/json"
	"fmt"
	"time"
)

const newsIndex = "news_index"

type NewsData struct {
	Index    int    `json:"index"`
	Title    string `json:"title"`
	HotValue string `json:"hotValue"`
	Link     string `json:"link"`
}

func (redis RedisService) SetNews(key string, newsData []NewsData) error {
	byteData, err := json.Marshal(newsData)                                                   //将对应数据序列化，也就是json化
	err = global.Redis.Set(fmt.Sprintf("%s_%s", newsIndex, key), byteData, 1*time.Hour).Err() //设置一个小时过期
	//err = global.Redis.HSet(newsIndex, key, byteData).Err()
	return err
}

func (redis RedisService) GetNews(key string) (newsData []NewsData, err error) {
	//res := global.Redis.HGet(newsIndex, key).Val()\
	res := global.Redis.Get(fmt.Sprintf("%s_%s", newsIndex, key)).Val()
	err = json.Unmarshal([]byte(res), &newsData) //将对应数据反序列化
	return
}
