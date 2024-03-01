package core

import (
	"GoRoLingG/config"
	"GoRoLingG/global"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// InitConfig 读取配置文件
func InitConfig() {
	const ConfigFile = "setting.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}
	log.Println("config yamlFile load Init success.")
	fmt.Println(c)
	//赋值全局变量
	global.Config = c
}
