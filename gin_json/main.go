package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func json(c *gin.Context) {
	// data := map[string]interface{}{
	// 	"name":    "小子",
	// 	"message": "hello world!",
	// 	"age":     "18",
	// }
	data := gin.H{"name": "xiaowang", "message": "hello world", "age": 18}
	c.JSON(http.StatusOK, data)
}

// 方法二：结构体

type msg struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Age     int    `json:"age"`
}

func json1(c *gin.Context) {
	data1 := msg{
		"xiaowangzi",
		"hello goland",
		18,
	}
	c.JSON(http.StatusOK, data1) // json的序列化
}

func ApiRoute(r *gin.Engine) {
	if r != nil {
		JSON := r.Group("/wang")
		JSON.GET("/json", json)
		JSON.GET("/json1", json1)
	}
}

func main() {
	r := gin.Default()
	ApiRoute(r)
	r.Run(":8080")
}
