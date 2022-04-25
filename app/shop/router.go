package shop

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.GET("/insert", insertShop)
	e.GET("/update", updateShop)
	e.GET("/select", selectShop)
}
