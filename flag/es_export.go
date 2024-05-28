package flag

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type ESRawMessage struct {
	Row json.RawMessage `json:"row"` //json.RawMessage是一个特殊的类型，用于存储JSON格式的原始数据。
	ID  string          `json:"id"`
}

type ESIndexResponse struct {
	Data    []ESRawMessage `json:"data"`
	Mapping string         `json:"mapping"`
	Index   string         `json:"index"`
}

func ESDump() {
	index := models.FullTextSearchModel{}.Index()
	mapping := models.FullTextSearchModel{}.Mapping()
	esDump(index, mapping)
	index = models.ArticleModel{}.Index()
	mapping = models.ArticleModel{}.Mapping()
	esDump(index, mapping)
}

// 将对应索引的数据进行导出
func esDump(index, mapping string) {
	result, err := global.ESClient.Search(index).Query(elastic.NewMatchAllQuery()).Size(10000).Do(context.Background())
	if err != nil {
		logrus.Fatalf("索引%s err: %s", index, err.Error())
	}
	var dataList []ESRawMessage
	for _, hit := range result.Hits.Hits {
		dataList = append(dataList, ESRawMessage{
			Row: hit.Source,
			ID:  hit.Id,
		})
	}
	response := ESIndexResponse{
		Data:    dataList, //Data为对应索引导出的素具
		Mapping: mapping,  //Mapping为索引内数据的存储结构
		Index:   index,    //Index为索引的名字
	}

	fileName := fmt.Sprintf("%s_%s.json", index, time.Now().Format("20060102")) //将数据写入对应的json文件中
	file, _ := os.Create(fileName)

	byteData, _ := json.Marshal(response) //将返回数据json化
	file.Write(byteData)
	file.Close()

	logrus.Infof("%s 索引内的数据导出成功，导出到了 %s", index, fileName)
}
