package user_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/models/ctype"
	"GoRoLingG/plugins/qq"
	"GoRoLingG/res"
	"GoRoLingG/utils"
	"GoRoLingG/utils/jwt"
	"GoRoLingG/utils/pwd"
	"GoRoLingG/utils/random"
	"fmt"
	"github.com/gin-gonic/gin"
)

// QQLoginView 用户QQ登录
// @Tags 用户管理
// @Summary 用户QQ登录
// @Description	用户QQ登录
// @Param data query string true	"用户QQ登录的一些参数"
// @Produce json
// @Router /api/qqLogin [post]
func (UserApi) QQLoginView(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		res.FailWithMsg("没有code传递过来", c)
		return
	}
	fmt.Println(code)
	qqInfo, err := qq.NewQQLogin(code) //这个获得的qqInfo只能用一次
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	openID := qqInfo.OpenID
	//根据openID判断用户是否存在
	var userModel models.UserModel
	err = global.DB.Take(&userModel, "token = ?", openID).Error
	if err != nil {
		//用户不存在于数据库中，则注册用户
		randomPwd, randomErr := random.GenerateRandomString(12) //生出12位的随机密码
		if randomErr != nil {
			global.Log.Error(randomErr)
			res.FailWithMsg("随机密码生成失败", c)
			return
		}
		ip, addr := utils.GetAddrByGin(c) //根据IP算地址
		hashedPwd := pwd.HashPwd(randomPwd)
		userModel = models.UserModel{
			NickName:   qqInfo.Nickname,
			UserName:   openID,    //qq登录，之后绑定邮箱可以用邮箱+密码登录
			Password:   hashedPwd, //随机生成的12位密码(哈希)
			Avatar:     qqInfo.Avatar,
			Token:      openID,
			Address:    addr,
			IP:         ip,
			Role:       ctype.PermissionUser,
			SignStatus: ctype.SignQQ,
		}
		err = global.DB.Create(&userModel).Error
		if err != nil {
			global.Log.Error(err)
			res.FailWithMsg("qq登录用户注册失败", c)
			return
		}
	}
	//登录操作
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

	ip, addr := utils.GetAddrByGin(c)
	global.DB.Create(&models.LoginDataModel{
		UserID:    userModel.ID,
		IP:        ip,
		NickName:  userModel.NickName,
		Token:     token,
		Device:    "",
		Addr:      addr,
		LoginType: ctype.SignQQ,
	})

	res.OKWithData(token, c)
}
