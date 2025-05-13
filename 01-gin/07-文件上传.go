package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main7() {
	r := gin.Default()

	// 1. 单个文件
	r.POST("/upload", func(c *gin.Context) {
		// 获取上传的文件
		file, _ := c.FormFile("file")
		fmt.Println(file.Filename)
		// 上传文件到指定的目录
		c.SaveUploadedFile(file, fmt.Sprintf("./%s", file.Filename))
		c.String(http.StatusOK, "ok")
	})

	// 2. 多个文件
	r.POST("/uploads", func(c *gin.Context) {
		// 获取上传的文件
		form, _ := c.MultipartForm()
		files := form.File["files"]

		for index, file := range files {
			fmt.Println(file.Filename)
			dst := fmt.Sprintf("./%s_%d", file.Filename, index)
			// 上传文件到指定的目录
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})

	r.Run(":8080")
}
