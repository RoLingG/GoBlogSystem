package request

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Post(url string, data any, headers map[string]interface{}, timeout time.Duration) (response *http.Response, err error) {
	reqParam, _ := json.Marshal(data) //json化数据
	reqBody := strings.NewReader(string(reqParam))
	//设置请求类型
	httpReq, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return
	}
	//设置header
	httpReq.Header.Add("Content-Type", "application/json")
	for key, val := range headers {
		switch v := val.(type) { //类型推断，headers里面一共也就两个数据类型
		case string: //设置header：Signaturekey
			httpReq.Header.Add(key, v)
		case int: //设置header：Version，并将其从string类型转换成int类型
			httpReq.Header.Add(key, strconv.Itoa(v))
		}
	}
	//创建http客户端，用于后面post请求操作
	client := http.Client{
		Timeout: timeout, //设置超时检测时间
	}
	//执行post请求操作，并实例化httpResponse存储post后传输过来的数据
	httpResponse, err := client.Do(httpReq)
	//返回post过去传回来的数据
	return httpResponse, err
}
