package config

type config struct {
	Mysql  Mysql  `yaml:"mysql"`
	Logger logger `yaml:"logger"`
	system system `yaml:"system"`
}
