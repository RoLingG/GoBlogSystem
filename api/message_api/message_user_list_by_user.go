package message_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
)

type MessageUserListByUserRequest struct {
	models.PageInfo
	UserID uint `json:"user_id" form:"user_id" binding:"required"`
}

// MessageUserListByUser 某个用户的聊天列表
// @Tags 消息管理
// @Summary 某个用户的聊天列表
// @Description 某个用户的聊天列表
// @Router /api/messageUserListByUser [get]
// @Param token header string  true  "Token"
// @Param data query MessageUserListByUserRequest   false  "查询参数"
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[MessageUserListResponse]}
func (MessageApi) MessageUserListByUser(c *gin.Context) {
	var cr MessageUserListByUserRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithMsg("参数错误", c)
		return
	}

	if cr.Limit == 0 {
		cr.Limit = 10
	}

	offset := (cr.Page - 1) * cr.Limit

	type resType struct {
		SendUserID uint
		RevUserID  uint
		Count      int
	}

	var _list []resType
	global.DB.Model(models.MessageModel{}).Where("send_user_id = ? or rev_user_id = ?", cr.UserID, cr.UserID).
		Group("send_user_id").
		Group("rev_user_id").Limit(cr.Limit).Offset(offset).Select("send_user_id", "rev_user_id", "count(id) as count").Scan(&_list)

	var userMessageMap = map[uint]int{}

	for _, r := range _list {
		sendVal, ok1 := userMessageMap[r.SendUserID]
		if !ok1 && cr.UserID != r.SendUserID {
			userMessageMap[r.SendUserID] = r.Count
		} else {
			if cr.UserID != r.SendUserID {
				userMessageMap[r.SendUserID] = r.Count + sendVal
			}
		}
		revVal, ok2 := userMessageMap[r.RevUserID]
		if !ok2 && cr.UserID != r.RevUserID {
			userMessageMap[r.RevUserID] = r.Count
		} else {
			if cr.UserID != r.RevUserID {
				userMessageMap[r.RevUserID] = r.Count + revVal
			}
		}
	}
	var userIDList []uint
	for uid, _ := range userMessageMap {
		userIDList = append(userIDList, uid)
	}
	var userList []models.UserModel
	global.DB.Find(&userList, userIDList)
	var userMap = map[uint]models.UserModel{}
	for _, model := range userList {
		userMap[model.ID] = model
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

	res.OKWithList(list, int64(len(list)), c)
}
