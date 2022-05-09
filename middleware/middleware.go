package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 定义全局中间 件
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("全局中间件开始执行了")
		// 设置变量到context的key中，可以通过Get()获取
		c.Set("request", "全局中间件")
		status := c.Writer.Status()
		fmt.Println("全局中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

// 定义局部中间件
func PartMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("局部中间件开始执行了")
		// 设置变量到context的key中，可以通过Get()获取
		c.Set("request", "局部中间件")
		status := c.Writer.Status()
		fmt.Println("局部中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

// 程序计时中间件
func CalculateTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		since := time.Since(start)
		fmt.Println("接口用时：", since)
	}
}

// 设计一个模仿cookie验证的权限验证的中间件
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1.获取客户端cookie并校验
		if cookie, err := c.Cookie("password"); err == nil {
			if cookie == "123" {
				return
			}
		}
		// 返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		// 若验证不通过，不再调用后续的函数处理
		c.Abort()
		return
	}
}
