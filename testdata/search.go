package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
	"strings"
)

func main() {
	var data = "## 环境搭建\n\n拉取镜像\n\n```Python\ndocker pull elasticsearch:7.12.0\n```\n\n\n\n创建docker容器挂在的目录：\n\n```Python\nmkdir -p /opt/elasticsearch/config & mkdir -p /opt/elasticsearch/data & mkdir -p /opt/elasticsearch/plugins\n\nchmod 777 /opt/elasticsearch/data\n\n```\n\n配置文件\n\n```Python\necho \"http.host: 0.0.0.0\" >> /opt/elasticsearch/config/elasticsearch.yml\n```\n\n\n\n创建容器\n\n```Python\n# linux\ndocker run --name es -p 9200:9200  -p 9300:9300 -e \"discovery.type=single-node\" -e ES_JAVA_OPTS=\"-Xms84m -Xmx512m\" -v /opt/elasticsearch/config/elasticsearch.yml:/usr/share/elasticsearch/config/elasticsearch.yml -v /opt/elasticsearch/data:/usr/share/elasticsearch/data -v /opt/elasticsearch/plugins:/usr/share/elasticsearch/plugins -d elasticsearch:7.12.0\n```\n\n\n\n访问ip:9200能看到东西\n\n![](http://python.fengfengzhidao.com/pic/20230129212040.png)\n\n就说明安装成功了\n\n\n\n浏览器可以下载一个 `Multi Elasticsearch Head` es插件\n\n\n\n第三方库\n\n```Go\ngithub.com/olivere/elastic/v7\n```\n\n## es连接\n\n```Go\nfunc EsConnect() *elastic.Client  {\n  var err error\n  sniffOpt := elastic.SetSniff(false)\n  host := \"http://127.0.0.1:9200\"\n  c, err := elastic.NewClient(\n    elastic.SetURL(host),\n    sniffOpt,\n    elastic.SetBasicAuth(\"\", \"\"),\n  )\n  if err != nil {\n    logrus.Fatalf(\"es连接失败 %s\", err.Error())\n  }\n  return c\n}\n```"
	GetSearchIndexDataByContent("/article/hd893bxGHD84", "es的环境搭建", data)
}

type SearchData struct {
	Body  string `json:"body"`  // 正文
	Slug  string `json:"slug"`  // 包含文章的id 的跳转地址
	Title string `json:"title"` // 标题
}

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
		})
	}
	b, _ := json.Marshal(searchDataList) //将对应实例转换成json好传给前端
	fmt.Println(string(b))
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
