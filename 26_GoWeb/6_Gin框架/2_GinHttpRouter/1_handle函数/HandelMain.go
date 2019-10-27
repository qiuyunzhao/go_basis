package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default() //Default 使用 Logger 和 Recovery 中间件 ; engine:=gin.New()不使用默认中间件

	//---------------------------------------- 1.Handle 方法请求---------------------------------------------------------
	//GET请求 localhost:8888/login?name=qyz
	engine.Handle("GET", "/login", login)
	//POST请求 localhost:8888/login1   请求体 username  password
	engine.Handle("POST", "/login1", login1)

	engine.Run(":8888") // 监听并在 0.0.0.0:8888 上启动服务
}

//POST请求
func login1(context *gin.Context) {
	//获取解析接口
	path := context.FullPath()
	log.Println("解析的接口：" + path)
	//获取请求参数
	userName := context.PostForm("username")
	password := context.PostForm("password")
	log.Println("请求参数：" + userName + "--" + password)
	//给请求端返回数据
	context.JSON(200, gin.H{
		"username": userName,
		"password": password,
	})
}

//GET请求
func login(context *gin.Context) { //context上下文环境 请求参数在其中
	//获取解析接口
	path := context.FullPath()
	log.Println("解析的接口：" + path)
	//获取请求参数
	requestParam := context.DefaultQuery("name", "获取不到时候的默认值")
	log.Println("请求参数：" + requestParam)
	//给请求端返回数据
	context.Writer.Write([]byte("解析接口:" + path + "\t 获取到请求参数为：" + requestParam))
}
