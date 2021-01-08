package main

import (
	"github.com/gin-gonic/gin"
	result "go_basis/26_Gin框架/ginResult"
)

func main() {
	engine := gin.Default()

	engine.GET("/ok", func(c *gin.Context) {
		c.JSON(200, result.OK)
	})

	engine.GET("/data", func(c *gin.Context) {
		res := struct {
			Name  string `json:"name"`
			Age   int    `json:"age"`
			Email string `json:"email"`
		}{
			Name:  "小明",
			Age:   18,
			Email: "110@qq.com",
		}
		c.JSON(200, result.OK.WithData(res))
	})

	engine.GET("/err", func(c *gin.Context) {
		c.JSON(200, result.ErrOrderOutTime)
	})

	engine.Run(":8888") // 监听并在 0.0.0.0:8888 上启动服务

}
