/*
@ Time : 2021/1/8 14:42
@ Author : qyz
@ File : main
@ Software: GoLand
@ Description:
*/

package main

import (
	"github.com/gin-gonic/gin"
	"go_basis/30_微服务/01_nacos/config"
	"go_basis/30_微服务/01_nacos/nacosUitls"
	"log"
	"strconv"
)

func main() {
	// 读取本地配置文件
	config.LoadConfig()

	// 创建nacos配置中心客户端
	if err := nacosUitls.CreateConfigClient(); err != nil {
		log.Panic("创建nacos配置中心客户端错误：", err)
	}
	// 读取nacos远程配置中心的配置
	if err := nacosUitls.GetNacosRemoteConfig(); err != nil {
		log.Panic("读取nacos远程配置中心的配置错误：", err)
	}

	// 创建nacos注册中心客户端
	if err := nacosUitls.CreateDiscoveryClient(); err != nil {
		log.Panic("创建nacos注册中心客户端错误：", err)
	}
	// 注册该服务到nacos注册中心
	if err := nacosUitls.RegisterInstance(); err != nil {
		log.Panic("注册该服务到nacos注册中心错误：", err)
	}

	// 从nacos注册中心注销该实例
	defer nacosUitls.DeregisterInstance()

	// Gin
	engine := gin.Default()
	engine.Handle("GET", "/getConfig", getConfig)
	_ = engine.Run(":" + strconv.Itoa(int(config.ServerConf.Server.Port)))
}

func getConfig(context *gin.Context) {
	//给请求端返回数据
	context.JSON(200, gin.H{
		"serverConf":   config.ServerConf,
		"postgresConf": config.PostgresConf,
		"redisConf":    config.RedisConf,
	})
}
