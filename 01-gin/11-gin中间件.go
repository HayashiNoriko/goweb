package main

import (
	"bytes"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// 1. 定义中间件
// 1.1 StatCost 是一个统计耗时请求耗时的中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		c.Set("name", "zhangsan")

		// 处理请求【阻塞，直到 Gin 完成所有请求处理（包括生成响应数据）】
		// 也就是等待 c.JSON()、c.String()、c.HTML() 等响应方法
		c.Next()

		// 处理响应
		cost := time.Since(start)
		log.Printf("cost=%v\n", cost)
	}
}

// 1.2 BodyLog 是一个记录返回给客户端响应体的中间件
type bodyLogWriter struct {
	gin.ResponseWriter               // 嵌入gin框架ResponseWriter（继承）
	body               *bytes.Buffer // 我们记录用的response
}

// Write 写入响应体数据
// 重写 Write 方法，实现双写
// c 本来会调用它自己的原生 c.Writer，但这里我们“派生”了一个 bodyLogWriter ，它具有原 Writer 的所有功能
// 并且我们重写了 bodyLogWriter 的 Write方法，所以可以直接把 blw 赋值给 c.Writer的位置
// 可以正常调用，并且达到双写效果
func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)                  // 先将响应体数据写入w.body中（我们自己记录的）
	return w.ResponseWriter.Write(b) // 再将响应体数据写入客户端（真正返回的）
}

func BodyLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &bodyLogWriter{
			body:           bytes.NewBuffer([]byte{}),
			ResponseWriter: c.Writer,
		}

		// 使用我们自定义的类型替换默认的
		c.Writer = blw
		log.Println("Next 之前的响应体：", blw.body.String()) // 这里输出空字符串，因为还没有执行业务逻辑、生成响应体

		// 阻塞等待执行业务处理函数
		// 它们最终都会调用 c.Writer.Write() 方法
		// 我们通过替换 c.Writer 拦截了这个调用过程
		c.Next()

		// 事后按需记录返回的响应
		log.Println("响应体：", blw.body.String())
	}
}

func main11() {
	// 2. 注册中间件

	// 新建一个没有任何默认中间件的路由
	r := gin.New()

	// 2.1 注册一个全局中间件
	r.Use(StatCost())

	r.GET("/test", func(c *gin.Context) {
		name := c.MustGet("name") // 从上下文中取值
		log.Println(name)
		c.String(200, "okokok")
	})

	// 2.2 为某个路由单独注册
	r.GET("/test2", StatCost(), func(c *gin.Context) {
		c.String(200, "okokok2")
	})

	// 2.3 为路由组注册中间件
	// 第一种写法，作为 Group 函数的第二个参数
	shopGroup := r.Group("/shop", BodyLog())
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.String(200, "shop")
		})
	}

	// 第二种写法，使用 Use 函数
	//shopGroup.Use(BodyLog())
	r.Run()
}
