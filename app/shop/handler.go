package shop

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func insertShop(c *gin.Context) {
	host := c.Request.Host
	scheme := c.Request.URL.Scheme
	fmt.Println("host:", host)
	fmt.Println("scheme:", scheme)
	fmt.Println("insert Shop")
}

func updateShop(c *gin.Context) {
	fmt.Println("update Shop")
}

func selectShop(c *gin.Context) {
	fmt.Println("select Shop")
}
