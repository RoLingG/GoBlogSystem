package user_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (UserApi) UserRemove(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	//批量删除
	var userList []models.UserModel
	count := global.DB.Find(&userList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("所要删除的用户不存在", c)
		return
	}
	//批量删除用户事务(成功就一起成功，失败就一起失败)
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// TODO:删除用户，消息表，评论表，用户收藏的文章，用户发表的文章
		err = global.DB.Delete(&userList).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("删除用户事务失败", c)
		return
	}
	res.OKWithMsg(fmt.Sprintf("共删除 %d 个用户", count), c)
}
