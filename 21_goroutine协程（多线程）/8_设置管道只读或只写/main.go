package main

import (
	"fmt"
)

//管道可以声明为只读或者只写
//使用场景：作为协程的形参，可以防止协程内读写权限的误操作
func main() {
	//1. 在默认情况下下，管道是双向
	var chan1 chan int //可读可写
	chan1 = make(chan int, 3)
	chan1 <- 1
	num := <-chan1
	fmt.Println("num=", num)

	//2 声明为只写
	var chan2 chan<- int
	chan2 = make(chan int, 3)
	chan2 <- 20
	//num1 := <-chan2 //编译报错

	//3. 声明为只读
	var chan3 <-chan int
	num2 := <-chan3
	//chan3<- 30 //编译报错
	fmt.Println("num2", num2)

}

//在协程中 go send(ch1,ch2)中，ch1只写,ch2只读
func send(ch1 chan<- int, ch2 <-chan int) {

}
