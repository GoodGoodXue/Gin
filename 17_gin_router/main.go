package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Any(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"msg": "laowang",
	})
}

func main() {
	engin := gin.Default()
	// g1 := engin.Group("/api")
	engin.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "laowang",
		})
	})
	engin.Any("/any", Any)
	engin.Run(":9090")
}
