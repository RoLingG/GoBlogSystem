package news_api

import (
	"GoRoLingG/res"
	"GoRoLingG/utils/request"
	"encoding/json"
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

// NewsData这个类型对应的是NewsResponse的data，var response NewsResponse将解除json化的httpresponse的正文接手以后，里面就变成了
//{200 [{1 那英别说了 汪苏泷快碎了 128.4万 https://s.weibo.com/weibo?q=那英别说了 汪苏泷快碎了}] 请求成功}的[{1 那英别说了 汪苏泷快碎了 128.4万 https://s.weibo.com/weibo?q=那英别说了 汪苏泷快碎了}]这部分就是正文，也就是Newdata结构类型对应的
//这样做的目的就是为了简化NewsResponse里的代码

type NewsData struct {
	Index    int    `json:"index"`
	Title    string `json:"title"`
	HotValue string `json:"hotValue"`
	Link     string `json:"link"`
}

type NewsResponse struct {
	Code int        `json:"code"`
	Data []NewsData `json:"data"`
	Msg  string     `json:"msg"`
}

const newsAPI = "https://api.codelife.cc/api/top/list"
const timeout = 2 * time.Second

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

	//将header的数据map化
	headersMap := structs.Map(headers)
	//执行post请求，获取post过来的数据
	httpResponse, err := request.Post(newsAPI, cr, headersMap, timeout)
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}

	//实例化返回来的参数的对应类型NewsResponse
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
	return
}
