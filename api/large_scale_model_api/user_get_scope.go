package large_scale_model_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/utils/jwt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserGetScopeRequest struct {
	Status bool `json:"status"` //状态有分领取和不领取
}

// UserGetScopeView 用户获取当天的积分
func (LargeScaleModelApi) UserGetScopeView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwt.CustomClaims)
	userID := claims.UserID

	var cr UserGetScopeRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMsg("参数错误", c)
		return
	}

	// 查这个用户当天能不能领取这个积分
	var userScopeModel models.UserScopeModel
	err = global.DB.Take(&userScopeModel, "user_id = ? and to_days(create_at)=to_days(now())", userID).Error
	if err == nil {
		res.FailWithMsg("今日已领取积分啦", c)
		return
	}

	// 用户获取当天积分
	var user models.UserModel
	err = global.DB.Take(&user, userID).Error
	if err != nil {
		res.FailWithMsg("用户不存在", c)
		return
	}
	scope := global.Config.LargeScaleModel.ModelSessionSetting.DailyScope
	global.DB.Model(&user).Update("scope", gorm.Expr("scope + ?", scope))
	err = global.DB.Create(&models.UserScopeModel{
		UserID: userID,
		Scope:  scope,
		Status: cr.Status,
	}).Error
	if err != nil {
		res.FailWithMsg("用户积分获取失败", c)
		return
	}

	res.OKWithMsg("积分领取成功", c)
}
