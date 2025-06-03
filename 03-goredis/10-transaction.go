package main

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main10() {
	// basicTransaction()
	// transactionWithErrors()
	watchExample()
}

// 1、基础事务（和之前的 pipe.go 文件中讲述的一样）
func basicTransaction() {
	rdb, ctx := initRedis()
	defer rdb.Close()

	// 开始事务
	pipe := rdb.TxPipeline()

	// 将命令放入队列
	pipe.Incr(ctx, "counter")
	pipe.Expire(ctx, "counter", time.Minute)

	// 执行事务
	_, err := pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("事务执行成功")
}

// 2、事务中的错误处理
func transactionWithErrors() {
	rdb, ctx := initRedis()
	defer rdb.Close()

	// 错误示例：在事务中执行不支持的命令
	_, err := rdb.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Set(ctx, "key1", "value1", 0)
		pipe.HSet(ctx, "key1", "value1", "value2") // HSET 对字符串类型会报错
		pipe.Set(ctx, "key2", "value2", 0)
		return nil
	})

	if err != nil {
		fmt.Println("事务执行出错: ", err)
	}

}

// 3、乐观锁实现---CAS 模式
func watchExample() {
	rdb, ctx := initRedis()
	defer rdb.Close()

	// 要监视的 key
	key := "balance"
	// 重试次数
	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		// 监视 key
		err := rdb.Watch(ctx, func(tx *redis.Tx) error {
			// 获取当前值
			val, err := tx.Get(ctx, key).Int()
			if err != nil && err != redis.Nil {
				return err
			}

			// 模拟处理时间
			time.Sleep(5 * time.Second)

			// 开始事务
			_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
				fmt.Println("即将执行业务逻辑")
				// 业务逻辑：余额-10
				pipe.Set(ctx, key, val-10, 0)
				return nil
			})

			return err

		}, key)

		switch err {
		case nil:
			// 成功
			fmt.Println("事务执行成功")
			return
		case redis.TxFailedErr:
			// 被其他客户端修改，重试
			fmt.Printf("第 %d 尝试失败，重试中...\n", i+1)
			continue
		default:
			// 其他错误
			panic(err)
		}

	}

	fmt.Println("达到最大重试次数，操作失败")

}
