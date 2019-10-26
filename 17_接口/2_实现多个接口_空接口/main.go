package main

import (
	"fmt"
)

type AInterface interface {
	test01()
	//test02()  //报错
}

type BInterface interface {
	test02()
}

//接口继承多个接口，被继承的接口不能有相同的方法
type CInterface interface {
	AInterface
	BInterface
	test03()
}

//如果需要实现CInterface,就需要将AInterface BInterface的方法都实现
type Stu struct {
}

func (stu Stu) test01() {
	fmt.Println("test01")
}
func (stu Stu) test02() {
	fmt.Println("test02")
}
func (stu Stu) test03() {
	fmt.Println("test03")
}

//------------------------------------------------------ 空接口 --------------------------------------------------------
//空接口没有任何方法，所有类型都实现了空接口（即所有类型都可以传参或者赋值给空接口）
type EnptyInterface interface {
}

func main() {
	var stu Stu

	var a AInterface = stu
	a.test01()
	var b BInterface = stu
	b.test02()
	var c CInterface = stu
	c.test01()
	c.test02()
	c.test03()

	//所有类型都可以传参或者赋值给空接口
	var t EnptyInterface = stu
	fmt.Println(t)

	var t2 interface{} = stu
	fmt.Println(t2)

	var num1 float64 = 8.8
	t = num1
	t2 = "字符串类型赋值给空接口"
	fmt.Println(t, t2)
}
