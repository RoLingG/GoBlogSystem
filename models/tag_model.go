package models

type TagModel struct {
	Model
	Title string `gorm:"size:16" json:"title"`
}
