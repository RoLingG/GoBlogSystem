package article_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

// IDListRequest 老样子，用list获取批量删除的id
type IDListRequest struct {
	IDList []string `json:"id_list"`
}

func (ArticleApi) ArticleRemoveView(c *gin.Context) {
	var cr IDListRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err)
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	//es中要批量操作得用bulk桶，将文章索引装入桶中，这样桶中就有文章索引的所有内部数据
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
	res.OKWithMsg(fmt.Sprintf("成功删除 %d 篇文章", len(result.Succeeded())), c)
	return

}
