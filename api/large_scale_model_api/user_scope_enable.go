package large_scale_model_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/utils/jwt"
	"github.com/gin-gonic/gin"
)

type UserScopeEnableResponse struct {
	Enable bool `json:"enable"`
	Scope  int  `json:"scope"`
}

// UserScopeEnableView 用户是否可以领取积分
func (LargeScaleModelApi) UserScopeEnableView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	userID := claims.UserID

	// 查这个用户，今天能不能领取这个积分
	var userScopeModel models.UserScopeModel
	err := global.DB.Take(&userScopeModel, "user_id = ? and to_days(create_at)=to_days(now())", userID).Error
	var response UserScopeEnableResponse
	if err == nil {
		// 查到了，即当天获取过积分
		res.OKWithData(response, c)
		return
	}
	response.Enable = true
	response.Scope = global.Config.LargeScaleModel.ModelSessionSetting.DailyScope
	res.OKWithData(response, c)
}
