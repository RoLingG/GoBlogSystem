package es_service

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/olivere/elastic/v7"
	"github.com/russross/blackfriday"
	"github.com/sirupsen/logrus"
	"strings"
)

type SearchData struct {
	Body  string `json:"body"`  // 正文
	Key   string `json:"key"`   //文章关联的id
	Slug  string `json:"slug"`  // 包含文章的id 的跳转地址
	Title string `json:"title"` // 标题
}

// AsyncArticleByFullTextSearch es同步文章数据到全文搜索
func AsyncArticleByFullTextSearch(id, title, content string) {
	indexList := GetSearchIndexDataByContent(id, title, content)
	bulk := global.ESClient.Bulk()
	for _, indexData := range indexList {
		request := elastic.NewBulkIndexRequest().Index(models.FullTextSearchModel{}.Index()).Doc(indexData)
		bulk.Add(request)
	}
	fullTextSearchRes, err := bulk.Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Infof("%s 添加成功,共添加了 %d 条", title, len(fullTextSearchRes.Succeeded()))
}

// DeleteFullTextSearchByID 删除全文搜索中对应id的文章数据
func DeleteFullTextSearchByID(id string) {
	boolSearch := elastic.NewTermQuery("key", id)
	ftsRes, _ := global.ESClient.DeleteByQuery().Index(models.FullTextSearchModel{}.Index()).Query(boolSearch).Size(1000).Do(context.Background())
	logrus.Infof("删除成功,共删除了 %d 条索引记录", ftsRes.Deleted)
}

// GetSearchIndexDataByContent 全文搜索
func GetSearchIndexDataByContent(id, title, content string) (searchDataList []SearchData) {
	dataList := strings.Split(content, "\n") //按行分割内容
	var isCode bool = false
	var headList, bodyList []string
	var body string
	headList = append(headList, getHeader(title)) //将文章标题作为内容标题之一加进去，以防出现没有内容标题直接写内容的情况导致无法划分好内容与标题
	for _, s := range dataList {
		// #{1,6}
		// 判断当前行是否是代码块始、末
		if strings.HasPrefix(s, "```") {
			isCode = !isCode //第一次遇到代码块起始则变为true，再次遇到代码块末则变为false，很好的解决了isCode转换的问题
		}
		if strings.HasPrefix(s, "#") && !isCode {
			headList = append(headList, getHeader(s))  //将该行内容加进标题列表里
			bodyList = append(bodyList, getBody(body)) //将之前组合起来的正文内容传进正文列表里分隔好，
			body = ""                                  //将暂存正文清空，备用于下一次遇到标题时正文存储进正文列表里分隔
			continue
		}
		body += s //如果该行不是标题，则作为正文添加进当前标题下的正文中
	}
	bodyList = append(bodyList, getBody(body)) //最后没有标题，所以存储的当前正文要加进正文列表中
	ln := len(headList)
	for i := 0; i < ln; i++ {
		//将分隔好的文章标题、正文组合起来
		searchDataList = append(searchDataList, SearchData{
			Title: headList[i],
			Body:  bodyList[i],
			Slug:  id + getSlug(headList[i]),
			Key:   id,
		})
	}
	//b, _ := json.Marshal(searchDataList) //将对应实例转换成json好传给前端
	//fmt.Println(string(b))
	return searchDataList
}

func getHeader(head string) string {
	//过滤掉#和空格
	head = strings.ReplaceAll(head, "#", "") //将#替换成无字符
	head = strings.ReplaceAll(head, " ", "") //将空格替换成无字符
	return head
}

func getBody(body string) string {
	unsafe := blackfriday.MarkdownCommon([]byte(body))                         //将markdown正文转化成html
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe))) //goquery 是一个HTML查询库，NewDocumentFromReader 函数用于从 io.Reader 接口创建一个 goquery 文档
	return doc.Text()
}

func getSlug(slug string) string {
	return "#" + slug
}
