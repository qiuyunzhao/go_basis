package main

import (
	"fmt"
	"time"
)

//函数1（没问题）
func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello,world")
	}
}

//函数2（没有make集合map就对其赋值，会panic异常，如果不处理会导致整个程序崩溃）
func test() {
	//这里我们可以使用defer + recover
	defer func() {
		//捕获test抛出的panic
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()

	var myMap map[int]string
	myMap[0] = "golang" //panic: assignment to entry in nil map
}

func main() {

	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("1_常用时间日期函数() ok=", i)
		time.Sleep(time.Second)
	}

}
