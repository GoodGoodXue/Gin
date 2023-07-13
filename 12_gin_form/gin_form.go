package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Form(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func Post(c *gin.Context) {
	// 获取form表单提交的数据
	// username := c.PostForm("username")
	// password := c.PostForm("password")

	// 取的到key返回原，取不到返回默认值
	// username := c.DefaultPostForm("username", "somebody")
	// password := c.DefaultPostForm("password", "***")

	username, ok := c.GetPostForm("username")
	if !ok {
		username = "sb"
	}

	password, ok := c.GetPostForm("password")
	if !ok {
		password = "***"
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Name":     username,
		"Password": password,
	})
}

// 一次请求对应一个响应

func main() {
	engin := gin.Default()
	engin.LoadHTMLFiles("./login.html", "./index.html")
	// if engin != nil {
	// 	g1 := engin.Group("/api")
	// 	g1.GET("/login", Form)
	// 	g1.POST("/login", Post)
	// 	engin.Run(":9090")
	// }
	engin.GET("/login", Form)
	engin.POST("/login", Post)
	engin.Run(":9090")
}
