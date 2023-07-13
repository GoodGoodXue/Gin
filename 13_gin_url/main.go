package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Url(c *gin.Context) {
	// 获取路径参数
	name := c.Param("name") // 返回string类型
	age := c.Param("age")
	c.JSON(http.StatusOK, gin.H{
		"msg":  "haol",
		"name": name,
		"age":  age,
	})
}
func Url1(c *gin.Context) {
	// 获取路径参数
	year := c.Param("year")
	month := c.Param("month")
	c.JSON(http.StatusOK, gin.H{
		"msg":   "haol",
		"year":  year,
		"month": month,
	})
}

// 获取请求的path（URL）参数，返回的都是字符串

func main() {
	engin := gin.Default()
	engin.GET("/user/:name/:age", Url)
	engin.GET("/blog/:year/:month", Url1)
	engin.Run(":9090")
}
