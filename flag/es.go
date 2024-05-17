package flag

import "GoRoLingG/models"

func EsCreateIndex() {
	models.ArticleModel{}.CreateIndex() //生成表结构
}
