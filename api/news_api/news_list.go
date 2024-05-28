package news_api

import (
	"GoRoLingG/res"
	"GoRoLingG/service"
	"GoRoLingG/service/redis_service"
	"GoRoLingG/utils/request"
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

type params struct {
	ID   string `json:"id"`
	Size int    `json:"size"`
}

type Header struct {
	Signaturekey string `form:"signaturekey" structs:"signaturekey"`
	Version      string `form:"version" structs:"version"`
	UserAgent    string `form:"User-Agent" structs:"User-Agent"`
}

// NewsData这个类型对应的是NewsResponse的Data，var response NewsResponse将解除json化的httpresponse的正文接手以后，里面就变成了
//{200 [{1 那英别说了 汪苏泷快碎了 128.4万 https://s.weibo.com/weibo?q=那英别说了 汪苏泷快碎了}] 请求成功}的[{1 那英别说了 汪苏泷快碎了 128.4万 https://s.weibo.com/weibo?q=那英别说了 汪苏泷快碎了}]这部分就是正文，也就是Newdata结构类型对应的
//这样做的目的就是为了简化NewsResponse里的代码

type NewsResponse struct {
	Code int                      `json:"code"`
	Data []redis_service.NewsData `json:"data"`
	Msg  string                   `json:"msg"`
}

const newsAPI = "https://api.codelife.cc/api/top/list"
const timeout = 2 * time.Second

// NewsListView 新闻列表
// @Tags 新闻管理
// @Summary 新闻列表
// @Description	查询新闻列表
// @Param Signaturekey header string true "itab新闻接口密钥"
// @Param version header string true "itab新闻接口版本号"
// @Param data body params true	"查询itab新闻的荷载ID和显示新闻多少Size"
// @Router /api/newsList [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[NewsResponse]}
func (NewsApi) NewsListView(c *gin.Context) {
	var cr params      //params是为了给予itab需要的id以及传过来新闻的多少数量设置
	var headers Header //Header是为了设置post过去时要用的header
	err := c.ShouldBindJSON(&cr)
	err = c.ShouldBindHeader(&headers)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	if cr.Size == 0 {
		cr.Size = 1
	}
	//设置redis内缓存存储新闻key名字
	key := fmt.Sprintf("%s-%d", cr.ID, cr.Size)
	//查看redis内是否有对应ID的文章(但是这样有个问题，就是每次查询新闻列表都用的缓存的数据)
	//后话：更新了，RedisService.GetNews()和RedisService.SetNews()内redis的操作从HGet/HSet换成了Get/Set，这样就能够设置其在redis内的过期时间了
	newsData, _ := service.Service.RedisService.GetNews(key)
	if len(newsData) != 0 {
		res.OKWithData(newsData, c)
		return
	}
	//将header的数据map化
	headersMap := structs.Map(headers)
	//执行post请求，获取post过来的数据
	httpResponse, err := request.Post(newsAPI, cr, headersMap, timeout)
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}

	//实例化返回来的参数的对应结构体NewsResponse
	var response NewsResponse
	//获取请求的正文
	byteData, err := io.ReadAll(httpResponse.Body)
	err = json.Unmarshal(byteData, &response)
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	if response.Code != 200 {
		res.FailWithMsg(response.Msg, c)
		return
	}
	res.OKWithData(response.Data, c)
	service.Service.RedisService.SetNews(key, response.Data)
	return
}
