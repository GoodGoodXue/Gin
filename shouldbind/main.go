package main

import (
	"Gin/shouldbind/pkg"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// 需要全局定义
var db *gorm.DB

func main() {
	engine := gin.Default()

	engine.Use(pkg.Cors())
	db, err := gorm.Open(mysql.Open("root:123@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal("lianjieshibai", err)
	}
	sql, _ := db.DB()
	defer sql.Close()

	engine.POST("/Index", Index)
	engine.Run(":8090")
}

func Index(c *gin.Context) {
	var user Student
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		// db.Create(&user)
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "创建成功",
			"User": user,
		})
	}

}
