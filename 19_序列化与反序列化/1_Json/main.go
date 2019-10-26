package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := serialStruct()      //演示对结构体的序列化
	b := serialtestMap()     //演示对map的序列化
	c := serialtestSlice()   //演示对切片的序列化
	d := serialtestFloat64() //演示对基本数据类型的序列化

	//反序列化后数据类型要和序列化之前的数据类型一致，对于结构体，结构体类型无所谓，但是结构体字段要完全一致
	unSerialStruct(a)      //演示对结构体的反序列化
	unSerialtestMap(b)     //演示对map的反序列化
	unSerialtestSlice(c)   //演示对切片的反序列化
	unSerialtestFloat64(d) //演示对基本数据类型的反序列化
}

//----------------------------------------------对结构体的序列化----------------------------------------------
//序列化为json的结构体首字母必须是大写，对外可见，否则转换成json时首字母小写的字段会被抛弃
type Monster struct {
	Name     string  //没有tag则json的key为字段名称
	Age      int     `json:"monster_age"`               //反射机制 将tag映射成json的key
	Birthday string  `json:"-"`                         //该字段会被忽略
	Skill    string  `json:"monster_skill,omitempty"`   //将tag映射成json的key,如果字段值为空则省略掉
	Address  string  `json:"monster_address,omitempty"` //初始化时 Address:  "" 或 不写 Address: 就会被忽略掉
	Sal      float64 `json:",omitempty"`
}

func serialStruct() []byte {

	monster := Monster{
		Name:     "牛魔王",
		Age:      500,
		Birthday: "2011-11-11",
		Skill:    "牛魔拳",
		Address:  "", //或者直接不写
		Sal:      8000.0,
	}

	//将monster序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("monster序列化后=%v\n", string(data))

	return data

}

//----------------------------------------------对map的序列化----------------------------------------------
func serialtestMap() []byte {
	var a map[string]interface{} //定义一个map

	a = make(map[string]interface{}) //使用map,需要make
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "洪崖洞"

	//将map进行序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("a map 序列化后=%v\n", string(data))

	return data
}

//----------------------------------------------对切片的序列化----------------------------------------------
func serialtestSlice() []byte {
	var slice []map[string]interface{} //切片 []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}) //使用map前，需要先make
	m1["name"] = "jack"
	m1["age"] = "7"
	m1["address"] = "北京"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = "20"
	m2["address"] = [2]string{"墨西哥", "夏威夷"}
	slice = append(slice, m2)

	//将切片进行序列化操作
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("slice 序列化后=%v\n", string(data))

	return data
}

//----------------------------------------------对基本数据类型序列化----------------------------------------------
// 对基本数据类型进行序列化意义不大（结构就是一个字符串）
func serialtestFloat64() []byte {
	var num1 float64 = 2345.67

	//对num1进行序列化
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	//输出序列化后的结果
	fmt.Printf("num1 序列化后=%v\n", string(data))

	return data
}

//----------------------------------------------对结构体的反序列化----------------------------------------------
func unSerialStruct(bytes []byte) {
	//定义一个Monster实例
	var monster Monster

	//反序列化
	err := json.Unmarshal(bytes, &monster)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}

	fmt.Printf("反序列化后 monster=%v \n", monster)
}

//----------------------------------------------对map的反序列化----------------------------------------------
func unSerialtestMap(bytes []byte) {

	//定义一个map
	var a map[string]interface{}

	//反序列化
	//注意：反序列化map,不需要make,因为make操作被封装到 Unmarshal函数
	err := json.Unmarshal(bytes, &a)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}

	fmt.Printf("反序列化后 a=%v\n", a)
}

//----------------------------------------------对切片的反序列化----------------------------------------------
func unSerialtestSlice(bytes []byte) {
	//定义一个slice
	var slice []map[string]interface{}

	//反序列化，不需要make,因为make操作被封装到 Unmarshal函数
	err := json.Unmarshal(bytes, &slice)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 slice=%v\n", slice)
}

//----------------------------------------------对基本数据类型反序列化----------------------------------------------
func unSerialtestFloat64(bytes []byte) {
	var num float64

	//反序列化
	err := json.Unmarshal(bytes, &num)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}

	fmt.Printf("反序列化后 num=%v\n", num)
}
