package redis_service

import (
	"GoRoLingG/global"
	"fmt"
	"time"
)

// Logout 针对注销的操作
func (redis RedisService) Logout(token string, diff time.Duration) error {
	//用户注销账户，则将token放入redis中，便于后续用户注销账户后操作检测
	err := global.Redis.Set(fmt.Sprintf("%s%s", logoutPrefix, token), "", diff).Err()
	return err
}
