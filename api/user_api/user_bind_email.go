package user_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/plugins/email"
	"GoRoLingG/res"
	"GoRoLingG/utils/jwt"
	"GoRoLingG/utils/pwd"
	"GoRoLingG/utils/random"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type BindEmailView struct {
	Email    string  `json:"email" binding:"required,email" msg:"邮箱非法"`
	Password string  `json:"password" binding:"required" msg:"请输入密码"`
	Code     *string `json:"code"`
}

func (UserApi) UserBindEmail(c *gin.Context) {
	//用户绑定邮箱 第一次输入邮箱
	var cr BindEmailView
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//创建session对象，设置会话
	session := sessions.Default(c)
	//后台会给邮箱发验证码
	if cr.Code == nil {
		//第一次，后台给邮箱发验证码
		//生成四位验证码，将验证码存入session，保证前后一致
		code := random.RandCode(4)
		//写入session
		session.Set("valid_code", code)
		session.Set("email_for_binding", cr.Email)
		//保存会话
		err = session.Save()
		if err != nil {
			global.Log.Error(err)
			res.FailWithMsg("session出错", c)
			return
		}
		err = email.NewCode().Send(cr.Email, "你的验证码是:"+code)
		if err != nil {
			global.Log.Error(err)
			res.FailWithMsg("验证码发送错误", c)
			return
		}
		res.OKWithMsg("验证法发送成功，请到邮箱中查收", c)
		return
	}
	//获取session中保存的验证码
	code := session.Get("valid_code")
	email := session.Get("email_for_binding")
	//第二次输入邮箱、验证码、密码、使其完整绑定
	//第一次的邮箱和第二次的邮箱也要做一致性校验
	//因为第一次post拿验证码，第二次带验证码post确认并绑定邮箱
	//当第一次拿了验证码后，对应的邮箱和验证码都在session里没被删除和用，当第二次换个不存在的邮箱和正确的验证码，如果不做一致性匹配，就会出现绑定不存在的邮箱成功操作
	//并且还要预防用第一次邮箱拿到验证码后，因为没有删除session内的code，获得到验证码后再用不存在的邮箱获取一次验证码，session里就有两个邮箱和两个code，因为没有删除，导致不存在的邮箱可以用存在的邮箱的验证码进行绑定。
	if email != cr.Email {
		res.FailWithMsg("两次要绑定的邮箱不一致，请重新输入邮箱", c)
		return
	}
	//校验验证码
	if code != *cr.Code {
		res.FailWithMsg("验证码错误", c)
		return
	}
	//修改用户邮箱和密码
	_claims, _ := c.Get("claims")         //从jwt.auth中获取claims
	claims := _claims.(*jwt.CustomClaims) //断言
	var userModel models.UserModel
	err = global.DB.Take(&userModel, claims.UserID).Error //从数据库中获取jwt对应用户的id(准确来说是用户登录后就会有jwt的JwtPayLoad存储用户的信息，后端可以通过里面的用户信息去进行操作)
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}
	if len(cr.Password) < 6 {
		res.FailWithMsg("用户密码过短", c)
		return
	}
	hashedPwd := pwd.HashPwd(cr.Password)
	err = global.DB.Model(&userModel).Updates(map[string]any{
		"email":    cr.Email,
		"password": hashedPwd,
	}).Error
	if err != nil {
		res.FailWithMsg("绑定邮箱失败", c)
		return
	}

	//完成绑定
	// 邮箱更新成功后，从session中删除邮箱和验证码
	session.Delete("email_for_binding")
	session.Delete("valid_code")
	err = session.Save()
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("session出错", c)
		return
	}
	res.OKWithMsg("邮箱绑定成功", c)
	return
}
