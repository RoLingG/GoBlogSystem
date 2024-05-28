package user_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/utils/jwt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strings"
)

type UserUpdateRequest struct {
	NickName  string `json:"nick_name" structs:"nick_name"`
	Signature string `json:"signature" structs:"signature"`
	Telephone string `json:"telephone" structs:"telephone"`
}

// UserInfoUpdateView 用户信息更新
// @Tags 用户管理
// @Summary 用户信息更新
// @Description 用户信息更新，修改当前登录人的昵称，签名，手机号
// @Param token header string  true  "token"
// @Param data body UserUpdateNicknameRequest  true  "昵称，签名，手机号"
// @Produce json
// @Router /api/userInfoUpdate [put]
// @Success 200 {object} res.Response{}
func (UserApi) UserInfoUpdateView(c *gin.Context) {
	var cr UserUpdateRequest
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	if len(cr.Telephone) != 11 {
		cr.Telephone = ""
		logrus.Errorf("输入的手机号非法，已将用户手机号重置为空号")
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
	err = global.DB.Model(&userModel).Updates(newMaps).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("修改用户信息失败", c)
		return
	}
	res.OKWithMsg("修改个人信息成功", c)
}
