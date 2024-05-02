package core

import (
	"GoRoLingG/config"
	"GoRoLingG/global"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/fs"
	"io/ioutil"
	"log"
)

const ConfigFile = "setting.yaml"

// InitConfig 读取配置文件
func InitConfig() {
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

// 实现修改yaml文件
func SetYaml() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(ConfigFile, byteData, fs.ModePerm)
	if err != nil {
		return err
	}
	global.Log.Info("配置文件修改成功")
	return nil
}
