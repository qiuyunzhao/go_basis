package main

import (
	"fmt"
)

func main() {

	//1. 创建一个可以存放3个int类型的管道
	var intChan chan int        //声明
	intChan = make(chan int, 3) //管道必须make初始化才能使用（管道的容量是固定的，不能自动增长）

	//2. intChan是什么（指向队列的地址，所以channel是引用类型的）
	fmt.Printf("intChan的值=%v intChan本身的地址=%v\n", intChan, &intChan) //intChan的值=0xc000096080 intChan本身的地址=0xc00008e018

	//3. 向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 50
	//如果从channel取出数据后，可以继续放入
	<-intChan
	intChan <- 98 //注意：给管道写入数据时，不能超过其容量，超出容量会报错：all goroutines are asleep - deadlock!

	//4. 看看管道的长度和cap(容量)
	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan)) // 3, 3

	//5. 从管道中读取数据
	num2 := <-intChan
	fmt.Println("num2=", num2)                                          //211
	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan)) // 2, 3

	//6. 在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报错：all goroutines are asleep - deadlock!
	num3, num4 := <-intChan
	//num5 := <-intChan

	fmt.Println("num3=", num3, "num4=", num4 /*, "num5=", num5*/)

}
