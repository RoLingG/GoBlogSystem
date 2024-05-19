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

	//根据当前登录用户的id获取他收藏的文章，并有分页
	list, count, err := common.CommonList(models.UserCollectModel{UserID: claims.UserID}, common.Option{
		PageInfo: cr,
	})

	var articleIDList []interface{} //用于存储当前用户收藏的所有文章id
	var collectMap = map[string]string{}
	//获取当前登录用户收藏的文章id列表
	for _, model := range list {
		articleIDList = append(articleIDList, model.ArticleID)
		collectMap[model.ArticleID] = model.CreateAt.Format("2006-01-02 15:04:05") //获取收藏文章的map，只需要文章id和创建时间，key为文章id，与key对应的value为文章创建的时间
	}

	//传id列表查es，去获取文章
	var collectList = make([]CollectResponse, 0)
	boolSearch := elastic.NewTermsQuery("_id", articleIDList...) //根据字段精确匹配要用NewTermsQuery()
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	//fmt.Println(result.Hits.TotalHits.Value, articleIDList)
	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		fmt.Println(article)
		if err != nil {
			global.Log.Error(err)
			continue
		}
		//hit.Source里有除了_id之外的所有数据信息
		//因为这里ArticleModel的ID不再是mysql那种自增的，而是es的文章_id，所以要进行赋值获取
		article.ID = hit.Id
		collectList = append(collectList, CollectResponse{
			ArticleModel: article,
			CreateAt:     collectMap[hit.Id],
		})
	}
	res.OKWithList(collectList, count, c)
}

//没写完
