package chat_api

import (
	"GoRoLingG/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

//不校验用户是否登录，大家都可以聊天的聊天室

// ConnGroupMap map[连接ip]连接对象
var ConnGroupMap = map[string]*websocket.Conn{}

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
			break
		}
		SendGroupMsg(string(msg))
		// 发送消息
		//conn.WriteMessage(websocket.TextMessage, []byte("xxx")) //发送消息后对方回复"xxx"
	}
	defer conn.Close()
	delete(ConnGroupMap, addr)
}

// SendGroupMsg 群聊功能
func SendGroupMsg(text string) {
	//用for循环轮询连接对象，给所有连接对象都发送某个用户发送的消息，以达到全局群聊的效果
	for _, conn := range ConnGroupMap {
		conn.WriteMessage(websocket.TextMessage, []byte(text))
	}
}
