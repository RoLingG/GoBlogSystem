package article_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/es_service"
	"GoRoLingG/utils/jwt"
	"github.com/gin-gonic/gin"
)

// ArticleUserCollectView 收藏文章
// @Tags 文章管理
// @Summary 收藏文章
// @Description	收藏文章
// @Param token header string true "Authorization token"
// @Param data body models.ESIDRequest true	"收藏文章的一些参数"
// @Produce json
// @Router /api/articleCollect [post]
// @Success 200 {object} res.Response{}
func (ArticleApi) ArticleUserCollectView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	var cr models.ESIDRequest
	err := c.ShouldBindUri(&cr) //通过uri去进行获取文章id
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	article, err := es_service.CommonDetail(cr.ID)
	if err != nil {
		res.FailWithMsg("文章不存在", c)
		return
	}

	var collect models.UserCollectModel
	err = global.DB.Take(&collect, "user_id = ? and article_id = ?", claims.UserID, cr.ID).Error
	var num = -1
	if err != nil {
		// 没有找到 收藏文章
		global.DB.Create(&models.UserCollectModel{
			UserID:    claims.UserID,
			ArticleID: cr.ID,
		})
		// 给文章的收藏数 +1
		num = 1
	} else {
		// 找到文章 取消收藏文章
		global.DB.Delete(&collect)
	}

	// 更新文章收藏数
	err = es_service.ArticleUpdate(cr.ID, map[string]any{
		"collect_count": article.CollectCount + num,
	})
	if num == 1 {
		res.OKWithMsg("收藏文章成功", c)
	} else {
		res.OKWithMsg("取消收藏成功", c)
	}
}
