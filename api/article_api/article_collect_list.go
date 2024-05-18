package article_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/common"
	"GoRoLingG/utils/jwt"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

// CollectResponse 用于返回给前端收藏文章的数据
type CollectResponse struct {
	models.ArticleModel
	CreateAt string `json:"create_at"`
}

func (ArticleApi) ArticleUserCollectListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	var cr models.PageInfo
	c.ShouldBindQuery(&cr)

	var articleIDList []interface{}
	//分页
	list, count, err := common.CommonList(models.UserCollectModel{UserID: claims.UserID}, common.Option{
		PageInfo: cr,
	})
	var collectMap = map[string]string{}
	//获取文章id列表
	for _, model := range list {
		articleIDList = append(articleIDList, model.ArticleID)
		collectMap[model.ArticleID] = model.CreateAt.Format("2006-01-02 15:04:05") //获取收藏文章的map，只需要文章id和创建时间，key为文章id，对应的value为文章创建的时间
	}

	//传id列表查es
	var collectList = make([]CollectResponse, 0)
	boolSearch := elastic.NewTermsQuery("_id", articleIDList...)
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	fmt.Println(result.Hits.TotalHits.Value, articleIDList)
	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error(err)
			continue
		}
		article.ID = hit.Id
		collectList = append(collectList, CollectResponse{
			ArticleModel: article,
			CreateAt:     collectMap[hit.Id],
		})
	}
	res.OKWithList(collectList, count, c)
}

//没写完
