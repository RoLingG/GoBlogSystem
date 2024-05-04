package config

type ImageUpload struct {
	Size int    `json:"size" yaml:"size"` //图片上传的大小，单位为MB
	Path string `json:"path" yaml:"path"` //图片上传的路径
}
