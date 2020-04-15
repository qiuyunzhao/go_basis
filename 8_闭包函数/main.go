package main

import (
	"fmt"
	"strings"
)

func main() {
	////闭包说明
	//test1()

	//使用案例1
	test2()
}

func test1() {
	f := AddUpper()
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))
}

//用闭包实现累加器
//解释：1.AddUpper是一个函数，返回的数据类型是匿名函数  func(int) int 即闭包
//      2.返回的匿名函数 func(x int) int 和 其引用的函数外的变量n 构成闭包
//      3.可以将理解为：闭包-->类  变量n-->字段  匿名函数-->操作
//      4.反复调用f函数时，因为n只初始化一次，所以每次调用都会进行累加
func AddUpper() func(int) int {
	n := 10
	str := "hello"
	return func(x int) int {
		n += x
		str += string(36) //36-->'$'
		fmt.Println("str=", str)
		return n
	}
}

// 功能要求：
// 1)编写一个函数 makeSuffix(suffix string)  可以接收一个文件后缀名(比如.jpg)，并返回一个闭包
// 2)调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(比如.jpg) ,则返回 文件名.jpg , 如果已经有.jpg后缀，则返回原文件名。
// 3)要求使用闭包的方式完成
// 4)strings.HasSuffix , 该函数可以判断某个字符串是否有指定的后缀。

//测试makeSuffix 的使用
func test2() {
	//不使用闭包实现
	fmt.Println("文件名处理后=", makeSuffix2("jpg", "winter"))   // winter.jgp
	fmt.Println("文件名处理后=", makeSuffix2("jpg", "bird.jpg")) // bird.jpg
	//得到一个闭包
	f2 := makeSuffix(".jpg")               //如果使用闭包完成，好处是只需要传入一次后缀。
	fmt.Println("文件名处理后=", f2("winter"))   // winter.jgp
	fmt.Println("文件名处理后=", f2("bird.jpg")) // bird.jpg
}

//不用闭包实现功能
func makeSuffix2(suffix string, name string) string {
	if !strings.HasSuffix(name, suffix) {
		return name + suffix
	}
	return name
}

//使用闭包实现说明：
//1.返回的匿名函数func (name string) string 和 输入变量suffix构成闭包
//2.优势是不用每次调用都传入 ".jpg"
func makeSuffix(suffix string) func(string) string {
	return func(name string) string {
		//如果 name 没有指定后缀，则加上，否则就返回原来的名字
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
