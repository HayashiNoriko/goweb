package main

import (
	"fmt"
	"time"

	// "github.com/robfig/cron"
	"github.com/robfig/cron/v3"
)

func main() {
	// 创建一个 cron 实例
	scheduler := cron.New(cron.WithSeconds()) // 启用秒字段支持

	// 执行定时任务（每 5 秒执行一次）
	// entryID 是任务的唯一标识，可用于后续删除任务
	_, err := scheduler.AddFunc("*/1 * * * * *", taskFunc)
	if err != nil {
		fmt.Println(err)
	}

	// 启动 scheduler
	// go scheduler.Run() // run 会阻塞
	scheduler.Start()      // start 不阻塞
	defer scheduler.Stop() // stop 优雅停止

	time.Sleep(10 * time.Second)

}

// 任务函数
func taskFunc() {
	fmt.Println("每 1 秒，执行定时任务...")
}
