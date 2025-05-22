// 事务处理
package main

import (
	"fmt"
	"time"
)

func main7() {
	rdb, ctx := initRedis()
	defer rdb.Close()

	// 使用事务
	pipe := rdb.TxPipeline()

	pipe.Set(ctx, "key1", "value1", 0)
	pipe.Set(ctx, "key2", "value2", 0)
	pipe.Incr(ctx, "counter")
	pipe.Expire(ctx, "key1", 10*time.Second)

	// 执行事务
	_, err := pipe.Exec(ctx)
	if err != nil {
		fmt.Println("事务执行失败:", err)
	} else {
		fmt.Println("事务执行成功")
	}
}
