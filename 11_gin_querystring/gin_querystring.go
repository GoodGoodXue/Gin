package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// type My struct {
// 	Id   int
// 	Name string
// 	Age  string
// }

func Ta(c *gin.Context) {
	// id := c.Query("id")
	name := c.Query("name")
	age := c.Query("age")
	// name := c.DefaultQuery("name", "some") // 取不到key返回默认值

	// name, ok := c.GetQuery("name") // 多了一个返回bool值（false)
	// if !ok {
	// 	// 取不到
	// 	name = "somebody"
	// }

	c.JSON(http.StatusOK, gin.H{
		"status_code": http.StatusOK,
		"msg":         "zhaodaol",
		"name":        name,
		"age":         age,
	})
}

func main() {
	r := gin.Default()
	gin := r.Group("/api")
	gin.GET("/find", Ta)
	r.Run(":9090")
}
