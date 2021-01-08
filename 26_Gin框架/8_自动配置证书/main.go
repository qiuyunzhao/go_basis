package main

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	//GET请求 http://127.0.0.1:8888/login1?name=qyz&password=1234&phone=17615110273&age=18
	engine.GET("/login1", login1)

	// 流程：生成本地秘钥->秘钥发给证书颁发机构获取私钥->私钥验证->私钥加密
	autotls.Run(engine, "www.itpp.tk")
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
