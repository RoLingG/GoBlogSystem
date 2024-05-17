package models

type CommentModel struct {
	Model
	SubComments        []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"subComments"`          //子评论列表
	ParentCommentModel *CommentModel   `gorm:"foreignKey:ParentCommentID" json:"parent_comment_model"` //父评论
	ParentCommentID    *uint           `json:"parent_comment_id"`                                      //父评论ID
	Content            string          `gorm:"size:256" json:"content"`                                //评论内容
	DiggCount          int             `gorm:"size:8;default:0;" json:"digg_count"`                    //点赞数
	CommentCount       int             `gorm:"size:8;default:0" json:"comment_count"`                  //子评论数
	ArticleID          string          `gorm:"size:32" json:"article_id"`                              //关联文章的ID
	User               UserModel       `json:"user"`                                                   //关联的用户
	UserID             uint            `json:"user_id"`                                                //评论的用户
}
