package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age  int
}

func main() {

	fmt.Println("=============================== 1. 存放任意数据类型的管道 ========================================")

	var allTypeChan chan interface{}
	allTypeChan = make(chan interface{}, 3)

	fmt.Println("=============================== 2. 管道写入数据、取出数据 ========================================")

	allTypeChan <- 10

	allTypeChan <- "tom jack"

	cat := Cat{"小花猫", 4}
	allTypeChan <- cat

	//希望获得到管道中的第三个元素，需先将前2个取出
	<-allTypeChan
	<-allTypeChan
	newCat := <-allTypeChan
	fmt.Printf("newCat类型=%T , newCat值=%v\n", newCat, newCat) // newCat类型=main.Cat , newCat值={小花猫 4}

	fmt.Println("=============================== 3. 取出数据类型转换 ==============================================")

	//fmt.Printf("newCat.Name=%v", newCat.Name)	 // 类型是interface{}，不是具体结构体类型! 编译不通过

	//使用类型断言
	a := newCat.(Cat)
	fmt.Printf("newCat.Name=%v", a.Name) // newCat.Name=小花猫
}
