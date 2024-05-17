package article_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/models/ctype"
	"GoRoLingG/res"
	"GoRoLingG/utils/jwt"
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday"
	"math/rand"
	"strings"
	"time"
)

type ArticleRequest struct {
	Title    string      `json:"title" binding:"required" msg:"文章标题必填"`   //文章标题
	Abstract string      `json:"abstract"`                                //文章简介，不填就要根据正文内容摘选
	Content  string      `json:"content" binding:"required" msg:"文章正文必填"` //文章正文
	Category string      `json:"category"`                                //文章分类
	Source   string      `json:"source"`                                  //资源来源
	Link     string      `json:"link"`                                    //原文链接
	ImageID  uint        `json:"image_id"`                                //文章封面ID
	Tags     ctype.Array `json:"tags"`                                    //文章标签
}

func (ArticleApi) ArticleCreateView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var cr ArticleRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	userID := claims.UserID
	userNickName := claims.NickName
	//校验content xss防范

	//处理content markdown转html
	contentHTML := blackfriday.MarkdownCommon([]byte(cr.Content))
	//是不是有script标签
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(contentHTML))) //获取html的内容
	nodes := doc.Find("script").Nodes
	if len(nodes) > 0 {
		//有script标签
		doc.Find("script").Remove() //去除掉标签及其内容
		//html转换回markdown
		converter := md.NewConverter("", true, nil)
		html, _ := doc.Html()
		markdown, _ := converter.ConvertString(html)
		cr.Content = markdown
	}

	//没有写文章简介的逻辑处理
	if cr.Abstract == "" {
		// 汉字的截取不一样
		//abs := []rune(cr.Content)
		abs := []rune(doc.Text())
		// 将content转为html，并且过滤xss，以及获取中文内容
		if len(abs) > 100 {
			cr.Abstract = string(abs[:100])
		} else {
			cr.Abstract = string(abs)
		}
	}

	// 不传banner_id,后台就随机去选择一张
	if cr.ImageID == 0 {
		var imageIDList []uint
		global.DB.Model(models.ImageModel{}).Select("id").Scan(&imageIDList) //将数据库中所有图片的id扫出来
		if len(imageIDList) == 0 {
			//数据库中没有图片
			res.FailWithMsg("没有对应的图片数据", c)
			return
		}
		//数据库中有图片，则随机挑
		fmt.Println(imageIDList)
		rand.Seed(time.Now().UnixNano())                      //设置随机种子
		cr.ImageID = imageIDList[rand.Intn(len(imageIDList))] //随机挑一张，Intn()为设置随机最大范围
	}

	//查询image_id对应的image_url
	var imageUrl string
	err = global.DB.Model(&models.ImageModel{}).Where("id = ?", cr.ImageID).Select("path").Scan(&imageUrl).Error
	if err != nil {
		res.FailWithMsg("图片不存在", c)
		return
	}

	// 查找用户头像
	var avatar string
	err = global.DB.Model(models.UserModel{}).Where("id = ?", userID).Select("avatar").Scan(&avatar).Error //将对应id的用户的头像扫描出来
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	article := models.ArticleModel{
		CreateAt:     now,
		UpdateAt:     now,
		Title:        cr.Title,
		Keyword:      cr.Title,
		Abstract:     cr.Abstract,
		Content:      cr.Content,
		UserID:       userID,
		UserNickName: userNickName,
		UserAvatar:   avatar,
		Category:     cr.Category,
		Source:       cr.Source,
		Link:         cr.Link,
		ImageID:      cr.ImageID,
		ImageUrl:     imageUrl,
		Tags:         cr.Tags,
	}
	if article.ISExistData() {
		//如果文章存在，则不添加文章
		global.Log.Error(err)
		res.FailWithMsg("文章已存在", c)
		return
	}

	err = article.Create()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.OKWithMsg("文章发布成功", c)
}
