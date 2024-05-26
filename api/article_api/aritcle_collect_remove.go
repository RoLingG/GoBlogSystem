package article_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/es_service"
	"GoRoLingG/utils/jwt"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

// ArticleUserCollectRemoveView 用户取消文章收藏
// @Tags 文章管理
// @Summary 用户取消文章收藏
// @Description	用户取消文章收藏
// @Param token header string true "Authorization token"
// @Param data body models.ESIDListRequest true	"当前用户取消文章收藏的一些参数"
// @Produce json
// @Router /api/articleCollectRemove [delete]
// @Success 200 {object} res.Response{}
func (ArticleApi) ArticleUserCollectRemoveView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	var cr models.ESIDListRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var collectList []models.UserCollectModel
	var articleIDList []string
	//先找出当前登录用户收藏过的文章，再将文章id找出来到articleIDList中，这样能规避乱填文章id可能显示取消收藏成功的风险
	err = global.DB.Find(&collectList, "user_id = ? and article_id in ?", claims.UserID, cr.IDList).Select("article_id").Scan(&articleIDList).Error
	if len(articleIDList) == 0 {
		res.FailWithMsg("删除请求非法", c)
		return
	}

	//这里实例化一个[]interface{}类型的IDList，而不是直接在下面用articleIDList是因为NewTermsQuery()第二个传过去的参数得是[]interface{}类型，因此要周转一步
	var IDList []interface{}
	for _, articleID := range articleIDList {
		IDList = append(IDList, articleID)
	}
	//更新文章数
	boolSearch := elastic.NewTermsQuery("_id", IDList...) //根据字段精确匹配要用NewTermsQuery()
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(boolSearch).
		Size(1000).
		Do(context.Background())
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		if err != nil {
			global.Log.Error(err)
			continue
		}
		//es中对应文章的收藏数-1
		collectCount := article.CollectCount - 1
		//更新文章数据
		err = es_service.ArticleUpdate(hit.Id, map[string]any{
			"collect_count": collectCount,
		})
		if err != nil {
			global.Log.Error(err)
			continue
		}
	}
	//别忘了删除数据库中用户对应取消收藏文章的数据
	global.DB.Delete(&collectList)
	res.OKWithMsg(fmt.Sprintf("成功取消收藏 %d 篇文章", len(articleIDList)), c)
}
