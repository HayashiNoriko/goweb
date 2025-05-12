package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	// 1. 绑定 JSON
	r.POST("/loginJSON", handler)

	// 2. 绑定表单
	r.POST("/loginForm", handler)

	// 3. 绑定 querystring
	r.GET("/loginForm", handler)

	r.Run()
}

func handler(c *gin.Context) {
	fmt.Println("Content-Type:", c.GetHeader("Content-Type"))
	var login Login

	// ShouldBind()会根据请求的Content-Type自行选择绑定器
	// 这里代码可能没太体现出来，但总之，它会自动识别请求的Content-Type，规则为：
	// 对于 GET 请求，强制只从 querystring 绑定数据
	// 对于 POST 请求，根据 Content-Type 来判断：
	//   1. 若为"application/json"，则使用JSON绑定器，同理若为"application/xml"，则使用XML绑定器...
	//   2. 若不是上述几种，则使用Form绑定器（包括 form-data 和 x-www-form-urlencoded）
	// 可以多用 postman 试试
	if err := c.ShouldBind(&login); err != nil {
		// 绑定出错
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		// 绑定成功
		fmt.Println("login info:", login)
		c.JSON(http.StatusOK, gin.H{
			"status":   "ok",
			"username": login.Username,
			"password": login.Password,
		})
	}
}
