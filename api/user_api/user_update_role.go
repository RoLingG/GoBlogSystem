package user_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/models/ctype"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
)

type UserRole struct {
	Role   ctype.Role `json:"role" binding:"required,oneof=1 2 3 4" msg:"权限不足，操作失败"`
	UserID uint       `json:"user_id" binding:"required" msg:"用户id错误"`
}

// UserUpdateRole 用户权限变更
func (UserApi) UserUpdateRoleView(c *gin.Context) {
	var cr UserRole
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var userModel models.UserModel
	err = global.DB.Take(&userModel, cr.UserID).Error
	if err != nil {
		res.FailWithMsg("对应ID的用户不存在", c)
		return
	}
	err = global.DB.Model(&userModel).Update("role", cr.Role).Error
	if err != nil {
		res.FailWithMsg("权限修改失败", c)
		return
	}
	res.OKWithMsg("权限修改成功", c)
}
