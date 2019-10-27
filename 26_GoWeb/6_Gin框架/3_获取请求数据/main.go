package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

//将请求内容绑定到自定义结构体类型
//    1. form标签用于和参数对应
//    2. binding标签表示该请求参数是必须的
//    3. json标签为JSON序列化后字段名称命名
//       注意：BindJSON()能解析前端JSON请求体 要求标签json与form同名 否则解析失败
type User struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Password string `form:"password" binding:"required"`
	Phone    string `form:"phone"`
	Age      int    `form:"age"`
}

func main() {
	engine := gin.Default()

	//GET请求 localhost:8888/login?name=qyz&password=1234&phone=17615110273&age=18
	engine.GET("/login", login)
	//POST请求 localhost:8888/login      请求体：name=qyz password=1234 phone=17615110273 age=18
	engine.POST("/login1", login1)

	engine.Run(":8888")
}

//GET请求 获取 普通请求 的参数
func login(context *gin.Context) {

	var user User
	//ShouldBindQuery() 用于GET请求绑定请求参数到结构体
	err := context.ShouldBindQuery(&user)
	if err != nil {
		log.Println(err.Error())
		return
	}

	res, _ := json.Marshal(user)
	//给请求端返回数据
	context.Writer.Write(res)
}

//POST请求 获取 普通请求和Json请求 的参数
func login1(context *gin.Context) {
	var user User
	//ShouldBind() 用于解析普通POST请求，并绑定请求参数到结构体
	//err := context.ShouldBind(&user) //解析普通请求

	//ShouldBind() 用于解析JSON格式POST请求，并绑定请求参数到结构体
	err := context.BindJSON(&user) //解析JSON请求
	if err != nil {
		log.Println(err.Error())
		return
	}

	res, _ := json.Marshal(user)
	//给请求端返回数据
	context.Writer.Write(res)
}
