package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main3() {
	r := gin.Default()

	r.GET("/user/search", func(c *gin.Context) {
		// 1. 普通用法
		address := c.Query("address")

		// 2. 默认值
		username := c.DefaultQuery("username", "tina")

		// 3. 有 age 时，返回 age，否则返回 666
		age, ok := c.GetQuery("age")
		if !ok {
			age = "666"
		}

		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
			"age":      age,
		})

	})

	r.Run()
}
