package config

type Email struct {
	Host             string `json:"host" yaml:"host"`
	Port             int    `json:"port" yaml:"port"`
	User             string `json:"user" yaml:"user"` //发送人邮箱
	Password         string `json:"password" yaml:"password"`
	DefaultFromEmail string `json:"default_from_email" yaml:"default_from_email"` //默认发件人名称
	UseSSL           bool   `json:"use_ssl" yaml:"use_ssl"`                       //是否使用ssl
	UseTls           bool   `json:"use_tls" yaml:"use_tls"`                       //是否使用tls
}
