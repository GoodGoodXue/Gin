package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 建立一个结构体来存储数据
type UserInfo struct {
	Id       int
	Username string
	Password string
}

// var db *gorm.DB

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.SetTrustedProxies(nil)
	db, err := gorm.Open(mysql.Open("root:123@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal("lianjieshibai", err)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	r.GET("/user", func(c *gin.Context) {
		//username := c.Query("username")
		//passsword := c.Query("password")
		//u := UserInfo{
		//    username: username,
		//    password: passsword,
		//}
		//声明一个UserInfo类型的变量u
		var u UserInfo
		//这里把地址传过去
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			db.Create(&u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})
	r.Run(":8090")
}
