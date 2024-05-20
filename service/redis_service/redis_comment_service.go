package redis_service

import (
	"GoRoLingG/global"
	"strconv"
)

const commentPrefix = "comment"

// Comment 评论某一篇文章
func (RedisService) Comment(id string) error {
	commentCount, _ := global.Redis.HGet(commentPrefix, id).Int() //获取当前文章的点赞数
	commentCount++
	err := global.Redis.HSet(commentPrefix, id, commentCount).Err()
	return err
}

// GetComment 获取某一篇文章下的文章
func (RedisService) GetComment(id string) int {
	commentCount, _ := global.Redis.HGet(commentPrefix, id).Int()
	return commentCount
}

// GetCommentInfo 取出点赞数据 为了每隔一段时间去同步数据到es中
func (RedisService) GetCommentInfo() map[string]int {
	var CommentInfo = map[string]int{}
	maps := global.Redis.HGetAll(commentPrefix).Val() //maps是map[string]string类型的实例
	for id, val := range maps {
		//在redis的哈希里，每个id作为key都对应着自身的点赞数
		num, _ := strconv.Atoi(val) //因为id和val都是string类型，所以要通过strconv.Atoi()将val变成int类型
		CommentInfo[id] = num
	}
	return CommentInfo
}

// CommentClear 清楚文章评论数据
func (RedisService) CommentClear() {
	//直接删索引进行清空对应的数据
	global.Redis.Del(commentPrefix)
}

// CommentClearByID 删除对应id文章redis内的评论数据，用于同步某篇文章的es评论数
func (RedisService) CommentClearByID(id string) {
	global.Redis.HDel(commentPrefix, id)
}
