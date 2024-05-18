package redis_service

import (
	"GoRoLingG/global"
	"strconv"
)

const lookPrefix = "look"

// Look  浏览某一篇文章
func (RedisService) Look(id string) error {
	lookCount, _ := global.Redis.HGet(lookPrefix, id).Int() //获取当前文章的点赞数
	lookCount++
	err := global.Redis.HSet(lookPrefix, id, lookCount).Err()
	return err
}

// GetLook 获取某一篇文章下的浏览数
func (RedisService) GetLook(id string) int {
	lookCount, _ := global.Redis.HGet(lookPrefix, id).Int()
	return lookCount
}

// GetLookInfo 取出浏览量数据 为了每隔一段时间去同步数据到es中
func (RedisService) GetLookInfo() map[string]int {
	var LookInfo = map[string]int{}
	maps := global.Redis.HGetAll(lookPrefix).Val() //maps是map[string]string类型的实例
	for id, val := range maps {
		//在redis的哈希里，每个id作为key都对应着自身的点赞数
		num, _ := strconv.Atoi(val) //因为id和val都是string类型，所以要通过strconv.Atoi()将val变成int类型
		LookInfo[id] = num
	}
	return LookInfo
}

// LookClear 同步es后要清空redis内浏览量数据，因为浏览量数据是 同步前es + redis内现有的数据 才是同步后es的数据，所以每次同步后都要清空redis中的浏览量数据
func (RedisService) LookClear() {
	//直接删索引进行清空对应的数据
	global.Redis.Del(lookPrefix)
}

// LookClearByID 删除对应id文章redis内的浏览量数据，用于同步某篇文章的es浏览量
func (RedisService) LookClearByID(id string) {
	global.Redis.HDel(lookPrefix, id)
}
