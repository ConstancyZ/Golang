package routers

import (
	"github.com/gin-gonic/gin"
	"test_gin/middleware"
)

// routers/routers.go中根据需要定义Include函数用来注册子app中定义的路由，
// Init函数用来进行路由的初始化
type Option func(*gin.Engine)

var options = []Option{}

//注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

// 初始化
func Init() *gin.Engine {
	r := gin.New()
	// 注册中间件，在这里定义相当于全局路由都会有这个中间件
	r.Use(middleware.Middleware())
	// 注册程序计时中间件
	r.Use(middleware.CalculateTime())
	for _, opt := range options {
		opt(r)
	}
	return r
}
