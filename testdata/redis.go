package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"time"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,   //0就用默认的DB
		PoolSize: 100, //连接池的大小
	})
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result() //ping一下redis测试是否连接成功，失败则会通过result返回一个结果到err
	if err != nil {
		logrus.Error(err)
		return
	}
}

func main() {
	err := rdb.Set("test1", "test1", 10*time.Second).Err() //添加set类型的新kv
	rdb.Set("test2", "test2", 10*time.Second)
	rdb.Set("test3", "test3", 10*time.Second)
	fmt.Println(err)
	keys := rdb.Keys("*") //使用正则去查所需要的key
	result, err := keys.Result()
	fmt.Println(result, err)
}
