package main

import (
	"fmt"
)

func main() {
	fmt.Println("=============================== 1. 创建管道 =====================================================")

	// 创建一个可以存放3个int类型的管道
	var intChan chan int        //管道声明
	intChan = make(chan int, 4) //管道赋值 - 管道必须make初始化才能使用（管道的容量是固定的，不能自动增长）

	fmt.Println("=============================== 2. 管道是什么 ===================================================")

	//2. intChan是 指向队列的地址，所以channel是引用类型的
	fmt.Printf("intChan的值=%v intChan本身的地址=%v\n", intChan, &intChan) //intChan的值=0xc000096080 intChan本身的地址=0xc00008e018

	fmt.Println("=============================== 3. 管道写入数据 ==================================================")

	//写入数据
	//注意：给管道写入数据时，不能超过其容量，超出容量会报错：all goroutines are asleep - deadlock!
	intChan <- 10
	num := 211
	intChan <- num
	intChan <- 50
	intChan <- 66

	fmt.Println("=============================== 4. 管道的长度和cap(容量) ========================================")

	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan)) // 4, 4

	fmt.Println("=============================== 5. 管道中读取数据 ===============================================")

	//从channel取出数据
	<-intChan
	_ = <-intChan

	item3 := <-intChan
	fmt.Println("从管道取出的第3个数据：", item3) // 50

	item4, isExit := <-intChan
	fmt.Println("从管道取出的第4个数据：", item4, "标志：", isExit) // 66 true

	// close(intChan) // 关闭管道后写数据报错 panic: send on closed channel

	//未关闭的管道，从channel取出数据后，可以继续写入
	intChan <- 98
	fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan)) // 1, 4
	item := <-intChan
	fmt.Println("item=", item) // 98

	//在没有使用协程的情况下，如果管道数据已经被全部取出，再取就会报错：all goroutines are asleep - deadlock!
	item = <-intChan
	fmt.Println("item=", item) // fatal error: all goroutines are asleep - deadlock!
}
