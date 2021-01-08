// nacos配置中心的配置:
//	spring:
//		datasource:
//			url: jdbc:postgresql://10.110.84.138:54321/postgres
//			username: postgres
//			password: aaaaaa

package config

var PostgresConf PostgresConfiguration // 全局配置

type PostgresConfiguration struct {
	Spring PostgresSpring `yaml:"spring"`
}

type PostgresSpring struct {
	DataSource PostgresDataSource `yaml:"datasource"`
}

type PostgresDataSource struct {
	Url      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
