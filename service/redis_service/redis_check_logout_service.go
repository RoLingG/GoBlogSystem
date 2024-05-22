package redis_service

import (
	"GoRoLingG/global"
	"GoRoLingG/utils"
)

// CheckLogout 检测token是否在redis，是则为对应用户已注销过账号
func (RedisService) CheckLogout(token string) bool {
	keys := global.Redis.Keys(logoutPrefix + "*").Val() //普通的keys返回的是对应条件的所有键值的集合指针(?不知道这样说对不对)，要加上.result()或者.val()才能获取到keys集合'
	if utils.InList(logoutPrefix+token, keys) {
		//真实开发环境别用这种方法，会出现阻塞Redis的情况，损耗redis性能
		return true
	}
	return false
}
