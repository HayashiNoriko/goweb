package main

import (
	// "fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 01 中已经定义了 Product 结构体
// type Product struct {
// 	gorm.Model // 内嵌 gorm.Model，包含 ID, CreatedAt, UpdatedAt, DeletedAt 字段
// 	Code       string
// 	Price      uint
// }

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动创建或更新 Product 表结构
	// 建议每次启动时都运行，可以检测模型结构变更，自动修改表结构
	// 对于生产环境，更严谨的做法是​使用专门的迁移工具​
	db.AutoMigrate(&Product{})

	// 插入记录
	db.Create(&Product{Code: "A42", Price: 100})
	db.Create(&Product{Code: "D42", Price: 100})

}
