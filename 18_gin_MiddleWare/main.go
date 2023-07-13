// Gin框架允许开发者在处理请求的过程中，加入用户自己的钩子（Hook）函数。
// 这个钩子函数就叫中间件，中间件适合处理一些公共的业务逻辑，比如登录认证、权限校验、数据分页、记录日志、耗时统计等。

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func indexHander(c *gin.Context) {
	fmt.Println("index")
	// 从上下文取值（跨越中间件m1存取值）
	name, ok := c.Get("name")
	if !ok {
		name = "匿名函数"
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": name,
	})
}

// 定义一个中间件m1：统计请求处理函数的耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in...")

	// gin中间件中使用goroutine
	// 不能使用原始的上下文（c *gin.Context），必须使用其只读副本（c.Copy()）
	// go funcxx(c.Copy()) // 在funcxx中只能使用c的拷贝
	now := time.Now() // 计时
	c.Next()          // 跳过该函数后续，调用后续处理函数
	// c.Abort() // 阻止调用后续的处理函数
	cost := time.Since(now) // since记录开始到调用完函数，到现在的时间
	fmt.Printf("COST:%v\n", cost)
	fmt.Println("m1 out...")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in...")
	c.Next()
	// c.Abort() // 阻止调用后续的处理函数，处理完该函数
	// return // 直接到此结束该函数处理
	c.Set("name", "xiao")
	fmt.Println("m2 out...")
}

func user(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "user",
	})
}

func vider(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "vider",
	})
}

func user1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "g1",
	})
}

func user2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "g2",
	})
}

// 闭包做用户验证登录
func authMiddleware(auCheck bool) gin.HandlerFunc {
	// 连接数据库
	// 其他的

	return func(c *gin.Context) {
		// 判定参数是否登录
		if auCheck {
			// 存放具体的逻辑
			// 是否登录的判定
			c.Next()
		} else {
			c.Next()
		}
	}

}

func main() {
	// 当代码运行时出错，而又没有在编码时显式返回错误时，Go 语言会抛出 panic，
	// 中文译作「运行时恐慌」，我们也可以将其看作 Go 语言版的异常。

	engin := gin.Default() // 默认使用了Logger（记录日志）和Recovery（恢复错误，返回500响应码）
	// engin := gin.New() // 无默认中间件
	engin.Use(m1, m2) // 全局注册中间件函数m1
	engin.GET("/index", indexHander)
	engin.GET("/user", user)
	engin.GET("/vider", vider)

	// 方法一，在路由组中注册局部中间件
	g1 := engin.Group("/api", authMiddleware(true)) // true 登录成功
	g1.GET("/oouser", user1)

	// 方法二，单独注册
	g2 := engin.Group("/api1")
	g2.Use(authMiddleware(true))
	g2.GET("/oouser", user2)

	engin.Run(":8090")
}
