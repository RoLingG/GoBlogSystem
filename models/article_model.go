package models

import (
	"GoRoLingG/global"
	"GoRoLingG/models/ctype"
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

type ArticleModel struct {
	ID       string `json:"id"`        //es的ID
	CreateAt string `json:"create_at"` //创建时间
	UpdateAt string `json:"update_at"` //更新时间

	Title    string `json:"title"`              //文章标题
	Keyword  string `json:"keyword,omit(list)"` // 关键字，用于检测文章是否存在，值一般与title一致
	Abstract string `json:"abstract"`           //文章简介
	Content  string `json:"content,omit(list)"` //文章正文

	LookCount    int `json:"look_count"`    //文章观看数
	CommentCount int `json:"comment_count"` //文章评论数
	DiggCount    int `json:"digg_count"`    //文章点赞数
	CollectCount int `json:"collect_count"` //文章收藏数

	CommentModel []CommentModel `json:"-"` //文章评论列表

	UserID       uint   `json:"user_id"`        //文章作者ID
	UserNickName string `json:"user_nick_name"` //文章用户昵称
	UserAvatar   string `json:"user_avatar"`    //文章用户头像

	Category string `json:"category"`          //文章分类
	Source   string `json:"source,omit(list)"` //资源来源
	Link     string `json:"link,omit(list)"`   //原文链接
	//Words    int    `json:"words"`             //文章总字数

	ImageID  uint   `json:"image_id"`  //文章封面ID
	ImageUrl string `json:"image_url"` //文章封面url

	Tags ctype.Array `json:"tags"` //文章标签，这里的tags分成两个也是和上面用户名同理
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
      "collects_count": {
        "type": "integer"
      },
      "user_id": {
        "type": "integer"
      },
      "user_nick_name": { 
        "type": "text"
      },
      "user_avatar": { 
        "type": "text"
      },
      "category": { 
        "type": "text"
      },
      "source": { 
        "type": "text"
      },
      "link": { 
        "type": "text"
      },
      "image_id": {
        "type": "integer"
      },
      "image_url": { 
        "type": "text"
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
func (article ArticleModel) Create() (err error) {
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
