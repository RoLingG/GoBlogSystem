package article_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/es_serivce"
	"GoRoLingG/utils/jwt"
	"github.com/gin-gonic/gin"
)

// 用户收藏或取消收藏文章
func (ArticleApi) ArticleUserCollectView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	//实例化es的id列表
	var cr models.ESIDRequest
	err := c.ShouldBindJSON(&cr) //通过uri去进行获取文章id
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	article, err := es_serivce.CommonDetail(cr.ID)
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
	err = es_serivce.ArticleUpdate(cr.ID, map[string]any{
		"collect_count": article.CollectCount + num,
	})
	if num == 1 {
		res.OKWithMsg("收藏文章成功", c)
	} else {
		res.OKWithMsg("取消收藏成功", c)
	}
}
