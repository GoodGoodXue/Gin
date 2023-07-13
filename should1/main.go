package main

import (
	"Gin/should1/pkg"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	Id       int
	Username string
	Possword string
}

var db *gorm.DB

func main() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	engine.SetTrustedProxies(nil)
	engine.Use(pkg.Cors())

	db, err := gorm.Open(mysql.Open("root:123@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal("lianjieshibai", err)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	// engine.GET("/user", Create)
	ApiRoute(engine)
	engine.Run(":8090")

}

func Create(c *gin.Context) {
	var u UserInfo
	err := c.ShouldBind(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"msg":   "生成失败",
			"error": err.Error(),
		})
	} else {
		fmt.Printf("%#v\n", u)
		db.Create(&u)
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "生成成功",
		})

	}

	// c.ShouldBind(&u)
	// if db.Create(&u) != nil {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"code":    http.StatusOK,
	// 		"msg":     "生成成功",
	// 		"Studnet": u,
	// 	})
	// }

}

func ApiRoute(engine *gin.Engine) {
	if engine != nil {
		g := engine.Group("/should")

		g.GET("/user", Create)
	}
}
