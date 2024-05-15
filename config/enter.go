package config

type Config struct {
	Mysql        Mysql        `yaml:"mysql"`
	Redis        Redis        `yaml:"redis"`
	Logger       Logger       `yaml:"logger"`
	System       System       `yaml:"system"`
	SiteInfo     SiteInfo     `yaml:"site_info"`
	JWT          JWT          `yaml:"JWT"`
	QiNiu        QiNiu        `yaml:"QiNiu"`
	Email        Email        `yaml:"Email"`
	QQ           QQ           `yaml:"QQ"`
	ImagesUpload ImagesUpload `yaml:"images_upload"`
}
