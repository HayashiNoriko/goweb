package main

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main6() {
	rdb, ctx := initRedis()
	defer rdb.Close()

	// 添加带分数的元素
	rdb.ZAdd(ctx, "myzset", redis.Z{Score: 90, Member: "Alice"})
	rdb.ZAdd(ctx, "myzset", redis.Z{Score: 87, Member: "Bob"})
	rdb.ZAdd(ctx, "myzset", redis.Z{Score: 96, Member: "Tom"})
	rdb.ZAdd(ctx, "myzset", redis.Z{Score: 80, Member: "Jack"})
	rdb.ZAdd(ctx, "myzset", redis.Z{Score: 70, Member: "Tim"})

	// 获取 Alice 的名次(从高到低)
	// ZRank 是从低到高
	rank, err := rdb.ZRevRank(ctx, "myzset", "Alice").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Alice rank:", rank)

	// 分数从低到高打印所有元素
	students, err := rdb.ZRangeWithScores(ctx, "myzset", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("all zset members:", students)

	// 打印分数最小的3个元素
	students, err = rdb.ZRangeWithScores(ctx, "myzset", 0, 2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("zset members:", students)

	// 打印分数最大的3个元素
	students, err = rdb.ZRevRangeWithScores(ctx, "myzset", 0, 2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("zset members:", students)
}
