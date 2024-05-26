package user_api

import (
	"GoRoLingG/models"
	"GoRoLingG/models/ctype"
	"GoRoLingG/res"
	"GoRoLingG/service/common"
	"GoRoLingG/utils/desensitization"
	"GoRoLingG/utils/jwt"
	"github.com/gin-gonic/gin"
)

// UserListView 用户列表
// @Tags 用户管理
// @Summary 用户列表
// @Description	用户列表，用于展示目前所有用户
// @Param token header string true "Authorization token"
// @Param data query models.PageInfo false	"查询用户列表的一些参数"
// @Router /api/userList [get]
// @Produce json
// @Success 200 {object} res.Response{data=res.ListResponse[models.UserModel]}
func (UserApi) UserListView(c *gin.Context) {
	_claims, _ := c.Get("claims")         //从jwt.auth中获取claims
	claims := _claims.(*jwt.CustomClaims) //断言
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
	}
	var userList []models.UserModel
	list, count, err := common.CommonList(models.UserModel{}, common.Option{
		PageInfo: cr,
		Debug:    false,
	})
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	for _, user := range list {
		//根据token解析的内容判断用户
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			//非管理员
			user.UserName = ""
		}
		//脱敏
		user.Telephone = desensitization.DesensitizationTel(user.Telephone)
		user.Email = desensitization.DesensitizationEmail(user.Email)
		userList = append(userList, user)
	}

	res.OKWithList(userList, count, c)
	return
}
