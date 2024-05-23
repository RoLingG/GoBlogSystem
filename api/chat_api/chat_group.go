package chat_api

import (
	"GoRoLingG/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

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
	for {
		// 消息类型，消息，错误
		_, msg, err := conn.ReadMessage()
		if err != nil {
			// 用户断开聊天
			break
		}
		fmt.Println(string(msg))
		// 发送消息
		conn.WriteMessage(websocket.TextMessage, []byte("xxx")) //发送消息后对方回复"xxx"
	}
	defer conn.Close()
}
