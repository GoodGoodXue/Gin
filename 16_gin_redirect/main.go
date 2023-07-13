package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Request(c *gin.Context) {
	// 跳转链接
	c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
}

// func A(c *gin.Context) {
// 	// 跳转到 /b对应的路由处理函数
// 	c.Request.URL.Path = "/b" // 把请求的url修改了
// 	r.HandleContext(c)        // 继续后续处理
// }

func B(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "b",
	})
}

func main() {
	r := gin.Default()
	r.GET("/index", Request)
	r.GET("/a", func(c *gin.Context) {
		// 跳转到 /b对应的路由处理函数
		c.Request.URL.Path = "/b" // 把请求的url修改了
		r.HandleContext(c)        // 继续后续处理
	})
	r.GET("/b", B)
	r.Run(":9080")
}
