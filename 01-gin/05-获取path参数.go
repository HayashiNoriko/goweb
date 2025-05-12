package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main5() {
	r := gin.Default()

	// 会优先匹配
	r.GET("/user/search/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")

		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	// 把前两个地址也纳入参数中
	r.GET("/:p1/:p2/:username/:address", func(c *gin.Context) {
		p1 := c.Param("p1")
		p2 := c.Param("p2")
		username := c.Param("username")
		address := c.Param("address")

		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"p1":       p1,
			"p2":       p2,
			"username": username,
			"address":  address,
		})
	})

	r.Run()
}
