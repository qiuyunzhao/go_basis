package main

import (
	"github.com/gin-gonic/gin"
	result "go_basis/26_Gin框架/ginResult"
	"net/http"
)

func main() {
	engine := gin.Default()

	engine.GET("/ok", func(c *gin.Context) {
		c.JSON(http.StatusOK, result.OK)
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
		c.JSON(http.StatusOK, result.OK.WithData(res))
	})

	engine.GET("/common/err", func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, result.Err.WithMsg("通用错误"))
	})

	engine.GET("/specified/err", func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, result.ErrOrderOutTime)
	})

	engine.Run(":8888") // 监听并在 0.0.0.0:8888 上启动服务

}
