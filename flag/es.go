package flag

import "GoRoLingG/models"

func EsCreateIndex() {
	models.ArticleModel{}.CreateIndex() //生成文章表结构
}

func EsRemoveIndex() {
	models.ArticleModel{}.RemoveIndex() //删除文章表结构
}
