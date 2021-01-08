/*
@ Time : 2021/1/8 15:29
@ Author : qyz
@ File : config
@ Software: GoLand
@ Description:
*/

package config

import (
	"go_basis/30_微服务/01_nacos/ipUtils"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var ServiceIp = ""
var AppConf AppConfiguration

type AppConfiguration struct {
	Nacos Nacos `yaml:"nacos"`
}

type Nacos struct {
	IpAddr      string        `yaml:"ipAddr"`
	Port        uint64        `yaml:"port"`
	NamespaceId string        `yaml:"namespaceId"`
	LogDir      string        `yaml:"logDir"`
	CacheDir    string        `yaml:"cacheDir"`
	Discovery   DiscoveryConf `yaml:"discovery"`
	Config      ConfigConf    `yaml:"config"`
}

type DiscoveryConf struct {
	Group       string `yaml:"group"`
	Cluster     string `yaml:"cluster"`
	ServiceName string `yaml:"serviceName"`
}

type ConfigConf struct {
	Group      string `yaml:"group"`
	ServerId   string `yaml:"serverId"`
	PostgresId string `yaml:"postgresId"`
	RedisId    string `yaml:"redisId"`
}

func LoadConfig() {
	// 读取配置文件
	yamlFile, err := ioutil.ReadFile("30_微服务/01_nacos/config/application.yaml")
	if err != nil {
		log.Panic(err.Error())
	}
	err = yaml.UnmarshalStrict(yamlFile, &AppConf)
	if err != nil {
		log.Panic(err.Error())
	}

	// 获取本机ip
	if ip, err := ipUtils.GetExternalIP(); err == nil {
		ServiceIp = ip.String()
	} else {
		log.Panic("获取本机IP失败")
	}
}
