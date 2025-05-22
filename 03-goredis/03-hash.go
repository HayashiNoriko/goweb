package main

import "fmt"

func main() {
	rdb, ctx := initRedis()
	defer rdb.Close()

	// 设置哈希字段
	rdb.HSet(ctx, "user1", "name", "Tina", "age", 20)

	// 获取单个字段
	name, err := rdb.HGet(ctx, "user1", "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("[user1] name:", name)

	// 获取所有字段
	userData, err := rdb.HGetAll(ctx, "user1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("[user1] %#v\n", userData)

	// 自增哈希字段
	rdb.HIncrBy(ctx, "user1", "age", 1)
}
