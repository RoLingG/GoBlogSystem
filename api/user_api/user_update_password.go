package user_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/utils/jwt"
	"GoRoLingG/utils/pwd"
	"github.com/gin-gonic/gin"
)

type UpdatePasswordRequest struct {
	OldPwd string `json:"old_pwd" binding:"required" msg:"请输入旧密码"`
	NewPwd string `json:"new_pwd" binding:"required" msg:"请输入新密码"`
}

// UserUpdatePasswordView 用户密码修改
// @Tags 用户管理
// @Summary 用户密码修改
// @Description	用户密码修改
// @Param token header string true "Authorization token"
// @Param data body UpdatePasswordRequest true	"添加用户修改的一些参数"
// @Produce json
// @Router /api/userUpdatePassword [put]
// @Success 200 {object} res.Response{}
func (UserApi) UserUpdatePasswordView(c *gin.Context) {
	var cr UpdatePasswordRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")         //从jwt.auth中获取claims
	claims := _claims.(*jwt.CustomClaims) //断言
	var userModel models.UserModel
	err = global.DB.Take(&userModel, claims.UserID).Error
	if err != nil {
		res.FailWithMsg("对应id的用户不存在", c)
		return
	}
	//判断密码是否一致
	if !pwd.CheckPwd(userModel.Password, cr.OldPwd) {
		res.FailWithMsg("密码错误", c)
		return
	}
	hashPwd := pwd.HashPwd(cr.NewPwd)
	err = global.DB.Model(&userModel).Update("password", hashPwd).Error
	if err != nil {
		res.FailWithMsg("密码修改失败", c)
		return
	}
	res.OKWithMsg("密码修改成功", c)
	return
}
