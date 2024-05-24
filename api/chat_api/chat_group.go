package chat_api

import (
	"GoRoLingG/models/ctype"
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

type ChatUser struct {
	Conn     *websocket.Conn
	NickName string
	Avatar   string
}

//var ConnGroupMap = map[string]ChatUser{}

const (
	TextMsg    ctype.MsgType = 1
	ImageMsg   ctype.MsgType = 2
	SystemMsg  ctype.MsgType = 3
	InRoomMsg  ctype.MsgType = 4
	OutRoomMsg ctype.MsgType = 5
)

//使用预先生成好的头像和名字
//type GroupRequest struct {
//	Msg      string  `json:"msg"`       //聊天的内容，暂不支持发送图片
//	MsgType  MsgType `json:"msg_type"`  //聊天类型
//}
//type GroupResponse struct {
//	NickName string  `json:"nick_name"` //前端生成
//	Avatar   string  `json:"avatar"`    //头像
//	Msg      string  `json:"msg"`       //聊天的内容，暂不支持发送图片
//	MsgType  MsgType `json:"msg_type"`  //聊天类型
//	Date time.Time `json:"date"` //消息的时间
//}

// GroupRequest 入参
type GroupRequest struct {
	NickName string        `json:"nick_name"` //前端生成
	Avatar   string        `json:"avatar"`    //头像
	Msg      string        `json:"msg"`       //聊天的内容，暂不支持发送图片
	MsgType  ctype.MsgType `json:"msg_type"`  //聊天类型
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
	//nickName := randomname.GenerateName()
	//nickNameFirst := string([]rune(nickName)[0])
	//avatar := fmt.Sprintf("upload/chat_avatar/%s.png", nickNameFirst)
	//chatUser := ChatUser{
	//	Conn:     conn,
	//	NickName: nickName,
	//	Avatar:   avatar,
	//}
	//ConnGroupMap[addr] = chatUser
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
		request.Avatar = "upload/avatar/avatar.png"
		err = json.Unmarshal(msg, &request)
		if err != nil {
			//参数绑定失败
			request.MsgType = SystemMsg
			request.Msg = "参数绑定失败"
			SendMsg(addr, GroupResponse{
				GroupRequest: request,
				Date:         time.Now(),
			})
			conn.WriteMessage(websocket.TextMessage, []byte("参数绑定失败"))
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
				request.MsgType = SystemMsg
				request.Msg = "发送消息不能为空"
				SendMsg(addr, GroupResponse{
					GroupRequest: request,
					Date:         time.Now(),
				})
				continue
			}
			request.MsgType = TextMsg
			SendGroupMsg(GroupResponse{
				GroupRequest: request,
				Date:         time.Now(),
			})
		case InRoomMsg:
			request.MsgType = InRoomMsg
			request.Msg = request.NickName + " 进入聊天室"
			SendGroupMsg(GroupResponse{
				GroupRequest: request,
				Date:         time.Now(),
			})
		default:
			request.MsgType = SystemMsg
			request.Msg = "消息类型未知"
			SendMsg(addr, GroupResponse{
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

// SendMsg 给某个用户发消息
func SendMsg(addr string, response GroupResponse) {
	byteData, _ := json.Marshal(response)
	chatUser := ConnGroupMap[addr]
	chatUser.WriteMessage(websocket.TextMessage, byteData)
}
