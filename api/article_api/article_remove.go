package article_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/es_service"
	"GoRoLingG/utils/jwt"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

// IDListRequest 老样子，用list获取批量删除的id
type IDListRequest struct {
	IDList []string `json:"id_list"`
}

// ArticleRemoveView 文章删除
// @Tags 文章管理
// @Summary 文章删除
// @Description	删除文章
// @Param token header string true "Authorization token"
// @Param data body IDListRequest true	"删除文章的一些参数"
// @Produce json
// @Router /api/articleRemove [delete]
// @Success 200 {object} res.Response{}
func (ArticleApi) ArticleRemoveView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var cr IDListRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	//es中要批量操作得用bulk桶，将文章索引装入桶中，这样桶中就有文章索引的所有内部数据
	//如果删除了用户收藏过的文章，该怎么办(应该顺带把文章关联的收藏也删了)
	bulkService := global.ESClient.Bulk().
		Index(models.ArticleModel{}.Index()).
		Refresh("true")
	for _, id := range cr.IDList {
		//给桶返送删除请求，根据对应id删除桶中的对应文章数据
		req := elastic.NewBulkDeleteRequest().Id(id)
		bulkService.Add(req)
	}
	//执行桶中操作
	result, err := bulkService.Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("删除失败", c)
		return
	}
	//删除成功，同步全文搜索索引数据，顺便把有收藏过要删除文章的用户收藏给取消掉
	for _, articleID := range cr.IDList {
		var collect models.UserCollectModel
		err = global.DB.Take(&collect, "user_id = ? and article_id = ?", claims.UserID, articleID).Error
		if err == nil {
			// 找到则取消收藏文章
			global.DB.Delete(&collect)
		}
		es_service.DeleteFullTextSearchByID(articleID)
	}
	res.OKWithMsg(fmt.Sprintf("成功删除 %d 篇文章", len(result.Succeeded())), c)
	return
}
