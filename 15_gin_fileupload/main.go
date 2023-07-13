package main

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func Submit(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func Upload(c *gin.Context) {
	// 读取上传的文件
	f, err := c.FormFile("f1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		// dst 保存的目的地
		dst := path.Join("./", f.Filename)
		// 保存到本地服务器端
		c.SaveUploadedFile(f, dst)
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	}
}

func Sums(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]

	for index, file := range files {
		log.Panicln(file.Filename)
		dst := fmt.Sprintf("./%s_%d", file.Filename, index)
		// 上传文件到指定目录
		c.SaveUploadedFile(file, dst)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%d files uploades!", len(files)),
	})
}

func main() {
	engin := gin.Default()
	engin.LoadHTMLFiles("./index.html")
	engin.GET("/index", Submit)
	// engin.POST("/upload", Upload)
	engin.POST("/upload", Sums)
	engin.Run(":9080")
}
