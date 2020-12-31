package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	egine := gin.Default()

	//请求静态资源 第1步
	egine.LoadHTMLGlob("26_Gin框架/html/*") //设置可加载HTML目录，要返回HTML类型需要先设置
	//请求静态资源 第2步
	egine.Static("/static", "26_Gin框架/static") //加载图片等静态资源要设置加载路径 relativePath:请求路由 root:项目静态资源路径
	//egine.StaticFS("/static",http.Dir("26_Gin框架/static")) //同上
	//egine.StaticFile("/static","26_Gin框架/static/赵敏.jpg") //请求单个静态文件

	//GET请求： http://localhost:8888/returnDifferentType
	egine.GET("returnDifferentType", returnDifferentType)

	egine.Run(":8888")

}

//请求返回 字节数组 和 字符串
func returnDifferentType(context *gin.Context) {
	fullPath := "请求路径：" + context.FullPath()

	////--------------------------------------------  1 返回字节数组 --------------------------------------------------
	//context.Writer.Write([]byte("返回字节数组\n"))
	//
	////--------------------------------------------  2 返回字符串 ----------------------------------------------------
	//context.Writer.WriteString("返回字符串\n")                 //方式1
	//context.String(http.StatusOK,"%s,%s","字符串1","字符串2")  //方式2
	//
	////--------------------------------------------  3 返回 状态码+JSON ----------------------------------------------
	////3.1 Map类型JSON
	//context.JSON(200, map[string]interface{}{
	//	//常用自定义格式：
	//	"code":    1,        // code请求状态码 1:请求成功 0:请求失败
	//	"message": "OK",     // message请求消息 成功：OK  失败:具体错误信息
	//	"data":    fullPath, // data 具体返回数据
	//})
	//
	////3.2 Struct类型JSON
	//resp := Respone{
	//	Code:    1,
	//	Message: "OK",
	//	Data:    fullPath,
	//}
	//context.JSON(200, &resp)

	//--------------------------------------------  4 返回 HTML ---------------------------------------------------
	//把前边几种返回方式注释掉,在浏览器中发出请求即可实现解析后的html展示

	//请求静态资源 第3步
	context.HTML(http.StatusOK, "index.html", gin.H{
		"title":    "Gin返回Html格式数据",
		"fullPath": fullPath,
	})

}

//响应结构体
type Respone struct {
	Code    int         `json:"resp_code"`
	Message string      `json:"resp_name"`
	Data    interface{} `json:"resp_data"`
}
