package article_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"time"
)

type CalendarResponse struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

// BucketsType 用于处理日历聚合后的输出格式转换
type BucketsType struct {
	Buckets []struct {
		KeyAsString string `json:"key_as_string"`
		Key         int64  `json:"key"`
		DocCount    int    `json:"doc_count"`
	} `json:"buckets"`
}

var DateCount = map[string]int{}

// ArticleCalendarView 文章日历
// @Tags 文章管理
// @Summary 文章日历
// @Description	查看近期文章发布的数量日历
// @Produce json
// @Router /api/articleCalendar [get]
// @Success 200 {object} res.Response{data=CalendarResponse}
func (ArticleApi) ArticleCalendarView(c *gin.Context) {
	//es内数据以时间聚合，也就是以create_at进行聚合
	agg := elastic.NewDateHistogramAggregation().Field("create_at").CalendarInterval("day") //小时发就是hour，分钟就是minute，天就是day

	//时间段搜索
	//从今天开始到去年今天这个时间段
	now := time.Now()
	oneYearAgo := now.AddDate(-1, 0, 0)                                                                //一年前的时间
	format := "2006-01-02 15:04:05"                                                                    //时间格式
	query := elastic.NewRangeQuery("create_at").Gte(oneYearAgo.Format(format)).Lte(now.Format(format)) //lt小于，gt大于

	//查询通用写法
	result, err := global.ESClient.
		Search(models.ArticleModel{}.Index()).
		Query(query).
		Aggregation("calendar", agg).
		Size(0).
		Do(context.Background())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("查询失败", c)
		return
	}

	var data BucketsType
	_ = json.Unmarshal(result.Aggregations["calendar"], &data) //将时间聚合的结果json解码并传给data(data的类型是一个同等于时间聚合json格式的type)
	//fmt.Println(string(result.Aggregations["calendar"]))
	//这个不能直接用，因为格式是{"buckets":[{"key_as_string":"2024-05-17 00:00:00","key":1715904000000,"doc_count":1}]}，但我们可以看出这是一个json嵌套，所以我们自己写一个对应的type去接收就可以获取对应的数据了
	//所以上面就有一个BucketsType结构体的data去接收

	var resultList = make([]CalendarResponse, 0) //实例化，并防止为空时输出null

	for _, bucket := range data.Buckets {
		Time, _ := time.Parse(format, bucket.KeyAsString)      //将时间转换格式，bucket中的KeyAsString就是对应的时间，然后赋值给Time
		DateCount[Time.Format("2006-01-02")] = bucket.DocCount //当天发布的文章数也就是bucket中的DocCount
	}

	days := int(now.Sub(oneYearAgo).Hours() / 24) //获得一年多少天
	//获取去年今天到今天的每一天的时间以及当天发布的文章数
	for i := 0; i <= days; i++ {
		day := oneYearAgo.AddDate(0, 0, i).Format("2006-01-02") //从去年的当前时间开始循环+1天，并且时间格式改成年-月-日

		count, _ := DateCount[day]
		resultList = append(resultList, CalendarResponse{
			Date:  day,
			Count: count,
		})
	}

	res.OKWithData(resultList, c)
}
