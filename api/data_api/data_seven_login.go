package data_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
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

// SevenLogin 七日内登录/注册数据
// @Tags 数据收集管理
// @Summary 七日内登录/注册数据
// @Description	七日内登录/注册数据
// @Router /api/dateLogin [get]
// @Produce json
// @Success 200 {object} res.Response{data=DateCountResponse}
func (DataApi) SevenLogin(c *gin.Context) {
	var loginDateCount, signDateCount []DateCount

	var loginDateCountMap = map[string]int{}
	var signDateCountMap = map[string]int{}
	var loginCountList, signCountList []int
	now := time.Now()

	global.DB.Model(models.LoginDataModel{}).
		Where("date_sub(curdate(), interval 7 day) <= create_at").
		Select("date_format(create_at, '%Y-%m-%d') as date", "count(id) as count").
		Group("date").
		Scan(&loginDateCount)
	global.DB.Model(models.UserModel{}).
		Where("date_sub(curdate(), interval 7 day) <= create_at").
		Select("date_format(create_at, '%Y-%m-%d') as date", "count(id) as count").
		Group("date").
		Scan(&signDateCount)

	//因为loginDateCount里面的数据都是[{日期 计数}]，所以将loginDateCount内的数据存进map里，方便操作
	for _, count := range loginDateCount {
		loginDateCountMap[count.Date] = count.Count
	}
	//同上
	for _, count := range signDateCount {
		signDateCountMap[count.Date] = count.Count
	}

	var dateList []string
	//当天之前的七天之内的数据统计存储
	for i := -6; i <= 0; i++ {
		day := now.AddDate(0, 0, i).Format("2006-01-02") //获取当天的日期
		loginCount := loginDateCountMap[day]             //获取当天登录计数
		signCount := signDateCountMap[day]               //获取当天注册计数
		dateList = append(dateList, day)
		loginCountList = append(loginCountList, loginCount)
		signCountList = append(signCountList, signCount)
	}

	res.OKWithData(DateCountResponse{
		DateList:  dateList,
		LoginData: loginCountList,
		SignData:  signCountList,
	}, c)

}
