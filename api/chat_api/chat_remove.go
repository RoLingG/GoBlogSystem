package chat_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

// ChatRemoveView 删除群聊的聊天记录
// @Tags 聊天管理
// @Summary 删除群聊的聊天记录
// @Description 删除群聊的聊天记录
// @Param data body models.RemoveRequest   false  "表示多个参数"
// @Router /api/chatRemove [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (ChatApi) ChatRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	var chatList []models.ChatModel
	global.DB.Find(&chatList, cr.IDList)
	if len(chatList) > 0 {
		err = global.DB.Delete(&chatList).Error
		if err != nil {
			res.FailWithMsg("群聊记录删除失败", c)
			return
		}
	}
	res.OKWithMsg(fmt.Sprintf("删除群聊聊天记录成功，共删除 %d 条聊天记录", len(chatList)), c)
}
