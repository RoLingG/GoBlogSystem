package main

import (
	"GoRoLingG/core"
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"time"
)

var client *elastic.Client

type DemoModel struct {
	ID       string `json:"id"` //一般来说不传这个，可以忽略掉
	Title    string `json:"title"`
	UserID   uint   `json:"user_id"`
	CreateAt string `json:"create_at"`
}

// Index 索引名称
func (DemoModel) Index() string {
	return "demo_index"
}

// EsConnect es连接
func EsConnect() *elastic.Client {
	var err error
	sniffOpt := elastic.SetSniff(false)
	host := "http://192.168.31.201:9200"
	c, err := elastic.NewClient(
		elastic.SetURL(host),
		sniffOpt,
		elastic.SetBasicAuth("", ""),
	)
	if err != nil {
		logrus.Fatalf("es连接失败 %s", err.Error())
	} else {
		logrus.Info("es连接成功")
	}
	return c
}

// FindList 列表查询
func FindList(key string, page, limit int) (demoList []DemoModel, count int) {
	boolSearch := elastic.NewBoolQuery() //实例化查询条件
	from := page
	if key != "" {
		boolSearch.Must(
			elastic.NewMatchQuery("title", key), //对title进行匹配搜索
		)
	}
	if limit == 0 {
		limit = 10
	}
	if from == 0 {
		from = 1
	}

	res, err := client.
		Search(DemoModel{}.Index()).
		Query(boolSearch). //查询条件
		From((from - 1) * limit).
		Size(limit).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	count = int(res.Hits.TotalHits.Value) //搜索到结果总条数
	for _, hit := range res.Hits.Hits {
		var demo DemoModel
		/*
			例如：
			...
			"_source": {
				"id": "",
				"title": "测试",
				"user_id": 6,
				"create_at": "2024-05-17 11:01:03"
			}
			...
		*/
		data, err := hit.Source.MarshalJSON() //从该条结果的Source获取数据，并将其json化
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		err = json.Unmarshal(data, &demo) //将json划的数据解码进对应的结构体里
		if err != nil {
			logrus.Error(err)
			continue
		}
		demo.ID = hit.Id
		demoList = append(demoList, demo)
	}
	return demoList, count
}

// FindSourceList 返回指定的字段
func FindSourceList(key string, page, limit int) (demoList []DemoModel, count int) {
	boolSearch := elastic.NewBoolQuery()
	from := page
	if key != "" {
		boolSearch.Must(
			elastic.NewMatchQuery("title", key),
		)
	}
	if limit == 0 {
		limit = 10
	}
	if from == 0 {
		from = 1
	}

	res, err := client.
		Search(DemoModel{}.Index()).
		Query(boolSearch).
		Source(`{"_source": ["title"]}`). //相比于简单的关键字列表搜索，这里
		From((from - 1) * limit).
		Size(limit).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	count = int(res.Hits.TotalHits.Value) //搜索到结果总条数
	for _, hit := range res.Hits.Hits {
		var demo DemoModel
		data, err := hit.Source.MarshalJSON()
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		err = json.Unmarshal(data, &demo)
		if err != nil {
			logrus.Error(err)
			continue
		}
		demo.ID = hit.Id
		demoList = append(demoList, demo)
	}
	return demoList, count
}

// Update 更新
func Update(id string, data *DemoModel) error {
	_, err := client.
		Update().
		Index(DemoModel{}.Index()).
		Id(id).
		Doc(map[string]string{
			"title": data.Title,
		}).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	logrus.Info("更新demo成功")
	return nil
}

// Remove 批量删除
func Remove(idList []string) (count int, err error) {
	bulkService := client.Bulk().Index(DemoModel{}.Index()).Refresh("true") //使用桶，用桶装删除请求，批量放入以达到批量删除的效果
	for _, id := range idList {
		request := elastic.NewBulkDeleteRequest().Id(id) //对应id的删除请求
		bulkService.Add(request)                         //装入桶中
	}
	res, err := bulkService.Do(context.Background())
	return len(res.Succeeded()), err
}

// Create 索引创建
func Create(data *DemoModel) (err error) {
	indexResponse, err := client.Index().Index(data.Index()).BodyJson(data).Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	logrus.Infof("%#v", indexResponse)
	data.ID = indexResponse.Id //因为data是指针类型，所以会修改数据
	return nil
}

// es实现
func init() {
	core.InitConfig()
	core.InitLogger()
	client = EsConnect()
}

func main() {
	//DemoModel{}.CreateIndex()
	//DemoModel{}.RemoveIndex()

	//索引底下的数据创建
	Create(&DemoModel{
		Title:    "测试",
		UserID:   6,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
	})

	//列表查询
	//list, count := FindList("", 1, 5) //添加之后马上去查是查不到的，要查得到需要去设置一下es的查询开关
	//fmt.Println(list, count)

	//关键字查询
	//list, count := FindSourceList("测试", 1, 5) //暂时有问题
	//fmt.Println(list, count)

	//更新
	//Update("UjF-hI8B4tZsh_kv9S8w", &DemoModel{Title: "更新测试"}) //这里传的id是es索引的的_id

	//批量删除
	//count, err := Remove([]string{"UzGAhI8B4tZsh_kv6i_k", "UjF-hI8B4tZsh_kv9S8w"})
	//fmt.Println(count, err)
}
