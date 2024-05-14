package user_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/utils/jwt"
	"GoRoLingG/utils/pwd"
	"github.com/gin-gonic/gin"
)

type EmailLoginRequest struct {
	UserName string `json:"user_name" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入用户密码"`
}

func (UserApi) EmailLoginView(c *gin.Context) {
	//通用参数绑定格式
	var cr EmailLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ? or email = ?", cr.UserName, cr.UserName).Error
	if err != nil {
		//没找到
		global.Log.Warn("用户名不存在")
		res.FailWithMsg("用户名或密码错误", c)
		return
	}
	//校验密码
	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	//如果密码正确isCheck就为true
	if !isCheck {
		global.Log.Warn("用户密码错误")
		res.FailWithMsg("用户名或密码错误", c)
		return
	}
	//登录成功，生成Token
	token, tokenErr := jwt.GenToken(jwt.JwtPayLoad{
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
	})
	if tokenErr != nil {
		global.Log.Error(tokenErr)
		res.FailWithMsg("生成Token失败", c)
		return
	}
	res.OKWithData(token, c)
}
