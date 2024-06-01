package main

import (
	"GoRoLingG/core"
	_ "GoRoLingG/docs"
	"GoRoLingG/flag"
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/routers"
	"GoRoLingG/service/redis_service"
	"GoRoLingG/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"strings"
)

type Job struct {
}

func (Job) Run() {
	result, err := global.ESClient.Search(models.ArticleModel{}.Index()).Query(elastic.NewMatchAllQuery()).Size(10000).Do(context.Background())
	if err != nil {
		logrus.Error(err)
		return
	}
	diggInfo := redis_service.NewArticleDiggIndex().GetInfo()
	lookInfo := redis_service.NewArticleLookIndex().GetInfo()
	commentInfo := redis_service.NewArticleCommentIndex().GetInfo()
	for _, hit := range result.Hits.Hits {
		var article models.ArticleModel
		err = json.Unmarshal(hit.Source, &article)
		digg := diggInfo[hit.Id]
		look := lookInfo[hit.Id]
		comment := commentInfo[hit.Id]

		newDigg := article.DiggCount + digg
		newLook := article.LookCount + look
		newComment := article.CommentCount + comment
		if newDigg == article.DiggCount && newLook == article.LookCount && newComment == article.CommentCount {
			logrus.Info(article.Title, " 点赞数、浏览数和评论数无变化")
			continue
		}
		_, err = global.ESClient.Update().Index(models.ArticleModel{}.Index()).Id(hit.Id).Doc(map[string]int{
			"digg_count":    newDigg,
			"look_count":    newLook,
			"comment_count": newComment,
		}).Do(context.Background())
		if err != nil {
			logrus.Error(err)
			continue
		}
		logrus.Infof("%s, 点赞数据和浏览数据同步成功，点赞数为 %d; 浏览数为 %d; 评论数为 %d", article.Title, newDigg, newLook, newComment)
	}
	//数据同步成功后将redis内的缓存数据清空，重新计数
	redis_service.NewArticleDiggIndex().Clear()
	redis_service.NewArticleLookIndex().Clear()
	redis_service.NewArticleCommentIndex().Clear()
	fmt.Println("定时任务完成！")
}

// @title GoRoLingG API文档
// @version	1.0
// @description GoRoLingG API文档
// @host 127.0.0.01:8080
// @BasePath /
func main() {
	//swagger文档host
	//docs.SwaggerInfo.Host = ""

	//读取配置文件，main中调用InitConfig
	core.InitConfig()
	//连接IP地址数据库
	core.InitAddrDB()
	defer global.AddrDB.Close() //主程序结束，则关闭IP地址数据库

	//初始化日志
	global.Log = core.InitLogger()

	//连接数据库
	global.DB = core.InitGorm()
	//连接Redis
	global.Redis = core.ConnectRedis()
	//连接ES
	global.ESClient = core.ConnectES()

	//定时任务
	Cron := cron.New(cron.WithSeconds())
	Cron.AddJob("0 */10 * * * *", Job{}) //每1分钟同步一次Redis缓存中的文章点赞数和文章浏览数
	Cron.Start()

	//flag迁移数据库肯定是在连接数据库之后，路由连接之前
	//命令行参数绑定
	option := flag.Parse()
	//如果需要web项目停止运行，则后面的操作都不能执行，立刻停止web项目
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	//路由连接
	router := routers.InitRouter()
	host := global.Config.System.Host
	port := global.Config.System.Port
	addr := global.Config.System.Addr()

	if host == "0.0.0.0" {
		ipList := utils.GetIPList()
		for _, ip := range ipList {
			parts := strings.Split(ip, ".")
			if parts[0] == "169" { //过滤掉访问不了的内网ip
				continue
			}
			global.Log.Infof("gvb_Server 运行在: http://%s:%d/api", ip, port) //传输路由连接log
			global.Log.Infof("gvb_Server api文档 运行在: http://%s:%d/swagger/index.html#", ip, port)
		}
	} else {
		global.Log.Infof("gvb_Server 运行在: http://%s:%d/api", host, port) //传输路由连接log
		global.Log.Infof("gvb_Server api文档 运行在: http://%s:%d/swagger/index.html#", host, port)
	}
	err := router.Run(addr) //将路由运行到该地址下
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
