package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main11() {
	ctx := context.Background()

	// 创建哨兵模式客户端
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName: "mymaster", // 主节点名称（在哨兵配置中指定）
		SentinelAddrs: []string{ // 哨兵节点地址列表
			"127.0.0.1:26379",
			"127.0.0.1:26380",
			"127.0.0.1:26381",
		},
		Password:     "",
		DB:           0,
		PoolSize:     10,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})

	// 测试连接
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("连接失败:", err)
		return
	}
	fmt.Println("哨兵模式连接成功:", pong)

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for i := 0; i < 999; i++ { // 循环5次，方便演示
		<-ticker.C
		// 获取主节点信息
		info, err := rdb.Info(ctx, "replication").Result()
		if err != nil {
			fmt.Println("获取主节点信息失败:", err)
			continue
		}
		fmt.Println("主节点信息:", info)

		// 获取当前主节点地址
		sentinel := redis.NewSentinelClient(&redis.Options{
			Addr: "127.0.0.1:26379", // 任选一个哨兵节点
		})
		addr, err := sentinel.GetMasterAddrByName(ctx, "mymaster").Result()
		if err != nil {
			fmt.Println("获取主节点地址失败:", err)
			continue
		}
		fmt.Printf("当前主节点地址: %v\n", addr)
	}

	if err := rdb.Close(); err != nil {
		fmt.Println("关闭连接出错:", err)
	}

}
