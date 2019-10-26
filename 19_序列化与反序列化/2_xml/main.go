package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
)

//序列化为xml的结构体首字母必须是大写，对外可见，否则转换时首字母小写的字段会被抛弃
type Persons struct {
	//XMLName xml.Name `xml:"persons"` //父标签
	Persons []Person `xml:"person"` //子标签
}
type Person struct {
	Name      string   `xml:"name,attr"` //标签的属性要跟attr标签
	Age       int      `xml:"age,attr"`
	Career    string   `xml:"career"`
	Interests []string `xml:"interests>interest"` //不写 > 当子标签为一个的时候会把它当做对象解析
}

func main() {
	//---------------------------------------- xml序列化 -----------------------------------------
	peolpes := new(Persons) //结构体是值类型，序列化需要引用类型，所以这样创建

	content, err := ioutil.ReadFile("19_序列化与反序列化/2_xml/input.xml") //读文件
	if err != nil {
		log.Println(err)
	}

	err = xml.Unmarshal(content, peolpes) //将文件转化成对象(要传入引用类型)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(peolpes)
	//---------------------------------------- xml反序列化 -----------------------------------------

	//反序列化
	//bytes, err := xml.Marshal(peolpes)
	//参数说明：prefix:每行开头添加内容  indent:缩进关系开头添加内容
	bytes, err := xml.MarshalIndent(peolpes, "", "	")
	if err != nil {
		log.Println(err)
	}

	//添加xml头
	bytes = append([]byte(xml.Header), bytes...)

	//写文件
	err = ioutil.WriteFile("19_序列化与反序列化/2_xml/output.xml", bytes, 0777)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("执行完毕")
}
