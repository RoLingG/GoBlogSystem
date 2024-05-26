package user_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models/ctype"
	"GoRoLingG/res"
	"GoRoLingG/service/user_service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	NickName string     `json:"nick_name" binding:"required" msg:"请输入昵称"`  // 昵称
	UserName string     `json:"user_name" binding:"required" msg:"请输入用户名"` // 用户名
	Password string     `json:"password" binding:"required" msg:"请输入密码"`   // 密码
	Role     ctype.Role `json:"role" binding:"required" msg:"请选择权限"`       // 权限  1 管理员  2 普通用户  3 游客
}

// UserCreateView 添加用户
// @Tags 用户管理
// @Summary 添加用户
// @Description	添加新用户
// @Param token header string true "Authorization token"
// @Param data body UserCreateRequest true	"添加新用户的一些参数"
// @Produce json
// @Router /api/userCreate [post]
// @Success 200 {object} res.Response{}
func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	err := user_service.UserService{}.CreateUser(cr.UserName, cr.NickName, cr.Password, cr.Role, "", c.ClientIP())
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.OKWithMsg(fmt.Sprintf("用户 %s 创建成功!", cr.UserName), c)
	return
}
