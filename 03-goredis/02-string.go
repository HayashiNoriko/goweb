package main

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main2() {
	rdb, ctx := initRedis()
	defer rdb.Close()

	// 设置键值
	err := rdb.Set(ctx, "name", "张三", 0).Err()
	if err != nil {
		panic(err)
	}

	// 获取值
	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name:", val)

	// 获取一个不存在的键的值
	val, err = rdb.Get(ctx, "addr").Result()
	if err == redis.Nil {
		fmt.Println("not exists")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("addr:", val)
	}

	// 设置带过期时间的键值
	rdb.SetEx(ctx, "temp_key", "临时值", 10*time.Second)

	// 自增操作（会先创建键值）
	rdb.Incr(ctx, "counter")      // +1
	rdb.IncrBy(ctx, "counter", 5) // +5

}
