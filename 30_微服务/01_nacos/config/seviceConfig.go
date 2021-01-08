//nacos配置中心的配置:
//	server:
//		port: 30002

package config

var ServerConf ServerConfiguration // 全局配置

type ServerConfiguration struct {
	Server Server `yaml:"server"`
}

type Server struct {
	Port uint64 `yaml:"port"`
}
