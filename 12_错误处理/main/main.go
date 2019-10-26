package main

import (
	"errors"
	"fmt"
)

//GO 不支持 try-catch-finally的异常捕获机制，采用panic,defer,recover方式处理
//GO可以抛出panic异常然后通过recover捕获这个异常，然后经行处理
func main() {
	test1() //测试异常捕获
	test2() //测试自定义错误
	fmt.Println("执行到main结尾")
}

//----------------------------------------- 使用defer + recover 来捕获和处理异常 ----------------------------------------
func test1() {
	//会在 test1() 函数执行后再执行
	defer func() {
		err := recover() // recover()是内置函数，可以捕获到异常
		if err != nil {
			//这里就可以进行错误信息处理
			fmt.Println("err=", err)
		}
	}() //匿名函数调用执行

	num1 := 10
	num2 := 0
	res := num1 / num2       //发生异常
	fmt.Println("res=", res) //发生异常后不会往下执行
}

//-------------------------------------------- 自定义错误 errors.New + panic --------------------------------------------
//1. errors.New("错误说明")返回一个error类型的值，表示一个错误
//2. panic内置函数，接收一个 interface{} 空接口类型的值（即任何值）作为一个参数，
//   接收error类型的变量，会 输出错误信息并退出程序

func test2() {
	err := readConf("config2.ini")
	if err != nil {
		panic(err) //如果读取文件发送错误，就输出这个错误，并终止程序
	}
	fmt.Println("test02()继续执行....")
}

//模拟读取配置文件init.conf的信息,如果文件名传入不正确，我们就返回一个自定义的错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		//读取...
		return nil
	} else {
		return errors.New("读取文件错误...") //返回一个自定义错误
	}
}
