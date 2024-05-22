package redis_service

import (
	"GoRoLingG/global"
	"strconv"
)

func (redis RedisService) Set(id string) error {
	lookCount, _ := global.Redis.HGet(redis.CountIndex, id).Int() //获取当前文章的点赞数
	lookCount++
	err := global.Redis.HSet(redis.CountIndex, id, lookCount).Err()
	return err
}

func (redis RedisService) Get(id string) int {
	lookCount, _ := global.Redis.HGet(redis.CountIndex, id).Int()
	return lookCount
}

func (redis RedisService) GetInfo() map[string]int {
	var LookInfo = map[string]int{}
	maps := global.Redis.HGetAll(redis.CountIndex).Val() //maps是map[string]string类型的实例
	for id, val := range maps {
		//在redis的哈希里，每个id作为key都对应着自身的点赞数
		num, _ := strconv.Atoi(val) //因为id和val都是string类型，所以要通过strconv.Atoi()将val变成int类型
		LookInfo[id] = num
	}
	return LookInfo
}

func (redis RedisService) Clear() {
	//直接删索引进行清空对应的数据
	global.Redis.Del(redis.CountIndex)
}

func (redis RedisService) ClearByID(id string) {
	global.Redis.HDel(redis.CountIndex, id)
}
