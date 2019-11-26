package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/ini.v1"
	"log"
)

func main() {
	config, _ := ReadConfig()

	//将对象，转换成json格式展示
	data, err := json.Marshal(config)
	if err != nil {
		fmt.Println("err:\t", err.Error())
	}
	fmt.Println(string(data))
}

//reference:https://blog.csdn.net/skh2015java/article/details/78505731

//配置文件要通过tag来指定配置文件中的名称
type Config struct {
	MqttHostname           string   `ini:"mqtt_hostname"`
	MqttPort               string   `ini:"mqtt_port"`
	MqttClientID           string   `ini:"mqtt_clientid"`
	MqttUser               string   `ini:"mqtt_username"`
	MqttPassWord           string   `ini:"mqtt_password"`
	MqttQos                int      `ini:"mqtt_qos"`
	MqttTopic              []string `ini:"mqtt_topic"`
	MqttCAcertFilePath     string   `ini:"mqtt_CAcertFilePath"`
	MqttClientcertFilePath string   `ini:"mqtt_ClientcertFilePath"`
	MqttClientkeyFilePath  string   `ini:"mqtt_ClientkeyFilePath"`
}

//读取配置文件并转成结构体
func ReadConfig() (Config, error) {
	var config Config
	conf, err := ini.Load("36_读取配置文件/2_comf配置文件/config.conf") //加载配置文件
	if err != nil {
		log.Println("load config file fail!")
		return config, err
	}
	conf.BlockMode = false
	err = conf.MapTo(&config) //解析成结构体
	if err != nil {
		log.Println("mapto config file fail!")
		return config, err
	}
	return config, nil
}
