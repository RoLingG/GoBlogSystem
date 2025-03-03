package models

import (
	"GoRoLingG/global"
	"gorm.io/gorm"
)

type CommentModel struct {
	Model              `json:",select(comment)"`
	SubComments        []CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments,select(comment)"` // 子评论列表
	ParentCommentModel *CommentModel  `gorm:"foreignKey:ParentCommentID" json:"comment_model,omit(comment)"`  // 父级评论
	ParentCommentID    *uint          `json:"parent_comment_id,select(comment)"`                              // 父评论id
	Content            string         `gorm:"size:256" json:"content,select(comment)"`                        // 评论内容
	DiggCount          int            `gorm:"size:8;default:0;" json:"digg_count,select(comment)"`            // 点赞数
	CommentCount       int            `gorm:"size:8;default:0;" json:"comment_count,select(comment)"`         // 子评论数
	ArticleID          string         `gorm:"size:32" json:"article_id,select(comment)"`                      // 文章id
	User               UserModel      `json:"user,select(comment)"`                                           // 关联的用户
	UserID             uint           `json:"user_id,select(comment)"`                                        // 评论的用户
}

// BeforeDelete 钩子函数
func (c *CommentModel) BeforeDelete(tx *gorm.DB) (err error) {
	// 先把子评论删掉
	return nil
}

// FindAllSubCommentList 找一个评论的所有子评论,一维化
func FindAllSubCommentList(com CommentModel) (subList []CommentModel) {
	global.DB.Preload("SubComments").Preload("User").Take(&com)
	for _, model := range com.SubComments {
		subList = append(subList, model)
		subList = append(subList, FindAllSubCommentList(model)...)
	}
	return
}

// GetCommentTree 获取评论树
func GetCommentTree(rootComment *CommentModel) *CommentModel {
	var subComments []*CommentModel
	global.DB.Preload("User").Where("parent_comment_id = ?", rootComment.ID).Find(&subComments)
	// 递归获取子评论树
	// 将子评论指针数组转换为值数组
	rootComment.SubComments = make([]CommentModel, len(subComments))
	for i, subComment := range subComments {
		rootComment.SubComments[i] = *subComment
		GetCommentTree(&rootComment.SubComments[i]) // 递归调用
	}
	return rootComment
}
