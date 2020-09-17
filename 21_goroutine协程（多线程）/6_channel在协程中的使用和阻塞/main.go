package main

import (
	"fmt"
	"time"
)

func main() {

	//创建两个管道
	intChan := make(chan int, 5)       // 存放数据
	finishedChan := make(chan bool, 1) // 标志协程是否结束

	//两个自子协程通信
	go writeData(intChan)
	go readData(intChan, finishedChan)

	//等待协程执行完毕。不建议使用 time.Sleep(time.Second * 10) 这种方来增加主函数开销
	for {
		// 如果后续有协程往管道中写数据，但管道中目前没有数据可取,就会阻塞在这里循环等待；
		// 如果后续没有协程往该管道写入数据，执行到这里就会报死锁错误
		_, ok := <-finishedChan
		if !ok {
			fmt.Println("协程都执行完成")
			break
		}
	}

}

//协程：write Data
func writeData(intChan chan int) {
	for i := 1; i <= 30; i++ {
		intChan <- i //放入数据  如果放满了整个程序有读取操作但还没有执行取出就会阻塞在这里，如果整个程序没有读取操作就会死锁报错
		fmt.Println("writeData ", i)
		time.Sleep(time.Millisecond * 100) //为了演示效果
	}
	close(intChan) // 关闭channel时会写入一个false标志; 循环遍历时，会读取channel关闭的标志。
}

//协程：read data
func readData(intChan chan int, finishedChan chan bool) {

	for {
		v, ok := <-intChan // ok是管道是否已关闭的标志，对于已经关闭的通道，ok为false
		if !ok {           //通道已经关闭
			break
		}
		// 如果写入速度比读取速度快，管道存满后不会死锁报错，而是写入会等待读取后再写入
		// 如果不读取，只写入；管道存满后会死锁报错
		time.Sleep(time.Millisecond * 350)
		fmt.Printf("readData 读到数据=%v\n", v)
	}

	//读取数据完成，即整个任务完成
	finishedChan <- true // 往exitChan管道写入true告诉主程序读取结束
	close(finishedChan)
}
