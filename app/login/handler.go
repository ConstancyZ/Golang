package login

import "github.com/gin-gonic/gin"

// 只有登陆了这个接口才会开放权限
func login(c *gin.Context) {
	// 设置cookie
	c.SetCookie("password","123",60,"/","localhost",false,true)
	// 返回信息
	c.String(200, "Login success!")
}
