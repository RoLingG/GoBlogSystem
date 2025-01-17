package config

import (
	"fmt"
)

type ES struct {
	Host                string `yaml:"host"`
	Port                int    `yaml:"port"`
	ArticleIndex        string `yaml:"article_Index"`
	FullTextSearchIndex string `yaml:"full_text_search_index"`
	User                string `yaml:"user"`
	Password            string `yaml:"password"`
}

func (es ES) ConnectUrl() string {
	return fmt.Sprintf("http://%s:%d", es.Host, es.Port)
}
