package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main1() {
	r := gin.Default()

	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method":  "get",
			"message": "查询书籍",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method":  "post",
			"message": "创建书籍",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method":  "put",
			"message": "修改书籍",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method":  "delete",
			"message": "删除书籍",
		})
	})

	r.Run()
}
