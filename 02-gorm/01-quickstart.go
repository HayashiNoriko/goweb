package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model // 内嵌 gorm.Model，包含 ID, CreatedAt, UpdatedAt, DeletedAt 字段
	Code       string
	Price      uint
}

func main1() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
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

	// 查询记录
	var product1 Product
	var product2 Product
	db.First(&product1, 1) // 通过主键ID=1查询
	fmt.Println("查询到的记录", product1)
	db.First(&product2, "code = ?", "D42") // 通过code字段查询
	fmt.Println("查询到的记录", product2)

	// 修改记录 - 将第一个 product 的 price 更新为 200
	db.Model(&product1).Update("Price", 200)
	// Update - 更新多个字段
	db.Model(&product1).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	db.Model(&product1).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	db.Delete(&product1, 1)
}
