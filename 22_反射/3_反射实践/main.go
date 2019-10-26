package main

import (
	"fmt"
	"reflect"
)

func main() {
	//创建了一个Monster实例
	var a Monster = Monster{
		Name:  "黄鼠狼精",
		Age:   400,
		Score: 30.8,
	}

	//将Monster实例传递给TestStruct函数，进行反射调用
	TestStruct(a)
}

//----------------------------------------- 定义了一个Monster结构体 ----------------------------------------------------
type Monster struct {
	Name  string  `json:"name"`
	Age   int     `json:"monster_age"`
	Score float32 `json:"成绩"`
	Sex   string
}

//方法，返回两个数的和
func (monster Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

//方法， 接收四个值，给monster赋值
func (monster Monster) Set(name string, age int, score float32, sex string) {
	monster.Name = name
	monster.Age = age
	monster.Score = score
	monster.Sex = sex
}

//方法，打印monster的值
func (monster Monster) Print() {
	fmt.Println("---start~----")
	fmt.Println(monster)
	fmt.Println("---end~----")
}

//------------------------------------------------ 反射 ---------------------------------------------------------------
func TestStruct(a interface{}) {

	//获取reflect.Type 类型
	rType := reflect.TypeOf(a)

	//获取reflect.Value 类型
	rVal := reflect.ValueOf(a)

	//获取到a对应的类别
	kd := rVal.Kind()

	//如果传入的不是struct，就退出
	if kd != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	//获取到该结构体有几个字段
	num := rVal.NumField()
	fmt.Printf("struct has %d fields \n", num) //4

	//变量结构体的所有字段
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: 值为=%v\n", i, rVal.Field(i))
		//获取到struct的标签, 注意需要通过reflect.Type来获取 名为"json:" 的tag标签的值
		tagVal := rType.Field(i).Tag.Get("json")
		//如果该字段于tag标签就显示，否则就不显示
		if tagVal != "" {
			fmt.Printf("Field %d: tag为=%v\n", i, tagVal)
		}
	}

	//获取到该结构体有多少个方法
	numOfMethod := rVal.NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	//方法的排序默认是按照函数名进行排序（ASCII码）
	rVal.Method(1).Call(nil) //获取到第二个方法（i从0开始）。调用它

	//调用结构体的第1个方法Method(0)
	var params []reflect.Value //声明了切片： []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := rVal.Method(0).Call(params) //传入的参数是 []reflect.Value, 返回 []reflect.Value
	fmt.Println("res=", res[0].Int())  //返回结果, 返回的结果是 []reflect.Value*/

}
