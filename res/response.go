package res

import (
	"GoRoLingG/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

func (r Response) Json() string {
	byteData, _ := json.Marshal(r)
	return string(byteData)
}

// ListResponse 这里list用泛型T作为类型是因为list会被高频使用，每次使用的时候list的类型都是不确定的。
// ListResponse 如果不使用泛型作为类型，那每次都要定义list是什么结构体类型，达不到封装的效果
type ListResponse[T any] struct {
	Count int64 `json:"count"`
	List  T     `json:"list"`
}

const (
	Success = 0
	Error   = 7
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func OK(data any, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}

func OKWithData(data any, c *gin.Context) {
	Result(Success, data, "操作成功", c)
}

func OKWithDataSSE(data any, c *gin.Context) {
	content := Response{
		Code: Success,
		Data: data,
		Msg:  "成功",
	}.Json()
	c.SSEvent("", content)
}

func OKWithDataAndMsgSSE(data any, msg string, c *gin.Context) {
	content := Response{
		Code: Success,
		Data: data,
		Msg:  msg,
	}.Json()
	c.SSEvent("", content)
}

func OKWithMsg(msg string, c *gin.Context) {
	Result(Success, map[string]any{}, msg, c)
}

func OKWithDataAndMsg(data any, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}

func OKWithoutData(c *gin.Context) {
	Result(Success, map[string]any{}, "操作成功", c)
}

func OKWithList(list any, count int64, c *gin.Context) {
	OKWithData(ListResponse[any]{
		count,
		list,
	}, c)
}

func Fail(data any, msg string, c *gin.Context) {
	Result(Error, data, msg, c)
}

func FailWithMsg(msg string, c *gin.Context) {
	Result(Error, map[string]any{}, msg, c)
}

func FailWithMsgSSE(msg string, c *gin.Context) {
	data := Response{
		Code: Error,
		Data: map[string]any{},
		Msg:  msg,
	}.Json()
	c.SSEvent("", data)
}

func FailWithError(err error, obj any, c *gin.Context) {
	msg := utils.GetValidMsg(err, obj)
	FailWithMsg(msg, c)
}

func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := ErrorMap[code]
	if ok {
		Result(int(code), map[string]any{}, msg, c)
	}
	Result(Error, map[string]any{}, "未知错误", c)
}
