package config

import "fmt"

type QQ struct {
	AppID    string `json:"app_id" yaml:"app_id"`
	Key      string `json:"key" yaml:"key"`
	Redirect string `json:"redirect" yaml:"redirect"`
}

func (q QQ) GetPath() string {
	if q.Key == "" || q.AppID == "" || q.Redirect == "" {
		return ""
	}
	return fmt.Sprintf("https://qraph.qq.com/oauth2.0/show?which=Login&display=pc&respone_type=code&client_id=%s&redirect_uri=%s", q.AppID, q.Redirect)
}
