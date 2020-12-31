package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

//Go中基本数据类型转换
func main() {
	//trans()
	//fmt.Println("---------------------------------  基本类型 -> String  -----------------------------------------")
	//otherToString()
	fmt.Println("---------------------------------  String -> 基本类型  -----------------------------------------")
	stringToOther()

}

func stringToOther() {
	fmt.Println("-----------------  利用strconv包中的函数  ----------------------")
	var str1 string = "true"
	b, _ := strconv.ParseBool(str1)
	fmt.Printf("类型:%T---值：%v \n", b, b) //类型:bool---值：true

	var str2 string = "12345"
	num1, _ := strconv.ParseInt(str2, 10, 64) //只能返回64位的，想要别的需要在转换一次
	fmt.Printf("类型:%T---值：%v \n", num1, num1) //类型:bool---值：true

	var str3 string = "12.345"
	num2, _ := strconv.ParseFloat(str3, 64)   //只能返回64位的，想要别的需要在转换一次
	fmt.Printf("类型:%T---值：%v \n", num2, num2) //类型:bool---值：true
}

func otherToString() {
	var v1 int = 99
	var v2 float64 = 23.456
	var v3 bool = true
	var v4 byte = 'h'
	var v5 string

	fmt.Println("-----------------  方式1: fmt.Sprintf()  ----------------------")
	v5 = fmt.Sprintf("%d", v1)
	fmt.Printf("类型:%T---值：%q \n", v5, v5) //类型:string---值："99"

	v5 = fmt.Sprintf("%f", v2)
	fmt.Printf("类型:%T---值：%q \n", v5, v5) //类型:string---值："23.456000"

	v5 = fmt.Sprintf("%t", v3)
	fmt.Printf("类型:%T---值：%q \n", v5, v5) //类型:string---值："true"

	v5 = fmt.Sprintf("%c", v4)
	fmt.Printf("类型:%T---值：%q \n", v5, v5) //类型:string---值："h"

	fmt.Println("-------------  方式2：利用strconv包中的函数  ------------------")
	v5 = strconv.Itoa(v1)
	fmt.Printf("类型:%T---值：%q \n", v5, v5) //类型:string---值："99"

	v5 = strconv.FormatInt(int64(v1), 10)
	fmt.Printf("类型:%T---值：%q \n", v5, v5) //类型:string---值："99"

	v5 = strconv.FormatFloat(v2, 'f', 5, 64)
	fmt.Printf("类型:%T---值：%q \n", v5, v5) //类型:string---值："23.456000"

	v5 = strconv.FormatBool(v3)
	fmt.Printf("类型:%T---值：%q \n", v5, v5) //类型:string---值："true"
}

//	注意：
//		1.数据类型转换必须要显示强转
//      2.范围大-->范围小 转换时可能会按照溢出结果处理
func trans() {
	var a int = 100
	println("a的占位大小", unsafe.Sizeof(a)) //a的占位大小 8

	//int --> float
	var b float32 = float32(a)
	fmt.Printf("b的类型：%T \n", b) //b的类型：float32

	//int --> int8
	var c int8 = int8(a)
	fmt.Printf("c的类型：%T \n", c) //c的类型：int8

	//int --> string
	var d = string(a)
	fmt.Printf("d的类型：%T \n", d) //d的类型：string

	//2.范围大-->范围小 转换时可能会按照溢出结果处理
	var num1 int64 = 999999
	var num2 int8 = int8(num1)
	fmt.Printf("num2的值会按照溢出结果处理：%d \n", num2) //num2的值会按照溢出结果处理：63
}
