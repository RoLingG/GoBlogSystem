package ctype

import "encoding/json"

type ImageType int

const (
	Local ImageType = 1 //本地
	QiNiu ImageType = 2 //七牛云
)

// 枚举，算是gorm里面的一个重要知识
func (image ImageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(image.String())
}

func (image ImageType) String() string {
	var str string
	switch image {
	case Local:
		str = "本地存储"
	case QiNiu:
		str = "七牛云存储"
	default:
		str = "未知存储类型"
	}
	return str
}
