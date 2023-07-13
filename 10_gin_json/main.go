package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Search(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "you finded it",
	})
}
func Create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Created successfully!",
	})
}
func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delected!",
	})
}
func Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Updated!",
	})
}
func main() {
	r := gin.Default()
	Book := r.Group("/user")
	Book.GET("/get", Search)
	Book.POST("/post", Create)
	Book.DELETE("/delete", Delete)
	Book.PUT("/put", Update)
	r.Run(":9000")
}
