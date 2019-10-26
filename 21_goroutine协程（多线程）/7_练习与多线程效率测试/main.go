package main

import (
	"fmt"
	"time"
)

//向 intChan放入 1-100000个数
func putNum(intChan chan int) {
	for i := 1; i <= 100000; i++ {
		intChan <- i
	}
	close(intChan) //关闭intChan
}

// 从intChan取出数据，并判断是否为素数,如果是，就放入到primeChan
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool //素数标志

	for {
		num, ok := <-intChan //intChan 取不到（说明管道intChan执行完后已关闭）
		if !ok {
			break
		}
		flag = true //假设是素数
		//判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { //说明该num不是素数
				flag = false
				break
			}
		}

		if flag {
			primeChan <- num //是素数：将这个数就放入到primeChan
		}
	}

	fmt.Println("有一个primeNum 协程因为取不到数据，退出")

	//这里我们还不能关闭 primeChan因为别的协程可能还在用
	exitChan <- true //向 exitChan 写入true标志当前协程结束
}

func main() {

	intChan := make(chan int, 50000)   //1-100000输入的管道
	primeChan := make(chan int, 50000) //放入素数的管道
	exitChan := make(chan bool, 4)     //标识退出的管道

	start := time.Now().UnixNano()

	//开启一个协程，向 intChan放入 1-100000个数
	go putNum(intChan)

	//开启4个协程，从 intChan取出数据，并判断是否为素数,如果是，就放入到primeChan
	for i := 0; i < 8; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	////匿名函数协程，判断获取素数的协程是否都执行完了
	//go func() {
	//	//协程读取时阻塞式的，没有可以拿出的数据就阻塞等待
	//	for i := 0; i < 4; i++ {
	//		<-exitChan
	//	}
	//	end := time.Now().UnixNano()
	//	fmt.Println("使用协程耗时=", end-start)
	//
	//	//当我们从exitChan取到了4个结束标志，说明获取素数的协程都执行完了，关闭 primeChan
	//	close(primeChan)
	//}()

	////遍历我们的 primeChan ,把结果取出
	//for {
	//	res, ok := <-primeChan
	//	if !ok {
	//		break
	//	}
	//	//将结果输出
	//	fmt.Printf("素数=%d\n", res)
	//}

	for i := 0; i < 8; i++ {
		<-exitChan
	}
	end := time.Now().UnixNano()
	fmt.Println("使用协程耗时=", end-start)

	//------------------------------------------------------------------------------------------------------------------
	start1 := time.Now().UnixNano()
	for num := 1; num <= 100000; num++ {

		flag := true //假设是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { //说明该num不是素数
				flag = false
				break
			}
		}

		if flag {

		}

	}
	end1 := time.Now().UnixNano()
	fmt.Println("普通方法耗时=", end1-start1)

	fmt.Println("main线程退出")

}
