package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//----------------------------------------------- 0.设置打印位置为文件 ----------------------------------------------
	//file, _ := os.Create("26_Gin框架/7_中间件/gin.log")
	//gin.DefaultWriter = io.MultiWriter(file)      //日志信息
	//gin.DefaultErrorWriter = io.MultiWriter(file) //错误信息

	//----------------------------------------------- 1.默认中间件 -----------------------------------------------------
	//看源码可知，默认使用的两个中间件：engine.Use(Logger(), Recovery())
	//   中间件Logger()作用: 控制台输出 请求时间和请求路由
	//   中间件Recovery()作用: 当出现错误panic是不至于导致整个线程挂掉
	//engine := gin.Default()

	//----------------------------------------------- 2.使用Logger()中间件 ---------------------------------------------
	//engine := gin.New()
	//engine.Use(gin.Logger())

	//----------------------------------------------- 3.使用自定义中间件 ---------------------------------------------
	engine := gin.Default()
	engine.Use(IPAuthMiddleware())

	//GET请求 http://127.0.0.2:8888/login1?name=qyz&password=1234&phone=17615110273&age=18
	engine.GET("/login1", login1)
	engine.Run("127.0.0.2:8888")
}

//GET请求
func login1(context *gin.Context) {

	context.JSON(200, gin.H{
		"name":     context.Query("name"),
		"password": context.Query("password"),
		"phone":    context.Query("phone"),
		"age":      context.DefaultQuery("age", "18"),
	})

}

//自定义中间件(Ip白名单)
func IPAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipList := []string{
			//"127.0.0.1",
			"127.0.0.2",
		}
		flag := false
		clientIP := c.ClientIP()
		for _, value := range ipList {
			if value == clientIP {
				flag = true
				break
			}
		}
		if !flag {
			c.String(http.StatusUnauthorized, "%s , not in ipList", clientIP)
			c.Abort()
		}
	}
}
