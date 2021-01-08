// nacos配置中心的配置:
//	spring:
//		redis:
//			database: 0
//			host: 10.110.84.138
//			port: 6931
//			password: Test6530

package config

var RedisConf RedisConfiguration // 全局配置

type RedisConfiguration struct {
	Spring RedisSpring `yaml:"spring"`
}

type RedisSpring struct {
	Data Redis `yaml:"redis"`
}

type Redis struct {
	Database int    `yaml:"database"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
}
