/*
@ Time : 2021/1/8 14:35
@ Author : qyz
@ File : nacosDiscovery
@ Software: GoLand
@ Description:
*/

//nacos官网 https://github.com/nacos-group/nacos-sdk-go/blob/master/README_CN.md

package nacosUitls

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"go_basis/30_微服务/01_nacos/config"
	"log"
)

var NacosDiscoveryClient naming_client.INamingClient // 全局Nacos注册中心客户端

/**
 * @ Time:  2021/1/8 17:57
 * @ Author: qyz
 * @ Description: 创建服务发现客户端
 * @ Param:
 *          :
 * @ return:
 *          :
**/
func CreateDiscoveryClient() error {

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

	// 创建服务发现客户端
	discoveryClient, err := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	if err == nil {
		NacosDiscoveryClient = discoveryClient
	}

	return err
}

/**
 * @ Time:  2021/1/8 17:58
 * @ Author: qyz
 * @ Description: 注册实例
 * @ Param:
 *          :
 * @ return:
 *          :
**/
func RegisterInstance() error {
	param := vo.RegisterInstanceParam{
		Ip:          config.ServiceIp,
		Port:        config.ServerConf.Server.Port,
		ServiceName: config.AppConf.Nacos.Discovery.ServiceName,
		ClusterName: config.AppConf.Nacos.Discovery.Cluster, // 默认值 DEFAULT
		GroupName:   config.AppConf.Nacos.Discovery.Group,   // 默认值 DEFAULT_GROUP
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"preserved.register.source": "SPRING_CLOUD"},
	}

	if success, err := NacosDiscoveryClient.RegisterInstance(param); !success || err != nil {
		log.Println("注册实例到nacos失败")
		return err
	}

	return nil
}

/**
 * @ Time:  2021/1/8 18:18
 * @ Author: qyz
 * @ Description: 注销实例
 * @ Param:
 *          :
 * @ return:
 *          :
**/
func DeregisterInstance() error {
	param := vo.DeregisterInstanceParam{
		Ip:          config.ServiceIp,
		Port:        config.ServerConf.Server.Port,
		ServiceName: config.AppConf.Nacos.Discovery.ServiceName,
		Cluster:     config.AppConf.Nacos.Discovery.Cluster, // 默认值 DEFAULT
		GroupName:   config.AppConf.Nacos.Discovery.Group,   // 默认值 DEFAULT_GROUP
		Ephemeral:   true,
	}

	if success, err := NacosDiscoveryClient.DeregisterInstance(param); !success || err != nil {
		log.Println("注册实例到nacos失败")
		return err
	}

	return nil
}
