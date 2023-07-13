package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 反射取到字段，需要首字母大写，tag反射查找对应
type UserInfo struct {
	Username  string `form:"username" json:"user"`
	Passsword string `form:"password" json:"pwd"`
}

func Can(c *gin.Context) {
	// username := c.Query("username")
	// password := c.Query("password")
	// u := UerInfo{
	// 	username:  username,
	// 	passsword: password,
	// }

	var u UserInfo // 声明一个UserIfo结构体的变量u
	//
	err := c.ShouldBind(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		fmt.Printf("%#v\n", u)
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}

}

// 表单提交
func Form(c *gin.Context) {
	var u UserInfo // 声明一个UserIfo结构体的变量u
	err := c.ShouldBind(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		fmt.Printf("%#v\n", u)
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func Json(c *gin.Context) {
	var u UserInfo // 声明一个UserIfo结构体的变量u
	err := c.ShouldBind(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		fmt.Printf("%#v\n", u)
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}

func main() {
	engin := gin.Default()
	engin.LoadHTMLFiles("./index.html")
	engin.GET("/user", Can)
	engin.GET("/index", Index)
	engin.POST("/form", Form)
	engin.POST("/json", Json)
	engin.Run(":9080")
}
