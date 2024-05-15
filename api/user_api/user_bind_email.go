package user_api

import (
	"GoRoLingG/plugins/email"
	"GoRoLingG/res"
	"GoRoLingG/utils/random"
	"github.com/gin-gonic/gin"
)

type BindEmailView struct {
	Email    string  `json:"email" binding:"required,email" msg:"邮箱非法"`
	Password string  `json:"password" binding:"required" msg:"请输入密码"`
	Code     *string `json:"code"`
}

func (UserApi) UserBindEmail(c *gin.Context) {
	//用户绑定邮箱 第一次输入是邮箱
	var cr BindEmailView
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//后台会给邮箱发验证码
	if cr.Code == nil {
		//第一次，后台给邮箱发验证码
		//生成四位验证码，将验证码存入session，保证前后一致
		code := random.RandCode(4)
		//写入session
		email.NewCode().Send(cr.Email, "你的验证码是:"+code)
		res.OKWithMsg("验证法发送成功", c)
	}

	//第二次输入邮箱、验证码、密码、使其完整绑定

	//完成绑定
}
