/*
@ Time : 2021/1/8 14:36
@ Author : qyz
@ File : nacosConfiguration
@ Software: GoLand
@ Description:
*/

//nacos官网 https://github.com/nacos-group/nacos-sdk-go/blob/master/README_CN.md

package nacosUitls

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go_basis/30_微服务/01_nacos/config"
	"gopkg.in/yaml.v2"
	"log"
)

var NacosConfigClient config_client.IConfigClient // 全局Nacos配置中心客户端

/**
 * @ Time:  2021/1/8 16:48
 * @ Author: qyz
 * @ Description: 创建动态配置客户端
 * @ Param:
 *          :
 * @ return:
 *          :
**/
func CreateConfigClient() error {

	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         config.AppConf.Nacos.NamespaceId, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              config.AppConf.Nacos.LogDir,
		CacheDir:            config.AppConf.Nacos.CacheDir,
		RotateTime:          "1h",
		MaxAge:              3,
	}

	// 至少一个ServerConfig(多个是集群配置)
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: config.AppConf.Nacos.IpAddr,
			Port:   config.AppConf.Nacos.Port,
		},
	}

	// 创建动态配置客户端
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	if err == nil {
		NacosConfigClient = configClient
	}

	return err
}

/**
 * @ Time:  2021/1/8 16:54
 * @ Author: qyz
 * @ Description: 获取nacos远程配置新的配置
 * @ Param:
 *          :
 * @ return:
 *          :
**/
func GetNacosRemoteConfig() error {
	if err := getServerConfig(); err != nil {
		log.Println("读取 nacos 配置中心的", config.AppConf.Nacos.Config.ServerId, "配置信息失败：", err)
		return err
	}
	if err := getPostgresConfig(); err != nil {
		log.Println("读取 nacos 配置中心的", config.AppConf.Nacos.Config.PostgresId, "配置信息失败：", err)
		return err
	}
	if err := getRedisConfig(); err != nil {
		log.Println("读取 nacos 配置中心的", config.AppConf.Nacos.Config.RedisId, "配置信息失败：", err)
		return err
	}
	return nil
}

// 获取 本服务 配置
func getServerConfig() error {
	content, err := NacosConfigClient.GetConfig(vo.ConfigParam{
		DataId: config.AppConf.Nacos.Config.ServerId,
		Group:  config.AppConf.Nacos.Config.Group})
	if err != nil {
		return err
	}

	err = yaml.Unmarshal([]byte(content), &config.ServerConf)
	if err != nil {
		return err
	}

	return nil
}

// 获取 Postgress 数据库配置
func getPostgresConfig() error {
	content, err := NacosConfigClient.GetConfig(vo.ConfigParam{
		DataId: config.AppConf.Nacos.Config.PostgresId,
		Group:  config.AppConf.Nacos.Config.Group})
	if err != nil {
		return err
	}

	err = yaml.Unmarshal([]byte(content), &config.PostgresConf)
	if err != nil {
		return err
	}

	return nil
}

//获取Redis配置
func getRedisConfig() error {
	content, err := NacosConfigClient.GetConfig(vo.ConfigParam{
		DataId: config.AppConf.Nacos.Config.RedisId,
		Group:  config.AppConf.Nacos.Config.Group})
	if err != nil {
		log.Println("读取 nacos 配置中心的 Redis 配置信息失败：", err)

		return err
	}

	err = yaml.Unmarshal([]byte(content), &config.RedisConf)
	if err != nil {
		return err
	}

	return nil
}
