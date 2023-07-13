package main

import (
	"Gin/URL/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 匹配 /user/geektutu
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	r.Use(pkg.Cors())
	r.Run(":8010")
}
