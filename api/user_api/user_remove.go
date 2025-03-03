package user_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/plugins/log_stash_v1"
	"GoRoLingG/res"
	"GoRoLingG/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserRemove 用户删除(管理员)
// @Tags 用户管理
// @Summary 管理员删除用户
// @Description	管理员删除用户
// @Param token header string true "Authorization token"
// @Param data body models.RemoveRequest true	"用户删除的一些参数"
// @Produce json
// @Router /api/userRemove [delete]
// @Success 200 {object} res.Response{}
func (UserApi) UserRemove(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}

	token := c.GetHeader("token")
	log := log_stash_v1.NewLogByGin(c)
	ip, _ := utils.GetAddrByGin(c)
	log = log_stash_v1.New(ip, token)

	//批量删除
	var userList []models.UserModel
	count := global.DB.Find(&userList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("所要删除的用户不存在", c)
		log.Error("删除的用户不存在")
		return
	}

	//批量删除用户事务(成功就一起成功，失败就一起失败)
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// TODO:删除用户，消息表，评论表，用户收藏的文章，用户发表的文章
		// 删除相关的评论
		err = global.DB.Where("user_id IN (?)", cr.IDList).Delete(&models.CommentModel{}).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		// 删除用户收藏的文章
		err = global.DB.Where("user_id IN (?)", cr.IDList).Delete(&models.UserCollectModel{}).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		// 删除登录数据
		err = global.DB.Where("user_id IN (?)", cr.IDList).Delete(&models.LoginDataModel{}).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		// 删除用户
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
		log.Error("用户删除失败")
		return
	}
	res.OKWithMsg(fmt.Sprintf("共删除 %d 个用户", count), c)
	log.Info("用户删除成功")
}
