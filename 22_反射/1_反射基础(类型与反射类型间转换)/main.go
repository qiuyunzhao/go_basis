package main

import (
	"fmt"
	"reflect"
)

//注意：反射是运行时的，所以类型在编译过程中没法确定，需要转换
func main() {
	//演示基本数据类型、interface{}、reflect.Value反射类型的转换
	var num int = 100
	reflectTest01(num)

	fmt.Println("------------------------------------------")
	//演示结构体类型、interface{}、reflect.Value反射类型的转换
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	reflectTest02(stu)

}

//-------------------------------- 演示基本数据类型、interface{}、reflect.Value反射类型的转换 ----------------------------
func reflectTest01(b interface{}) {
	//1. 获取到 reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Printf("rType=%v  type：%T\n", rType, rType) //rType=int  type：*reflect.rtype

	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal=%v  type：%T\n", rVal, rVal) //rVal=100 type：reflect.Value

	//n1 := 2 + rVal //编译不通过，类型不匹配
	n := 2 + rVal.Int()
	fmt.Println("n2=", n) //n2= 102

	//3.两种方式获取反射类型对应的 kind(一个表示类型的常量)
	kind1 := rVal.Kind()
	kind2 := rType.Kind()
	fmt.Printf("kind1=%v kind2=%v\n", kind1, kind2) //kind1=int kind2=int

	//4.将 rVal 转成 interface{}
	iVal := rVal.Interface()
	fmt.Printf("iVal=%v type：%T\n", iVal, iVal) //iVal=100 type：int

	//5.将interface{} 通过断言转成原来的类型
	num2 := iVal.(int)
	fmt.Printf("num2=%v type：%T\n", num2, num2) //num2=100 type：int

}

//-------------------------------- 演示结构体类型、interface{}、reflect.Value反射类型的转换 ----------------------------
type Student struct {
	Name string
	Age  int
}

type Monster struct {
	Name string
	Age  int
}

func reflectTest02(b interface{}) {

	//1. 获取到 reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Printf("rType=%v  type：%T\n", rType, rType) //rType=1_常用时间日期函数.Student  type：*reflect.rtype

	//2. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal=%v   type：%T\n", rVal, rVal) //rVal={tom 20}   type：reflect.Value

	//3. 获取变量对应的Kind
	kind1 := rVal.Kind()
	kind2 := rType.Kind()
	fmt.Printf("kind=%v   kind=%v\n", kind1, kind2) //kind=struct   kind=struct

	//4.将 rVal 转成 interface{}
	iVal := rVal.Interface()
	fmt.Printf("iVal=%v   type：%T \n", iVal, iVal) //iVal={tom 20}   type：1_常用时间日期函数.Student

	//5.将 interface{} 通过断言转成需要的类型
	//这里，我们就简单使用了一带检测的类型断言.
	//也可以使用 swtich 的断言形式来做的更加的灵活
	stu, ok := iVal.(Student)
	if ok {
		fmt.Printf("stu.Name = %v\n", stu.Name) //stu.Name=tom
	}

}
