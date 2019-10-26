package main

import (
	"fmt"
	"reflect"
)

//通过反射，修改数据的值

func reflect01(b interface{}) {
	//1. 获取到 reflect.Value
	rVal := reflect.ValueOf(b)

	// 看看rVal的Kind是
	fmt.Printf("rVal kind=%v\n", rVal.Kind())

	//3. rVal
	//Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装（相当于指针用 *ptr 取值）
	rVal.Elem().SetInt(20)
}

func main() {

	var num int = 10
	reflect01(&num)
	fmt.Println("num=", num) // 20

}
