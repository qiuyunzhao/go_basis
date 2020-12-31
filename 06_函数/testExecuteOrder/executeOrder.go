package testExecuteOrder

import "fmt"

var Name string
var Age int

//用到的别的包下的init函数也会在 基础用法() 函数执行前执行
func init() {
	Name = "小明"
	Age = 24
	fmt.Println("package testExecuteOrder.init()")
}

func GetInfo() {
	println("package testExecuteOrder.GetInfo()", "Name:", Name, "---Age:", Age)
}
