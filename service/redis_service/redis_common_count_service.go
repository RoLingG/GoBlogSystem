package redis_service

import (
	"GoRoLingG/global"
	"strconv"
)

func (redis RedisService) Set(id string) error {
	Count, _ := global.Redis.HGet(redis.CountIndex, id).Int()
	Count++
	err := global.Redis.HSet(redis.CountIndex, id, Count).Err()
	return err
}

// SetCount 在原有基础上增加多少
func (redis RedisService) SetCount(id string, num int) error {
	oldCount, _ := global.Redis.HGet(redis.CountIndex, id).Int()
	newCount := oldCount + num
	err := global.Redis.HSet(redis.CountIndex, id, newCount).Err()
	return err
}

func (redis RedisService) Get(id string) int {
	Count, _ := global.Redis.HGet(redis.CountIndex, id).Int()
	return Count
}

func (redis RedisService) GetInfo() map[string]int {
	var Info = map[string]int{}
	maps := global.Redis.HGetAll(redis.CountIndex).Val() //maps是map[string]string类型的实例
	for id, val := range maps {
		//在redis的哈希里，每个id作为key都对应着自身的点赞数
		num, _ := strconv.Atoi(val) //因为id和val都是string类型，所以要通过strconv.Atoi()将val变成int类型
		Info[id] = num
	}
	return Info
}

func (redis RedisService) Clear() {
	//直接删索引进行清空对应的数据
	global.Redis.Del(redis.CountIndex)
}

func (redis RedisService) ClearByID(id string) {
	global.Redis.HDel(redis.CountIndex, id)
}
