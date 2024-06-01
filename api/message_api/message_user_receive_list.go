package message_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
)

type MessageUserListRequest struct {
	models.PageInfo
	NickName string `json:"nick_name" form:"nick_name"`
}

type MessageUserListResponse struct {
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	UserID   uint   `json:"user_id"`
	Avatar   string `json:"avatar"`
	Count    int    `json:"count"`
}

// MessageUserReceiveList 有消息的用户列表
// @Tags 消息管理
// @Summary 有消息的用户列表
// @Description 有消息的用户列表
// @Router /api/receiveNewMessageList [get]
// @Param token header string  true  "token"
// @Param data query MessageUserListRequest   false  "查询参数"
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[MessageUserListResponse]}
func (MessageApi) MessageUserReceiveList(c *gin.Context) {
	var cr MessageUserListRequest
	c.ShouldBindQuery(&cr)

	var count int64
	global.DB.Model(models.MessageModel{}).Where(models.MessageModel{SendUserNickName: cr.NickName}).Group("send_user_id").Count(&count)

	type resType struct {
		SendUserID uint
		Count      int // 发送人的个数2
	}
	offset := (cr.Page - 1) * cr.Limit

	var _list []resType
	global.DB.Model(models.MessageModel{}).Where(models.MessageModel{SendUserNickName: cr.NickName}).
		Group("send_user_id").Limit(cr.Limit).Offset(offset).Select("send_user_id", "count(distinct rev_user_id) as count").Scan(&_list)

	var userMessageMap = map[uint]int{}

	for _, r := range _list {
		userMessageMap[r.SendUserID] = r.Count //userMessageMap存储每个发送者的用户id及他们发送消息的数量
	}
	var userIDList []uint
	for uid, _ := range userMessageMap {
		userIDList = append(userIDList, uid) //获取所有发送人的id
	}
	var userList []models.UserModel
	global.DB.Find(&userList, userIDList) //将userIDList对应的所有用户数据获取出来
	var userMap = map[uint]models.UserModel{}
	for _, user := range userList {
		userMap[user.ID] = user //userMap根据用户的id去对应用户的数据
	}

	var list = make([]MessageUserListResponse, 0)
	for uid, count := range userMessageMap {
		user := userMap[uid]
		list = append(list, MessageUserListResponse{
			UserName: user.UserName,
			NickName: user.NickName,
			UserID:   user.ID,
			Avatar:   user.Avatar,
			Count:    count,
		})
	}

	res.OKWithList(list, count, c)
}
