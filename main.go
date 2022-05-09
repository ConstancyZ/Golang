package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"test_gin/app/blog"
	"test_gin/app/login"
	"test_gin/app/shop"
	"test_gin/app/test"
	"test_gin/routers"
)

func main() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "fight the loneliness!"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello friend!")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	// 加载多个APP路由配置
	routers.Include(shop.Routers, blog.Routers, test.Routers, login.Routers)
	// 初始化路由
	r := routers.Init()
	if err := r.Run("127.0.0.1:8082"); err != nil {
		fmt.Println("startup service failed", err)
	}
}
