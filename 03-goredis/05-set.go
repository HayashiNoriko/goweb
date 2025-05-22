package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func printLanguages(rdb *redis.Client, ctx context.Context) {
	languages, err := rdb.SMembers(ctx, "myset").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("languages:", languages)
}

func main5() {
	rdb, ctx := initRedis()
	defer rdb.Close()

	// 添加元素
	rdb.SAdd(ctx, "myset", "c", "cpp", "java", "golang")

	// 打印所有元素
	printLanguages(rdb, ctx)

	// 检查元素是否存在
	isMember, err := rdb.SIsMember(ctx, "myset", "c").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("c isMember:", isMember)

	isMember, err = rdb.SIsMember(ctx, "myset", "python").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("python isMember:", isMember)

	// 删除元素
	rdb.SRem(ctx, "myset", "c")
	printLanguages(rdb, ctx)

	// 移动元素
	rdb.SMove(ctx, "myset", "myset2", "java")
	printLanguages(rdb, ctx)

	// 随机获取元素
	random, err := rdb.SRandMember(ctx, "myset").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("random:", random)

	// 求交集
	rdb.SAdd(ctx, "myset3", "c", "cpp", "golang")
	inter, err := rdb.SInter(ctx, "myset", "myset3").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("inter:", inter)

	// 求并集
	union, err := rdb.SUnion(ctx, "myset", "myset3").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("union:", union)

}
