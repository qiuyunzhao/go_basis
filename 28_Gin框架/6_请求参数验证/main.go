package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//验证规则参考文档： https://godoc.org/gopkg.in/go-playground/validator.v9

//将请求内容绑定到自定义结构体类型
//    1. form         ：用于和请求参数对应
//    2. binding      ：required-表示该请求参数是必须的  gt-大于某值 gtefield-大于等于某个字段
//    3. json         ：为JSON序列化后字段名称命名
//    4. time_format  ：日期格式

//    注意：BindJSON()能解析前端JSON请求体 要求标签json与form同名 否则解析失败
type User struct {
	Name     string    `form:"name" binding:"required" json:"_name" `
	Password string    `form:"password"`
	Phone    string    `form:"phone"`
	Age      int       `form:"age" binding:"required,gt=18"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02 15:04:05"`
	Today    time.Time `form:"today" binding:"gtefield=Birthday" time_format:"2006-01-02 15:04:05"`
}

func main() {
	engine := gin.Default()

	//GET请求 http://localhost:8888/login1?name=qyz&password=1234&phone=17615110273&age=19&birthday=2019-11-10 15:04:05&today=2019-11-10 15:04:05
	engine.GET("/login1", login1)

	engine.Run(":8888")
}

//GET请求 将请求参数绑定到结构体
func login1(context *gin.Context) {
	var user User
	err := context.ShouldBindQuery(&user) //ShouldBindQuery() 用于GET请求绑定请求参数到结构体
	if err != nil {
		context.String(http.StatusBadRequest, "%v", err)
		return
	}
	res, _ := json.Marshal(user)
	//返回数据
	context.Writer.Write(res)
}
