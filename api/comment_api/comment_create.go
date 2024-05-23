package comment_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/es_service"
	"GoRoLingG/service/redis_service"
	"GoRoLingG/utils/jwt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentRequest struct {
	ArticleID       string `json:"article_id" binding:"required" msg:"请选择文章"`
	Content         string `json:"content" binding:"required" msg:"请输入评论内容"`
	ParentCommentID *uint  `json:"parent_comment_id"` // 父评论id
}

func (CommentApi) CommentCreateView(c *gin.Context) {
	var cr CommentRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	//查对应id的文章详情，看文章是否存在，报错则不存在
	_, err = es_service.CommonDetail(cr.ArticleID)
	if err != nil {
		res.FailWithMsg("文章不存在", c)
		return
	}

	//判断是否是子评论
	if cr.ParentCommentID != nil {
		//有父评论则是子评论
		var parentComment models.CommentModel
		//找父评论
		err = global.DB.Take(&parentComment, cr.ParentCommentID).Error
		if err != nil {
			res.FailWithMsg("父评论不存在", c)
			return
		}
		//判断父评论的文章是否当前文章一致
		if parentComment.ArticleID != cr.ArticleID {
			res.FailWithMsg("评论文章不一致", c)
			return
		}
		//给父评论的子评论+1
		global.DB.Model(&parentComment).Update("comment_count", gorm.Expr("comment_count + 1"))
	}
	// 添加评论
	global.DB.Create(&models.CommentModel{
		ParentCommentID: cr.ParentCommentID,
		Content:         cr.Content,
		ArticleID:       cr.ArticleID,
		UserID:          claims.UserID,
	})
	// 拿到文章数，新的文章评论数存缓存里
	redis_service.NewArticleCommentIndex().Set(cr.ArticleID)
	//同步数据到es的对应文章中
	article, err := es_service.CommonDetail(cr.ArticleID)
	err = es_service.ArticleUpdate(cr.ArticleID, map[string]any{
		"comment_count": article.CollectCount + 1,
	})
	res.OKWithMsg("文章评论成功", c)
	return
}
