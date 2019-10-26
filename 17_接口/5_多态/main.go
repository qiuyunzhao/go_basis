package main

import (
	"fmt"
)

//接口
type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	name string
}

//Phone 实现 Usb接口的所有方法
func (p Phone) Start() {
	fmt.Println("手机开始工作。。。")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作。。。")
}

//自己的方法
func (p Phone) Call() {
	fmt.Println("手机打电话---")
}

type Camera struct {
	name string
}

//让Camera 实现 Usb接口的方法
func (c Camera) Start() {
	fmt.Println("相机开始工作。。。")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作。。。")
}

type Computer struct {
}

func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
	//如果usb是指向Phone结构体变量，则还需要调用Call方法(类型断言)
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
}

func main() {
	//定义一个Usb接口数组，可以存放Phone和Camera的结构体变量
	//多态数组
	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"小米"}
	usbArr[2] = Camera{"尼康"}

	fmt.Println(usbArr)

	var computer Computer
	for _, v := range usbArr {
		computer.Working(v)
		fmt.Println()
	}

}
