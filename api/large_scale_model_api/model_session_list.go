package large_scale_model_api

import (
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/service/common"
	"github.com/gin-gonic/gin"
)

type SessionListRequest struct {
	models.PageInfo
}

type SessionListResponse struct {
	models.Model
	UserID      uint   `json:"user_id"`      // 用户id
	NickName    string `json:"nick_name"`    // 用户名称
	SessionName string `json:"session_name"` // 会话名称
	RoleName    string `json:"role_name"`    // 角色名称
	ChatCount   int    `json:"chat_count"`   // 对话次数
	LastContent string `json:"last_content"` //最后一次对话内容
}

func (LargeScaleModelApi) ModelSessionListView(c *gin.Context) {
	var cr SessionListRequest
	c.ShouldBindQuery(&cr)
	_list, count, _ := common.CommonList(models.LargeScaleModelSessionModel{}, common.Option{
		Preload: []string{"UserModel", "RoleModel", "ChatList"},
	})

	var list = make([]SessionListResponse, 0)
	for _, item := range _list {
		var lastContent string
		if len(item.ChatList) > 0 {
			lastContent = item.ChatList[len(item.ChatList)-1].UserContent
		}
		list = append(list, SessionListResponse{
			Model:       item.Model,
			UserID:      item.UserID,
			NickName:    item.UserModel.NickName,
			SessionName: item.SessionName,
			RoleName:    item.RoleModel.Name,
			ChatCount:   len(item.ChatList),
			LastContent: lastContent,
		})
	}
	res.OKWithList(list, count, c)
}
