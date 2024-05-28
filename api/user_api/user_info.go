package user_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/utils/jwt"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

// UserInfoView 用户信息
// @Tags 用户管理
// @Summary 用户信息
// @Description 根据Token获取用户信息
// @Param token header string  true  "token"
// @Produce json
// @Router /api/userInfo [get]
// @Success 200 {object} res.Response{data=models.UserModel}
func (UserApi) UserInfoView(c *gin.Context) {

	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)

	var userInfo models.UserModel
	err := global.DB.Take(&userInfo, claims.UserID).Error
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}
	info := filter.Select("info", userInfo)
	res.OKWithData(info, c)

}
