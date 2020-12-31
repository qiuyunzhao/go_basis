package main

import "github.com/gin-gonic/gin"

func main() {
	engine := gin.Default()

	// 路由组: G1
	rg1 := engine.Group("/user")
	{
		rg1.POST("/register", register) //http://localhost:8888/user/register
		rg1.POST("/login", login)       //http://localhost:8888/user/login
		rg1.DELETE("/:id", deleteUser)  //http://localhost:8888/user/111
	}

	// 路由组: G2
	rg2 := engine.Group("/goods")
	{
		rg2.POST("/sell", sell) //http://localhost:8888/goods/sell
		rg2.POST("/buy", buy)   //http://localhost:8888/goods/buy
	}

	engine.Run(":8888") // 监听并在 0.0.0.0:8888 上启动服务

}

func register(context *gin.Context) {
	context.Writer.WriteString("用户注册成功")
}

func login(context *gin.Context) {
	context.Writer.WriteString("用户登陆成功")
}
func deleteUser(context *gin.Context) {
	userID := context.Param("id")
	context.Writer.WriteString("删除用户" + userID)
}

func sell(context *gin.Context) {
	context.Writer.WriteString("出售商品")
}

func buy(context *gin.Context) {
	context.Writer.WriteString("购买商品")
}
