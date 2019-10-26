package main

import (
	"fmt"
)

//注意：接口是引用类型
//-------------------------------------------------------声明、定义接口-------------------------------------------------
type Usb interface {
	//GO接口中不能包含属性，常量也不行
	//只能声明没有实现的方法（不能有方法体）
	Start()
	Stop()
}

type Usb2 interface {
	Start()
	Stop()
}

//-------------------------------------------------------结构体实现接口-------------------------------------------------
type Phone struct {
}

//让Phone 实现 Usb（和Usb2）接口的所有方法
func (p Phone) Start() {
	fmt.Println("手机开始工作。。。")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作。。。")
}

type Camera struct {
}

//让Camera 实现 Usb（和Usb2）接口的所有方法
func (c Camera) Start() {
	fmt.Println("相机开始工作~~~")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作~~~")
}

//------------------------------------------------------- 接口方法 -------------------------------------------------
type Computer struct {
}

//编写一个方法Working 方法，接收一个Usb接口类型变量
//只要是实现了 Usb接口的结构体 都可以传入实现多态
//所谓实现Usb接口，就是指实现了Usb接口声明的 “所有方法” ，如果两个接口所有方法都一样则同时实现这两个接口
func (c Computer) Working(usb Usb) {
	//通过usb接口变量来调用Start和Stop方法
	usb.Start()
	usb.Stop()
}

func (c Computer) Working2(usb Usb2) {
	usb.Start()
	usb.Stop()
}

//--------------------------------------- 自定义类型也可以实现接口（不只是结构体） --------------------------------------
type Integer int

func (i Integer) Start() {
	fmt.Println("自定义类型Start()")
}
func (i Integer) Stop() {
	fmt.Println("自定义类型Stop()")
}

func main() {

	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	//接口本身不能实例化，但是可以指向一个实现了该接口的自定义类型变量（多态）
	var usb Usb = phone
	usb.Start()

	//多态
	computer.Working(phone)
	computer.Working(camera)

	//下面代码证明结构体 Phone{}和Camera{}同时实现了Usb和Usb2接口
	computer.Working2(phone)
	computer.Working2(camera)

	//自定义类型实现接口后也可以赋值给该接口变量
	var i Integer
	var usb1 Usb = i
	usb1.Start()
}
