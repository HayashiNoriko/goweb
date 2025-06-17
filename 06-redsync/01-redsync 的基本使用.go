package main

import (
	"fmt"
	"time"

	"github.com/go-redsync/redsync/v4"
	redsyncgoredis "github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
)

func main() {
	// 创建 redis 客户端
	rclient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pool := redsyncgoredis.NewPool(rclient)

	// 初始化 Redsync
	rs := redsync.New(pool)

	// 加锁
	// 第一个参数是锁的名字，锁的唯一标识
	// 第二个参数及之后都是可选的
	// 这里第二个参数是过期时间，不传的话默认是 8 秒过期
	mutex := rs.NewMutex("my_lock", redsync.WithExpiry(1000*time.Second))
	if err := mutex.Lock(); err != nil {
		fmt.Println("get lock failed")
		return
	}
	defer mutex.Unlock()

	// 模拟处理业务逻辑
	time.Sleep(100 * time.Second)
}
