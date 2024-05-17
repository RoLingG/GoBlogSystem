package main

import (
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
	"strings"
)

func main() {
	//github.com/PuerkitoBio/goquery 	markdown转html显示
	unsafe := blackfriday.MarkdownCommon([]byte("### 你好\n ```go\nprint('hello')\n```\n - 测试测试 \n \n<script>alert(测试测试)</script>\n\n ![图片](http://xxx.com)"))
	fmt.Println(string(unsafe))

	//github.com/PuerkitoBio/goquery		html获取文本内容
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	fmt.Println(doc.Text())
	//doc.Find("script").Remove() //过滤对对应html标签的信息
	//fmt.Println(doc.Text())
	nodes1 := doc.Find("script").Nodes
	fmt.Println(nodes1)
	nodes2 := doc.Find("h1").Nodes
	fmt.Println(nodes2)

	//github.com/JohannesKaufmann/html-to-markdown		html转markdown
	converter := md.NewConverter("", true, nil)
	html, _ := doc.Html()
	markdown, err := converter.ConvertString(html)
	fmt.Println(markdown, err)
}
