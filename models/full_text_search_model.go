package models

import (
	"GoRoLingG/global"
	"context"
	"github.com/sirupsen/logrus"
)

type FullTextSearchModel struct {
	ID    string `json:"id" structs:"id"`       //es的ID
	Key   string `json:"key"`                   //文章关联的id
	Title string `json:"title" structs:"title"` //文章标题
	Slug  string `json:"slug" struct:"slug"`    //标题的跳转地址
	Body  string `json:"body" structs:"body"`   //文章正文
}

func (FullTextSearchModel) Index() string {
	return global.Config.ES.FullTextSearchIndex
}

func (FullTextSearchModel) Mapping() string {
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
      "key": { 
        "type": "keyword"
      },
      "slug": { 
        "type": "keyword"
      },
      "body": { 
        "type": "text"
      }
    }
  }
}
`
}

// IndexExists 索引是否存在
func (fts FullTextSearchModel) IndexExists() bool {
	exists, err := global.ESClient.
		IndexExists(fts.Index()).
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
func (fts FullTextSearchModel) CreateIndex() error {
	if fts.IndexExists() {
		// 有索引则删掉索引
		fts.RemoveIndex()
	}
	// 没有索引
	// 创建索引，相当于刷新索引
	createIndex, err := global.ESClient.
		CreateIndex(fts.Index()).
		BodyString(fts.Mapping()).
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
	logrus.Infof("索引 %s 创建成功", fts.Index())
	return nil
}

// RemoveIndex es索引删除
func (fts FullTextSearchModel) RemoveIndex() error {
	logrus.Info("索引存在，删除索引")
	// 删除索引
	deleteIndex, err := global.ESClient.
		DeleteIndex(fts.Index()).
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
