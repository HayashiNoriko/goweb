package main

import "fmt"

func main4() {
	rdb, ctx := initRedis()
	defer rdb.Close()

	// 从左侧推入元素
	rdb.LPush(ctx, "mylist", "a", "b", "c")

	// 从右侧推入元素
	rdb.RPush(ctx, "mylist", "d", "e", "f")

	// [c b a d e f]

	// 获取列表长度
	length, err := rdb.LLen(ctx, "mylist").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("length:", length)

	// 获取范围元素（没有 RRange 哈）
	// 0 <= start <= stop
	// stop 如果大于 length-1 的话，取 length-1
	// stop 为 -1 表示最后一个元素
	letters, err := rdb.LRange(ctx, "mylist", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("letter:", letters)

	// 从右侧弹出元素，LPOP 从左侧
	letter, err := rdb.RPop(ctx, "mylist").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("pop letter:", letter)

	// 再次获取范围元素
	letters, err = rdb.LRange(ctx, "mylist", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("letter:", letters)
}
