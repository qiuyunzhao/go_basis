package main

import (
	"fmt"
)

//包内初始化顺序，全局变量-->init()-->基础用法()
func init() {
	fmt.Println("package 基础用法.init()函数被执行")
}

func main() {
	////知识点1：在Go中，函数也是一种数据类型，可以将函数赋值给一个变量，该变量就是一个函数类型的变量了。通过该变量可以实现对函数的调用
	//test1()

	////知识点2：函数可以作为实参传入
	//test2()

	////知识点3：Go支持对函数返回值命名
	//test3()

	////知识点4：GO函数支持可变参数
	//test4()

	////知识点5：测试函数执行顺序
	//testExecuteOrder.GetInfo()

	//知识点6：匿名函数
	test6()

}

//知识点1：在Go中，函数也是一种数据类型，可以将函数赋值给一个变量，该变量就是一个函数类型的变量了。通过该变量可以实现对函数的调用
func test1() {
	a := getSum
	fmt.Printf("变量a的类型---%T \n", a)
	fmt.Printf("getSum函数的类型---%T \n", getSum)
	res := a(1, 2)
	fmt.Println("res=", res)
}

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

//知识点2：Go中函数是一种数据类型，因此在Go中，函数可以作为形参，并且调用
func test2() {
	res2 := funAsParam(getSum, 3, 6)
	fmt.Println("res2=", res2)
}

func funAsParam(funAdd func(int, int) int, num1 int, num2 int) int {
	return funAdd(num1, num2)
}

//知识点3：对函数返回值命名---多用于返回多个变量时，提前声明返回值变量，防止返回值对应关系混乱问题
func test3() {
	sum, sub := getSumAndSub(22, 11)
	fmt.Println("sum=", sum, "sub=", sub)
}

func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return //此时不用再写返回值
}

//知识点4：GO函数支持可变参数
//Go中函数支持 0-多个参数 写法 (其中args是slice切片，通过args[index]可以访问各个值)
//Go中函数支持 1-多个参数 写法 (注意：vars名字是任意的不过习惯用args，可变参数要放在参数列表最后)
func test4() {
	res3 := sum0(1, 2, 3, 4, 5)
	fmt.Println("res3=", res3)
	res4 := sum0(10, 1, 2, 3, 4, 5)
	fmt.Println("res4=", res4)
}

//0-多个参数
func sum0(args ...int) int {
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}

//1-多个参数
func sum1(num1 int, vars ...int) (sum int) {
	sum = num1
	for i := 0; i < len(vars); i++ {
		sum += vars[i]
	}
	return
}

//知识点6：匿名函数
func test6() {
	//只能使用一次的匿名函数
	res := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res=", res)

	//可以多次使用的匿名函数（将匿名函数赋给变量）
	MyFunc := func(n1 int, n2 int) int {
		return n1 * n2
	}
	res = MyFunc(3, 5)
	fmt.Println("res=", res)
}
