package core

import (
	"GoRoLingG/global"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

// ConnectES  es连接
func ConnectES() *elastic.Client {
	var err error
	sniffOpt := elastic.SetSniff(false)
	c, err := elastic.NewClient(
		elastic.SetURL(global.Config.ES.ConnectUrl()),
		sniffOpt,
		elastic.SetBasicAuth(global.Config.ES.User, global.Config.ES.Password),
	)
	if err != nil {
		logrus.Fatalf("es连接失败 %s", err.Error())
	} else {
		logrus.Info("es连接成功")
	}
	return c
}
