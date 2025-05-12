package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main4() {
	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		// 同理也有 DefaultPostForm、GetPostForm，略

		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"password": password,
		})
	})

	r.Run()
}
