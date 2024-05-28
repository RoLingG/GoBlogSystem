package flag

import (
	"GoRoLingG/global"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"os"
)

func ESImport(jsonPath string) {
	byteData, err := os.ReadFile(jsonPath)
	if err != nil {
		logrus.Fatalf("%s 文件读取出错，err: %s", jsonPath, err.Error())
		return
	}
	var response ESIndexResponse
	err = json.Unmarshal(byteData, &response)
	if err != nil {
		logrus.Fatalf("数据 %s 导入失败 err: %s", string(byteData), err.Error())
	}

	// 创建索引
	err = createIndexByJson(response.Index, response.Mapping)

	// 批量导入数据
	bulk := global.ESClient.Bulk().Index(response.Index).Refresh("true")
	for _, model := range response.Data {
		var mapData map[string]any
		_ = json.Unmarshal(model.Row, &mapData) //因为这里response.Data获取的是json文件的原始数据，带有换行
		row, _ := json.Marshal(mapData)         //所以这里经过Unmarshal之后又要Marshal回去，将换行去掉
		// 插入的数据，不能有换行
		req := elastic.NewBulkCreateRequest().Id(model.ID).Doc(string(row)) //string化是因为es的Doc()需要一个字符画的json数据切片
		bulk.Add(req)
	}
	res, err := bulk.Do(context.Background())
	if err != nil {
		logrus.Errorf("数据添加失败 err:%s", err.Error())
		return
	}
	logrus.Infof("数据添加成功， 共添加 %d 条", len(res.Succeeded()))
}

func createIndexByJson(index, mapping string) error {
	if indexExists(index) {
		logrus.Infof("索引 %s 已存在，无需创建", index)
		return nil
	}
	// 没有索引，则创建索引
	createIndex, err := global.ESClient.
		CreateIndex(index).
		BodyString(mapping).
		Do(context.Background())
	if err != nil {
		logrus.Errorf("创建索引失败, 原因是：%s", err.Error())
		return err
	}
	//确认索引的创建失败，则报错
	if !createIndex.Acknowledged {
		logrus.Errorf("%s 创建失败", index)
		return err
	}
	logrus.Infof("索引 %s 创建成功", index)
	return nil
}

func indexExists(index string) bool {
	exists, err := global.ESClient.
		IndexExists(index).
		Do(context.Background())
	if err != nil {
		//存在则报错
		logrus.Error(err.Error())
		return exists //exists为true
	}
	//不存在则直接返回
	return exists //exists为false
}
