package main

import (
	"context"
	"github.com/sirupsen/logrus"
)

// Mapping text字段可以进行模糊匹配
func (DemoModel) Mapping() string {
	return `
{
  "settings": {
    "index":{
      "max_result_window": "100000"
    }
  }, 
  "mappings": {
    "properties": {
      "title": { 
        "type": "text"
      },
      "user_id": {
        "type": "integer"
      },
      "create_at":{
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      }
    }
  }
}
`
}

// IndexExists 索引是否存在
func (demo DemoModel) IndexExists() bool {
	exists, err := client.
		IndexExists(demo.Index()).
		Do(context.Background())
	if err != nil {
		//存在则报错
		logrus.Error(err.Error())
		return exists //exists为true
	}
	//不存在则直接返回
	return exists //exists为false
}

// CreateIndex es索引添加
func (demo DemoModel) CreateIndex() error {
	if demo.IndexExists() {
		// 有索引则删掉索引
		demo.RemoveIndex()
	}
	// 没有索引
	// 创建索引，相当于刷新索引
	createIndex, err := client.
		CreateIndex(demo.Index()).
		BodyString(demo.Mapping()).
		Do(context.Background())
	if err != nil {
		logrus.Error("创建索引失败")
		logrus.Error(err.Error())
		return err
	}
	//确认索引的创建失败，则报错
	if !createIndex.Acknowledged {
		logrus.Error("创建失败")
		return err
	}
	logrus.Infof("索引 %s 创建成功", demo.Index())
	return nil
}

// RemoveIndex es索引删除
func (demo DemoModel) RemoveIndex() error {
	logrus.Info("索引存在，删除索引")
	// 删除索引
	deleteIndex, err := client.
		DeleteIndex(demo.Index()).
		Do(context.Background())
	if err != nil {
		logrus.Error("删除索引失败")
		logrus.Error(err.Error())
		return err
	}
	//确认索引的删除失败，则报错
	if !deleteIndex.Acknowledged {
		logrus.Error("删除索引失败")
		return err
	}
	logrus.Info("索引删除成功")
	return nil
}
