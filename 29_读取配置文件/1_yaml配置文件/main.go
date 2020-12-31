package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func main() {
	//// --------------------------读取普通类型配置文件
	//var confSimple ConfSimple
	//confSimple.getConfSimple()
	//
	////将对象，转换成json格式展示
	//data, err := json.Marshal(confSimple)
	//if err != nil {
	//	fmt.Println("err:\t", err.Error())
	//}
	//fmt.Println(string(data))
	//
	// --------------------------读取结构体类型配置文件
	var conf1 Myconf
	conf1.getConfStruct()

	//将对象，转换成json格式展示
	data, err := json.Marshal(conf1)
	if err != nil {
		fmt.Println("err:\t", err.Error())
	}
	fmt.Println(string(data))

	//// --------------------------读取复杂类型配置文件
	//var conf Conf
	//conf.getConf()
	//
	////将对象，转换成json格式展示
	//data, err := json.Marshal(conf)
	//if err != nil {
	//	fmt.Println("err:\t", err.Error())
	//}
	//fmt.Println(string(data))
}

// ------------------------------------ 普通类型配置文件 ------------------------------------
type ConfSimple struct {
	Host     string   `yaml:"host"`
	Port     int      `yaml:"port"`
	SendMail []string `yaml:"sendMail"`
	Password string   `yaml:"password"`
}

func (c *ConfSimple) getConfSimple() {
	yamlFile, err := ioutil.ReadFile("29_读取配置文件/1_yaml配置文件/conf.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// ------------------------------------ 结构体类型配置文件 ------------------------------------
//配置文件中所有字母要小写，结构体属性首字母要大写
type Myconf struct {
	IpPort             string
	StartSendTime      string
	SendMaxCountPerDay int
	Devices            []Device
	WarnFrequency      int
	SendFrequency      int
}

type Device struct {
	DevId string
	Nodes []Node
}

type Node struct {
	PkId     string
	BkId     string
	Index    string
	MinValue float32
	MaxValue float32
	DataType string
}

func (c *Myconf) getConfStruct() {
	yamlFile, err := ioutil.ReadFile("29_读取配置文件/1_yaml配置文件/conf1.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// ------------------------------------ 复杂类型配置文件 ------------------------------------
type Conf struct {
	ApiVersion string   `yaml:"apiVersion"`
	Kind       string   `yaml:"kind"`
	Metadata   Metadata `yaml:"metadata"`
	Spec       Spec     `yaml:"spec"`
}

type Metadata struct {
	Name string `yaml:"name"`
	//map类型
	Labels map[string]*NodeServer `yaml:"labels"`
}

type NodeServer struct {
	Address string `yaml:"address"`
	Id      string `yaml:"id"`
	Name    string `yaml:"name"`
	//注意，属性里，如果有大写的话，tag里不能存在空格
	//如yaml: "nodeName" 格式是错误的，中间多了一个空格，不能识别的
	NodeName string `yaml:"nodeName"`
	Role     string `yaml:"role"`
}

type Spec struct {
	Replicas int    `yaml:"replicas"`
	Name     string `yaml:"name"`
	Image    string `yaml:"image"`
	Ports    int    `yaml:"ports"`
	//slice类型
	Conditions []Conditions `yaml:"conditions"`
}

type Conditions struct {
	ContainerPort string   `yaml:"containerPort"`
	Requests      Requests `yaml:"requests"`
	Limits        Limits   `yaml:"limits"`
}

type Requests struct {
	CPU    string `yaml:"cpu"`
	MEMORY string `yaml:"memory"`
}

type Limits struct {
	CPU    string `yaml:"cpu"`
	MEMORY string `yaml:"memory"`
}

func (c *Conf) getConf() {
	yamlFile, err := ioutil.ReadFile("29_读取配置文件/1_yaml配置文件/conf2.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.UnmarshalStrict(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
}
