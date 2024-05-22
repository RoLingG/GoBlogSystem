package main

import (
	"GoRoLingG/core"
	"GoRoLingG/global"
	"GoRoLingG/models"
	"fmt"
)

func main() {
	core.InitConfig()
	global.Log = core.InitLogger()
	global.DB = core.InitGorm()
	FindArticleCommmentList("iDA_lo8BgM_PmuvUtu50")
}

func FindArticleCommmentList(articleID string) (rootCommentList []*models.CommentModel) {
	//先把文章下的父论根查出来
	global.DB.Find(&rootCommentList, "article_id = ? and parent_comment_id is null", articleID)
	//遍历父评论，将对应的子评论给递归出来
	for _, model := range rootCommentList {
		var subCommentList []models.CommentModel
		FindSubComment(*model, &subCommentList)
		model.SubComments = subCommentList
		fmt.Println(model.Content, subCommentList)
	}
	return
}

// FindSubComment 递归查评论下的子评论
func FindSubComment(model models.CommentModel, subCommentList *[]models.CommentModel) {
	global.DB.Preload("SubComments.User").Take(&model)
	for _, sub := range model.SubComments {
		*subCommentList = append(*subCommentList, sub)
		FindSubComment(sub, subCommentList) //子评论的子评论也是在根评论的下一级，属于同级关系，不会出现一直套娃的现象
	}
	return
}
