package main

import "fmt"

type A struct {
	Num int
}
type B struct {
	Num int
}

func main() {
	var a A
	var b B
	a = A(b) //可以转换，但是要求结构体的的字段要完全一样(包括:名字、个数、类型)
	fmt.Println(a, b)

	type Stu A //给结构体A取别名Stu
	var c Stu
	//a = c 不可以， go中重命名后不再是相同类型
	a = A(c)
	fmt.Println(a, c)

	//同理
	type Integer int
	var i Integer = 10
	var j int = int(i)
	fmt.Println(j)
}
