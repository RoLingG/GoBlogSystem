package article_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/models/ctype"
	"GoRoLingG/res"
	"GoRoLingG/service/es_service"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"time"
)

type ArticleUpdateRequest struct {
	Title    string   `json:"title"`    // 文章标题
	Abstract string   `json:"abstract"` // 文章简介
	Content  string   `json:"content"`  // 文章内容
	Category string   `json:"category"` // 文章分类
	Source   string   `json:"source"`   // 文章来源
	Link     string   `json:"link"`     // 原文链接
	ImageID  uint     `json:"image_id"` // 文章封面id
	Tags     []string `json:"tags"`     // 文章标签	tags在model里是ctype.Array类型，ctype.Array类型本质就是一个[]string，所以这里能用[]string接收
	ID       string   `json:"id"`
}

// ArticleUpdateView 文章更新
// @Tags 文章管理
// @Summary 文章更新
// @Description	更新文章
// @Param token header string true "Authorization token"
// @Param data body ArticleUpdateRequest true	"更新文章的一些参数"
// @Produce json
// @Router /api/articleUpdate [put]
// @Success 200 {object} res.Response{}
func (ArticleApi) ArticleUpdateView(c *gin.Context) {
	//_claims, _ := c.Get("claims")
	//claims := _claims.(*jwt.CustomClaims)
	var cr ArticleUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithError(err, &cr, c)
		return
	}

	//如果要更新文章的话，获取文章图片更新后的图片路径
	var imageUrl string
	if cr.ImageID != 0 {
		err = global.DB.Model(models.ImageModel{}).Where("id = ?", cr.ImageID).Select("path").Scan(&imageUrl).Error //将图片表中的图片路径获取出来
		if err != nil {
			res.FailWithMsg("图片不存在", c)
			return
		}
	}

	//更新后的文章数据
	article := models.ArticleModel{
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
		Title:    cr.Title,
		Keyword:  cr.Title,
		Abstract: cr.Abstract,
		Content:  cr.Content,
		Category: cr.Category,
		Source:   cr.Source,
		Link:     cr.Link,
		ImageID:  cr.ImageID,
		ImageUrl: imageUrl,
		Tags:     cr.Tags,
	}

	maps := structs.Map(&article)  //将article map化，好进行添加，这里因为用了structs，所以ArticleModel里面的对应参数要加上structs标签
	var DataMap = map[string]any{} //用于获取map化后的将article实例的所有非空参数
	// 去掉map内相关参数的空值，并将非空的value传进map里
	for key, value := range maps {
		switch val := value.(type) {
		case string:
			if val == "" {
				continue
			}
		case uint:
			if val == 0 {
				continue
			}
		case int:
			if val == 0 {
				continue
			}
		case ctype.Array: //虽然ctype.Array本质是[]string，但还是得这样写才生效
			if len(val) == 0 {
				continue
			}
		case []string:
			if len(val) == 0 {
				continue
			}
		}
		//将非空参数的key的值赋值进去，为的就是将不修改的参数不会修改为空，只修改要修改的参数
		DataMap[key] = value
	}

	//检测对应id要更新的文章是否存在，注意，这一步会作用于article
	err = article.GetDataByID(cr.ID)
	if err != nil {
		res.FailWithMsg("对应id的文章不存在", c)
		return
	}

	err = es_service.ArticleUpdate(cr.ID, DataMap)
	//更新失败
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("文章更新失败", c)
		return
	}

	//更新成功，同步数据到全文搜索
	newArticle, _ := es_service.CommonDetail(cr.ID)
	if article.Content != newArticle.Content || article.Title != newArticle.Title {
		es_service.DeleteFullTextSearchByID(cr.ID)
		es_service.AsyncArticleByFullTextSearch(cr.ID, article.Title, newArticle.Content)
	}

	res.OKWithMsg("文章更新成功", c)
}
