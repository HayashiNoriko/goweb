package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main1() {
	ctx := context.Background()

	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{ // 集群节点地址列表（不需要全部，会自动发现）
			"127.0.0.1:7000",
			// "127.0.0.1:7001",
			// "127.0.0.1:7002",
		},
		Password:     "",
		PoolSize:     10,
		MaxRetries:   2,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})

	// 测试连接
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("集群连接成功:", pong)

	// 集群信息
	clusterInfo, err := rdb.ClusterInfo(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("集群信息:", clusterInfo)

	// 集群节点信息
	nodes, err := rdb.ClusterNodes(ctx).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("集群节点信息:", nodes)

	// 键操作示例（会自动路由到正确的节点）
	err = rdb.Set(ctx, "foo", "bar666", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "foo").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("foo的值:", val)

	// 关闭连接
	if err := rdb.Close(); err != nil {
		fmt.Println("关闭连接出错:", err)
	}
}
