package main

import (
	"fmt"
	"time"
)

func main() {
	//------------------------------------------------ channel关闭 -----------------------------------------------------
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan) // close关闭管道，管道关闭后不能再写入数据，但是可以取出数据

	//关闭后，不能够再写入数到channel
	//intChan<- 300

	//当管道关闭后，读取数据是可以的
	n1 := <-intChan
	fmt.Println("n1=", n1)

	//------------------------------------------------ channel遍历 -----------------------------------------------------
	intChan2 := make(chan int, 30)
	for i := 1; i <= 30; i++ {
		intChan2 <- i //放入30个数据到管道
	}

	////遍历管道不能使用普通的 for 循环
	//for i := 1; i <= len(intChan2); i++ {
	//	a := <-intChan2 //取出后len(intChan2)会减一，会导致后边的数据无法从通道取出
	//	fmt.Printf("v%d=%d \n", i, a)
	//}

	go func() {
		for {
			time.Sleep(time.Second * 5)
			intChan2 <- 100000000
		}
	}()

	//关闭channel时会写入一个标志。
	//在遍历时，会读取channel关闭的标志，读不到关闭的标志，且程序有地方可能会往管道中写入数据，就会读取出所有数据后阻塞在这里；
	//在遍历时，如果channel没有关闭，并且已经不会有在网管道中写入数据，则会出现deadlock的错误；
	//在遍历时，如果channel已经关闭，则会正常遍历取出所有数据，遍历完后，就会退出遍历；

	close(intChan2)
	index := 1
	for v := range intChan2 {
		fmt.Printf("V%d=%d \n", index, v)
		index++
	}

}
