package models

type TagModel struct {
	Model
	Title   string         `gorm:"size:16" json:"title"`
	Article []ArticleModel `gorm:"many2many:article_tag_models" json:"-"`
}
