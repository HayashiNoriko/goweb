package main

import (
	"github.com/gin-gonic/gin"
)

func main0() {
	// 1.创建一个默认的路由引擎
	r := gin.Default()

	// 2.绑定路由规则，执行的函数
	// gin.Engine 组合了 gin.RouterGroup，所以可以直接调用 GET() 方法
	// gin.H 就是 map[string]interface{}
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	// 3.启动服务
	r.Run() // 默认监听 :8080
}
