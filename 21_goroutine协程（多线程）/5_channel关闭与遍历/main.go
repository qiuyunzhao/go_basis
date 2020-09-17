package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("=============================== 1. channel关闭 ==================================================")

	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200

	close(intChan) //关闭管道 -- 管道关闭后不能再写入数据，但是可以取出数据

	//intChan<- 300 // 管道关闭后在写入报错：panic: send on closed channel
	n1 := <-intChan // 管道关闭后，可以读取数据
	fmt.Println("n1=", n1)

	fmt.Println("=============================== 2. channel遍历 ==================================================")

	intChan2 := make(chan int, 30)
	for i := 1; i <= 30; i++ {
		intChan2 <- i
	}

	////遍历管道不能使用普通的 for 循环
	//for i := 1; i <= len(intChan2); i++ {
	//	a := <-intChan2       //取出后len(intChan2)会减一，会导致后边的数据无法从通道取出
	//	fmt.Printf("v%d=%d \n", i, a)
	//}

	// 关闭channel时会写入一个false标志; range遍历时，会读取channel关闭的标志。
	// (1) 在遍历时，如果channel没有关闭，且有协程可能会往管道中写入数据，就会读取出所有数据后阻塞在这里，一直循环等待新数据写入管道后继续遍历；
	// (2) 在遍历时，如果channel没有关闭，并且已经不会有在网管道中写入数据，则会出现 all goroutines are asleep - deadlock! 的错误；
	// (3) 在遍历时，如果channel已经关闭，则会正常遍历取出所有数据，遍历完后就会退出遍历（即使还有协程会可能会写入数据）；

	go func() {
		for {
			time.Sleep(time.Second * 5)
			intChan2 <- 100000000
		}
	}()

	close(intChan2)

	index := 1
	for v := range intChan2 {
		fmt.Printf("V%d=%d \n", index, v)
		index++
	}
	fmt.Println("***********遍历结束**********")
}
