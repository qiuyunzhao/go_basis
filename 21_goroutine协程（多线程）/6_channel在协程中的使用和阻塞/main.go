package main

import (
	"fmt"
	"time"
)

//管道写入了没有

//协程：write Data
func writeData(intChan chan int) {
	for i := 1; i <= 30; i++ {
		intChan <- i //放入数据  如果放满了整个程序有读取操作但还没有执行取出就会阻塞在这里，如果整个程序没有读取操作就会死锁报错
		fmt.Println("writeData ", i)
		time.Sleep(time.Millisecond * 100) //为了演示效果
	}
	close(intChan) //关闭后，才可以用for range遍历
}

//协程：read data
func readData(intChan chan int, exitChan chan bool) {

	for {
		v, ok := <-intChan //对于已经关闭的通道，语句： v,ok := <-c ，会将ok置为false
		if !ok {
			//通道已经关闭
			break
		}
		//如果写入速度比读取速度快，管道存满后不会死锁报错，而是写入会等待读取后再写入
		//如果不读取，只写入；管道存满后会死锁报错
		time.Sleep(time.Millisecond * 350)
		fmt.Printf("readData 读到数据=%v\n", v)
	}

	////读取数据完成，即任务完成
	exitChan <- true //往exitChan管道写入true告诉主程序读取结束
	close(exitChan)

}

func main() {

	//创建两个管道
	intChan := make(chan int, 5)   //存放数据
	exitChan := make(chan bool, 1) //标志协程是否结束

	//两个自子协程通信
	go writeData(intChan)
	go readData(intChan, exitChan)

	//等待协程执行完毕。  不使用 time.Sleep(time.Second * 10) 这种方来增加主函数开销
	for {
		//程序中有地方往管道中写数据，但管道中目前没有数据可取,就会阻塞在这里；
		// 如果整个程序没有地方往该管道写入数据，执行到这里就会报死锁错误
		_, ok := <-exitChan
		if !ok {
			fmt.Println("协程都执行完成")
			break
		}
	}

}
