package data_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type DateCount struct {
	Date  string `json:"date"`
	Count int    `json:"count"`
}

type DateCountResponse struct {
	DateList  []string `json:"date_list"`
	LoginData []int    `json:"login_data"`
	SignData  []int    `json:"sign_data"`
}

type LoginRequest struct {
	Date int `json:"date" form:"date"` // 1 七天 2 一个月 3 两个月 4 三个月 5 六个月  6 一年
}

// SevenLogin 七日内登录/注册数据
// @Tags 数据收集管理
// @Summary 七日内登录/注册数据
// @Description	七日内登录/注册数据
// @Router /api/dateLogin [get]
// @Produce json
// @Success 200 {object} res.Response{data=DateCountResponse}
func (DataApi) SevenLogin(c *gin.Context) {
	var cr LoginRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithMsg("参数绑定出错", c)
		return
	}

	var dateMap = map[int]int{
		0: 7,
		1: 7,
		2: 30,
		3: 60,
		4: 90,
		5: 180,
		6: 365,
	}

	var response DateCountResponse

	var dateType = dateMap[cr.Date]
	global.DB.Where("").Where(fmt.Sprintf("date_sub(curdate(), interval %d day) <= create_at", dateType))
	//获取时间范围最早的那一天→现在	addDay也就是时间范围最早第一天，例如我要找七天前的，那么addDay就是七天前的第一天
	preDay := time.Now().AddDate(0, 0, -dateType)
	for i := 1; i <= dateType; i++ {
		//获取需要统计数量的时间范围
		response.DateList = append(response.DateList, preDay.AddDate(0, 0, i).Format("2006-01-02"))
	}

	type dateCountType struct {
		Date  string `json:"date"`
		Count int    `json:"count"`
	}

	//统计登录用户数(日为单位)
	var dateLoginCountList []dateCountType
	//以日为单位获取出用户登录数量
	global.DB.Model(models.LoginDataModel{}).Where(global.DB.Where("")).
		Select(
			"date_format(create_at, '%Y-%m-%d') as date",
			"count(id) as count").
		Group("date").Scan(&dateLoginCountList)
	var dateLoginCountMap = map[string]int{}
	for _, countType := range dateLoginCountList {
		//以日为key，以数量为value进行存储
		dateLoginCountMap[countType.Date] = countType.Count
	}
	for _, s := range response.DateList {
		//在时间范围内，获取范围内每天的数据计数
		count, _ := dateLoginCountMap[s]
		response.LoginData = append(response.LoginData, count)
	}

	//统计注册用户数(日为单位)
	var dateSignCountList []dateCountType
	//以日为单位获取出用户注册数量
	global.DB.Model(models.UserModel{}).Where(global.DB.Where("")).
		Select(
			"date_format(create_at, '%Y-%m-%d') as date",
			"count(id) as count").
		Group("date").Scan(&dateSignCountList)
	var dateSignCountMap = map[string]int{}
	for _, countType := range dateSignCountList {
		//以日为key，以数量为value进行存储
		dateSignCountMap[countType.Date] = countType.Count
	}
	for _, s := range response.DateList {
		//在时间范围内，获取范围内每天的数据计数
		count, _ := dateSignCountMap[s]
		//顺序计数
		response.SignData = append(response.SignData, count)
	}

	res.OKWithData(response, c)
}
