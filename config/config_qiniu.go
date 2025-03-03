package config

type QiNiu struct {
	IsEnable  bool    `json:"is_enable" yaml:"is_enable"` //是否启用七牛存储，默认不启用
	AccessKey string  `json:"access_key" yaml:"access_key"`
	SecretKey string  `json:"secret_key" yaml:"secret_key"`
	Bucket    string  `json:"bucket" yaml:"bucket"` //存储桶
	CDN       string  `json:"cdn" yaml:"cdn"`       //访问图片的地址前缀
	Zone      string  `json:"zone" yaml:"zone"`     //存储地区
	Size      float64 `json:"size" yaml:"size"`     //存储大小限制，单位大小为MB
	Prefix    string  `json:"prefix" yaml:"prefix"` //存储目录，或叫图片存储前缀
}
