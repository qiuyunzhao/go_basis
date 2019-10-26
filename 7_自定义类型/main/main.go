package main

import "fmt"

func main() {
	//自定义类型myInt(相当于给int类型起别名)，myInt和int都是int类型，但是Go中认为两者是不同类型
	type myInt int

	var num1 myInt
	var num2 int

	num1 = 40
	//num2 = num1	//报错：类型不匹配
	num2 = int(num1)
	fmt.Printf("num1类型 %T \n", num1)
	fmt.Printf("num2类型 %T \n", num2)

	res := myFunc(getSum, 10, 20)
	fmt.Println("res=", res)
}

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

//自定义类型:myFuncType.
//myFuncType为 func(int, int) int 的函数类型
type myFuncType func(int, int) int

func myFunc(funAdd myFuncType, num1 int, num2 int) int {
	return funAdd(num1, num2)
}
