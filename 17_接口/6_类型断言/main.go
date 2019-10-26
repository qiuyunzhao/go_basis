package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	point := Point{1, 2}
	a = point //空接口可以赋值任何类型（多态）

	// 如何将 a 赋给一个Point变量?
	var b Point
	// b = a 不可以
	b = a.(Point) //类型断言（判断a是否指向Point类型的变量，如果是则进行转换并赋值，不是则报错panic）
	fmt.Println(b)

	//类型断言(带检测的)
	var x interface{}
	var b2 float32 = 2.1
	x = b2 //空接口，可以接收任意类型
	//类型断言检测机制(不会报panic终止程序执行)
	if y, ok := x.(float64); ok {
		fmt.Println("convert success")
		fmt.Printf("y 的类型是 %T 值是=%v", y, y)
	} else {
		fmt.Println("convert fail")
	}

	fmt.Println("继续执行...")

}
