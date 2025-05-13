package main

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main9() {
	r := gin.Default()

	// 1. 普通路由（之前的代码中使用的）

	// 2. 匹配所有请求方法的路由
	r.Any("/test", func(c *gin.Context) {
		method := c.Request.Method
		c.String(http.StatusOK, "your method: %s", method)
	})

	// 3. 为没有配置程序的路由配置默认程序
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "no route!!!")
	})

	r.Run(":8080")
}
