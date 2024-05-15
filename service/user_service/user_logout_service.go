package user_service

import (
	"GoRoLingG/service/redis_service"
	"GoRoLingG/utils/jwt"
	"time"
)

func (UserService) UserLogoutService(claims *jwt.CustomClaims, token string) error {
	//需要计算距离当前时间的token还有多久过期
	exp := claims.ExpiresAt   //获取token过期时间
	now := time.Now()         //获取当前时间
	diff := exp.Time.Sub(now) //用token过期时间-当前时间就算出距离当前还有多久过期
	//这里之所以不能这么写是因为会导致导入依赖循环的问题，user_logout代码内调用了"GoRoLingG/service"依赖，而像下面这样写也会导入"GoRoLingG/service"依赖，user_logout代码调用UserLogoutService()即循环导入
	//err := service.Service.RedisService.Logout(token, diff)
	err := redis_service.RedisService{}.Logout(token, diff)
	return err
}
