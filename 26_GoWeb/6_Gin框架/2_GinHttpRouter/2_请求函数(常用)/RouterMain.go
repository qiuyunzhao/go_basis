package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default() //Default 使用 Logger 和 Recovery 中间件 ; engine:=gin.New()不使用默认中间件

	//------------------------------------- 2.该方式比 Handle 方法更常用-------------------------------------------------
	//GET请求 localhost:8888/login?name=qyz
	engine.GET("/login", login)
	//POST请求 localhost:8888/login1   请求体 username  password
	engine.POST("/login1", login1)
	//DELETE请求 localhost:8888/user/123456
	engine.DELETE("/user/:id", deleteUser)

	engine.Run(":8888") // 监听并在 0.0.0.0:8888 上启动服务
}

//GET请求
func login(context *gin.Context) { //context上下文环境 请求参数在其中
	//获取解析接口
	path := context.FullPath()
	//获取请求参数
	requestParam := context.Query("name") //获取不到时不带默认值
	//requestParam := context.DefaultQuery("name", "获取不到时候的默认值") //获取不到时带默认值
	//给请求端返回数据
	context.Writer.Write([]byte("解析接口:" + path + "\t 获取到请求参数为：" + requestParam))
}

//POST请求
func login1(context *gin.Context) {
	//获取解析接口
	path := context.FullPath()
	//获取请求参数
	userName := context.PostForm("username")
	password, exist := context.GetPostForm("password") //返回password是否存在的标志
	if !exist {
		log.Println("获取不到请求参数-password")
	}
	//给请求端返回数据
	context.JSON(200, gin.H{
		"matchUrl": path,
		"username": userName,
		"password": password,
	})
}

//DELETE请求
func deleteUser(context *gin.Context) {
	//获取解析接口
	path := context.FullPath()
	//获取变量参数 id 的值
	userID := context.Param("id")
	//给请求端返回数据
	context.Writer.Write([]byte("解析接口:" + path + "\t 获取变量参数id的值：" + userID))
}
