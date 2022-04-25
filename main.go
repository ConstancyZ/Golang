package main

import (
	"fmt"
	"test_gin/app/blog"
	"test_gin/app/login"
	"test_gin/app/shop"
	"test_gin/app/test"
	"test_gin/routers"
)

func main() {
	// 加载多个APP路由配置
	routers.Include(shop.Routers,blog.Routers,test.Routers,login.Routers)
	// 初始化路由
	r := routers.Init()
	if err := r.Run(); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
