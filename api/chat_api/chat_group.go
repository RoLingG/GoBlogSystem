package chat_api

import (
	"GoRoLingG/res"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
	"time"
)

//不校验用户是否登录，大家都可以聊天的聊天室

// ConnGroupMap map[连接ip]连接对象
var ConnGroupMap = map[string]*websocket.Conn{}

type MsgType int

const (
	TextMsg    MsgType = 1
	ImageMsg   MsgType = 2
	SystemMsg  MsgType = 3
	InRoomMsg  MsgType = 4
	OutRoomMsg MsgType = 5
)

// GroupRequest 入参
type GroupRequest struct {
	NickName string  `json:"nick_name"` //前端生成
	Avatar   string  `json:"avatar"`    //头像
	Msg      string  `json:"msg"`       //聊天的内容，暂不支持发送图片
	MsgType  MsgType `json:"msg_type"`  //聊天类型
}

// GroupResponse 出参
type GroupResponse struct {
	GroupRequest
	Date time.Time `json:"date"` //消息的时间
}

// ChatGroupView 聊天的基本框架
func (ChatApi) ChatGroupView(c *gin.Context) {
	var upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// 鉴权 true表示放行，false表示拦截
			return true
		},
	}
	// 将http升级至websocket
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	fmt.Println(err)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	addr := conn.RemoteAddr().String()
	ConnGroupMap[addr] = conn
	for {
		// 消息类型，消息，错误
		_, msg, err := conn.ReadMessage()
		if err != nil {
			// 用户断开聊天
			SendGroupMsg(GroupResponse{
				GroupRequest: GroupRequest{Msg: addr + " 离开聊天室"},
				Date:         time.Now(),
			})
			break
		}
		//进行参数绑定
		var request GroupRequest
		err = json.Unmarshal(msg, &request)
		if err != nil {
			//参数绑定失败
			conn.WriteMessage(websocket.TextMessage, []byte("xxx"))
			continue
		}
		//用户信息不能为空
		if strings.TrimSpace(request.Avatar) == "" || strings.TrimSpace(request.NickName) == "" {
			continue
		}
		//判断前端传过来的消息类型
		switch request.MsgType {
		case TextMsg:
			//用户发送消息不能为空
			if strings.TrimSpace(request.Msg) == "" {
				continue
			}
			SendGroupMsg(GroupResponse{
				GroupRequest: request,
				Date:         time.Now(),
			})
		case InRoomMsg:
			request.Msg = request.NickName + " 进入聊天室"
			SendGroupMsg(GroupResponse{
				GroupRequest: request,
				Date:         time.Now(),
			})
		}
		// 发送消息
		//conn.WriteMessage(websocket.TextMessage, []byte("xxx")) //发送消息后对方回复"xxx"
	}
	defer conn.Close()
	delete(ConnGroupMap, addr)
}

// SendGroupMsg 群聊功能
func SendGroupMsg(response GroupResponse) {
	byteData, _ := json.Marshal(response)
	//用for循环轮询连接对象，给所有连接对象都发送某个用户发送的消息，以达到全局群聊的效果
	for _, conn := range ConnGroupMap {
		conn.WriteMessage(websocket.TextMessage, byteData)
	}
}
