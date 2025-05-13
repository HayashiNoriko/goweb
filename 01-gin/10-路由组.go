package main

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main10() {
	r := gin.Default()

	// 习惯性一对{}包裹同组的路由，这只是为了看着清晰，但用不用{}包裹功能上没什么区别
	userGroup := r.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {
			c.String(http.StatusOK, "user index")
		})
		userGroup.GET("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "user login")
		})
		userGroup.POST("/login", func(c *gin.Context) {
			c.String(http.StatusOK, "user login post")
		})

	}
	shopGroup := r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.String(http.StatusOK, "shop index")
		})
		shopGroup.GET("/cart", func(c *gin.Context) {
			c.String(http.StatusOK, "shop cart")
		})
		shopGroup.POST("/checkout", func(c *gin.Context) {
			c.String(http.StatusOK, "shop checkout")
		})
		// 路由组也可以嵌套路由组
		carGroup := shopGroup.Group("car")
		{
			carGroup.GET("/benchi", func(c *gin.Context) {
				c.String(http.StatusOK, "shop car benchi")
			})
			carGroup.GET("/baoma", func(c *gin.Context) {
				c.String(http.StatusOK, "shop car baoma")
			})
			carGroup.GET("/aodi", func(c *gin.Context) {
				c.String(http.StatusOK, "shop car aodi")
			})

		}
	}

	r.Run(":8080")
}
