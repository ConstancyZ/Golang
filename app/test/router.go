package test

import (
	"github.com/gin-gonic/gin"
	"test_gin/middleware"
)

func Routers(e *gin.Engine) {
	// 下面的路由都是通过AuthMiddleWare的，都需要先登陆接口登陆
	e.GET("/returnStruct", middleware.AuthMiddleWare(), returnStruct)
	e.GET("/returnXML", middleware.AuthMiddleWare(), returnXML)
	e.GET("/returnYAML", middleware.AuthMiddleWare(), returnYAML)
	e.GET("/returnProtoBuf", middleware.AuthMiddleWare(), returnProtoBuf)
	e.GET("/redirect", middleware.AuthMiddleWare(), redirect)
	e.GET("/async", middleware.AuthMiddleWare(), async)
	e.GET("/sync", middleware.AuthMiddleWare(), sync)
	e.GET("/testMiddleWare", middleware.AuthMiddleWare(), testMiddleWare)
	// 局部中间件使用
	e.GET("/partMiddleWare", middleware.PartMiddleware(), partMiddleWare)
	// 结构体验证
	e.GET("/dataVerify", middleware.PartMiddleware(), dataVerify)
	// 颁发token
	e.GET("/setToken", middleware.PartMiddleware(), setToken)
	// 解析token
	e.GET("/getToken", middleware.PartMiddleware(), getToken)
}
