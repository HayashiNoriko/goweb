package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func initRedis() (*redis.Client, context.Context) {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "123456",
		DB:       0, // 使用默认 DB
	})

	return rdb, ctx
}

func main1() {
	rdb, ctx := initRedis()
	defer rdb.Close()

	// 测试连接
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("redis ping: ", pong)
}
