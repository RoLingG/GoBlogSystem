package user_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/models/ctype"
	"GoRoLingG/plugins/log_stash"
	"GoRoLingG/res"
	"GoRoLingG/utils"
	"GoRoLingG/utils/jwt"
	"GoRoLingG/utils/pwd"
	"fmt"
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

	log := log_stash.NewLogByGin(c) //因为这里一开始没有生成token，所以里面New()的token也就是空的，会导致报错，但没事，已经解决了！

	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ? or email = ?", cr.UserName, cr.UserName).Error
	if err != nil {
		//没找到
		global.Log.Warn("用户名不存在")
		log.Warning(fmt.Sprintf("%s 用户名不存在", cr.UserName))
		res.FailWithMsg("用户名或密码错误", c)
		return
	}
	//校验密码
	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	//如果密码正确isCheck就为true
	if !isCheck {
		global.Log.Warn("用户密码错误")
		log.Warning(fmt.Sprintf("%s %s 用户密码错误", cr.UserName, cr.Password))
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
		log.Error(fmt.Sprintf("%s 生成Token失败", err.Error()))
		res.FailWithMsg("生成Token失败", c)
		return
	}

	ip, addr := utils.GetAddrByGin(c)

	log = log_stash.New(ip, token)
	log.Info("登录成功")

	global.DB.Create(&models.LoginDataModel{
		UserID:    userModel.ID,
		IP:        c.ClientIP(),
		NickName:  userModel.NickName,
		Token:     token,
		Device:    ip,
		Addr:      addr,
		LoginType: ctype.SignEmail,
	})

	res.OKWithData(token, c)
}
