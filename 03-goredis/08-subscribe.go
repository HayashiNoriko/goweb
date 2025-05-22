// 发布/订阅
package main

import (
	"fmt"
	"time"
)

func main8() {
	rdb, ctx := initRedis()
	defer rdb.Close()

	// 订阅
	pubsub := rdb.Subscribe(ctx, "mychannel")
	defer pubsub.Close()

	// 监听消息
	go func() {
		ch := pubsub.Channel()
		for msg := range ch {
			fmt.Printf("channel=%s message=%s\n", msg.Channel, msg.Payload)
		}
	}()

	// 确保已经监听了消息
	time.Sleep(1 * time.Second)

	// 发布
	// 也可以在 redis-cli 中执行 PUBLISH mychannel "hello"
	rdb.Publish(ctx, "mychannel", "hello world")

	// 阻塞不退出
	select {}
}
