package models

type ImageModel struct {
	Model
	Path string `json:"path"`                //图片路径
	Hash string `json:"hash"`                //图片哈希值，用于判断重复图片
	Name string `gorm:"size:32" json:"name"` //图片名称
}
