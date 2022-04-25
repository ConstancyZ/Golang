package shop

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func insertShop(c *gin.Context) {
	fmt.Println("insert Shop")
}

func updateShop(c *gin.Context) {
	fmt.Println("update Shop")
}

func selectShop(c *gin.Context) {
	fmt.Println("select Shop")
}
