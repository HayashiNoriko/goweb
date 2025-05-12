package main

import "github.com/gin-gonic/gin"

type Message struct {
	Name        string `json:"name"`
	Message     string
	privateinfo string
}

func main2() {
	r := gin.Default()
	// 使用结构体
	r.GET("/moreJSON", func(c *gin.Context) {
		msg := Message{"小明", "你好", "私有字段不传递"}
		c.JSON(200, msg)
	})
	r.Run()
}
