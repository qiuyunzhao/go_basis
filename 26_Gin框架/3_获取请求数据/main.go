package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
)

//将请求内容绑定到自定义结构体类型
//    1. form标签用于和参数对应
//    2. binding标签表示该请求参数是必须的
//    3. json标签为JSON序列化后字段名称命名
//       注意：BindJSON()能解析前端JSON请求体 要求标签json与form同名 否则解析失败
type User struct {
	Name     string    `form:"name" binding:"required" json:"name" `
	Password string    `form:"password" binding:"required"`
	Phone    string    `form:"phone"`
	Age      int       `form:"age"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02 15:04:05"`
}

func main() {
	engine := gin.Default()

	//GET请求 http://localhost:8888/login1?name=qyz&password=1234&phone=17615110273&age=18
	engine.GET("/login1", login1)
	//GET请求 http://localhost:8888/login2/array/12345678
	engine.GET("/login2/:name/:pwd", login2)

	//POST请求 http://localhost:8888/login3
	// 请求体：name=qyz password=1234 phone=17615110273 age=18 birthday=2019-11-10 01:02:11
	// {"name":"qyz","Password":"222","Phone":"333","Age":444,"Birthday":"2019-11-10T01:02:11+08:00"}
	engine.POST("/login3", login3)

	engine.Run(":8888")
}

//GET请求 获取 普通请求 的参数
func login1(context *gin.Context) {

	//1. 分别获取各参数
	context.JSON(200, gin.H{
		"name":     context.Query("name"),
		"password": context.Query("password"),
		"phone":    context.Query("phone"),
		"age":      context.DefaultQuery("age", "18"),
	})

	//2. 将请求参数绑定到结构体
	var user User
	_ = context.ShouldBindQuery(&user) //ShouldBindQuery() 用于GET请求绑定请求参数到结构体
	res, _ := json.Marshal(user)
	//返回数据
	context.Writer.Write(res)
}

//GET请求 获取请求路由中的参数
func login2(context *gin.Context) {
	//返回数据
	context.JSON(200, gin.H{
		"name":     context.Param("name"), //用于获取路由中的参数
		"password": context.Param("pwd"),
	})
}

//POST请求 获取 普通请求和Json请求 的参数
func login3(context *gin.Context) {

	//3. 以流的形式一次性获取 (读取后)
	bodyBytes, _ := ioutil.ReadAll(context.Request.Body)
	context.Writer.Write(bodyBytes)

	context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) //将读出的流数据写回，否则后边读不到了

	//1. 分别获取各参数
	context.JSON(200, gin.H{
		"name":     context.PostForm("name"),
		"password": context.PostForm("password"),
		"phone":    context.PostForm("phone"),
		"age":      context.PostForm("age"),
	})

	//2.将请求参数绑定到结构体
	var user User
	_ = context.ShouldBind(&user) //解析普通请求:ShouldBind() 用于解析普通POST请求，并绑定请求参数到结构体
	//_ = context.BindJSON(&user) //解析JSON请求:BindJSON() 用于解析JSON格式POST请求，并绑定请求参数到结构体
	res, _ := json.Marshal(user)
	//给请求端返回数据
	context.Writer.Write(res)
}
