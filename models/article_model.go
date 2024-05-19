package models

import (
	"GoRoLingG/global"
	"GoRoLingG/models/ctype"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

type ArticleModel struct {
	ID       string `json:"id" structs:"id"`               //es的ID
	CreateAt string `json:"create_at" structs:"create_at"` //创建时间
	UpdateAt string `json:"update_at" structs:"update_at"` //更新时间

	Title    string `json:"title" structs:"title"`                //文章标题
	Keyword  string `json:"keyword,omit(list)" structs:"keyword"` // 关键字，用于检测文章是否存在，值一般与title一致
	Abstract string `json:"abstract" structs:"abstract"`          //文章简介
	Content  string `json:"content,omit(list)" structs:"content"` //文章正文

	LookCount    int `json:"look_count" structs:"look_count"`       //文章观看数
	CommentCount int `json:"comment_count" structs:"comment_count"` //文章评论数
	DiggCount    int `json:"digg_count" structs:"digg_count"`       //文章点赞数
	CollectCount int `json:"collect_count" structs:"collect_count"` //文章收藏数

	CommentModel []CommentModel `json:"-"` //文章评论列表

	UserID       uint   `json:"user_id" structs:"user_id"`               //文章作者ID
	UserNickName string `json:"user_nick_name" structs:"user_nick_name"` //文章用户昵称
	UserAvatar   string `json:"user_avatar" structs:"user_avatar"`       //文章用户头像

	Category string `json:"category" structs:"category"`        //文章分类
	Source   string `json:"source,omit(list)" structs:"source"` //资源来源
	Link     string `json:"link,omit(list)" structs:"link"`     //原文链接
	//Words    int    `json:"words"`             //文章总字数

	ImageID  uint   `json:"image_id" structs:"image_id"`   //文章封面ID
	ImageUrl string `json:"image_url" structs:"image_url"` //文章封面url

	Tags ctype.Array `json:"tags" structs:"tags"` //文章标签，这里的tags分成两个也是和上面用户名同理
}

//var client *elastic.Client

func (ArticleModel) Index() string {
	return "article_index"
}

func (ArticleModel) Mapping() string {
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
      "keyword": { 
        "type": "keyword"
      },
      "abstract": { 
        "type": "text"
      },
      "content": { 
        "type": "text"
      },
      "look_count": {
        "type": "integer"
      },
      "comment_count": {
        "type": "integer"
      },
      "digg_count": {
        "type": "integer"
      },
      "collect_count": {
        "type": "integer"
      },
      "user_id": {
        "type": "integer"
      },
      "user_nick_name": { 
        "type": "keyword"
      },
      "user_avatar": { 
        "type": "keyword"
      },
      "category": { 
        "type": "keyword"
      },
      "source": { 
        "type": "keyword"
      },
      "link": { 
        "type": "keyword"
      },
      "image_id": {
        "type": "integer"
      },
      "image_url": { 
        "type": "keyword"
      },
      "tags": { 
        "type": "keyword"
      },
      "create_at":{
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      },
      "update_at":{
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
func (article ArticleModel) IndexExists() bool {
	exists, err := global.ESClient.
		IndexExists(article.Index()).
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
func (article ArticleModel) CreateIndex() error {
	if article.IndexExists() {
		// 有索引则删掉索引
		article.RemoveIndex()
	}
	// 没有索引
	// 创建索引，相当于刷新索引
	createIndex, err := global.ESClient.
		CreateIndex(article.Index()).
		BodyString(article.Mapping()).
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
	logrus.Infof("索引 %s 创建成功", article.Index())
	return nil
}

// RemoveIndex es索引删除
func (article ArticleModel) RemoveIndex() error {
	logrus.Info("索引存在，删除索引")
	// 删除索引
	deleteIndex, err := global.ESClient.
		DeleteIndex(article.Index()).
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

// Create 索引创建
func (article *ArticleModel) Create() (err error) {
	indexResponse, err := global.ESClient.Index().Index(article.Index()).BodyJson(article).Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	logrus.Infof("%#v", indexResponse)
	article.ID = indexResponse.Id //因为data是指针类型，所以会修改数据
	return nil
}

// ISExistData 是否存在该文章
func (article ArticleModel) ISExistData() bool {
	//查询通用写法
	res, err := global.ESClient.
		Search(article.Index()).
		Query(elastic.NewTermQuery("keyword", article.Title)).
		Size(1).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return false
	}
	if res.Hits.TotalHits.Value > 0 {
		//存在
		return true
	}
	return false
}

// GetDataByID 注意，这一步会作用于原来的值(因为是指针的原因)
func (article *ArticleModel) GetDataByID(id string) error {
	res, err := global.ESClient.Get().Index(article.Index()).Id(id).Do(context.Background())
	//判断对应id的文章是否存在
	if err != nil {
		return err
	}
	//检查在反序列化过程中是否出现任何问题
	//判断该article对应的结构体字段是否与es中存放的文章结构体匹配，不匹配则有err内容
	err = json.Unmarshal(res.Source, article)
	return err
}
