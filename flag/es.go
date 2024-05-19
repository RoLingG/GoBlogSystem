package flag

import "GoRoLingG/models"

func EsCreateIndex() {
	models.ArticleModel{}.CreateIndex()        //生成文章es索引
	models.FullTextSearchModel{}.CreateIndex() //生成全文搜索es索引
}

func EsRemoveIndex() {
	models.ArticleModel{}.RemoveIndex()        //删除文章es索引
	models.FullTextSearchModel{}.RemoveIndex() //删除全文搜索es索引
}
