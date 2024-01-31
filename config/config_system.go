package config

type system struct {
	host string `yaml:"host"`
	port int    `yaml:"port"`
	env  string `yaml:"env"`
}
