package core

import (
	"GoRoLingG/global"
	"context"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"time"
)

func ConnectRedis() *redis.Client {
	return ConnectRedisDB(0)
}

func ConnectRedisDB(db int) *redis.Client {
	redisConfig := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr(),
		Password: redisConfig.Password,
		DB:       0,                    //0就用默认的DB
		PoolSize: redisConfig.PoolSize, //连接池的大小
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result() //ping一下redis测试是否连接成功，失败则会通过result返回一个结果到err
	if err != nil {
		logrus.Errorf("redis连接失败 %s", err)
		return nil
	}
	return rdb
}
