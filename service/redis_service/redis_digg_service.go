package redis_service

import (
	"GoRoLingG/global"
	"strconv"
)

const diggPrefix = "digg"

// Digg 点赞某一篇文章
func (RedisService) Digg(id string) error {
	diggCount, _ := global.Redis.HGet(diggPrefix, id).Int() //获取当前文章的点赞数
	diggCount++
	err := global.Redis.HSet(diggPrefix, id, diggCount).Err()
	return err
}

// GetDigg 获取某一篇文章下的点赞数
func (RedisService) GetDigg(id string) int {
	diggCount, _ := global.Redis.HGet(diggPrefix, id).Int()
	return diggCount
}

// GetDiggInfo 取出点赞数据 为了每隔一段时间去同步数据到es中
func (RedisService) GetDiggInfo() map[string]int {
	var DiggInfo = map[string]int{}
	maps := global.Redis.HGetAll(diggPrefix).Val() //maps是map[string]string类型的实例
	for id, val := range maps {
		//在redis的哈希里，每个id作为key都对应着自身的点赞数
		num, _ := strconv.Atoi(val) //因为id和val都是string类型，所以要通过strconv.Atoi()将val变成int类型
		DiggInfo[id] = num
	}
	return DiggInfo
}

// DiggClear 同步es后要清空redis内点赞数据，因为点赞数据是 同步前es + redis内现有的数据 才是同步后es的数据，所以每次同步后都要清空redis中的点赞数据
func (RedisService) DiggClear() {
	//直接删索引进行清空对应的数据
	global.Redis.Del(diggPrefix)
}

// DiggClearByID 删除对应id文章redis内的点赞数据，用于同步某篇文章的es点赞数
func (RedisService) DiggClearByID(id string) {
	global.Redis.HDel(diggPrefix, id)
}
