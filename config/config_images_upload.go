package config

type ImagesUpload struct {
	Size int    `json:"size" yaml:"size"` //图片上传的大小，单位为MB
	Path string `json:"path" yaml:"path"` //图片上传的路径，用于本地图片上传
}
