package test

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"log"
	"net/http"
	"time"
)

func returnStruct(c *gin.Context) {
	var msg struct {
		Name    string
		Message string
		Number  int
	}
	msg.Name = "root"
	msg.Message = "message"
	msg.Number = 123
	c.JSON(200, msg)
}
func returnXML(c *gin.Context) {
	c.XML(200, gin.H{"message": "abc"})
}
func returnYAML(c *gin.Context) {
	c.YAML(200, gin.H{"name": "zhangsan"})
}

func returnProtoBuf(c *gin.Context) {
	reps := []int64{int64(1), int64(2)}
	// 定义数据
	label := "label"
	// 传protobuf格式数据
	data := &protoexample.Test{
		Label: &label,
		Reps:  reps,
	}
	c.ProtoBuf(200, data)
}

func redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
}

// goroutine机制可以方便地实现异步处理
// 另外，在启动新的goroutine时，不应该使用原始上下文，必须使用它的只读副本
func async(c *gin.Context) {
	// context副本
	copyContext := c.Copy()
	// 异步处理
	go func() {
		time.Sleep(3 * time.Second)
		log.Println("异步执行：" + copyContext.Request.URL.Path)
	}()
}

func sync(c *gin.Context) {
	time.Sleep(3 * time.Second)
	log.Println("同步执行：" + c.Request.URL.Path)
}

func testMiddleWare(c *gin.Context) {
	// 取context中的值
	req, _ := c.Get("request")
	fmt.Println("request:", req)
	// 页面接收
	c.JSON(200, gin.H{"request": req})
}

// 局部中间件
func partMiddleWare(c *gin.Context) {
	// 取context中的值
	req, _ := c.Get("request")
	fmt.Println("request:", req)
	// 页面接收
	c.JSON(200, gin.H{"request": req})
}

type Person struct {
	//不能为空并且大于10
	Age      int       `form:"age" binding:"required,gt=10"`
	Name     string    `form:"name" binding:"required"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

// 验证结构体
func dataVerify(c *gin.Context) {
	var person Person
	if err := c.ShouldBind(&person); err != nil {
		c.String(500, fmt.Sprint(err))
		return
	}
	c.String(200, fmt.Sprintf("%#v", person))
}

var jwtKey = []byte("www.topgoer.com")
var str string

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// 颁发token
func setToken(c *gin.Context) {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: 2,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",  //签名颁发者
			Subject:   "user token", // 签名主题
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Println(err)
	}
	str = tokenString
	c.JSON(200,gin.H{"token":tokenString})
}

// 解析token
func getToken(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	//vcalidate token formate
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		c.Abort()
		return
	}

	token, claims, err := ParseToken(tokenString)
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
		c.Abort()
		return
	}
	fmt.Println(111)
	fmt.Println(claims.UserId)
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, Claims, err
}