package models

import (
	"GoRoLingG/global"
	"GoRoLingG/models/ctype"
	"gorm.io/gorm"
	"os"
)

type ImageModel struct {
	Model
	Path      string          `json:"path"`                        //图片路径
	Hash      string          `json:"hash"`                        //图片哈希值，用于判断重复图片
	Name      string          `gorm:"size:128" json:"name"`        //图片名称
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"` //图片存储类型
}

// hook钩子函数,当要删除图片的前一段时间触发
func (image ImageModel) BeforeDelete(db *gorm.DB) (err error) {
	//判断是否是本地图片，如果是本地图片则既要删除数据库数据，也要删除本地图片资源
	if image.ImageType == ctype.Local {
		//删除本地图片存储资源
		err = os.Remove(image.Path)
		if err != nil {
			global.Log.Error(err)
			return
		}
	}
	return nil
}
