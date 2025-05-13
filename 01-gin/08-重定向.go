package main

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main8() {
	r := gin.Default()

	// HTTP 重定向
	r.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.qq.com/")
	})

	// 路由重定向
	r.GET("/oldrouter", func(c *gin.Context) {
		// 指定重定向的URL
		c.Request.URL.Path = "/newrouter"
		r.HandleContext(c)
	})
	r.GET("/newrouter", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "this is new router"})
	})

	r.Run(":8080")
}
