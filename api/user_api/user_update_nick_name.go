package user_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/models/ctype"
	"GoRoLingG/res"
	"GoRoLingG/utils/jwt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"strings"
)

type UserUpdateNickNameRequest struct {
	NickName string `json:"nick_name" structs:"nick_name"`
	Sign     string `json:"sign" structs:"sign"`
	Avatar   string `json:"avatar" structs:"avatar"`
	Link     string `json:"link" structs:"link"`
}

// UserUpdateNickNameView 修改当前登录人的昵称，签名，链接
// @Tags 用户管理
// @Summary 修改当前登录人的昵称，签名，链接
// @Description 修改当前登录人的昵称，签名，链接
// @Router /api/userUpdateNickName [put]
// @Param token header string  true  "token"
// @Param data body UserUpdateNickNameRequest  true  "昵称，签名，链接"
// @Produce json
// @Success 200 {object} res.Response{}
func (UserApi) UserUpdateNickNameView(c *gin.Context) {
	var cr UserUpdateNickNameRequest
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var newMaps = map[string]interface{}{}
	maps := structs.Map(cr)
	for key, v := range maps {
		if val, ok := v.(string); ok && strings.TrimSpace(val) != "" {
			newMaps[key] = val
		}
	}

	var userModel models.UserModel
	err = global.DB.Debug().Take(&userModel, claims.UserID).Error
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}

	// 如果改的是头像，则判断一下用户的注册来源
	_, ok := newMaps["avatar"]
	if ok && userModel.SignStatus != ctype.SignEmail {
		//如果是qq登录，则不能修改头像
		delete(newMaps, "avatar")
	}

	err = global.DB.Model(&userModel).Updates(newMaps).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("修改用户信息失败", c)
		return
	}
	res.OKWithMsg("修改个人信息成功", c)
}
